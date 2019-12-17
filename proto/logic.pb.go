// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: logic.proto

package proto

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type EmptyReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyReply) Reset()         { *m = EmptyReply{} }
func (m *EmptyReply) String() string { return proto.CompactTextString(m) }
func (*EmptyReply) ProtoMessage()    {}
func (*EmptyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_60207fea82c31ca8, []int{0}
}
func (m *EmptyReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EmptyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EmptyReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EmptyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyReply.Merge(m, src)
}
func (m *EmptyReply) XXX_Size() int {
	return m.Size()
}
func (m *EmptyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyReply.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyReply proto.InternalMessageInfo

type ConnectReq struct {
	Sid                  int64    `protobuf:"varint,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Uid                  int64    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectReq) Reset()         { *m = ConnectReq{} }
func (m *ConnectReq) String() string { return proto.CompactTextString(m) }
func (*ConnectReq) ProtoMessage()    {}
func (*ConnectReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_60207fea82c31ca8, []int{1}
}
func (m *ConnectReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConnectReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConnectReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConnectReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectReq.Merge(m, src)
}
func (m *ConnectReq) XXX_Size() int {
	return m.Size()
}
func (m *ConnectReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectReq.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectReq proto.InternalMessageInfo

func (m *ConnectReq) GetSid() int64 {
	if m != nil {
		return m.Sid
	}
	return 0
}

func (m *ConnectReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type ConnectReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectReply) Reset()         { *m = ConnectReply{} }
func (m *ConnectReply) String() string { return proto.CompactTextString(m) }
func (*ConnectReply) ProtoMessage()    {}
func (*ConnectReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_60207fea82c31ca8, []int{2}
}
func (m *ConnectReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConnectReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConnectReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConnectReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectReply.Merge(m, src)
}
func (m *ConnectReply) XXX_Size() int {
	return m.Size()
}
func (m *ConnectReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectReply.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectReply proto.InternalMessageInfo

type DisConnectReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisConnectReq) Reset()         { *m = DisConnectReq{} }
func (m *DisConnectReq) String() string { return proto.CompactTextString(m) }
func (*DisConnectReq) ProtoMessage()    {}
func (*DisConnectReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_60207fea82c31ca8, []int{3}
}
func (m *DisConnectReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DisConnectReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DisConnectReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DisConnectReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisConnectReq.Merge(m, src)
}
func (m *DisConnectReq) XXX_Size() int {
	return m.Size()
}
func (m *DisConnectReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DisConnectReq.DiscardUnknown(m)
}

var xxx_messageInfo_DisConnectReq proto.InternalMessageInfo

type DisConnectReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisConnectReply) Reset()         { *m = DisConnectReply{} }
func (m *DisConnectReply) String() string { return proto.CompactTextString(m) }
func (*DisConnectReply) ProtoMessage()    {}
func (*DisConnectReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_60207fea82c31ca8, []int{4}
}
func (m *DisConnectReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DisConnectReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DisConnectReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DisConnectReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisConnectReply.Merge(m, src)
}
func (m *DisConnectReply) XXX_Size() int {
	return m.Size()
}
func (m *DisConnectReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DisConnectReply.DiscardUnknown(m)
}

var xxx_messageInfo_DisConnectReply proto.InternalMessageInfo

type MessageReq struct {
	SenderId             int64    `protobuf:"varint,1,opt,name=senderId,proto3" json:"senderId,omitempty"`
	RecipientId          int64    `protobuf:"varint,2,opt,name=recipientId,proto3" json:"recipientId,omitempty"`
	Type                 int64    `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageReq) Reset()         { *m = MessageReq{} }
func (m *MessageReq) String() string { return proto.CompactTextString(m) }
func (*MessageReq) ProtoMessage()    {}
func (*MessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_60207fea82c31ca8, []int{5}
}
func (m *MessageReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MessageReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageReq.Merge(m, src)
}
func (m *MessageReq) XXX_Size() int {
	return m.Size()
}
func (m *MessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_MessageReq proto.InternalMessageInfo

func (m *MessageReq) GetSenderId() int64 {
	if m != nil {
		return m.SenderId
	}
	return 0
}

func (m *MessageReq) GetRecipientId() int64 {
	if m != nil {
		return m.RecipientId
	}
	return 0
}

func (m *MessageReq) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MessageReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*EmptyReply)(nil), "proto.EmptyReply")
	proto.RegisterType((*ConnectReq)(nil), "proto.connectReq")
	proto.RegisterType((*ConnectReply)(nil), "proto.connectReply")
	proto.RegisterType((*DisConnectReq)(nil), "proto.disConnectReq")
	proto.RegisterType((*DisConnectReply)(nil), "proto.disConnectReply")
	proto.RegisterType((*MessageReq)(nil), "proto.MessageReq")
}

