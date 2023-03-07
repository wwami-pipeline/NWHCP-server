package handlers

import "net/http"

func PostError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}
