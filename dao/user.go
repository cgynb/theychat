package dao

import (
	"theychat/model"
	"theychat/utils"
)

func GetUser(by, info string) (*model.User, bool) {
	var u model.User
	result := DB.Where(by+" = ?", info).First(&u)
	return &u, result.Error == nil
}

func CreateUser(name, email, pwd string) (u *model.User, ok bool, msg string) {
	hpwd, err := utils.GenHashPassword(pwd)
	if err != nil {
		return nil, false, "please register again"
	}
	u = &model.User{Name: name, Password: hpwd, Email: email}
	result := DB.Create(u)
	if result.Error != nil {
		return nil, false, "username has been used"
	}
	return u, true, "ok"
}
