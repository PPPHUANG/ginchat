// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-10
// Time: 23:47

package user

import (
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"

	"ginchat/httphandlerpack/user"
	"ginchat/model"
	"ginchat/util"
)

var userService user.UserService

func UserLogin(c *gin.Context) {
	mobile := c.PostForm("mobile")
	passwd := c.PostForm("passwd")
	userInfo, err := userService.Login(mobile, passwd)
	if err != nil {
		util.RespFail(c, err.Error())
	} else {
		util.RespOk(c, userInfo, "")
	}
}

func UserRegister(c *gin.Context) {
	mobile := c.PostForm("mobile")
	plainpwd := c.PostForm("passwd")
	nickname := fmt.Sprintf("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_UNKNOW
	userInfo, err := userService.Register(mobile, plainpwd, nickname, avatar, sex)
	if err != nil {
		util.RespFail(c, err.Error())
	} else {
		util.RespOk(c, userInfo, "")
	}
}