func init() { proto.RegisterFile("logic.proto", fileDescriptor_60207fea82c31ca8) }

var fileDescriptor_60207fea82c31ca8 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4e, 0x84, 0x30,
	0x14, 0xc6, 0x53, 0x99, 0xf1, 0xcf, 0x03, 0x1d, 0xa7, 0x1a, 0x43, 0xba, 0x20, 0x84, 0xd5, 0x6c,
	0x64, 0xd4, 0xd9, 0xb9, 0xd4, 0xb8, 0x98, 0x85, 0x1b, 0x6e, 0x20, 0xa5, 0x62, 0x13, 0xa6, 0xad,
	0x50, 0x12, 0x39, 0x93, 0x17, 0x71, 0xe9, 0x11, 0x0c, 0x27, 0x31, 0x2d, 0x03, 0x8c, 0x66, 0x56,
	0x7d, 0xdf, 0xaf, 0xef, 0xbd, 0xaf, 0xfd, 0xc0, 0x2d, 0x64, 0xce, 0x69, 0xac, 0x4a, 0xa9, 0x25,
	0x9e, 0xda, 0x83, 0x5c, 0xe7, 0x5c, 0xbf, 0xd5, 0x69, 0x4c, 0xe5, 0x66, 0x99, 0xcb, 0x5c, 0x2e,
	0x2d, 0x4e, 0xeb, 0x57, 0xab, 0xac, 0xb0, 0x55, 0x37, 0x45, 0x5c, 0x7b, 0x74, 0x22, 0xf2, 0x00,
	0x9e, 0x36, 0x4a, 0x37, 0x09, 0x53, 0x45, 0x13, 0xdd, 0x00, 0x50, 0x29, 0x04, 0xa3, 0x3a, 0x61,
	0xef, 0xf8, 0x1c, 0x9c, 0x8a, 0x67, 0x3e, 0x0a, 0xd1, 0xc2, 0x49, 0x4c, 0x69, 0x48, 0xcd, 0x33,
	0xff, 0xa0, 0x23, 0x35, 0xcf, 0xa2, 0x33, 0xf0, 0x86, 0x09, 0xb3, 0x61, 0x06, 0xa7, 0x19, 0xaf,
	0x1e, 0x87, 0x25, 0xd1, 0x1c, 0x66, 0xbb, 0xc0, 0xf4, 0x7c, 0x00, 0x3c, 0xb3, 0xaa, 0x7a, 0xc9,
	0x99, 0x71, 0x21, 0x70, 0x5c, 0x31, 0x91, 0xb1, 0x72, 0xdd, 0x5b, 0x0d, 0x1a, 0x87, 0xe0, 0x96,
	0x8c, 0x72, 0xc5, 0x99, 0xd0, 0xeb, 0xde, 0x77, 0x17, 0x61, 0x0c, 0x13, 0xdd, 0x28, 0xe6, 0x3b,
	0xf6, 0xca, 0xd6, 0xd8, 0x87, 0x23, 0x2a, 0x85, 0x66, 0x42, 0xfb, 0x93, 0x10, 0x2d, 0x4e, 0x92,
	0x5e, 0xde, 0x7d, 0x22, 0x98, 0xda, 0x00, 0xf1, 0x0a, 0x5c, 0xe3, 0xb2, 0x7d, 0x07, 0x9e, 0x77,
	0x71, 0xc4, 0xe3, 0xbb, 0x48, 0x8f, 0xc6, 0x78, 0xf0, 0xad, 0x5d, 0x6c, 0x3e, 0x32, 0x0c, 0x8c,
	0x71, 0x91, 0x8b, 0xff, 0xc8, 0x8c, 0xdc, 0x03, 0x8c, 0xdf, 0xc7, 0x97, 0xdb, 0x96, 0x3f, 0x11,
	0x91, 0xab, 0x3d, 0x54, 0x15, 0xcd, 0x83, 0xf7, 0xd5, 0x06, 0xe8, 0xbb, 0x0d, 0xd0, 0x4f, 0x1b,
	0xa0, 0xf4, 0xd0, 0x36, 0xad, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xa3, 0xfc, 0x9b, 0x02,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogicClient is the client API for Logic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogicClient interface {
	SendMessage(ctx context.Context, in *MessageReq, opts ...grpc.CallOption) (*EmptyReply, error)
	Connect(ctx context.Context, in *ConnectReq, opts ...grpc.CallOption) (*ConnectReply, error)
	DisConnect(ctx context.Context, in *DisConnectReq, opts ...grpc.CallOption) (*DisConnectReply, error)
}

