package routes

import (
	"github.com/gorilla/mux"
	"github.com/sebastiankennedy/go-web-skeleton/app/http/controllers"
	"net/http"
)

func RegisterWebRoutes(r *mux.Router) {
	hone := new(controllers.HomeController)
	r.HandleFunc("/", hone.Index).Methods("GET").Name("home.index")

	admin := new(controllers.AdminController)
	r.HandleFunc("/admin", admin.Index).Methods("GET").Name("admin.index")

	auth := new(controllers.AuthController)
	r.HandleFunc("/auth/login", auth.Login).Methods("GET").Name("auth.login")
	r.HandleFunc("/auth/register", auth.Register).Methods("GET").Name("auth.register")
	r.HandleFunc("/admin/auth/register", auth.RegisterView).Methods("GET").Name("admin.auth.register_view")
	r.HandleFunc("/admin/auth/register", auth.RegisterOperation).Methods("POST").Name("admin.auth.register_operation")
}

func RegisterResourceRoutes(r *mux.Router) {
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/img/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/webfonts/").Handler(http.FileServer(http.Dir("./public")))
}
