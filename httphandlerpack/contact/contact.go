// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-11
// Time: 01:00

package contact

import (
	"errors"
	"time"

	"ginchat/db_conn"
	"ginchat/model"

	log "github.com/sirupsen/logrus"
)

type ContactService struct {
}

//add Friend
func (service *ContactService) AddFriend(userId, dstId int64) error {
	if userId == dstId {
		return errors.New("can't add yourself as a friend")
	}
	tmp := model.Contact{}
	_, err := db_conn.DbClient.Where("ownerid = ?", userId).
		And("dstid = ?", dstId).
		And("cate = ?", model.CONCAT_CATE_USER).
		Get(&tmp)
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/contact.go/AddFriend/Get",
		}).Error(err.Error())
	}
	if tmp.Id > 0 {
		return errors.New("user has been added")
	}
	session := db_conn.DbClient.NewSession()
	err = session.Begin()
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/contact.go/AddFriend/sessionBegin",
		}).Error(err.Error())
		return err
	}
	_, e2 := session.InsertOne(model.Contact{
		Ownerid:  userId,
		Dstobj:   dstId,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})
	_, e3 := session.InsertOne(model.Contact{
		Ownerid:  dstId,
		Dstobj:   userId,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})
	if e2 == nil && e3 == nil {
		err = session.Commit()
		if err != nil {
			log.WithFields(log.Fields{
				"filename": "/contact.go/AddFriend/sessionCommit",
			}).Error(err.Error())
		}
		return err
	} else {
		err = session.Rollback()
		if err != nil {
			log.WithFields(log.Fields{
				"filename": "/contact.go/AddFriend/sessionRollback",
			}).Error(err.Error())
			return err
		}
		if e2 != nil {
			return e2
		} else {
			return e3
		}
	}
}

func (service *ContactService) SearchCommunity(userId int64) []model.Community {
	conconts := make([]model.Contact, 0)
	comIds := make([]int64, 0)

	err := db_conn.DbClient.Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_COMUNITY).Find(&conconts)
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/contact.go/SearchCommunity/Find",
		}).Error(err.Error())
		return []model.Community{}
	}
	for _, v := range conconts {
		comIds = append(comIds, v.Dstobj)
	}
	coms := make([]model.Community, 0)
	if len(comIds) == 0 {
		return coms
	}
	err = db_conn.DbClient.In("id", comIds).Find(&coms)
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/contact.go/SearchComunity/Find",
		}).Error(err.Error())
		return []model.Community{}
	}
	return coms
}

func (service *ContactService) SearchCommunityIds(userId int64) []int64 {
	//获取用户全部群ID
	conconts := make([]model.Contact, 0)
	comIds := make([]int64, 0)

	err := db_conn.DbClient.Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_COMUNITY).Find(&conconts)
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/contact.go/SearchCommunity/Find",
		}).Error(err.Error())
		return comIds
	}
	for _, v := range conconts {
		comIds = append(comIds, v.Dstobj)
	}
	return comIds
}

//add Group
func (service *ContactService) JoinCommunity(userId, comId int64) error {
	cot := model.Contact{
		Ownerid: userId,
		Dstobj:  comId,
		Cate:    model.CONCAT_CATE_COMUNITY,
	}
	_, err := db_conn.DbClient.Get(&cot)
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/contact.go/JoinCommunity/Get",
		}).Error(err.Error())
		return err
	}
	if cot.Id == 0 {
		cot.Createat = time.Now()
		_, err := db_conn.DbClient.InsertOne(cot)
		return err
	} else {
		return nil
	}
}

//create Group
func (service *ContactService) CreateCommunity(comm model.Community) (model.Community, error) {
	if len(comm.Name) == 0 {
		err := errors.New("缺少群名称")
		return model.Community{}, err
	}
	if comm.Ownerid == 0 {
		err := errors.New("请先登录")
		return model.Community{}, err
	}
	com := model.Community{
		Ownerid: comm.Ownerid,
	}
	num, err := db_conn.DbClient.Count(&com)

	if num > 5 {
		err = errors.New("一个用户最多只能创见5个群")
		return com, err
	} else {
		comm.Createat = time.Now()
		session := db_conn.DbClient.NewSession()
		err = session.Begin()
		if err != nil {
			log.WithFields(log.Fields{
				"filename": "/contact.go/CreateCommunity/sessionBegin",
			}).Error(err.Error())
			return model.Community{}, err
		}
		_, err = session.InsertOne(&comm)
		if err != nil {
			err := session.Rollback()
			if err != nil {
				log.WithFields(log.Fields{
					"filename": "/contact.go/CreateCommunity/Rollback",
				}).Error(err.Error())
			}
			return com, err
		}
		_, err = session.InsertOne(
			model.Contact{
				Ownerid:  comm.Ownerid,
				Dstobj:   comm.Id,
				Cate:     model.CONCAT_CATE_COMUNITY,
				Createat: time.Now(),
			})
		if err != nil {
			err = session.Rollback()
			if err != nil {
				log.WithFields(log.Fields{
					"filename": "/contact.go/CreateCommunity/Rollback",
				}).Error(err.Error())
			}
		} else {
			err = session.Commit()
			if err != nil {
				log.WithFields(log.Fields{
					"filename": "/contact.go/CreateCommunity/Commit",
				}).Error(err.Error())
			}
		}
		//将用户添加到redis中群集合里
		_ = AddUserToCommunityRedis(comm.Id, comm.Ownerid)
		return com, err
	}
}

//find Friend
func (service *ContactService) SearchFriend(userId int64) []model.User {
	conconts := make([]model.Contact, 0)
	objIds := make([]int64, 0)
	err := db_conn.DbClient.Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_USER).Find(&conconts)
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/contact.go/SearchFriend/Find",
		}).Error(err.Error())
		return []model.User{}
	}
	for _, v := range conconts {
		objIds = append(objIds, v.Dstobj)
	}
	coms := make([]model.User, 0)
	if len(objIds) == 0 {
		return coms
	}
	err = db_conn.DbClient.In("id", objIds).Find(&coms)
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/contact.go/SearchFriend/Find",
		}).Error(err.Error())
	}
	return coms
}

func (service *ContactService) GetCommunityUsers(communityId int64) []string {
	//todo 从redis中获取群组set的的值返回
	setId := "Community" + string(communityId)
	userIds, err := db_conn.RedisClient.SMembers(setId).Result()
	if err != nil {
		log.WithFields(log.Fields{
			"filename":    "/contact.go/GetCommunityUsers/SMembers",
			"communityId": communityId,
		}).Error(err.Error())
		return []string{}
	}
	return userIds
}

func InitCommunityRedis() error {
	conconts := make([]model.Contact, 0)
	err := db_conn.DbClient.Where("cate = ?", model.CONCAT_CATE_COMUNITY).Find(&conconts)
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/contact.go/InitCommunityRedis/Where",
		}).Error(err.Error())
		return err
	}
	for _, v := range conconts {
		setId := "Community" + string(v.Dstobj)
		_, err = db_conn.RedisClient.SAdd(setId, v.Ownerid).Result()
		if err != nil {
			log.WithFields(log.Fields{
				"filename": "/contact.go/InitCommunityRedis/SAdd",
			}).Error(err.Error())
		}
	}
	return err
}

func AddUserToCommunityRedis(communityId int64, userId int64) error {
	setId := "Community" + string(communityId)
	_, err := db_conn.RedisClient.SAdd(setId, userId).Result()
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/contact.go/AddUserToCommunityRedis/SAdd",
		}).Error(err.Error())
	}
	return err
}
