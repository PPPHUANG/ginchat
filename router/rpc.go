// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-25
// Time: 20:19

package router

import (
	"fmt"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"ginchat/common"
	pb "ginchat/proto/chat"
	"ginchat/service"
)


func ListenRpc() {
	//启动GRPC监听
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", common.RpcPort))
	if err != nil {
		fmt.Printf("listen error:%v\n", err)
		os.Exit(1)
	}
	sp := keepalive.ServerParameters{
		Time: time.Minute * 1,
	}
	server := grpc.NewServer(grpc.KeepaliveParams(sp))
	pb.RegisterChatterServer(server,&service.ChatService{})
	if err := server.Serve(listener);err != nil {
		fmt.Printf("faild to serve: %v",err)
		os.Exit(1)
	}
}