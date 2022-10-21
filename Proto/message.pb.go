// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package main

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request_Category int32

const (
	Request_Logistico  Request_Category = 0
	Request_Financiero Request_Category = 1
	Request_Militar    Request_Category = 2
)

var Request_Category_name = map[int32]string{
	0: "Logistico",
	1: "Financiero",
	2: "Militar",
}

var Request_Category_value = map[string]int32{
	"Logistico":  0,
	"Financiero": 1,
	"Militar":    2,
}

func (x Request_Category) String() string {
	return proto.EnumName(Request_Category_name, int32(x))
}

func (Request_Category) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2, 0}
}

type Message struct {
	Category             string   `protobuf:"bytes,1,opt,name=Category,proto3" json:"Category,omitempty"`
	Id                   int32    `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Message) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Message) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Reply struct {
	Reply                bool     `protobuf:"varint,1,opt,name=Reply,proto3" json:"Reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetReply() bool {
	if m != nil {
		return m.Reply
	}
	return false
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type RequestNN struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestNN) Reset()         { *m = RequestNN{} }
func (m *RequestNN) String() string { return proto.CompactTextString(m) }
func (*RequestNN) ProtoMessage()    {}
func (*RequestNN) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *RequestNN) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestNN.Unmarshal(m, b)
}
func (m *RequestNN) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestNN.Marshal(b, m, deterministic)
}
func (m *RequestNN) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestNN.Merge(m, src)
}
func (m *RequestNN) XXX_Size() int {
	return xxx_messageInfo_RequestNN.Size(m)
}
func (m *RequestNN) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestNN.DiscardUnknown(m)
}

var xxx_messageInfo_RequestNN proto.InternalMessageInfo

func (m *RequestNN) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterEnum("main.Request_Category", Request_Category_name, Request_Category_value)
	proto.RegisterType((*Message)(nil), "main.Message")
	proto.RegisterType((*Reply)(nil), "main.Reply")
	proto.RegisterType((*Request)(nil), "main.Request")
	proto.RegisterType((*RequestNN)(nil), "main.RequestNN")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcf, 0x4e, 0x84, 0x30,
	0x10, 0xc6, 0x2d, 0xee, 0x2e, 0x30, 0x04, 0x24, 0x13, 0x0f, 0x04, 0x63, 0xb2, 0x69, 0x3c, 0x70,
	0x91, 0x83, 0x26, 0xde, 0x8d, 0x1b, 0x13, 0x12, 0x97, 0x43, 0x8d, 0x0f, 0x50, 0xa1, 0x21, 0x4d,
	0x58, 0x8a, 0xa5, 0x6a, 0xf6, 0x29, 0x7c, 0x65, 0xb3, 0x50, 0xff, 0xb0, 0xb7, 0xf9, 0x3a, 0xbf,
	0xce, 0x7c, 0xdf, 0x40, 0xb8, 0x13, 0xc3, 0xc0, 0x1b, 0x91, 0xf7, 0x5a, 0x19, 0x85, 0x8b, 0x1d,
	0x97, 0x1d, 0x2d, 0xc0, 0xdd, 0x4e, 0xcf, 0x98, 0x82, 0xf7, 0xc0, 0x8d, 0x68, 0x94, 0xde, 0x27,
	0x64, 0x4d, 0x32, 0x9f, 0xfd, 0x6a, 0x8c, 0xc0, 0x29, 0xea, 0xc4, 0x59, 0x93, 0x6c, 0xc9, 0x9c,
	0xa2, 0x46, 0x84, 0xc5, 0x86, 0x1b, 0x9e, 0x9c, 0x8e, 0xdc, 0x58, 0xd3, 0x4b, 0x58, 0x32, 0xd1,
	0xb7, 0x7b, 0x3c, 0xb7, 0xc5, 0x38, 0xc5, 0x63, 0x93, 0xa0, 0xf7, 0xe0, 0x32, 0xf1, 0xf6, 0x2e,
	0x06, 0x43, 0xef, 0xfe, 0x36, 0x61, 0x08, 0xfe, 0x93, 0x6a, 0xe4, 0x60, 0x64, 0xa5, 0xe2, 0x13,
	0x8c, 0x00, 0x1e, 0x65, 0xc7, 0xbb, 0x4a, 0x0a, 0xad, 0x62, 0x82, 0x01, 0xb8, 0x5b, 0xd9, 0x4a,
	0xc3, 0x75, 0xec, 0xd0, 0x0b, 0xf0, 0xed, 0x88, 0xb2, 0xb4, 0x96, 0xc8, 0x8f, 0xa5, 0x9b, 0x2f,
	0x02, 0x91, 0x8d, 0xf2, 0x2c, 0xf4, 0x87, 0xac, 0x04, 0x5e, 0xc1, 0xea, 0xa5, 0x6f, 0x15, 0xaf,
	0x31, 0xcc, 0x0f, 0x69, 0x73, 0xdb, 0x4f, 0x83, 0x49, 0x4e, 0x76, 0x33, 0xf0, 0x36, 0xea, 0xb3,
	0xfb, 0xcf, 0xd9, 0x2d, 0xe9, 0xfc, 0x1b, 0x5e, 0x43, 0x70, 0x48, 0x6a, 0xbb, 0x78, 0x36, 0x83,
	0xcb, 0xf2, 0x08, 0x7f, 0x5d, 0x8d, 0x87, 0xbe, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x53, 0x07,
	0x31, 0xa5, 0x79, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageServiceClient interface {
	Upload(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Reply, error)
	Download(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Message, error)
	DataRequest(ctx context.Context, in *RequestNN, opts ...grpc.CallOption) (*Message, error)
}

type messageServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) Upload(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/main.MessageService/Upload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) Download(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/main.MessageService/Download", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) DataRequest(ctx context.Context, in *RequestNN, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/main.MessageService/DataRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
type MessageServiceServer interface {
	Upload(context.Context, *Message) (*Reply, error)
	Download(context.Context, *Request) (*Message, error)
	DataRequest(context.Context, *RequestNN) (*Message, error)
}

// UnimplementedMessageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (*UnimplementedMessageServiceServer) Upload(ctx context.Context, req *Message) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (*UnimplementedMessageServiceServer) Download(ctx context.Context, req *Request) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (*UnimplementedMessageServiceServer) DataRequest(ctx context.Context, req *RequestNN) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataRequest not implemented")
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.MessageService/Upload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).Upload(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_Download_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).Download(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.MessageService/Download",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).Download(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_DataRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestNN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).DataRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.MessageService/DataRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).DataRequest(ctx, req.(*RequestNN))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _MessageService_Upload_Handler,
		},
		{
			MethodName: "Download",
			Handler:    _MessageService_Download_Handler,
		},
		{
			MethodName: "DataRequest",
			Handler:    _MessageService_DataRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
