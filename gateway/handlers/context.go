package handlers

import (
	"nwhcp/nwhcp-server/gateway/models/users"
	"nwhcp/nwhcp-server/gateway/sessions"
	// "github.com/nwhcp-server/gateway/models/users"
	// "github.com/nwhcp-server/gateway/sessions"
	// "NWHCP/NWHCP-server/gateway/models/users"
	// "NWHCP/NWHCP-server/gateway/sessions"
)

// Handler blah
type Handler struct {
	SessionKey   string         `json:"key"`
	SessionStore sessions.Store `json:"sessions"`
	UserStore    users.Store    `json:"users"`
}
