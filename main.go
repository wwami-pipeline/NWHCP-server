package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pipeline-db/handlers"
	"github.com/pipeline-db/stores"

	mgo "gopkg.in/mgo.v2"
)

func main() {

	// Set up for microservice information

	portAddr := os.Getenv("PORTADDR") //nwhealthcareerpath.uw.edu:4000 //capstone-test.andrewkan.me:4000
	print(portAddr)
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
		fmt.Println("Error dialing dbaddr")
	} else {
		fmt.Println("Success!")
	}
	schoolStore, err := stores.NewSchoolStore(mongoSession, "school", "mongodb")
	orgStore, err := stores.NewOrganizationStore(mongoSession, "organization", "mongodb")

	hctx := &handlers.Context{
		Store1: schoolStore,
		Store2: orgStore,
	}

	//ignore compiler
	log.Printf("%v", hctx)
	log.Printf("%v", schoolStore)
	log.Printf("%v", orgStore)

	// TLSCERT := os.Getenv("TLSCERT")
	// TLSKEY := os.Getenv("TLSKEY")

	if err != nil {
		fmt.Println("Error creating store")
	}

	mux := http.NewServeMux()
	fmt.Println("Pipeline-DB Microservice")
	mux.HandleFunc("/v1/pipeline-db/populate-school", hctx.PopulateDB)
	mux.HandleFunc("/v1/pipeline-db/populate-organizations", hctx.PopulateDB2)

	mux.HandleFunc("/v1/pipeline-db/getallschools", hctx.SchoolHandler)
	mux.HandleFunc("/v1/pipeline-db/getallorgs", hctx.OrgHandler)

	log.Printf("server listening at http://%s...", portAddr)
	// log.Fatal(http.ListenAndServeTLS(portAddr, TLSCERT, TLSKEY, mux))
	log.Fatal(http.ListenAndServe(portAddr, mux))
}
