// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quicksilver/interchainstaking/v1/messages.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgRequestRedemption represents a message type to request a burn of qAssets
// for native assets.
type MsgRequestRedemption struct {
	Value              types.Coin `protobuf:"bytes,1,opt,name=value,proto3" json:"value" yaml:"coin"`
	DestinationAddress string     `protobuf:"bytes,2,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	FromAddress        string     `protobuf:"bytes,3,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
}

func (m *MsgRequestRedemption) Reset()         { *m = MsgRequestRedemption{} }
func (m *MsgRequestRedemption) String() string { return proto.CompactTextString(m) }
func (*MsgRequestRedemption) ProtoMessage()    {}
func (*MsgRequestRedemption) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee484030fa140a82, []int{0}
}
func (m *MsgRequestRedemption) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestRedemption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestRedemption.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestRedemption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestRedemption.Merge(m, src)
}
func (m *MsgRequestRedemption) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestRedemption) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestRedemption.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestRedemption proto.InternalMessageInfo

// MsgSignalIntent represents a message type for signalling voting intent for
// one or more validators.
type MsgSignalIntent struct {
	ChainId     string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty" yaml:"chain_id"`
	Intents     string `protobuf:"bytes,2,opt,name=intents,proto3" json:"intents,omitempty" yaml:"intents"`
	FromAddress string `protobuf:"bytes,3,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
}

func (m *MsgSignalIntent) Reset()         { *m = MsgSignalIntent{} }
func (m *MsgSignalIntent) String() string { return proto.CompactTextString(m) }
func (*MsgSignalIntent) ProtoMessage()    {}
func (*MsgSignalIntent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee484030fa140a82, []int{1}
}
func (m *MsgSignalIntent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSignalIntent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSignalIntent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSignalIntent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSignalIntent.Merge(m, src)
}
func (m *MsgSignalIntent) XXX_Size() int {
	return m.Size()
}
func (m *MsgSignalIntent) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSignalIntent.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSignalIntent proto.InternalMessageInfo

// MsgRequestRedemptionResponse defines the MsgRequestRedemption response type.
type MsgRequestRedemptionResponse struct {
}

func (m *MsgRequestRedemptionResponse) Reset()         { *m = MsgRequestRedemptionResponse{} }
func (m *MsgRequestRedemptionResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRequestRedemptionResponse) ProtoMessage()    {}
func (*MsgRequestRedemptionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee484030fa140a82, []int{2}
}
func (m *MsgRequestRedemptionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestRedemptionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestRedemptionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestRedemptionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestRedemptionResponse.Merge(m, src)
}
func (m *MsgRequestRedemptionResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestRedemptionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestRedemptionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestRedemptionResponse proto.InternalMessageInfo

// MsgSignalIntentResponse defines the MsgSignalIntent response type.
type MsgSignalIntentResponse struct {
}

func (m *MsgSignalIntentResponse) Reset()         { *m = MsgSignalIntentResponse{} }
func (m *MsgSignalIntentResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSignalIntentResponse) ProtoMessage()    {}
func (*MsgSignalIntentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee484030fa140a82, []int{3}
}
func (m *MsgSignalIntentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSignalIntentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSignalIntentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSignalIntentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSignalIntentResponse.Merge(m, src)
}
func (m *MsgSignalIntentResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSignalIntentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSignalIntentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSignalIntentResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRequestRedemption)(nil), "quicksilver.interchainstaking.v1.MsgRequestRedemption")
	proto.RegisterType((*MsgSignalIntent)(nil), "quicksilver.interchainstaking.v1.MsgSignalIntent")
	proto.RegisterType((*MsgRequestRedemptionResponse)(nil), "quicksilver.interchainstaking.v1.MsgRequestRedemptionResponse")
	proto.RegisterType((*MsgSignalIntentResponse)(nil), "quicksilver.interchainstaking.v1.MsgSignalIntentResponse")
}

func init() {
	proto.RegisterFile("quicksilver/interchainstaking/v1/messages.proto", fileDescriptor_ee484030fa140a82)
}

var fileDescriptor_ee484030fa140a82 = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x4f, 0x4f, 0xd4, 0x4e,
	0x1c, 0xc6, 0x77, 0x20, 0x3f, 0xfe, 0x0c, 0xe4, 0x87, 0x16, 0x12, 0xa1, 0x21, 0x2d, 0xe9, 0x89,
	0xa8, 0xb4, 0x2c, 0x28, 0x2a, 0x0a, 0xd1, 0xc5, 0x84, 0x70, 0xe0, 0x52, 0x6e, 0x5c, 0x36, 0xb3,
	0xed, 0xd7, 0x61, 0x42, 0x3b, 0x53, 0x3a, 0xd3, 0x06, 0xae, 0x9e, 0x3c, 0x9a, 0xe8, 0x0b, 0xe0,
	0x45, 0x10, 0xdf, 0x80, 0x1e, 0x38, 0x12, 0xbd, 0x78, 0xda, 0x18, 0xf0, 0xe0, 0xc9, 0x03, 0xaf,
	0xc0, 0xf4, 0xcf, 0x6e, 0x76, 0x81, 0x84, 0x75, 0xe3, 0xad, 0x9d, 0x67, 0x3e, 0xcf, 0x3c, 0xcf,
	0xce, 0x77, 0x8b, 0x9d, 0x83, 0x84, 0x79, 0xfb, 0x92, 0x05, 0x29, 0xc4, 0x0e, 0xe3, 0x0a, 0x62,
	0x6f, 0x8f, 0x30, 0x2e, 0x15, 0xd9, 0x67, 0x9c, 0x3a, 0x69, 0xd5, 0x09, 0x41, 0x4a, 0x42, 0x41,
	0xda, 0x51, 0x2c, 0x94, 0xd0, 0xe6, 0x3a, 0x00, 0xfb, 0x1a, 0x60, 0xa7, 0x55, 0xdd, 0xf0, 0x84,
	0x0c, 0x85, 0x74, 0x1a, 0x44, 0x82, 0x93, 0x56, 0x1b, 0xa0, 0x48, 0xd5, 0xf1, 0x04, 0xe3, 0x85,
	0x83, 0x3e, 0x53, 0xe8, 0xf5, 0xfc, 0xcd, 0x29, 0x5e, 0x4a, 0x69, 0x8a, 0x0a, 0x2a, 0x8a, 0xf5,
	0xec, 0xa9, 0x5c, 0x9d, 0xa5, 0x42, 0xd0, 0x00, 0x1c, 0x12, 0x31, 0x87, 0x70, 0x2e, 0x14, 0x51,
	0x4c, 0xf0, 0x16, 0xb3, 0x78, 0x6b, 0x83, 0x28, 0x16, 0x91, 0x90, 0x24, 0x28, 0x09, 0xeb, 0x37,
	0xc2, 0x53, 0xdb, 0x92, 0xba, 0x70, 0x90, 0x80, 0x54, 0x2e, 0xf8, 0x10, 0x46, 0x99, 0xa3, 0xf6,
	0x1a, 0xff, 0x97, 0x92, 0x20, 0x81, 0x69, 0x34, 0x87, 0xe6, 0xc7, 0x96, 0x66, 0xec, 0x32, 0x5c,
	0xd6, 0xc4, 0x2e, 0x9b, 0xd8, 0x1b, 0x82, 0xf1, 0xda, 0xe4, 0x69, 0xd3, 0xac, 0x5c, 0x36, 0xcd,
	0xb1, 0x23, 0x12, 0x06, 0xab, 0x56, 0xd6, 0xce, 0x72, 0x0b, 0x58, 0xdb, 0xc2, 0x93, 0x3e, 0x48,
	0xc5, 0x78, 0x1e, 0xb3, 0x4e, 0x7c, 0x3f, 0x06, 0x29, 0xa7, 0x07, 0xe6, 0xd0, 0xfc, 0x68, 0x6d,
	0xfa, 0xeb, 0xc9, 0xc2, 0x54, 0x69, 0xfb, 0xaa, 0x50, 0x76, 0x54, 0xcc, 0x38, 0x75, 0xb5, 0x0e,
	0xa8, 0x54, 0xb4, 0xe7, 0x78, 0xfc, 0x4d, 0x2c, 0xc2, 0xb6, 0xc7, 0xe0, 0x2d, 0x1e, 0x63, 0xd9,
	0xee, 0x72, 0x69, 0x75, 0xe4, 0xdd, 0xb1, 0x59, 0xf9, 0x75, 0x6c, 0x56, 0xac, 0x4f, 0x08, 0x4f,
	0x6c, 0x4b, 0xba, 0xc3, 0x28, 0x27, 0xc1, 0x16, 0x57, 0xc0, 0x95, 0x66, 0xe3, 0x91, 0xfc, 0x77,
	0xaa, 0x33, 0x3f, 0xaf, 0x3b, 0x5a, 0x9b, 0xbc, 0x6c, 0x9a, 0x13, 0x65, 0x9f, 0x52, 0xb1, 0xdc,
	0xe1, 0xfc, 0x71, 0xcb, 0xd7, 0x1e, 0xe2, 0x61, 0x96, 0x93, 0xad, 0x26, 0xda, 0x65, 0xd3, 0xfc,
	0xbf, 0xd8, 0x5e, 0x0a, 0x96, 0xdb, 0xda, 0xf2, 0xaf, 0x82, 0x1b, 0x78, 0xf6, 0xa6, 0x8b, 0x72,
	0x41, 0x46, 0x82, 0x4b, 0xb0, 0x66, 0xf0, 0xbd, 0x2b, 0xbd, 0x5a, 0xd2, 0xd2, 0xc7, 0x21, 0x3c,
	0xb8, 0x2d, 0xa9, 0xf6, 0x19, 0xe1, 0xbb, 0xd7, 0x6f, 0x7a, 0xc5, 0xbe, 0x6d, 0x8c, 0xed, 0x9b,
	0x0e, 0xd6, 0xd7, 0xfb, 0xe3, 0xda, 0x81, 0x57, 0xde, 0x7e, 0xfb, 0xf9, 0x61, 0x60, 0xd1, 0x7a,
	0xd0, 0xf5, 0xbf, 0x53, 0x87, 0xd9, 0x98, 0x5e, 0x9f, 0xdd, 0x18, 0x7c, 0x80, 0x70, 0x15, 0xdd,
	0xd7, 0x4e, 0x10, 0x1e, 0xef, 0xba, 0xbe, 0x6a, 0x4f, 0x41, 0x3a, 0x11, 0xfd, 0xd9, 0x5f, 0x23,
	0x7d, 0xc6, 0x2e, 0x86, 0x20, 0x8b, 0xfd, 0x05, 0xe1, 0x89, 0x4d, 0x91, 0x6e, 0x04, 0x42, 0xc2,
	0xc6, 0x1e, 0xe1, 0x1c, 0x02, 0xed, 0x51, 0x4f, 0x31, 0xae, 0x50, 0xfa, 0x8b, 0x7e, 0xa8, 0x76,
	0xfe, 0xb5, 0x3c, 0xff, 0x13, 0x6b, 0xa9, 0xa7, 0xfc, 0x5e, 0x66, 0x51, 0xf7, 0x0a, 0x8f, 0xac,
	0xc6, 0x29, 0xc2, 0x77, 0x36, 0x45, 0xea, 0x82, 0x88, 0x80, 0xb7, 0x7a, 0x3c, 0xee, 0x35, 0x51,
	0x17, 0xa6, 0xaf, 0xf5, 0x85, 0xb5, 0x9b, 0xac, 0xe7, 0x4d, 0x9e, 0x5a, 0xcb, 0x3d, 0x0e, 0x50,
	0xe6, 0xd1, 0x51, 0xa5, 0xb6, 0x7b, 0x7a, 0x6e, 0xa0, 0xb3, 0x73, 0x03, 0xfd, 0x38, 0x37, 0xd0,
	0xfb, 0x0b, 0xa3, 0x72, 0x76, 0x61, 0x54, 0xbe, 0x5f, 0x18, 0x95, 0xdd, 0x97, 0x94, 0xa9, 0xbd,
	0xa4, 0x61, 0x7b, 0x22, 0x74, 0x18, 0xa7, 0xc0, 0x13, 0xa6, 0x8e, 0x16, 0x1a, 0x09, 0x0b, 0xfc,
	0xae, 0xb3, 0x0e, 0x6f, 0x38, 0x47, 0x1d, 0x45, 0x20, 0x1b, 0x43, 0xf9, 0xe7, 0x75, 0xf9, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0x57, 0xae, 0x4c, 0x54, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// RequestRedemption defines a method for requesting burning of qAssets for
	// native assets.
	RequestRedemption(ctx context.Context, in *MsgRequestRedemption, opts ...grpc.CallOption) (*MsgRequestRedemptionResponse, error)
	// SignalIntent defines a method for signalling voting intent for one or more
	// validators.
	SignalIntent(ctx context.Context, in *MsgSignalIntent, opts ...grpc.CallOption) (*MsgSignalIntentResponse, error)
	// SignalIntent defines a method for signalling voting intent for one or more
	// validators.
	GovCloseChannel(ctx context.Context, in *MsgGovCloseChannel, opts ...grpc.CallOption) (*MsgGovCloseChannelResponse, error)
	GovReopenChannel(ctx context.Context, in *MsgGovReopenChannel, opts ...grpc.CallOption) (*MsgGovReopenChannelResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RequestRedemption(ctx context.Context, in *MsgRequestRedemption, opts ...grpc.CallOption) (*MsgRequestRedemptionResponse, error) {
	out := new(MsgRequestRedemptionResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.interchainstaking.v1.Msg/RequestRedemption", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SignalIntent(ctx context.Context, in *MsgSignalIntent, opts ...grpc.CallOption) (*MsgSignalIntentResponse, error) {
	out := new(MsgSignalIntentResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.interchainstaking.v1.Msg/SignalIntent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GovCloseChannel(ctx context.Context, in *MsgGovCloseChannel, opts ...grpc.CallOption) (*MsgGovCloseChannelResponse, error) {
	out := new(MsgGovCloseChannelResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.interchainstaking.v1.Msg/GovCloseChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GovReopenChannel(ctx context.Context, in *MsgGovReopenChannel, opts ...grpc.CallOption) (*MsgGovReopenChannelResponse, error) {
	out := new(MsgGovReopenChannelResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.interchainstaking.v1.Msg/GovReopenChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// RequestRedemption defines a method for requesting burning of qAssets for
	// native assets.
	RequestRedemption(context.Context, *MsgRequestRedemption) (*MsgRequestRedemptionResponse, error)
	// SignalIntent defines a method for signalling voting intent for one or more
	// validators.
	SignalIntent(context.Context, *MsgSignalIntent) (*MsgSignalIntentResponse, error)
	// SignalIntent defines a method for signalling voting intent for one or more
	// validators.
	GovCloseChannel(context.Context, *MsgGovCloseChannel) (*MsgGovCloseChannelResponse, error)
	GovReopenChannel(context.Context, *MsgGovReopenChannel) (*MsgGovReopenChannelResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RequestRedemption(ctx context.Context, req *MsgRequestRedemption) (*MsgRequestRedemptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestRedemption not implemented")
}
func (*UnimplementedMsgServer) SignalIntent(ctx context.Context, req *MsgSignalIntent) (*MsgSignalIntentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignalIntent not implemented")
}
func (*UnimplementedMsgServer) GovCloseChannel(ctx context.Context, req *MsgGovCloseChannel) (*MsgGovCloseChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovCloseChannel not implemented")
}
func (*UnimplementedMsgServer) GovReopenChannel(ctx context.Context, req *MsgGovReopenChannel) (*MsgGovReopenChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovReopenChannel not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RequestRedemption_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestRedemption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestRedemption(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.interchainstaking.v1.Msg/RequestRedemption",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestRedemption(ctx, req.(*MsgRequestRedemption))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SignalIntent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSignalIntent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SignalIntent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.interchainstaking.v1.Msg/SignalIntent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SignalIntent(ctx, req.(*MsgSignalIntent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GovCloseChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGovCloseChannel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GovCloseChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.interchainstaking.v1.Msg/GovCloseChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GovCloseChannel(ctx, req.(*MsgGovCloseChannel))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GovReopenChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGovReopenChannel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GovReopenChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.interchainstaking.v1.Msg/GovReopenChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GovReopenChannel(ctx, req.(*MsgGovReopenChannel))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quicksilver.interchainstaking.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestRedemption",
			Handler:    _Msg_RequestRedemption_Handler,
		},
		{
			MethodName: "SignalIntent",
			Handler:    _Msg_SignalIntent_Handler,
		},
		{
			MethodName: "GovCloseChannel",
			Handler:    _Msg_GovCloseChannel_Handler,
		},
		{
			MethodName: "GovReopenChannel",
			Handler:    _Msg_GovReopenChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quicksilver/interchainstaking/v1/messages.proto",
}

func (m *MsgRequestRedemption) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestRedemption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestRedemption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DestinationAddress) > 0 {
		i -= len(m.DestinationAddress)
		copy(dAtA[i:], m.DestinationAddress)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.DestinationAddress)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMessages(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgSignalIntent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSignalIntent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSignalIntent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Intents) > 0 {
		i -= len(m.Intents)
		copy(dAtA[i:], m.Intents)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Intents)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRequestRedemptionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestRedemptionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestRedemptionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSignalIntentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSignalIntentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSignalIntentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMessages(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessages(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRequestRedemption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Value.Size()
	n += 1 + l + sovMessages(uint64(l))
	l = len(m.DestinationAddress)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *MsgSignalIntent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Intents)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *MsgRequestRedemptionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSignalIntentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMessages(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessages(x uint64) (n int) {
	return sovMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRequestRedemption) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgRequestRedemption: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestRedemption: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSignalIntent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgSignalIntent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSignalIntent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Intents", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Intents = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRequestRedemptionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgRequestRedemptionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestRedemptionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSignalIntentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgSignalIntentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSignalIntentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessages(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessages
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
				return 0, ErrInvalidLengthMessages
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessages
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessages
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessages        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessages          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessages = fmt.Errorf("proto: unexpected end of group")
)
