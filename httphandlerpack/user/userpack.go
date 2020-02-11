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

	log "github.com/sirupsen/logrus"
)

type UserService struct {
}

func (s *UserService) Login(mobile, plainpwd string) (model.User, error) {
	tmp := model.User{}
	_, err := db_conn.DbClient.Where("mobile = ?", mobile).Get(&tmp)
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/userpack.go/Login",
		}).Error(err.Error())
		return tmp, err
	}
	if tmp.Id == 0 {
		return tmp, errors.New("not Exist User")
	}
	if !util.ValidatePasswd(plainpwd, tmp.Salt, tmp.Passwd) {
		return tmp, errors.New("wrong Pwd")
	}
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := util.MD5Encode(str)
	tmp.Token = token
	_, err = db_conn.DbClient.ID(tmp.Id).Cols("token").Update(&tmp)
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/userpack.go/Login",
		}).Error(err.Error())
	}
	return tmp, nil
}

func (s *UserService) Register(mobile, plainpwd, nickname, avatar, sex string) (model.User, error) {
	tmp := model.User{}
	_, err := db_conn.DbClient.Where("mobile=? ", mobile).Get(&tmp)
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/userpack.go/Register/Get",
		}).Error(err.Error())
		return tmp, err
	}
	if tmp.Id > 0 {
		return tmp, errors.New("the mobile number has been registered")
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
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/userpack.go/Register/InsertOne",
		}).Error(err.Error())
	}
	return tmp, err
}

func (s *UserService) Find(userId int64) model.User {
	tmp := model.User{}
	_, err := db_conn.DbClient.ID(userId).Get(&tmp)
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/userpack.go/Find/Get",
		}).Error(err.Error())
	}
	return tmp
}
