package requests

import (
	"github.com/sebastiankennedy/go-web-skeleton/app/models/user"
	"github.com/thedevsaddam/govalidator"
)

// ValidateLoginForm 验证表单，返回 errs 长度等于零即通过
func ValidateLoginForm(data user.User) map[string][]string {

	// 1. 定制认证规则
	rules := govalidator.MapData{
		"email":    []string{"required", "min:4", "max:30", "email"},
		"password": []string{"required", "min:6"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"password": []string{
			"required:密码为必填项",
			"min:长度需大于 6",
		},
	}

	// 3. 配置初始化
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	// 4. 开始验证
	errs := govalidator.New(opts).ValidateStruct()

	return errs
}
