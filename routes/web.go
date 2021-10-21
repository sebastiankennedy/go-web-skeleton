package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World.")
}

func RegisterWebRoutes(r *mux.Router)  {
	r.HandleFunc("/", defaultHandler).Methods("GET").Name("home")
}