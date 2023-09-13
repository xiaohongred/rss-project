package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, 200, "hi welcome")
}

func handlerErr(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 400, "err msg")
}
