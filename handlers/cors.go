package handlers

import (
	"net/http"
)

// CORS adds CORS header
type CORS struct {
	handler http.Handler
}

//ServeHTTP handles the request by passing it to the real
//handler and add cors header
func (c *CORS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, PUT, POST, PATCH, DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Add("Access-Control-Expose-Headers", "Authorization")
	w.Header().Add("Access-Control-Max-Age", "600")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	c.handler.ServeHTTP(w, r)

}

//AddCORS constructs a new middleware handler
func AddCORS(handlerToWrap http.Handler) *CORS {
	return &CORS{handlerToWrap}
}
