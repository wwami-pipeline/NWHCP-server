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
	schoolStore, err := stores.NewSchoolStore(mongoSession, "mongodb", "school")
	orgStore, err := stores.NewOrganizationStore(mongoSession, "mongodb", "organization")

	hctx := &handlers.Context{
		Store1: schoolStore,
		Store2: orgStore,
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

	apiEndpoint := ""
	if os.Getenv("PRODUCTION_MODE") != "production" {
		apiEndpoint = "/api/v1"
	}

	mux := http.NewServeMux()
	fmt.Println("Pipeline-DB Microservice")
	mux.HandleFunc(apiEndpoint+"/pipeline-db/populate-school", hctx.PopulateDB)
	mux.HandleFunc(apiEndpoint+"/pipeline-db/populate-organizations", hctx.PopulateDB2)

	mux.HandleFunc(apiEndpoint+"/pipeline-db/getallschools", hctx.SchoolHandler)
	mux.HandleFunc(apiEndpoint+"/pipeline-db/getallorgs", hctx.OrgHandler)

	mux.HandleFunc(apiEndpoint+"/post-test", handlers.HandlePost)
	log.Printf("server listening at http://%s...", portAddr)
	// log.Fatal(http.ListenAndServeTLS(portAddr, TLSCERT, TLSKEY, mux))
	log.Fatal(http.ListenAndServe(portAddr, mux))
}