type logicClient struct {
	cc *grpc.ClientConn
}

func NewLogicClient(cc *grpc.ClientConn) LogicClient {
	return &logicClient{cc}
}

func (c *logicClient) SendMessage(ctx context.Context, in *MessageReq, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, "/proto.logic/sendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) Connect(ctx context.Context, in *ConnectReq, opts ...grpc.CallOption) (*ConnectReply, error) {
	out := new(ConnectReply)
	err := c.cc.Invoke(ctx, "/proto.logic/connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) DisConnect(ctx context.Context, in *DisConnectReq, opts ...grpc.CallOption) (*DisConnectReply, error) {
	out := new(DisConnectReply)
	err := c.cc.Invoke(ctx, "/proto.logic/disConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogicServer is the server API for Logic service.
type LogicServer interface {
	SendMessage(context.Context, *MessageReq) (*EmptyReply, error)
	Connect(context.Context, *ConnectReq) (*ConnectReply, error)
	DisConnect(context.Context, *DisConnectReq) (*DisConnectReply, error)
}

func RegisterLogicServer(s *grpc.Server, srv LogicServer) {
	s.RegisterService(&_Logic_serviceDesc, srv)
}

func _Logic_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.logic/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).SendMessage(ctx, req.(*MessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.logic/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).Connect(ctx, req.(*ConnectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_DisConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisConnectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).DisConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.logic/DisConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).DisConnect(ctx, req.(*DisConnectReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Logic_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.logic",
	HandlerType: (*LogicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sendMessage",
			Handler:    _Logic_SendMessage_Handler,
		},
		{
			MethodName: "connect",
			Handler:    _Logic_Connect_Handler,
		},
		{
			MethodName: "disConnect",
			Handler:    _Logic_DisConnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logic.proto",
}

func (m *EmptyReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EmptyReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ConnectReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnectReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Sid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLogic(dAtA, i, uint64(m.Sid))
	}
	if m.Uid != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintLogic(dAtA, i, uint64(m.Uid))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ConnectReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnectReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *DisConnectReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DisConnectReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *DisConnectReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DisConnectReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *MessageReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MessageReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SenderId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLogic(dAtA, i, uint64(m.SenderId))
	}
	if m.RecipientId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintLogic(dAtA, i, uint64(m.RecipientId))
	}
	if m.Type != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintLogic(dAtA, i, uint64(m.Type))
	}
	if len(m.Content) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintLogic(dAtA, i, uint64(len(m.Content)))
		i += copy(dAtA[i:], m.Content)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintLogic(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EmptyReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ConnectReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sid != 0 {
		n += 1 + sovLogic(uint64(m.Sid))
	}
	if m.Uid != 0 {
		n += 1 + sovLogic(uint64(m.Uid))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ConnectReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DisConnectReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DisConnectReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MessageReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SenderId != 0 {
		n += 1 + sovLogic(uint64(m.SenderId))
	}
	if m.RecipientId != 0 {
		n += 1 + sovLogic(uint64(m.RecipientId))
	}
	if m.Type != 0 {
		n += 1 + sovLogic(uint64(m.Type))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovLogic(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLogic(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLogic(x uint64) (n int) {
	return sovLogic(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EmptyReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogic
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EmptyReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmptyReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipLogic(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogic
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogic
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConnectReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogic
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: connectReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: connectReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sid", wireType)
			}
			m.Sid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sid |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLogic(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogic
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogic
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConnectReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogic
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: connectReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: connectReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipLogic(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogic
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogic
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DisConnectReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogic
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: disConnectReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: disConnectReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipLogic(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogic
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogic
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DisConnectReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogic
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: disConnectReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: disConnectReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipLogic(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogic
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogic
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MessageReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogic
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MessageReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MessageReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderId", wireType)
			}
			m.SenderId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SenderId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipientId", wireType)
			}
			m.RecipientId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RecipientId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogic(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogic
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogic
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLogic(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLogic
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLogic
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLogic
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLogic
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthLogic
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLogic
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipLogic(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthLogic
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthLogic = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLogic   = fmt.Errorf("proto: integer overflow")
)
