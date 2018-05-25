package handlers

import "github.com/pipeline-db/stores"

type Context struct {
	Store1 stores.Store
	Store2 stores.Store2
}
