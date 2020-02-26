// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-25
// Time: 21:56

package client

import (
	"fmt"
	"log"

	"context"
	"errors"

	"google.golang.org/grpc"

	"ginchat/common"
	pb "ginchat/proto/chat"
)

func SendMessage(ip string, data []byte) error{
	//TODO 异常情况处理 redis挂了的问题 向所有节点发送，获取不到当前用户节点的连接 向所有节点发送
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, common.RpcPort), grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := pb.NewChatterClient(conn)
	r, err := c.SendMessage(context.Background(), &pb.ChatRequest{Mes: data})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	if r.Code != 200|| r.Mes != "" {
		return errors.New(r.Mes)
	}
	return nil
}
