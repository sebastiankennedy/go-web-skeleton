package controllers

import (
	"fmt"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/config"
	"net/http"
)

type AdminController struct {
	HttpController
}

func (*AdminController) Index(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello, "+config.GetString("app.name")+" Backend.")
	if err != nil {
		return
	}
}
