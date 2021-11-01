package controllers

import (
	"github.com/sebastiankennedy/go-web-skeleton/pkg/mvc"
	"net/http"
)

type AdminController struct {
	mvc.Controller
}

func (*AdminController) Index(w http.ResponseWriter, r *http.Request) {
	mvc.RenderAdmin(w, nil, "dashboard.index")
}
