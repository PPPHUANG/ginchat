// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-10
// Time: 23:51

package user

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"ginchat/db_conn"
	"ginchat/model"
	"ginchat/util"
)

type UserService struct {
}

func (s *UserService) Login(mobile, plainpwd string) (user model.User, err error) {
	//首先通过手机号查询用户
	tmp := model.User{}
	db_conn.DbClient.Where("mobile = ?", mobile).Get(&tmp)
	//如果没有找到
	if tmp.Id == 0 {
		return tmp, errors.New("该用户不存在")
	}
	//查询到了比对密码
	if !util.ValidatePasswd(plainpwd, tmp.Salt, tmp.Passwd) {
		return tmp, errors.New("密码不正确")
	}
	//刷新token,安全
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := util.MD5Encode(str)
	tmp.Token = token
	//返回数据
	db_conn.DbClient.ID(tmp.Id).Cols("token").Update(&tmp)
	return tmp, nil
}

func (s *UserService) Register(mobile, plainpwd, nickname, avatar, sex string) (user model.User, err error) {
	tmp := model.User{}
	_, err = db_conn.DbClient.Where("mobile=? ", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}
	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已经注册")
	}
	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nickname
	tmp.Sex = sex
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Passwd = util.MakePasswd(plainpwd, tmp.Salt)
	tmp.Createat = time.Now()
	tmp.Token = fmt.Sprintf("%08d", rand.Int31())
	_, err = db_conn.DbClient.InsertOne(&tmp)
	return tmp, err
}

//查找某个用户
func (s *UserService) Find(userId int64) (user model.User) {
	//首先通过手机号查询用户
	tmp := model.User{}
	db_conn.DbClient.ID(userId).Get(&tmp)
	return tmp
}
