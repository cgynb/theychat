package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"theychat/dao"
	"theychat/utils"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)
		userClaims, ok := utils.ParseToken(token)
		if !ok {
			c.Abort()
			c.JSON(http.StatusForbidden, "please login again")
		} else {
			id := strconv.Itoa(int(userClaims.ID))
			user, ok := dao.GetUser("id", id)
			if ok {
				c.Set("user", user)
				token, _ = utils.GenToken(&utils.UserClaims{
					Name:           user.Name,
					Password:       user.Password,
					Email:          user.Email,
					StandardClaims: jwt.StandardClaims{},
				})
				c.Header("Authorization", token)
				c.Next()
			} else {
				c.Abort()
				c.JSON(http.StatusForbidden, "please login again")
			}
		}
	}
}

func ChatAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)
		userClaims, ok := utils.ParseToken(token)
		if !ok {
			c.Abort()
			c.JSON(http.StatusForbidden, "please login again")
		} else {
			id := strconv.Itoa(int(userClaims.ID))
			user, ok := dao.GetUser("id", id)
			if ok {
				c.Set("user", user)
				c.Next()
			} else {
				c.Abort()
				c.JSON(http.StatusForbidden, "please login again")
			}
		}
	}
}
