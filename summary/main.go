package main

import (
	// "assignments-annazhoufast/servers/summary/"

	"log"
	"net/http"
	"os"
)

//main is the main entry point for the server
func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/summary", SummaryHandler)
	log.Printf("server is listening at http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
