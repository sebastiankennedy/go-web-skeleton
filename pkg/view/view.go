package view

import (
	"github.com/sebastiankennedy/go-web-skeleton/pkg/logger"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/router"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

type Data map[string]interface{}

func RenderAdmin(w io.Writer, data Data, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)
}

func RenderTemplate(w io.Writer, name string, data Data, tplFiles ...string) {
	var err error

	// 生成模板文件
	allFiles := getAdminTemplateFiles(tplFiles...)

	// 解析模板文件
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"RouteNameToUrl": router.NameToUrl,
	}).ParseFiles(allFiles...)
	logger.Error(err)

	// 渲染模板
	err = tmpl.ExecuteTemplate(w, name, data)
	if err != nil {
		return
	}
}

func getAdminTemplateFiles(tplFiles ...string) []string {
	return getTemplateFiles("admin", tplFiles...)
}

func getTemplateFiles(subViewDir string, tplFiles ...string) []string {
	// 设置模板相对路径
	viewDir := "resources/views/" + subViewDir

	// 判断目录是否以 / 字符结尾
	if !strings.HasSuffix(viewDir, "/") {
		viewDir = viewDir + "/"
	}

	// 遍历文件列表 Slice，设置正确的路径，支持 dir.filename 语法糖
	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	// 所有布局模板文件 Slice
	layoutFiles, err := filepath.Glob(viewDir + "/layouts/*.gohtml")
	logger.Error(err)

	// 合并所有文件
	return append(layoutFiles, tplFiles...)
}
