// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-28
// Time: 13:41

package chat

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"

	"ginchat/common"
)

type Message struct {
	Id      int64  `json:"id,omitempty" form:"id"`           //消息ID
	Userid  int64  `json:"userid,omitempty" form:"userid"`   //谁发的
	Cmd     int    `json:"cmd,omitempty" form:"cmd"`         //群聊还是私聊
	Dstid   int64  `json:"dstid,omitempty" form:"dstid"`     //对端用户ID/群ID
	Media   int    `json:"media,omitempty" form:"media"`     //消息按照什么样式展示
	Content string `json:"content,omitempty" form:"content"` //消息的内容
	Pic     string `json:"pic,omitempty" form:"pic"`         //预览图片
	Url     string `json:"url,omitempty" form:"url"`         //服务的URL
	Memo    string `json:"memo,omitempty" form:"memo"`       //简单描述
	Amount  int    `json:"amount,omitempty" form:"amount"`   //其他和数字相关的
}

/**
消息发送结构体
1、MEDIA_TYPE_TEXT
{id:1,userid:2,dstid:3,cmd:10,media:1,content:"hello"}
2、MEDIA_TYPE_News
{id:1,userid:2,dstid:3,cmd:10,media:2,content:"标题",pic:"http://www.baidu.com/a/log,jpg",url:"http://www.a,com/dsturl","memo":"这是描述"}
3、MEDIA_TYPE_VOICE，amount单位秒
{id:1,userid:2,dstid:3,cmd:10,media:3,url:"http://www.a,com/dsturl.mp3",anount:40}
4、MEDIA_TYPE_IMG
{id:1,userid:2,dstid:3,cmd:10,media:4,url:"http://www.baidu.com/a/log,jpg"}
5、MEDIA_TYPE_REDPACKAGR //红包amount 单位分
{id:1,userid:2,dstid:3,cmd:10,media:5,url:"http://www.baidu.com/a/b/c/redpackageaddress?id=100000","amount":300,"memo":"恭喜发财"}
6、MEDIA_TYPE_EMOJ 6
{id:1,userid:2,dstid:3,cmd:10,media:6,"content":"cry"}
7、MEDIA_TYPE_Link 6
{id:1,userid:2,dstid:3,cmd:10,media:7,"url":"http://www.a,com/dsturl.html"}

7、MEDIA_TYPE_Link 6
{id:1,userid:2,dstid:3,cmd:10,media:7,"url":"http://www.a,com/dsturl.html"}

8、MEDIA_TYPE_VIDEO 8
{id:1,userid:2,dstid:3,cmd:10,media:8,pic:"http://www.baidu.com/a/log,jpg",url:"http://www.a,com/a.mp4"}

9、MEDIA_TYPE_CONTACT 9
{id:1,userid:2,dstid:3,cmd:10,media:9,"content":"10086","pic":"http://www.baidu.com/a/avatar,jpg","memo":"胡大力"}

*/
//本核心在于形成userid和Node的映射关系
type Node struct {
	Uid           string          //用户ID
	Conn          *websocket.Conn //连接
	DataQueue     chan []byte     //发送数据chan
	GroupSets     set.Interface   //所在组集合
	HeartBeatChan chan bool       //心跳chan
	WRChan        chan bool       //此chan用于读协程异常退出之后通知写协程退出
}

func (node *Node) HeartBeat() {
	//for {
	//	select {
	//	case <-time.After(time.Second * time.Duration(30)):  //select里使用time.After会住占用很高的内存
	//		node.Release()
	//		return
	//	case <-node.HeartBeatChan:
	//	}
	//}
	BeatDuration := time.Second * time.Duration(common.HeartBeat)
	BeatDelay := time.NewTimer(BeatDuration)
	defer BeatDelay.Stop()
	for {
		select {
		case <-BeatDelay.C: //此处优化 select里使用time.After会住占用很高的内存
			node.Release()
			return
		case <-node.HeartBeatChan:
		}
	}
}

func (node *Node) Release() {
	_ = node.Conn.Close()
}

//ws发送协程
func (node *Node) SendProc() {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		case <-node.WRChan:
			log.Println("read goroutine error or network connection closed,write goroutine return")
			node.Release()
			userService.RemoveHost(node.Uid) //移除登录信息 不论是客户端异常断开，还是心跳超时断开都会走到这一步，否则有可能会调用两次remove操作
			return
		}
	}
}

//ws接收协程
func (node *Node) RecvProc() {
	defer func() {
		node.WRChan <- true //读协程异常退出之后 通知写协程退出高并发时降低资源消耗
	}()
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}
		//Dispatch(data)
		BroadMsg(data)
		node.HeartBeatChan <- true
		log.Printf("[ws]<=%s\n", data)
	}
}

type SendData struct {
	IP   string
	Data []byte
}
