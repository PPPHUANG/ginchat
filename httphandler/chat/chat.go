// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-11
// Time: 01:11

package chat

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"

	"ginchat/client"
	"ginchat/common"
	"ginchat/httphandlerpack/contact"
	"ginchat/httphandlerpack/user"
)

var contactService contact.ContactService
var userService user.UserService

//映射关系表
var clientMap = make(map[int64]*Node, 0)

var rwlocker sync.RWMutex

// ws://127.0.0.1/chat?id=1&token=xxxx
func Chat(c *gin.Context) {
	id := c.Query("id")
	token := c.Query("token")
	userId, _ := strconv.ParseInt(id, 10, 64)
	isValidate := checkToken(userId, token)
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isValidate
		},
	}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	node := &Node{
		Uid:           userId,
		Conn:          conn,
		DataQueue:     make(chan *Data, 50),
		GroupSets:     set.New(set.ThreadSafe),
		HeartBeatChan: make(chan bool, 50),
		WRChan:        make(chan bool, 1),
	}
	//获取用户全部群Id
	//comIds := contactService.SearchCommunityIds(userId)
	//for _, v := range comIds {
	//	node.GroupSets.Add(v)
	//}
	//userId和node形成绑定关系
	rwlocker.Lock()
	clientMap[userId] = node
	rwlocker.Unlock()
	//存储用户id => host
	userService.SaveHost(id)
	//完成发送逻辑
	go node.SendProc()
	//完成接收逻辑
	go node.RecvProc()
	// 心跳检测
	go node.HeartBeat()
	log.Printf("<-%d\n", userId)
}

//添加新的群ID到用户的groupSet中
func AddGroupId(userId, gid int64) {
	//取得node
	rwlocker.Lock()
	node, ok := clientMap[userId]
	if ok {
		node.GroupSets.Add(gid)
	}
	//clientMap[userId] = node
	rwlocker.Unlock()
	//添加gid到set
}

func init() {
	go rpcSendProc()
}

var rpcSendChan = make(chan *RpcSendData, 1024)

//推送消息或者发送到对应的节点上
func BroadMsg(data []byte, fromAnotherNode bool) {
	//解析data为message json解码可以优化为滴滴开源的 Json-iterator
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//写消息到消息表
	if !fromAnotherNode {
		mesId := msg.Save()
		msg.Id = mesId
	}

	Dispatch(&Data{&msg, data})
}

func RpcDispatch(rpcSendData *RpcSendData) {
	msg := &Message{}
	err := json.Unmarshal(rpcSendData.Data, msg)
	if err != nil {
		log.Println(err.Error())
		return
	}
	msg.Id = rpcSendData.MessageId
	//根据cmd对逻辑进行处理
	switch msg.Cmd {
	case CMD_SINGLE_MSG:
		sendMsg(&Data{msg, rpcSendData.Data})
	case CMD_ROOM_MSG:
		//群聊转发逻辑 获取群用户列表
		for _, uid := range rpcSendData.ToUIds {
			msg.Dstid = uid
			sendMsg(&Data{msg, rpcSendData.Data})
		}
	}
}
func localHost(usrId string) (string, bool) {
	ip := userService.GetHost(usrId)
	if common.ServerIp != ip {
		return ip, false
	}
	return ip, true
}

//完成rpc数据的发送协程
func rpcSendProc() {
	for {
		select {
		case data := <-rpcSendChan:
			err := client.SendMessage(data.MessageId, data.ToUIds, data.IP, data.Data)
			if err != nil {
				//TODO 重传或者其他处理
				log.Println(err.Error())
				return
			}
		}
	}
}

//后端调度逻辑处理
func Dispatch(data *Data) {
	//根据cmd对逻辑进行处理
	switch data.DataFormat.Cmd {
	case CMD_SINGLE_MSG:
		if ip, isLocal := localHost(strconv.FormatInt(data.DataFormat.Dstid, 10)); isLocal {
			sendMsg(data)
		} else {
			rpcSendChan <- &RpcSendData{
				ToUIds:    []int64{data.DataFormat.Dstid},
				MessageId: data.DataFormat.Id,
				IP:        ip,
				Data:      data.Data,
			}
		}
	case CMD_ROOM_MSG:
		//群聊转发逻辑 获取群用户列表
		CId := data.DataFormat.Dstid
		UIds := contactService.GetCommunityUsers(CId)
		sendIp := make(map[string][]int64)
		for _, userIdString := range UIds {
			userIdInt, _ := strconv.ParseInt(userIdString, 10, 64)
			if ip, isLocal := localHost(userIdString); isLocal {
				data.DataFormat.Dstid = userIdInt
				sendMsg(data)
			} else { //不能按人去推送给各个机器 要按机器推送
				sendIp[ip] = append(sendIp[ip], userIdInt)
			}
		}
		for ip, uIds := range sendIp {
			rpcSendChan <- &RpcSendData{
				MessageId: data.DataFormat.Id,
				ToUIds:    uIds,
				IP:        ip,
				Data:      data.Data,
			}
		}
	case CMD_HEART:
		//一般啥都不做
	}
}

//发送消息
func sendMsg(data *Data) {
	rwlocker.RLock()
	node, ok := clientMap[data.DataFormat.Dstid]
	rwlocker.RUnlock()
	if ok {
		node.DataQueue <- data
	}
}

//var userService user.UserService

func checkToken(userId int64, token string) bool {
	//从数据库里面查询并比对
	userInfo, ok := userService.Find(userId)
	return ok && userInfo.Token == token
}
