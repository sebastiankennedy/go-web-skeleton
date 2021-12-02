package user

import (
	"github.com/sebastiankennedy/go-web-skeleton/pkg/logger"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/model"
	"github.com/sebastiankennedy/go-web-skeleton/pkg/types"
)

func (user *User) Create() (err error) {
	if err = model.DB.Create(&user).Error; err != nil {
		logger.Error(err)
		return err
	} else {
		logger.Info("Create User Success.")
		return nil
	}
}

func Get(str string) (User, error)  {
	var user User
	id := types.StringToInt(str)
	if err := model.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GetByEmail(email string) (User, error) {
	var user User
	if err := model.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
