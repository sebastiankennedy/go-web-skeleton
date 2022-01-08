package view

import (
	"github.com/sebastiankennedy/go-web-skeleton/pkg/auth"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/config"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/flash"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/logger"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/router"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

type Data map[string]interface{}

func RenderSingle(w io.Writer, data Data, tplFile string) {
	var err error

	// 读取单一视图文件
	file := getSingleFile(tplFile)
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		logger.Error(err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		logger.Error(err)
	}
}

func getSingleFile(tplFile string) string {
	// 设置视图相对路径
	viewDir := "resources/views/"

	// 返回视图完整路径
	return viewDir + strings.Replace(tplFile, ".", "/", -1) + ".gohtml"
}

func RenderAdmin(w io.Writer, data Data, tplFiles ...string) {
	RenderAdminTemplate(w, "app", data, tplFiles...)
}

func RenderAdminTemplate(w io.Writer, name string, data Data, tplFiles ...string) {
	var err error

	// 通用模板数据
	data["isLogined"] = auth.Check()
	data["loginUser"] = auth.User()
	data["flash"] = flash.All()

	// 生成模板文件
	allFiles := getAdminTemplateFiles(tplFiles...)

	// 解析模板文件
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"RouteNameToUrl": router.NameToUrl,
		"Env":            config.Env,
	}).ParseFiles(allFiles...)
	logger.Error(err)

	// 渲染模板
	err = tmpl.ExecuteTemplate(w, name, data)
	if err != nil {
		logger.Error(err)
	}
}

func getHomeTemplateFiles(tplFiles ...string) []string {
	return getTemplateFiles("home", tplFiles...)
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
