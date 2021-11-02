package routes

import (
	"github.com/gorilla/mux"
	"github.com/sebastiankennedy/go-web-skeleton/app/http/controllers"
	"net/http"
)

func RegisterWebRoutes(r *mux.Router) {
	hc := new(controllers.HomeController)
	r.HandleFunc("/", hc.Index).Methods("GET").Name("home.index")

	ac := new(controllers.AdminController)
	r.HandleFunc("/admin", ac.Index).Methods("GET").Name("admin.index")
}

func RegisterResourceRoutes(r *mux.Router) {
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/img/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/webfonts/").Handler(http.FileServer(http.Dir("./public")))
}
