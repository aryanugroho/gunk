// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testdata.tld/util/echo.gunk

package util

/*
package util contains a simple Echo service.
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import empty "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/gunk/opt/http"
import imported "testdata.tld/util/imported"

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

// CheckStatusResponse is the response for a check status.
type CheckStatusResponse struct {
	Status               Status   `protobuf:"varint,1,,name=Status,proto3,enum=testdata.tld/util.Status" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckStatusResponse) Reset()         { *m = CheckStatusResponse{} }
func (m *CheckStatusResponse) String() string { return proto.CompactTextString(m) }
func (*CheckStatusResponse) ProtoMessage()    {}
func (*CheckStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_151e3c79747eb0f9, []int{0}
}
func (m *CheckStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckStatusResponse.Unmarshal(m, b)
}
func (m *CheckStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckStatusResponse.Marshal(b, m, deterministic)
}
func (dst *CheckStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckStatusResponse.Merge(dst, src)
}
func (m *CheckStatusResponse) XXX_Size() int {
	return xxx_messageInfo_CheckStatusResponse.Size(m)
}
func (m *CheckStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckStatusResponse proto.InternalMessageInfo

func (m *CheckStatusResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_Unknown
}

type UtilTestRequest struct {
	Ints                 []int32              `protobuf:"varint,1,rep,packed,name=Ints,proto3" json:"Ints,omitempty"`
	Structs              []*imported.Message  `protobuf:"bytes,2,rep,name=Structs,proto3" json:"Structs,omitempty"`
	Bool                 bool                 `protobuf:"varint,3,,name=Bool,proto3" json:"Bool,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,4,,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Duration             *duration.Duration   `protobuf:"bytes,5,,name=Duration,proto3" json:"Duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UtilTestRequest) Reset()         { *m = UtilTestRequest{} }
func (m *UtilTestRequest) String() string { return proto.CompactTextString(m) }
func (*UtilTestRequest) ProtoMessage()    {}
func (*UtilTestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_151e3c79747eb0f9, []int{1}
}
func (m *UtilTestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UtilTestRequest.Unmarshal(m, b)
}
func (m *UtilTestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UtilTestRequest.Marshal(b, m, deterministic)
}
func (dst *UtilTestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UtilTestRequest.Merge(dst, src)
}
func (m *UtilTestRequest) XXX_Size() int {
	return xxx_messageInfo_UtilTestRequest.Size(m)
}
func (m *UtilTestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UtilTestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UtilTestRequest proto.InternalMessageInfo

func (m *UtilTestRequest) GetInts() []int32 {
	if m != nil {
		return m.Ints
	}
	return nil
}

func (m *UtilTestRequest) GetStructs() []*imported.Message {
	if m != nil {
		return m.Structs
	}
	return nil
}

func (m *UtilTestRequest) GetBool() bool {
	if m != nil {
		return m.Bool
	}
	return false
}

func (m *UtilTestRequest) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *UtilTestRequest) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func init() {
	proto.RegisterType((*CheckStatusResponse)(nil), "testdata.tld/util.CheckStatusResponse")
	proto.RegisterType((*UtilTestRequest)(nil), "testdata.tld/util.UtilTestRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UtilClient is the client API for Util service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UtilClient interface {
	// Echo echoes a message.
	Echo(ctx context.Context, in *imported.Message, opts ...grpc.CallOption) (*imported.Message, error)
	// CheckStatus sends the server health status.
	CheckStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CheckStatusResponse, error)
}

type utilClient struct {
	cc *grpc.ClientConn
}

func NewUtilClient(cc *grpc.ClientConn) UtilClient {
	return &utilClient{cc}
}

func (c *utilClient) Echo(ctx context.Context, in *imported.Message, opts ...grpc.CallOption) (*imported.Message, error) {
	out := new(imported.Message)
	err := c.cc.Invoke(ctx, "/testdata.tld/util.Util/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *utilClient) CheckStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CheckStatusResponse, error) {
	out := new(CheckStatusResponse)
	err := c.cc.Invoke(ctx, "/testdata.tld/util.Util/CheckStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UtilServer is the server API for Util service.
type UtilServer interface {
	// Echo echoes a message.
	Echo(context.Context, *imported.Message) (*imported.Message, error)
	// CheckStatus sends the server health status.
	CheckStatus(context.Context, *empty.Empty) (*CheckStatusResponse, error)
}

func RegisterUtilServer(s *grpc.Server, srv UtilServer) {
	s.RegisterService(&_Util_serviceDesc, srv)
}

func _Util_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(imported.Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdata.tld/util.Util/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilServer).Echo(ctx, req.(*imported.Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Util_CheckStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilServer).CheckStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdata.tld/util.Util/CheckStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilServer).CheckStatus(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Util_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testdata.tld/util.Util",
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
	Metadata: "testdata.tld/util/echo.gunk",
}

// UtilTestsClient is the client API for UtilTests service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UtilTestsClient interface {
	UtilTest(ctx context.Context, in *UtilTestRequest, opts ...grpc.CallOption) (*CheckStatusResponse, error)
}

type utilTestsClient struct {
	cc *grpc.ClientConn
}

func NewUtilTestsClient(cc *grpc.ClientConn) UtilTestsClient {
	return &utilTestsClient{cc}
}

func (c *utilTestsClient) UtilTest(ctx context.Context, in *UtilTestRequest, opts ...grpc.CallOption) (*CheckStatusResponse, error) {
	out := new(CheckStatusResponse)
	err := c.cc.Invoke(ctx, "/testdata.tld/util.UtilTests/UtilTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UtilTestsServer is the server API for UtilTests service.
type UtilTestsServer interface {
	UtilTest(context.Context, *UtilTestRequest) (*CheckStatusResponse, error)
}

func RegisterUtilTestsServer(s *grpc.Server, srv UtilTestsServer) {
	s.RegisterService(&_UtilTests_serviceDesc, srv)
}

func _UtilTests_UtilTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UtilTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilTestsServer).UtilTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdata.tld/util.UtilTests/UtilTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilTestsServer).UtilTest(ctx, req.(*UtilTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UtilTests_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testdata.tld/util.UtilTests",
	HandlerType: (*UtilTestsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UtilTest",
			Handler:    _UtilTests_UtilTest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testdata.tld/util/echo.gunk",
}

func init() { proto.RegisterFile("testdata.tld/util/echo.gunk", fileDescriptor_echo_151e3c79747eb0f9) }

var fileDescriptor_echo_151e3c79747eb0f9 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x8e, 0x94, 0x40,
	0x10, 0xc6, 0x87, 0x1d, 0x76, 0x64, 0x7b, 0xfc, 0xdb, 0x9b, 0x18, 0x96, 0x35, 0x4a, 0xd8, 0xc4,
	0xa0, 0xc6, 0xee, 0x88, 0x9e, 0x3c, 0x99, 0xd5, 0x3d, 0x78, 0xf0, 0x32, 0x33, 0x7a, 0xf0, 0x62,
	0x18, 0x28, 0x81, 0x0c, 0xd0, 0x2d, 0x5d, 0x6d, 0x32, 0x57, 0x5f, 0xc1, 0x07, 0xf2, 0x0d, 0xbc,
	0xf8, 0x0a, 0x3e, 0x88, 0xa1, 0x19, 0x26, 0x66, 0x98, 0x4d, 0xe6, 0x42, 0xe8, 0xae, 0x1f, 0x5f,
	0x7d, 0x7c, 0x55, 0xe4, 0x1c, 0x41, 0x61, 0x1a, 0x63, 0xcc, 0xb0, 0x4c, 0xb9, 0xc6, 0xa2, 0xe4,
	0x90, 0xe4, 0x82, 0x65, 0xba, 0x5e, 0xd1, 0x7b, 0x83, 0xa2, 0x77, 0x9e, 0x09, 0x91, 0x95, 0xc0,
	0x65, 0x23, 0x50, 0x2c, 0xf5, 0x57, 0x0e, 0x95, 0xc4, 0x35, 0x33, 0x47, 0xef, 0xd1, 0x6e, 0x11,
	0x8b, 0x0a, 0x14, 0xc6, 0x95, 0xdc, 0x00, 0x0f, 0x77, 0x81, 0x54, 0x37, 0x31, 0x16, 0xa2, 0xde,
	0xd4, 0x1f, 0x0c, 0xdd, 0xe0, 0x5a, 0x82, 0x32, 0x76, 0xbc, 0x20, 0x2b, 0x30, 0xd7, 0x4b, 0x96,
	0x88, 0x8a, 0xb7, 0x17, 0x5c, 0x48, 0xe4, 0x39, 0xa2, 0x34, 0x8f, 0x8e, 0xb9, 0x18, 0x2a, 0x14,
	0x95, 0x14, 0x0d, 0x42, 0xda, 0xbe, 0x18, 0x28, 0x78, 0x43, 0x4e, 0xdf, 0xe6, 0x90, 0xac, 0xe6,
	0x18, 0xa3, 0x56, 0x33, 0x50, 0x52, 0xd4, 0x0a, 0xe8, 0x13, 0x32, 0xe9, 0x6e, 0x5c, 0xcb, 0x1f,
	0x85, 0xb7, 0xa3, 0x33, 0x36, 0x10, 0x63, 0x1d, 0x10, 0xfc, 0xb2, 0xc8, 0x9d, 0x8f, 0x58, 0x94,
	0x0b, 0x50, 0x38, 0x83, 0x6f, 0x1a, 0x14, 0xd2, 0x9b, 0xc4, 0x7e, 0x5f, 0x63, 0xfb, 0xf1, 0x38,
	0x3c, 0xa6, 0xaf, 0xc8, 0x8d, 0x39, 0x36, 0x3a, 0x41, 0xe5, 0x1e, 0xf9, 0xe3, 0x70, 0x1a, 0x5d,
	0xb0, 0xeb, 0xad, 0xb1, 0x0f, 0xa0, 0x54, 0x9c, 0x41, 0xab, 0x71, 0x29, 0x44, 0xe9, 0x8e, 0xfd,
	0x51, 0xe8, 0xd0, 0xe7, 0xe4, 0x64, 0xd1, 0x27, 0xe8, 0xda, 0xfe, 0x28, 0x9c, 0x46, 0x1e, 0xeb,
	0x22, 0x64, 0x7d, 0x84, 0x6c, 0x4b, 0xd0, 0x67, 0xc4, 0x79, 0xb7, 0xc9, 0xd3, 0x3d, 0x36, 0xf4,
	0xd9, 0x80, 0xee, 0x81, 0xe8, 0xb7, 0x45, 0xec, 0xf6, 0x0f, 0x28, 0x10, 0xfb, 0x2a, 0xc9, 0x05,
	0x3d, 0xc4, 0x9f, 0x77, 0x08, 0x14, 0x9c, 0xfe, 0xf8, 0xf3, 0xf7, 0xe7, 0xd1, 0xad, 0xd7, 0xd6,
	0xd3, 0xc0, 0xe1, 0xdf, 0x5f, 0x98, 0x85, 0xa2, 0x5f, 0xc8, 0xf4, 0xbf, 0xcc, 0xe9, 0xfd, 0x81,
	0xb3, 0xab, 0x76, 0x91, 0xbc, 0xc7, 0x7b, 0x32, 0xdf, 0x33, 0xab, 0xe0, 0xae, 0xe9, 0x41, 0xe8,
	0xb6, 0x41, 0x94, 0x90, 0x93, 0x7e, 0x22, 0x8a, 0x7e, 0x22, 0x4e, 0x7f, 0xa0, 0xc1, 0x1e, 0xc9,
	0x9d, 0xd9, 0x1d, 0xda, 0xf6, 0x72, 0xf2, 0xd9, 0x6e, 0x6b, 0xcb, 0x89, 0xb1, 0xfd, 0xf2, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x18, 0x1c, 0x0d, 0x6b, 0x3f, 0x03, 0x00, 0x00,
}
