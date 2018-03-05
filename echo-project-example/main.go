package main

import (
	"net/http"
	"os"

	logrus_stack "github.com/Gurpartap/logrus-stack"
	nfmiddleware "github.com/har07/go-training/echo-project-example/middleware"
	"github.com/har07/go-training/echo-project-example/model"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"gopkg.in/resty.v1"
)

var nfLog *log.Entry

func init() {
	// set logger Format
	log.SetFormatter(&log.JSONFormatter{})
	log.AddHook(logrus_stack.StandardHook())

	// Log default fields
	nfLog = log.WithFields(log.Fields{
		"app":  "hello-api",
		"type": "backend",
	})
}

func main() {

	// Its possible to specify env file
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}

	// Load env data
	err := godotenv.Load(envFile)
	if err != nil {
		nfLog.Fatal(err.Error())
		panic(err)
	}

	e := echo.New()
	e.HideBanner = true
	e.Validator = model.NewValidator()

	e.Use(nfmiddleware.Logger(os.Getenv("APP_NAME")))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	setupHandlers(e)

	servicePort := os.Getenv("PORT")
	nfLog.Info("Training API Started at Port " + servicePort)
	e.Start(servicePort)
}

func setupHandlers(e *echo.Echo) {
	r := e.Group("/:tenant/api")

	r.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello "+c.Param("tenant")+"! This is API version: "+os.Getenv("VERSION"))
	})

	r.GET("/panic", func(c echo.Context) error {
		panic("Waaaaaa!!!!!")
	})

	r.GET("/error", func(c echo.Context) error {
		return echo.NewHTTPError(http.StatusBadRequest, "Please correct your input")
	})

	r.POST("/binding", func(c echo.Context) error {
		u := new(model.User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, u)
	})

	r.POST("/hello-upstream", helloUpstream)

	r.POST("/hello-validate", helloValidate)

	r.POST("/hello-log", helloLog)
}

func helloUpstream(c echo.Context) (err error) {
	user := new(model.User)
	if err = c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Request body is not a valid User data")
	}
	resp, err := resty.R().
		SetBody(`{}`).
		SetResult(model.HelloMessage{}).
		Post("https://api-gateway.poc.apigateway.us/hello/v1.1/Hello2")

	if err != nil {
		return err
	}
	result := resp.Result().(*model.HelloMessage)
	result.Hello = "Hello " + user.Name
	return c.JSON(http.StatusOK, result)
}

func helloValidate(c echo.Context) (err error) {
	user := new(model.User)
	if err = c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Request body is not a valid User data")
	}
	if err = c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	resp, err := resty.R().
		SetBody(`{}`).
		SetResult(model.HelloMessage{}).
		Post("https://api-gateway.poc.apigateway.us/hello/v1.1/Hello2")

	if err != nil {
		return err
	}
	result := resp.Result().(*model.HelloMessage)
	result.Hello = "Hello " + user.Name
	return c.JSON(http.StatusOK, result)
}

func helloLog(c echo.Context) (err error) {
	logFields := log.Fields{
		"app":     "hello-api",
		"type":    "backend",
		"package": "main",
		"method":  "helloLog",
	}
	user := new(model.User)
	if err = c.Bind(user); err != nil {
		logFields["error"] = err.Error()
		log.WithFields(logFields).Info("Failed to bind request body to model")
		return echo.NewHTTPError(http.StatusBadRequest, "Request body is not a valid User data")
	}
	if err = c.Validate(user); err != nil {
		logFields["error"] = err.Error()
		logFields["user"] = user
		log.WithFields(logFields).Info("Validation failed")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	upstream := "https://api-gateway.poc.apigateway.us/hello/v1.1/Hello2"
	resp, err := resty.R().
		SetBody(`{}`).
		SetResult(model.HelloMessage{}).
		Post(upstream)

	if err != nil {
		logFields["error"] = err.Error()
		logFields["url"] = user
		log.WithFields(logFields).Info("Upstream error")
		return echo.NewHTTPError(http.StatusInternalServerError, "Upstream error")
	}
	result := resp.Result().(*model.HelloMessage)
	result.Hello = "Hello " + user.Name
	return c.JSON(http.StatusOK, result)
}
