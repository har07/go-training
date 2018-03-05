package helpers

import (
	"os"
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var (
	nfLog *log.Entry
	regex *regexp.Regexp
)

func init() {

	// Log default fields
	nfLog = log.WithFields(log.Fields{
		"app":    os.Getenv("APP_NAME"),
		"type":   "backend",
		"tenant": "none",
	})

	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	regex = reg
}

// Response JSONAPI object
type Response struct {
	Errors    []Error     `json:"errors,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Links     interface{} `json:"links,omitempty"`
	Meta      interface{} `json:"meta,omitempty"`
	Successes []Success   `json:"success,omitempty"`
}

// Error object
type Error struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail,omitempty"`
}

// Links object
type Links struct {
	Self  string `json:"self,omitempty"`
	First string `json:"first"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Last  string `json:"last"`
}

// Success object
type Success struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail,omitempty"`
	// Port        string `json:"port,omitempty"`
	// Node        string `json:"node,omitempty"`
	// IP          string `json:"ip,omitempty"`
}

// MakeErrorResponse is function to generate error response
func MakeErrorResponse(tenant string, message string, err error, status int) *Response {
	nfLog.WithFields(log.Fields{
		"tenant": tenant,
	}).Errorln(err)

	r := new(Response)
	es := make([]Error, 1)
	es[0] = Error{Status: status, Title: err.Error(), Detail: message}
	r.Errors = es
	return r
}

// GenerateLinks is
func GenerateLinks(host string, totalPages int, currentPage int, pageSize int, filterType string) *Links {
	l := new(Links)
	filterStr := ""
	if len(filterType) > 0 {
		filterStr = "&filter%5Btype%5D=" + filterType
	}

	l.Self = host + "?page%5Bnumber%5D=" + strconv.Itoa(currentPage) + "&page%5Bsize%5D=" + strconv.Itoa(pageSize) + filterStr
	l.First = host + "?page%5Bnumber%5D=1&page%5Bsize%5D=" + strconv.Itoa(pageSize) + filterStr
	l.Last = host + "?page%5Bnumber%5D=" + strconv.Itoa(totalPages) + "&page%5Bsize%5D=" + strconv.Itoa(pageSize) + filterStr
	if currentPage < totalPages {
		l.Next = host + "?page%5Bnumber%5D=" + strconv.Itoa(currentPage+1) + "&page%5Bsize%5D=" + strconv.Itoa(pageSize) + filterStr
	}
	if currentPage > 1 {
		l.Prev = host + "?page%5Bnumber%5D=" + strconv.Itoa(currentPage-1) + "&page%5Bsize%5D=" + strconv.Itoa(pageSize) + filterStr
	}
	return l
}

// MakeSuccessResponse is function to generate success response
func MakeSuccessResponse(tenant string, message string, status int) *Response {

	nfLog.WithFields(log.Fields{
		"tenant": tenant,
	}).Infoln(message)

	r := new(Response)
	es := make([]Success, 1)
	es[0] = Success{Status: status, Title: "Success", Detail: message}
	r.Successes = es
	return r
}
