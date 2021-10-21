package main

import (
	"github.com/gorilla/mux"
	"github.com/sebastiankennedy/go-web-skeleton/bootstrap"
	"net/http"
)

var router *mux.Router

func init() {

}


func main() {
	router = bootstrap.SetupRouter()

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		return
	}
}