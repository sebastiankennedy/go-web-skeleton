package middlewares

import (
	"github.com/sebastiankennedy/go-web-skeleton/pkg/auth"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/flash"
	"net/http"
)

func Guest(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if auth.Check() {
			flash.Warning("登录用户无法访问此页面")
			http.Redirect(writer, request, "/admin", http.StatusFound)
			return
		}

		next(writer, request)
	}
}
