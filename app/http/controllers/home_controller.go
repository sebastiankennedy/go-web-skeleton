package controllers

import (
	"fmt"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/config"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/mvc"
	"net/http"
)

type HomeController struct {
	mvc.Controller
}

func (*HomeController) Index(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello, "+config.GetString("app.name"))
	if err != nil {
		return
	}
}
