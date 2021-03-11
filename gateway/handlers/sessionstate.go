package handlers

import (
	// "info441-finalproj/servers/gateway/models/users"
	"NWHCP-server/gateway/models/users"
	"time"
)

// SessionState blah
type SessionState struct {
	Time time.Time  `json:"time"`
	User users.User `json:"user"`
}
