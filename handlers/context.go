package handlers

import "pipeline-db/stores"

type HandlerContext struct {
	OrgStore stores.Store
}
