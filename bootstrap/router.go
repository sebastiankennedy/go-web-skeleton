package bootstrap

import (
	"github.com/gorilla/mux"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/router"
	"github.com/sebastiankennedy/go-web-skeleton/routes"
)

func SetupRouter() *mux.Router {
	/*
	 * 等同
	 * type Router struct {}
	 * var router *Router
	 * router = new(Router)
	 */
	muxRouter := mux.NewRouter()
	routes.RegisterWebRoutes(muxRouter)
	router.SetRouter(muxRouter)

	return muxRouter
}
