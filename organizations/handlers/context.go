package handlers

import "pipeline-db/models/orgs"

type HandlerContext struct {
	OrgStore orgs.Store
}
