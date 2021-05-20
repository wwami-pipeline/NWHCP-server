package handlers

import (
	"database/sql"
	"pipeline-db/models/orgs"
)

type HandlerContext struct {
	OrgStore orgs.Store
	Db       *sql.DB
	UserID   int64
}

type User struct {
	UserID    int64
	FirstName string
	LastName  string
	Orgs      []UserOrgs
}

type Creds struct {
	ID        int64
	FirstName string
	LastName  string
	JoinDate  string
}

type UserOrgs struct {
	UserID     int64
	OrgID      int64
	OrgTitle   string
	OrgWebsite string
	OrgCity    string
	OrgState   string
	OrgZipcode string
	OrgPhone   string
}
