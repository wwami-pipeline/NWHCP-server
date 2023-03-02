package handlers

// "github.com/nwhcp-server/gateway/models/users"

// "github.com/nwhcp-server/gateway/sessions"

// "NWHCP/NWHCP-server/gateway/sessions"

// Handler blah
type Handler struct {
	SessionKey       string        `json:"key"`
	ThisSessionStore *SessionStore `json:"sessions"`
	UserStore        *UserStore    `json:"users"`
}

type HandlerContext struct {
	OrgStore *OrgStore
}
