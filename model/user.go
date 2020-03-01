// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-10
// Time: 23:51

package model

import "time"

const (
	SEX_WOMEN  = "W"
	SEX_MEN    = "M"
	SEX_UNKNOW = "U"
)

type User struct {
	//用户ID
	Id       int64  `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Mobile   string `xorm:"varchar(20)" form:"mobile" json:"mobile"`
	Passwd   string `xorm:"varchar(40)" form:"passwd" json:"-"` // 什么角色
	Avatar   string `xorm:"varchar(150)" form:"avatar" json:"avatar"`
	Sex      string `xorm:"varchar(2)" form:"sex" json:"sex"`            // 什么角色
	Nickname string `xorm:"varchar(20)" form:"nickname" json:"nickname"` // 什么角色
	//加盐随机字符串6
	Salt   string `xorm:"varchar(10)" form:"salt" json:"-"`    // 什么角色
	Online int    `xorm:"int(10)" form:"online" json:"online"` //是否在线
	//前端鉴权因子,
	Token    string    `xorm:"varchar(40)" form:"token" json:"token"`    // 什么角色
	Memo     string    `xorm:"varchar(140)" form:"memo" json:"memo"`     // 什么角色
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"` // 什么角色
}

//好友聊天信息表
type MessageUserToUser struct {
	Id           int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	FromUserId   int64     `xorm:"bigint(20)" form:"fromuserid" json:"fromuserid"`
	FromUserName string    `xorm:"varchar(30)" form:"fromusername" json:"fromusername"`
	ToUserId     int64     `xorm:"bigint(20)" form:"touserid" json:"touserid"`
	Content      string    `xorm:"text" form:"content" json:"content"`
	State        uint8     `xorm:"tinyint" form:"state" json:"state"`
	CreateTime   time.Time `xorm:"datetime" form:"creattime" json:"creattime"`
	TypeId       uint8     `xorm:"tinyint" form:"typeid" json:"typeid"`
}

//消息类型表
type MessageType struct {
	Id   int64  `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Name string `xorm:"varchar(30)" form:"name" json:"name"`
}
