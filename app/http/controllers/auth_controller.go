package controllers

import (
	"fmt"
	"github.com/sebastiankennedy/go-web-skeleton/app/http/requests"
	"github.com/sebastiankennedy/go-web-skeleton/app/models/user"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/auth"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/config"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/controller"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/router"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/view"
	"net/http"
)

type AuthController struct {
	controller.Controller
}

func (*AuthController) LoginView(w http.ResponseWriter, r *http.Request) {
	data := view.Data{
		"RegisterViewUrl":   router.NameToUrl("admin.auth.register_view"),
		"LoginOperationUrl": router.NameToUrl("admin.auth.login_operation"),
	}

	view.RenderSingle(w, data, "admin.auth.login")
}

func (*AuthController) LoginOperation(w http.ResponseWriter, r *http.Request) {
	// 获取数据
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	if err := auth.Attempt(email, password); err == nil {
		// 登录成功
		http.Redirect(w, r, "/admin", http.StatusFound)
	} else {
		data := view.Data{
			"Error":    err.Error(),
			"Email":    email,
			"Password": password,
		}
		view.RenderSingle(w, data, "admin.auth.login")
	}
}

func (*AuthController) RegisterView(w http.ResponseWriter, r *http.Request) {
	data := view.Data{
		"LoginViewUrl":         router.NameToUrl("admin.auth.login_view"),
		"RegisterOperationUrl": router.NameToUrl("admin.auth.register_operation"),
		"AppName":              config.Env("APP_NAME"),
	}

	view.RenderSingle(w, data, "admin.auth.register")
}

func (*AuthController) RegisterOperation(w http.ResponseWriter, r *http.Request) {
	// 获取数据
	_user := user.User{
		Username:             r.PostFormValue("username"),
		Email:                r.PostFormValue("email"),
		Password:             r.PostFormValue("password"),
		PasswordConfirmation: r.PostFormValue("password_confirmation"),
	}

	// 表单规则验证数据
	errs := requests.ValidateRegistrationForm(_user)
	if len(errs) > 0 {
		// 表单验证失败，重新显示表单
		view.RenderSingle(w, view.Data{
			"Errors":  errs,
			"User":    _user,
			"AppName": config.Env("APP_NAME"),
		}, "admin.auth.register")
	} else {
		err := _user.Create()
		if err != nil {
			return
		}

		if _user.ID > 0 {
			// 跳转至首页
			http.Redirect(w, r, "/admin", http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "注册失败，请联系管理员")
		}

	}

	// 验证失败,重新显示表单
}
