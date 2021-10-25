package controllers

import (
	"fmt"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/config"
	"net/http"
)

type HomeController struct {
	HttpController
}

func (*HomeController) Index(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello, "+config.GetString("app.name"))
	if err != nil {
		return
	}
}
