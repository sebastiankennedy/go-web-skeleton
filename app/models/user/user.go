package user

import (
	"github.com/sebastiankennedy/go-web-skeleton/app/models"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/password"
)

type User struct {
	models.Model

	Email    string `gorm:"type:varchar(255);not null;unique;" valid:"email"`
	Username string `gorm:"type:varchar(255);unique;" valid:"username"`
	Password string `gorm:"type:varchar(255);" valid:"password"`

	// gorm:"-" == 设置 GORM 在读写时略过当前字段，仅仅用于表单验证
	PasswordConfirmation string `gorm:"-" valid:"password_confirmation"`
}

// ComparePassword 对比密码是否匹配
func (user *User) ComparePassword(_password string) bool {
	return password.CheckHash(_password, user.Password)
}
