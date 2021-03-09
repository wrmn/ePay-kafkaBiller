package main

import "github.com/gorilla/mux"

// return HTTP handler
func server() *mux.Router {

	// create new handler instance
	router := mux.NewRouter()

	return router
}
