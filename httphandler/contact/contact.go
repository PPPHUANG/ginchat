// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-11
// Time: 00:39

package contact

import (
	"github.com/gin-gonic/gin"

	"ginchat/args"
	"ginchat/httphandlerpack/contact"
	"ginchat/model"
	"ginchat/util"
)

var contactService contact.ContactService

func LoadFriend(c *gin.Context) {
	var arg args.ContactArg
	err := c.Bind(&arg)
	if err != nil {
		util.RespFail(c, "not Enough Params")
		return
	}
	users := contactService.SearchFriend(arg.Userid)
	util.RespOkList(c, users, len(users))
}

func LoadCommunity(c *gin.Context) {
	var arg args.ContactArg
	err := c.Bind(&arg)
	if err != nil {
		util.RespFail(c, "not Enough Params")
		return
	}
	comunitys := contactService.SearchCommunity(arg.Userid)
	util.RespOkList(c, comunitys, len(comunitys))
}
func JoinCommunity(c *gin.Context) {
	var arg args.ContactArg
	err := c.Bind(&arg)
	if err != nil {
		util.RespFail(c, "not Enough Params")
		return
	}
	err = contactService.JoinCommunity(arg.Userid, arg.Dstid)
	//todo 刷洗redis中群组的成员信息
	//chat.AddGroupId(arg.Userid, arg.Dstid)
	_ = contact.AddUserToCommunityRedis(arg.Dstid, arg.Userid)
	if err != nil {
		util.RespFail(c, err.Error())
	} else {
		util.RespOk(c, nil, "")
	}
}
func CreateCommunity(c *gin.Context) {
	var arg model.Community
	err := c.Bind(&arg)
	if err != nil {
		util.RespFail(c, "not Enough Params")
		return
	}
	com, err := contactService.CreateCommunity(arg)
	if err != nil {
		util.RespFail(c, err.Error())
	} else {
		util.RespOk(c, com, "")
	}
}
func AddFriend(c *gin.Context) {
	var arg args.ContactArg
	err := c.Bind(&arg)
	if err != nil {
		util.RespFail(c, "not Enough Params")
		return
	}
	err = contactService.AddFriend(arg.Userid, arg.Dstid)
	if err != nil {
		util.RespFail(c, err.Error())
	} else {
		util.RespOk(c, nil, "add Friend Successfully")
	}
}
