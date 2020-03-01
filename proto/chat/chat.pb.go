// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package chat

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ChatRequest struct {
	MessageId            int64    `protobuf:"varint,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
	UserIds              []int64  `protobuf:"varint,2,rep,packed,name=userIds,proto3" json:"userIds,omitempty"`
	Mes                  []byte   `protobuf:"bytes,3,opt,name=mes,proto3" json:"mes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatRequest) Reset()         { *m = ChatRequest{} }
func (m *ChatRequest) String() string { return proto.CompactTextString(m) }
func (*ChatRequest) ProtoMessage()    {}
func (*ChatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

func (m *ChatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatRequest.Unmarshal(m, b)
}
func (m *ChatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatRequest.Marshal(b, m, deterministic)
}
func (m *ChatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatRequest.Merge(m, src)
}
func (m *ChatRequest) XXX_Size() int {
	return xxx_messageInfo_ChatRequest.Size(m)
}
func (m *ChatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChatRequest proto.InternalMessageInfo

func (m *ChatRequest) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *ChatRequest) GetUserIds() []int64 {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *ChatRequest) GetMes() []byte {
	if m != nil {
		return m.Mes
	}
	return nil
}

type ChatReply struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Mes                  string   `protobuf:"bytes,2,opt,name=mes,proto3" json:"mes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatReply) Reset()         { *m = ChatReply{} }
func (m *ChatReply) String() string { return proto.CompactTextString(m) }
func (*ChatReply) ProtoMessage()    {}
func (*ChatReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}

func (m *ChatReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatReply.Unmarshal(m, b)
}
func (m *ChatReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatReply.Marshal(b, m, deterministic)
}
func (m *ChatReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatReply.Merge(m, src)
}
func (m *ChatReply) XXX_Size() int {
	return xxx_messageInfo_ChatReply.Size(m)
}
func (m *ChatReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatReply.DiscardUnknown(m)
}

var xxx_messageInfo_ChatReply proto.InternalMessageInfo

func (m *ChatReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ChatReply) GetMes() string {
	if m != nil {
		return m.Mes
	}
	return ""
}

func init() {
	proto.RegisterType((*ChatRequest)(nil), "chat.ChatRequest")
	proto.RegisterType((*ChatReply)(nil), "chat.ChatReply")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xc2, 0xb9, 0xb8, 0x9d, 0x33,
	0x12, 0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x64, 0xb8, 0x38, 0x73, 0x53, 0x8b,
	0x8b, 0x13, 0xd3, 0x53, 0x3d, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x10, 0x02, 0x42,
	0x12, 0x5c, 0xec, 0xa5, 0xc5, 0xa9, 0x45, 0x9e, 0x29, 0xc5, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0xcc,
	0x41, 0x30, 0xae, 0x90, 0x00, 0x17, 0x73, 0x6e, 0x6a, 0xb1, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x4f,
	0x10, 0x88, 0xa9, 0x64, 0xc8, 0xc5, 0x09, 0x31, 0xb8, 0x20, 0xa7, 0x52, 0x48, 0x88, 0x8b, 0x25,
	0x39, 0x3f, 0x25, 0x15, 0x6c, 0x22, 0x6b, 0x10, 0x98, 0x0d, 0xd3, 0xc2, 0xa4, 0xc0, 0xa8, 0xc1,
	0x09, 0xd6, 0x62, 0x64, 0xc7, 0xc5, 0x0e, 0xd2, 0x52, 0x92, 0x5a, 0x24, 0x64, 0xcc, 0xc5, 0x5d,
	0x9c, 0x9a, 0x97, 0xe2, 0x0b, 0xb1, 0x5a, 0x48, 0x50, 0x0f, 0xec, 0x70, 0x24, 0x97, 0x4a, 0xf1,
	0x23, 0x0b, 0x15, 0xe4, 0x54, 0x2a, 0x31, 0x24, 0xb1, 0x81, 0x3d, 0x66, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x8a, 0x35, 0xa5, 0x79, 0xe6, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatterClient is the client API for Chatter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatterClient interface {
	// Sends a chatMessage
	SendMessage(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatReply, error)
}

type chatterClient struct {
	cc *grpc.ClientConn
}

func NewChatterClient(cc *grpc.ClientConn) ChatterClient {
	return &chatterClient{cc}
}

func (c *chatterClient) SendMessage(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatReply, error) {
	out := new(ChatReply)
	err := c.cc.Invoke(ctx, "/chat.Chatter/sendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatterServer is the server API for Chatter service.
type ChatterServer interface {
	// Sends a chatMessage
	SendMessage(context.Context, *ChatRequest) (*ChatReply, error)
}

func RegisterChatterServer(s *grpc.Server, srv ChatterServer) {
	s.RegisterService(&_Chatter_serviceDesc, srv)
}

func _Chatter_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatterServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chatter/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatterServer).SendMessage(ctx, req.(*ChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chatter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.Chatter",
	HandlerType: (*ChatterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sendMessage",
			Handler:    _Chatter_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
