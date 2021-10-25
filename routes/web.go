package routes

import (
	"github.com/gorilla/mux"
	"github.com/sebastiankennedy/go-web-skeleton/app/http/controllers"
)

func RegisterWebRoutes(r *mux.Router) {
	hc := new(controllers.HomeController)
	r.HandleFunc("/", hc.Index).Methods("GET").Name("home.index")

	ac := new(controllers.AdminController)
	r.HandleFunc("/admin", ac.Index).Methods("GET").Name("admin.index")
}
