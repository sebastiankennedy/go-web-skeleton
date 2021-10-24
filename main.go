package main

import (
	"github.com/sebastiankennedy/go-web-skeleton/bootstrap"
	"github.com/sebastiankennedy/go-web-skeleton/config"
	"net/http"
)

// 初始化配置信息
func init() {
	config.Initialize()
}

func main() {
	router := bootstrap.SetupRouter()

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		return
	}
}
