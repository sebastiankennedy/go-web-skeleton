package router

import "github.com/gorilla/mux"

var router *mux.Router

func SetRouter(r *mux.Router) {
	router = r
}

