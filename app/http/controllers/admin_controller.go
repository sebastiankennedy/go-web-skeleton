package controllers

import (
	"github.com/sebastiankennedy/go-web-skeleton/pkg/controller"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/view"
	"net/http"
)

type AdminController struct {
	controller.Controller
}

func (*AdminController) Index(w http.ResponseWriter, r *http.Request) {
	view.RenderAdmin(w, nil, "dashboard.index")
}
