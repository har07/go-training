package model

// User is
type User struct {
	ID             string                 `json:"id,omitempty"`
	TenantID       string                 `json:"tenant_id,omitempty"`
	UserID         string                 `json:"user_id,omitempty"`
	Username       string                 `json:"username,omitempty"`
	Email          string                 `json:"email,omitempty"`
	Name           string                 `json:"name,omitempty" validate:"required"`
	PositionID     string                 `json:"position_id,omitempty"`
	Position       string                 `json:"position,omitempty"`
	Organization   string                 `json:"organization,omitempty"`
	Groups         []string               `json:"groups"`
	ProfilePicture string                 `json:"profile_picture,omitempty"`
	UserMetadata   map[string]interface{} `json:"user_metadata"`
	AppMetadata    map[string]interface{} `json:"app_metadata"`
	Suspended      bool                   `json:"is_suspended"`
}

// HelloMessage is
type HelloMessage struct {
	Hello   string  `json:"hello"`
	Status  string  `json:"status"`
	Counter float32 `json:"counter"`
}
