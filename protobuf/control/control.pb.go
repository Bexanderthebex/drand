// Code generated by protoc-gen-go. DO NOT EDIT.
// source: control/control.proto

/*
Package control is a generated protocol buffer package.

It is generated from these files:
	control/control.proto

It has these top-level messages:
	ShareRequest
	ShareResponse
*/
package control

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import element "github.com/dedis/drand/protobuf/crypto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Sharerequest requests the private share of a drand node
type ShareRequest struct {
}

func (m *ShareRequest) Reset()                    { *m = ShareRequest{} }
func (m *ShareRequest) String() string            { return proto.CompactTextString(m) }
func (*ShareRequest) ProtoMessage()               {}
func (*ShareRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// ShareResponse holds the private share of a srand node
type ShareResponse struct {
	Index uint32          `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	Share *element.Scalar `protobuf:"bytes,2,opt,name=share" json:"share,omitempty"`
}

func (m *ShareResponse) Reset()                    { *m = ShareResponse{} }
func (m *ShareResponse) String() string            { return proto.CompactTextString(m) }
func (*ShareResponse) ProtoMessage()               {}
func (*ShareResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ShareResponse) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ShareResponse) GetShare() *element.Scalar {
	if m != nil {
		return m.Share
	}
	return nil
}

func init() {
	proto.RegisterType((*ShareRequest)(nil), "control.ShareRequest")
	proto.RegisterType((*ShareResponse)(nil), "control.ShareResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Control service

type ControlClient interface {
	Share(ctx context.Context, in *ShareRequest, opts ...grpc.CallOption) (*ShareResponse, error)
}

type controlClient struct {
	cc *grpc.ClientConn
}

func NewControlClient(cc *grpc.ClientConn) ControlClient {
	return &controlClient{cc}
}

func (c *controlClient) Share(ctx context.Context, in *ShareRequest, opts ...grpc.CallOption) (*ShareResponse, error) {
	out := new(ShareResponse)
	err := grpc.Invoke(ctx, "/control.Control/Share", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Control service

type ControlServer interface {
	Share(context.Context, *ShareRequest) (*ShareResponse, error)
}

func RegisterControlServer(s *grpc.Server, srv ControlServer) {
	s.RegisterService(&_Control_serviceDesc, srv)
}

func _Control_Share_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).Share(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control.Control/Share",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).Share(ctx, req.(*ShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Control_serviceDesc = grpc.ServiceDesc{
	ServiceName: "control.Control",
	HandlerType: (*ControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Share",
			Handler:    _Control_Share_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "control/control.proto",
}

func init() { proto.RegisterFile("control/control.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0xcf, 0x2b,
	0x29, 0xca, 0xcf, 0xd1, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x94, 0x48, 0x72, 0x51, 0x65, 0x41, 0x49, 0xbe, 0x7e, 0x6a, 0x4e, 0x6a, 0x6e, 0x6a, 0x5e, 0x09,
	0x44, 0x5a, 0x89, 0x8f, 0x8b, 0x27, 0x38, 0x23, 0xb1, 0x28, 0x35, 0x28, 0xb5, 0xb0, 0x34, 0xb5,
	0xb8, 0x44, 0xc9, 0x87, 0x8b, 0x17, 0xca, 0x2f, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe1,
	0x62, 0xcd, 0xcc, 0x4b, 0x49, 0xad, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0d, 0x82, 0x70, 0x84,
	0x54, 0xb9, 0x58, 0x8b, 0x41, 0xca, 0x24, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0xf8, 0xf5, 0x60,
	0xa6, 0x06, 0x27, 0x27, 0xe6, 0x24, 0x16, 0x05, 0x41, 0x64, 0x8d, 0x9c, 0xb9, 0xd8, 0x9d, 0x21,
	0xd6, 0x0b, 0x59, 0x70, 0xb1, 0x82, 0x0d, 0x16, 0x12, 0xd5, 0x83, 0x39, 0x10, 0xd9, 0x62, 0x29,
	0x31, 0x74, 0x61, 0x88, 0xfd, 0x4a, 0x0c, 0x4e, 0x9a, 0x51, 0xea, 0xe9, 0x99, 0x25, 0x19, 0xa5,
	0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x29, 0xa9, 0x29, 0x99, 0xc5, 0xfa, 0x29, 0x45, 0x89, 0x79,
	0x29, 0xfa, 0x60, 0x2f, 0x24, 0x95, 0xa6, 0xc1, 0xbc, 0x9c, 0xc4, 0x06, 0x16, 0x31, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x80, 0xb5, 0x10, 0x15, 0x0c, 0x01, 0x00, 0x00,
}
