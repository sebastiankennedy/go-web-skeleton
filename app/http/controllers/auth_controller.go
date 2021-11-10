package controllers

import (
	"fmt"
	"github.com/sebastiankennedy/go-web-skeleton/app/models/user"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/controller"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/router"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/view"
	"net/http"
)

type AuthController struct {
	controller.Controller
}

func (*AuthController) LoginView(w http.ResponseWriter, r *http.Request) {
	view.RenderSingle(w, nil, "admin.auth.login")
}

func (*AuthController) LoginOperation(w http.ResponseWriter, r *http.Request) {

}

func (*AuthController) RegisterView(w http.ResponseWriter, r *http.Request) {
	data := view.Data{
		"LoginViewUrl":         router.NameToUrl("admin.auth.login_view"),
		"RegisterOperationUrl": router.NameToUrl("admin.auth.register_operation"),
	}

	view.RenderSingle(w, data, "admin.auth.register")
}

func (*AuthController) RegisterOperation(w http.ResponseWriter, r *http.Request) {
	// 表单验证
	email := r.PostFormValue("email")
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	// 验证通过 -- 入库
	_user := user.User{
		Username: username,
		Password: password,
		Email:    email,
	}
	err := _user.Create()
	if err != nil {
		return
	}

	if _user.ID > 0 {
		fmt.Fprint(w, "注册成功，电子邮箱为 为"+_user.Email)
	} else {
		fmt.Fprint(w, "创建用户失败，请联系管理员")
	}

	// 验证失败 -- 重新显示表单
}
