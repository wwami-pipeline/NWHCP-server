package handlers

import (
	"pipeline-db/models/orgs"
	"database/sql"
)

type HandlerContext struct {
	OrgStore orgs.Store
	dbStore *sql.DB	
}
