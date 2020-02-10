// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-11
// Time: 00:39

package contact

import (
	"github.com/gin-gonic/gin"

	"ginchat/args"
	"ginchat/httphandler/chat"
	"ginchat/httphandlerpack/contact"
	"ginchat/model"
	"ginchat/util"
)

var contactService contact.ContactService

func LoadFriend(c *gin.Context) {
	var arg args.ContactArg
	c.Bind(&arg)
	users := contactService.SearchFriend(arg.Userid)
	util.RespOkList(c, users, len(users))
}

func LoadCommunity(c *gin.Context) {
	var arg args.ContactArg
	c.Bind(&arg)
	comunitys := contactService.SearchComunity(arg.Userid)
	util.RespOkList(c, comunitys, len(comunitys))
}
func JoinCommunity(c *gin.Context) {
	var arg args.ContactArg
	c.Bind(&arg)
	err := contactService.JoinCommunity(arg.Userid, arg.Dstid)
	//todo 刷新用户的群组信息
	chat.AddGroupId(arg.Userid, arg.Dstid)
	if err != nil {
		util.RespFail(c, err.Error())
	} else {
		util.RespOk(c, nil, "")
	}
}
func CreateCommunity(c *gin.Context) {
	var arg model.Community
	c.Bind(&arg)
	com, err := contactService.CreateCommunity(arg)
	if err != nil {
		util.RespFail(c, err.Error())
	} else {
		util.RespOk(c, com, "")
	}
}
func AddFriend(c *gin.Context) {
	var arg args.ContactArg
	c.Bind(&arg)
	err := contactService.AddFriend(arg.Userid, arg.Dstid)
	if err != nil {
		util.RespFail(c, err.Error())
	} else {
		util.RespOk(c, nil, "好友添加成功")
	}
}
