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

func AddUser(name, password string) error {
	users := []model.User{}
	err := pkg.DB.Find(&users, "name = ?", name).Table("users").Error
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return fmt.Errorf("user name exist")
	}

	return pkg.DB.Table("users").Create(&model.User{
		Name:     name,
		Password: password,
	}).Error
}
