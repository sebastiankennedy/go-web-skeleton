package controllers

import (
	"github.com/sebastiankennedy/go-web-skeleton/pkg/logger"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/mvc"
	"html/template"
	"net/http"
	"path/filepath"
)

type AdminController struct {
	mvc.Controller
}

func (*AdminController) Index(w http.ResponseWriter, r *http.Request) {
	// 设置模板相对路径
	viewDir := "resources/views"

	// 所有布局模板文件 Slice
	files, err := filepath.Glob(viewDir + "/admin/layouts/*.gohtml")

	// 在 Slice 里新增我们的目标文件
	newFiles := append(files, viewDir+"/admin/dashboard/index.gohtml")

	// 解析模板文件
	tmpl, err := template.ParseFiles(newFiles...)
	logger.Error(err)

	// 渲染模板，将数据传输进去
	err = tmpl.ExecuteTemplate(w, "app", nil)
	if err != nil {
		logger.Error(err)
	}
}
