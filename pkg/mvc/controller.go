package mvc

import (
	"fmt"
	"net/http"
)

type Controller struct {
}

func (*Controller) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, err := fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
	if err != nil {
		return
	}
}
