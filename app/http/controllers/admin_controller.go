package controllers

import (
	"github.com/sebastiankennedy/go-web-skeleton/pkg/mvc"
	"net/http"
	"text/template"
)

type AdminController struct {
	mvc.Controller
}

func (*AdminController) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("resources/views/admin/layouts/app.gohtml")
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(w, "app", nil)
	if err != nil {
		return
	}
}
