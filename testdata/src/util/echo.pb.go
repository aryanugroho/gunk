// Code generated by protoc-gen-go. DO NOT EDIT.
// source: util/echo.gunk

/*
Package util is a generated protocol buffer package.

package util contains a simple Echo service.

It is generated from these files:
	util/echo.gunk
	util/types.gunk

It has these top-level messages:
	CheckStatusResponse
*/
package util

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import util_imp_arg "util/imp-arg"
import _ "util/imp-noarg"

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

type CheckStatusResponse struct {
	Status Status `protobuf:"varint,1,,name=Status,enum=util.Status" json:"Status,omitempty"`
}

func (m *CheckStatusResponse) Reset()                    { *m = CheckStatusResponse{} }
func (m *CheckStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckStatusResponse) ProtoMessage()               {}
func (*CheckStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CheckStatusResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_Unknown
}

func init() {
	proto.RegisterType((*CheckStatusResponse)(nil), "util.CheckStatusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Util service

type UtilClient interface {
	// Echo echoes a message.
	Echo(ctx context.Context, in *util_imp_arg.Message, opts ...grpc.CallOption) (*util_imp_arg.Message, error)
	// CheckStatus sends the server health status.
	CheckStatus(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*CheckStatusResponse, error)
}

type utilClient struct {
	cc *grpc.ClientConn
}

func NewUtilClient(cc *grpc.ClientConn) UtilClient {
	return &utilClient{cc}
}

func (c *utilClient) Echo(ctx context.Context, in *util_imp_arg.Message, opts ...grpc.CallOption) (*util_imp_arg.Message, error) {
	out := new(util_imp_arg.Message)
	err := grpc.Invoke(ctx, "/util.Util/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *utilClient) CheckStatus(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*CheckStatusResponse, error) {
	out := new(CheckStatusResponse)
	err := grpc.Invoke(ctx, "/util.Util/CheckStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Util service

type UtilServer interface {
	// Echo echoes a message.
	Echo(context.Context, *util_imp_arg.Message) (*util_imp_arg.Message, error)
	// CheckStatus sends the server health status.
	CheckStatus(context.Context, *google_protobuf.Empty) (*CheckStatusResponse, error)
}

func RegisterUtilServer(s *grpc.Server, srv UtilServer) {
	s.RegisterService(&_Util_serviceDesc, srv)
}

func _Util_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(util_imp_arg.Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/util.Util/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilServer).Echo(ctx, req.(*util_imp_arg.Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Util_CheckStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilServer).CheckStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/util.Util/CheckStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilServer).CheckStatus(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Util_serviceDesc = grpc.ServiceDesc{
	ServiceName: "util.Util",
	HandlerType: (*UtilServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Util_Echo_Handler,
		},
		{
			MethodName: "CheckStatus",
			Handler:    _Util_CheckStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "util/echo.gunk",
}

func init() { proto.RegisterFile("util/echo.gunk", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2d, 0xc9, 0xcc,
	0xd1, 0x4f, 0x4d, 0xce, 0xc8, 0xd7, 0x4b, 0x2f, 0xcd, 0xcb, 0x16, 0x62, 0x01, 0xf1, 0xa5, 0xa4,
	0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x0b, 0x8a, 0xf2, 0x4b, 0xf2, 0x93, 0x4a, 0xd3, 0xf4,
	0x53, 0x73, 0x0b, 0x4a, 0x2a, 0xf5, 0xc0, 0x5c, 0x29, 0x7e, 0xb0, 0x96, 0x92, 0xca, 0x82, 0xd4,
	0x62, 0xb0, 0x1e, 0x29, 0x51, 0xb0, 0x40, 0x66, 0x6e, 0x81, 0x6e, 0x62, 0x51, 0x3a, 0x88, 0x86,
	0x08, 0x8b, 0xc3, 0x85, 0xf3, 0xf2, 0x91, 0x25, 0x94, 0x0c, 0xb9, 0x84, 0x9d, 0x33, 0x52, 0x93,
	0xb3, 0x83, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x83, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85,
	0xa4, 0xb8, 0xd8, 0x20, 0x22, 0x12, 0x8c, 0x1a, 0x7c, 0x46, 0x3c, 0x7a, 0x20, 0xed, 0x7a, 0x10,
	0x11, 0xa3, 0x3a, 0x2e, 0x96, 0xd0, 0x92, 0xcc, 0x1c, 0x21, 0x13, 0x2e, 0x16, 0xd7, 0xe4, 0x8c,
	0x7c, 0x21, 0x51, 0x3d, 0x64, 0x3b, 0xf5, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0xa5, 0xb0,
	0x0b, 0x0b, 0x39, 0x70, 0x71, 0x23, 0x59, 0x28, 0x24, 0xa6, 0x07, 0xf1, 0x9e, 0x1e, 0xcc, 0x7b,
	0x7a, 0xae, 0x20, 0xef, 0x49, 0x49, 0x42, 0xac, 0xc4, 0xe2, 0x36, 0x27, 0xb6, 0x28, 0x70, 0xc0,
	0x24, 0xb1, 0x81, 0xb5, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xb4, 0x12, 0x9a, 0x37,
	0x01, 0x00, 0x00,
}
