package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"theychat/dao"
	"theychat/utils"
)

type UserLoginArgs struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}

type UserRegisterArgs struct {
	Name     string `form:"name"`
	Password string `form:"password"`
	Email    string `form:"email"`
}

func UserLogin(args *UserLoginArgs) *Resp {
	user, ok := dao.GetUser("name", args.Name)
	if ok && utils.CheckPassword(args.Password, user.Password) {
		token, _ := utils.GenToken(&utils.UserClaims{
			ID:             user.ID,
			Name:           user.Name,
			Email:          user.Email,
			StandardClaims: jwt.StandardClaims{},
		})
		//c.Header("Authorization", token)
		return RespOk(gin.H{
			"token": token,
		})
	} else {
		return RespErr(400, "error password")
	}
}

func UserRegister(args *UserRegisterArgs) *Resp {
	user, ok, msg := dao.CreateUser(args.Name, args.Email, args.Password)
	if ok {
		return RespOk(user)
	} else {
		return RespErr(403, msg)
	}
}
