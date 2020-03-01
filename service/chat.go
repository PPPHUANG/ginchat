// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-25
// Time: 21:29

package service

import (
	"golang.org/x/net/context"

	"ginchat/httphandler/chat"
	pb "ginchat/proto/chat"
)

type ChatService struct{}

func (s *ChatService) SendMessage(ctx context.Context, in *pb.ChatRequest) (*pb.ChatReply, error) {
	//接收到信息之后直接给服务进行分析
	rpcData := &chat.RpcSendData{
		MessageId: in.MessageId,
		ToUIds:    in.UserIds,
		Data:      in.Mes,
	}
	chat.RpcDispatch(rpcData)
	return &pb.ChatReply{Code: 200, Mes: ""}, nil
}
