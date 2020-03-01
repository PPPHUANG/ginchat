// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-10
// Time: 23:51

package model

import "time"

type Community struct {
	Id int64 `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	//名称
	Name string `xorm:"varchar(30)" form:"name" json:"name"`
	//群主ID
	Ownerid int64 `xorm:"bigint(20)" form:"ownerid" json:"ownerid"` // 什么角色
	//群logo
	Icon string `xorm:"varchar(250)" form:"icon" json:"icon"`
	//como
	Cate int `xorm:"int(11)" form:"cate" json:"cate"` // 什么角色
	//描述
	Memo string `xorm:"varchar(120)" form:"memo" json:"memo"` // 什么角色
	//
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"` // 什么角色
}

const (
	COMMUNITY_CATE_COM = 0x01
)

//群消息表
type CommunityMessage struct {
	Id              int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Content         string    `xorm:"text" form:"content" json:"content"`
	FromId          int64     `xorm:"bigint(20)" form:"fromid" json:"fromid"`
	FromUName       string    `xorm:"varchar(30)" form:"fromuname" json:"fromuname"`
	CreateTime      time.Time `xorm:"datetime" form:"creattime" json:"creattime"`
	UserCommunityId int64     `xorm:"bigint(20)" form:"usercommunityid" json:"usercommunityid"`
	TypeId          uint8     `xorm:"tinyint" form:"typeid" json:"typeid"`
}

//群消息关联表
type CommunityMessageToUser struct {
	Id         int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	UserId     int64     `xorm:"bigint(20)" form:"userid" json:"userid"`
	ComMesId   int64     `xorm:"bigint(20)" form:"commesid" json:"commesid"`
	State      uint8     `xorm:"tinyint" form:"state" json:"state"`
	CreateTime time.Time `xorm:"datetime" form:"creattime" json:"creattime"`
}

//群内私聊消息关联表
type CommunityMessageUserToUser struct {
	Id              int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	FromUserId      int64     `xorm:"bigint(20)" form:"fromuserid" json:"fromuserid"`
	FromUserName    string    `xorm:"varchar(30)" form:"fromusername" json:"fromusername"`
	ToUserId        int64     `xorm:"bigint(20)" form:"touserid" json:"touserid"`
	Content         string    `xorm:"text" form:"content" json:"content"`
	State           uint8     `xorm:"tinyint" form:"state" json:"state"`
	CreateTime      time.Time `xorm:"datetime" form:"creattime" json:"creattime"`
	UserCommunityId int64     `xorm:"bigint(20)" form:"usercommunityid" json:"usercommunityid"`
	TypeId          uint8     `xorm:"tinyint" form:"typeid" json:"typeid"`
}
