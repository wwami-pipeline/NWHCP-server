package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pipeline-db/handlers"
	"pipeline-db/stores"

	mgo "gopkg.in/mgo.v2"
)

func main() {

	// Set up for microservice information

	portAddr := os.Getenv("PORTADDR") //nwhealthcareerpath.uw.edu:4000 //capstone-test.andrewkan.me:4000
	if len(portAddr) == 0 {
		portAddr = "localhost:4002"
	}
	log.Printf("PORTADDR: %s", portAddr)

	internalPort := os.Getenv("INTERNAL_PORT")
	if len(internalPort) == 0 {
		internalPort = ":4003"
	}

	dbAddr := os.Getenv("DBADDR") //pipelineDB:27017
	if len(dbAddr) == 0 {
		dbAddr = "localhost:27017"
	}
	log.Printf("DBADDR: %s", dbAddr)

	mongoSession, err := mgo.Dial(dbAddr)
	if err != nil {
		fmt.Println("Error dialing dbaddr: ", err)
	} else {
		fmt.Println("Success!")
	}
	//schoolStore, err := stores.NewSchoolStore(mongoSession, "mongodb", "school")
	orgStore, err := stores.NewOrgStore(mongoSession, "mongodb", "organization")

	hctx := &handlers.HandlerContext{
		OrgStore: orgStore,
	}

	//ignore compiler
	// log.Printf("%v", hctx)
	// log.Printf("%v", schoolStore)
	// log.Printf("%v", orgStore)

	// TLSCERT := os.Getenv("TLSCERT")
	// TLSKEY := os.Getenv("TLSKEY")

	if err != nil {
		fmt.Println("Error creating store")
	}

	apiEndpoint := "/api/v1"

	mux := http.NewServeMux()
	fmt.Println("Pipeline-DB Microservice")

	mux.HandleFunc(apiEndpoint+"/search", hctx.SearchOrgsHandler)
	mux.HandleFunc(apiEndpoint+"/orgs", hctx.GetAllOrgs)
	mux.HandleFunc(apiEndpoint+"/org/", hctx.SpecificOrgHandler)

	mux2 := http.NewServeMux()
	mux2.HandleFunc(apiEndpoint+"/pipeline-db/truncate", hctx.DeleteAllOrgsHandler)
	mux2.HandleFunc(apiEndpoint+"/pipeline-db/poporgs", hctx.InsertOrgs)
	go serve(mux2, internalPort)

	// mux.HandleFunc(apiEndpoint+"/post-test", handlers.HandlePost)
	log.Printf("server listening at http://%s...", portAddr)
	// log.Fatal(http.ListenAndServeTLS(portAddr, TLSCERT, TLSKEY, mux))
	log.Fatal(http.ListenAndServe(portAddr, handlers.AddCORS(mux)))

}

func serve(mux *http.ServeMux, addr string) {
	log.Fatal(http.ListenAndServe(addr, handlers.AddCORS(mux)))
	log.Printf("server is listening at %s...", addr)
}