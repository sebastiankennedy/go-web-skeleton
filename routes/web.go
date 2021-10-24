package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/config"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, "+config.GetString("app.name"))
}

func RegisterWebRoutes(r *mux.Router) {
	r.HandleFunc("/", defaultHandler).Methods("GET").Name("home")
}
