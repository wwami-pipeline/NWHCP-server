package handlers

import (
	// "info441-finalproj/servers/gateway/models/users"
	// "NWHCP-server/gateway/models/users"
	// "NWHCP/NWHCP-server/gateway/models/users"
	"nwhcp/nwhcp-server/gateway/models/users"
	"time"
	// "github.com/nwhcp-server/gateway/models/users"
)

// SessionState blah
type SessionState struct {
	Time time.Time  `json:"time"`
	User users.User `json:"user"`
}
