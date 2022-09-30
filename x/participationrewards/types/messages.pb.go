// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quicksilver/participationrewards/v1/messages.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	crypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
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

// MsgSubmitClaim represents a message type for submitting a participation
// claim regarding the given zone (chain).
type MsgSubmitClaim struct {
	UserAddress string    `protobuf:"bytes,1,opt,name=user_address,json=userAddress,proto3" json:"user_address,omitempty"`
	Zone        string    `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	SrcZone     string    `protobuf:"bytes,3,opt,name=src_zone,json=srcZone,proto3" json:"src_zone,omitempty"`
	ClaimType   ClaimType `protobuf:"varint,4,opt,name=claim_type,json=claimType,proto3,enum=quicksilver.participationrewards.v1.ClaimType" json:"claim_type,omitempty"`
	Proofs      []*Proof  `protobuf:"bytes,5,rep,name=proofs,proto3" json:"proofs,omitempty" yaml:"proofs"`
}

func (m *MsgSubmitClaim) Reset()         { *m = MsgSubmitClaim{} }
func (m *MsgSubmitClaim) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitClaim) ProtoMessage()    {}
func (*MsgSubmitClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87e3ea017f90b50, []int{0}
}
func (m *MsgSubmitClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitClaim.Merge(m, src)
}
func (m *MsgSubmitClaim) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitClaim.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitClaim proto.InternalMessageInfo

// MsgSubmitClaimResponse defines the MsgSubmitClaim response type.
type MsgSubmitClaimResponse struct {
}

func (m *MsgSubmitClaimResponse) Reset()         { *m = MsgSubmitClaimResponse{} }
func (m *MsgSubmitClaimResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitClaimResponse) ProtoMessage()    {}
func (*MsgSubmitClaimResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87e3ea017f90b50, []int{1}
}
func (m *MsgSubmitClaimResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitClaimResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitClaimResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitClaimResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitClaimResponse.Merge(m, src)
}
func (m *MsgSubmitClaimResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitClaimResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitClaimResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitClaimResponse proto.InternalMessageInfo

// Proof defines a type used to cryptographically prove a claim.
type Proof struct {
	Key       []byte           `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" yaml:"key"`
	Data      []byte           `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty" yaml:"data"`
	ProofOps  *crypto.ProofOps `protobuf:"bytes,3,opt,name=proof_ops,json=proofOps,proto3" json:"proof_ops,omitempty" yaml:"proof_ops"`
	Height    int64            `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty" yaml:"height"`
	ProofType string           `protobuf:"bytes,5,opt,name=proof_type,json=proofType,proto3" json:"proof_type,omitempty" yaml:"proof_type"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87e3ea017f90b50, []int{2}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(m, src)
}
func (m *Proof) XXX_Size() int {
	return m.Size()
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSubmitClaim)(nil), "quicksilver.participationrewards.v1.MsgSubmitClaim")
	proto.RegisterType((*MsgSubmitClaimResponse)(nil), "quicksilver.participationrewards.v1.MsgSubmitClaimResponse")
	proto.RegisterType((*Proof)(nil), "quicksilver.participationrewards.v1.Proof")
}

func init() {
	proto.RegisterFile("quicksilver/participationrewards/v1/messages.proto", fileDescriptor_b87e3ea017f90b50)
}

var fileDescriptor_b87e3ea017f90b50 = []byte{
	// 611 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x41, 0x4f, 0xd4, 0x40,
	0x14, 0xde, 0xee, 0x02, 0xc2, 0x2c, 0xa2, 0x4c, 0xd0, 0x14, 0xd4, 0x76, 0x53, 0x2e, 0x48, 0x42,
	0x1b, 0x8a, 0x07, 0x03, 0x89, 0x89, 0xe5, 0x8c, 0x9a, 0xa2, 0x17, 0x12, 0xb3, 0xe9, 0xb6, 0x63,
	0x99, 0xb0, 0xed, 0x8c, 0x33, 0x53, 0xa4, 0x1e, 0x3d, 0xe9, 0xcd, 0xc4, 0x3f, 0xc0, 0x0f, 0xf0,
	0xe8, 0xc5, 0xa3, 0x37, 0x8f, 0x44, 0x2f, 0x9e, 0x36, 0x06, 0x3c, 0x78, 0xde, 0x5f, 0x60, 0x66,
	0xa6, 0x48, 0x49, 0xd6, 0x64, 0xe3, 0x6d, 0xfa, 0xbe, 0xf7, 0x7d, 0xef, 0xf5, 0x7b, 0xef, 0x01,
	0xff, 0x65, 0x81, 0xe3, 0x03, 0x8e, 0xfb, 0x87, 0x88, 0x79, 0x34, 0x62, 0x02, 0xc7, 0x98, 0x46,
	0x02, 0x93, 0x9c, 0xa1, 0x57, 0x11, 0x4b, 0xb8, 0x77, 0xb8, 0xee, 0x65, 0x88, 0xf3, 0x28, 0x45,
	0xdc, 0xa5, 0x8c, 0x08, 0x02, 0x97, 0x6b, 0x1c, 0x77, 0x14, 0xc7, 0x3d, 0x5c, 0x5f, 0x5a, 0x48,
	0x49, 0x4a, 0x54, 0xbe, 0x27, 0x5f, 0x9a, 0xba, 0xb4, 0x18, 0x13, 0x9e, 0x11, 0xde, 0xd5, 0x80,
	0xfe, 0xa8, 0xa0, 0xdb, 0x29, 0x21, 0x69, 0x1f, 0x79, 0x11, 0xc5, 0x5e, 0x94, 0xe7, 0x44, 0x28,
	0xc5, 0x73, 0xf4, 0x8e, 0x40, 0x79, 0x82, 0x58, 0x86, 0x73, 0xe1, 0xc5, 0xac, 0xa4, 0x82, 0x78,
	0x94, 0x11, 0xf2, 0xa2, 0x82, 0x1f, 0x8c, 0xf3, 0x1b, 0x23, 0x5b, 0x55, 0x7c, 0xe7, 0x63, 0x13,
	0xcc, 0xed, 0xf0, 0x74, 0xb7, 0xe8, 0x65, 0x58, 0x6c, 0xf7, 0x23, 0x9c, 0xc1, 0x2d, 0x30, 0x5b,
	0x70, 0xc4, 0xba, 0x51, 0x92, 0x30, 0xc4, 0xb9, 0x69, 0x74, 0x8c, 0x95, 0x99, 0xc0, 0xfc, 0xf6,
	0x69, 0x6d, 0xa1, 0xea, 0xfb, 0xa1, 0x46, 0x76, 0x05, 0xc3, 0x79, 0x1a, 0xb6, 0x65, 0x76, 0x15,
	0x82, 0x10, 0x4c, 0xbc, 0x26, 0x39, 0x32, 0x9b, 0x92, 0x14, 0xaa, 0x37, 0x5c, 0x04, 0xd3, 0x9c,
	0xc5, 0x5d, 0x15, 0x6f, 0xa9, 0xf8, 0x15, 0xce, 0xe2, 0x3d, 0x09, 0xed, 0x00, 0x10, 0xcb, 0xa2,
	0x5d, 0x51, 0x52, 0x64, 0x4e, 0x74, 0x8c, 0x95, 0x39, 0xdf, 0x75, 0xc7, 0xb0, 0xd9, 0x55, 0xbd,
	0x3e, 0x2d, 0x29, 0x0a, 0x67, 0xe2, 0xf3, 0x27, 0x7c, 0x06, 0xa6, 0x94, 0x39, 0xdc, 0x9c, 0xec,
	0xb4, 0x56, 0xda, 0xfe, 0xea, 0x58, 0x52, 0x4f, 0x24, 0x25, 0x98, 0x1f, 0x0e, 0xec, 0xab, 0x65,
	0x94, 0xf5, 0x37, 0x1d, 0xad, 0xe1, 0x84, 0x95, 0xd8, 0xe6, 0xf4, 0xdb, 0x63, 0xbb, 0xf1, 0xfb,
	0xd8, 0x6e, 0x38, 0x26, 0xb8, 0x79, 0xd9, 0xad, 0x10, 0x71, 0x4a, 0x72, 0x8e, 0x9c, 0x77, 0x4d,
	0x30, 0xa9, 0x84, 0x60, 0x07, 0xb4, 0x0e, 0x50, 0xa9, 0x6c, 0x9b, 0x0d, 0xe6, 0x86, 0x03, 0x1b,
	0x68, 0xd5, 0x03, 0x54, 0x3a, 0xa1, 0x84, 0xe0, 0x32, 0x98, 0x48, 0x22, 0x11, 0x29, 0x93, 0x66,
	0x83, 0x6b, 0xc3, 0x81, 0xdd, 0xd6, 0x29, 0x32, 0xea, 0x84, 0x0a, 0x84, 0x8f, 0xc0, 0x8c, 0x2a,
	0xdf, 0x25, 0x94, 0x2b, 0xdb, 0xda, 0xfe, 0x2d, 0xf7, 0x62, 0x19, 0x5c, 0xbd, 0x0c, 0xba, 0xf9,
	0xc7, 0x94, 0x07, 0x0b, 0xc3, 0x81, 0x7d, 0xbd, 0xd6, 0xbf, 0xe4, 0x39, 0xe1, 0x34, 0xad, 0x70,
	0x78, 0x17, 0x4c, 0xed, 0x23, 0x9c, 0xee, 0x0b, 0x65, 0x73, 0xab, 0xfe, 0xbf, 0x3a, 0xee, 0x84,
	0x55, 0x02, 0xbc, 0x07, 0x80, 0x96, 0x50, 0x53, 0x99, 0x54, 0xf3, 0xbf, 0x31, 0x1c, 0xd8, 0xf3,
	0x75, 0x79, 0x89, 0x39, 0xa1, 0xee, 0x51, 0x9a, 0x7f, 0xe1, 0x92, 0xff, 0xc5, 0x00, 0xad, 0x1d,
	0x9e, 0xc2, 0xcf, 0x06, 0x68, 0xd7, 0x37, 0x6b, 0x63, 0xac, 0x71, 0x5c, 0x36, 0x78, 0x69, 0xeb,
	0x3f, 0x48, 0x7f, 0xa7, 0x72, 0xff, 0xcd, 0xf7, 0x5f, 0x1f, 0x9a, 0xbe, 0xb3, 0xe6, 0xd5, 0xef,
	0x44, 0x1c, 0xfd, 0xeb, 0x2a, 0x3c, 0xb5, 0x4b, 0x9b, 0xc6, 0x6a, 0xf0, 0xfc, 0xeb, 0xa9, 0x65,
	0x9c, 0x9c, 0x5a, 0xc6, 0xcf, 0x53, 0xcb, 0x78, 0x7f, 0x66, 0x35, 0x4e, 0xce, 0xac, 0xc6, 0x8f,
	0x33, 0xab, 0xb1, 0xb7, 0x9d, 0x62, 0xb1, 0x5f, 0xf4, 0xdc, 0x98, 0x64, 0x1e, 0xce, 0x53, 0x94,
	0x17, 0x58, 0x94, 0x6b, 0xbd, 0x02, 0xf7, 0x93, 0x4b, 0x55, 0x8e, 0x46, 0x57, 0x90, 0xbe, 0xf1,
	0xde, 0x94, 0x3a, 0xbf, 0x8d, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x56, 0x20, 0xd5, 0xdd, 0x87,
	0x04, 0x00, 0x00,
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
	SubmitClaim(ctx context.Context, in *MsgSubmitClaim, opts ...grpc.CallOption) (*MsgSubmitClaimResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitClaim(ctx context.Context, in *MsgSubmitClaim, opts ...grpc.CallOption) (*MsgSubmitClaimResponse, error) {
	out := new(MsgSubmitClaimResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.participationrewards.v1.Msg/SubmitClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SubmitClaim(context.Context, *MsgSubmitClaim) (*MsgSubmitClaimResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SubmitClaim(ctx context.Context, req *MsgSubmitClaim) (*MsgSubmitClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitClaim not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SubmitClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitClaim)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.participationrewards.v1.Msg/SubmitClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitClaim(ctx, req.(*MsgSubmitClaim))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quicksilver.participationrewards.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitClaim",
			Handler:    _Msg_SubmitClaim_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quicksilver/participationrewards/v1/messages.proto",
}

func (m *MsgSubmitClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proofs) > 0 {
		for iNdEx := len(m.Proofs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proofs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMessages(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.ClaimType != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.ClaimType))
		i--
		dAtA[i] = 0x20
	}
	if len(m.SrcZone) > 0 {
		i -= len(m.SrcZone)
		copy(dAtA[i:], m.SrcZone)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.SrcZone)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Zone) > 0 {
		i -= len(m.Zone)
		copy(dAtA[i:], m.Zone)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Zone)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UserAddress) > 0 {
		i -= len(m.UserAddress)
		copy(dAtA[i:], m.UserAddress)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.UserAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitClaimResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitClaimResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitClaimResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Proof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProofType) > 0 {
		i -= len(m.ProofType)
		copy(dAtA[i:], m.ProofType)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.ProofType)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Height != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x20
	}
	if m.ProofOps != nil {
		{
			size, err := m.ProofOps.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessages(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
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
func (m *MsgSubmitClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserAddress)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Zone)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.SrcZone)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.ClaimType != 0 {
		n += 1 + sovMessages(uint64(m.ClaimType))
	}
	if len(m.Proofs) > 0 {
		for _, e := range m.Proofs {
			l = e.Size()
			n += 1 + l + sovMessages(uint64(l))
		}
	}
	return n
}

func (m *MsgSubmitClaimResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Proof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.ProofOps != nil {
		l = m.ProofOps.Size()
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovMessages(uint64(m.Height))
	}
	l = len(m.ProofType)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func sovMessages(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessages(x uint64) (n int) {
	return sovMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSubmitClaim) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSubmitClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserAddress", wireType)
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
			m.UserAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Zone", wireType)
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
			m.Zone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcZone", wireType)
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
			m.SrcZone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimType", wireType)
			}
			m.ClaimType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimType |= ClaimType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proofs", wireType)
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
			m.Proofs = append(m.Proofs, &Proof{})
			if err := m.Proofs[len(m.Proofs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgSubmitClaimResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSubmitClaimResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitClaimResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *Proof) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Proof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofOps", wireType)
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
			if m.ProofOps == nil {
				m.ProofOps = &crypto.ProofOps{}
			}
			if err := m.ProofOps.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofType", wireType)
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
			m.ProofType = string(dAtA[iNdEx:postIndex])
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
