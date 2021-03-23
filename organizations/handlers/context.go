package handlers

import (
	"database/sql"
	"pipeline-db/models/orgs"
)

type HandlerContext struct {
	OrgStore orgs.Store
	Db       *sql.DB
}
