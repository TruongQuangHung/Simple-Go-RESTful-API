package repo

import (
	"fmt"

	"tekdraw-bff/model"
	"tekdraw-bff/pkg"
)

func GetUser(name, password string) (*model.User, error) {
	user := model.User{}
	err := pkg.DB.Table("users").Where(
		"name = ? and password = ?", name, password,
	).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func AddUser(name, password string) (int, error) {
	var count int64
	err := pkg.DB.Table("users").Where("name = ?", name).Count(&count).Error
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, fmt.Errorf("user name exists")
	}

	user := model.User{
		Name:     name,
		Password: password,
	}
	result := pkg.DB.Table("users").Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.Id, nil
}
