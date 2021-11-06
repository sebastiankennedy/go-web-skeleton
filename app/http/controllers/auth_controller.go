package controllers

import (
	"github.com/sebastiankennedy/go-web-skeleton/pkg/controller"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/view"
	"net/http"
)

type AuthController struct {
	controller.Controller
}

func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {
	view.RenderSingle(w, nil, "admin.auth.login")
}

func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSingle(w, nil, "admin.auth.register")
}
