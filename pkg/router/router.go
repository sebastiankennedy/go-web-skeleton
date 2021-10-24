package router

import (
	"github.com/gorilla/mux"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/config"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/logger"
)

var router *mux.Router

func SetRouter(r *mux.Router) {
	router = r
}

func NameToUrl(routeName string, pairs ...string) string {
	url, err := router.Get(routeName).URL(pairs...)
	if err != nil {
		logger.Error(err)
	}

	return config.GetString("app.url") + url.String()
}
