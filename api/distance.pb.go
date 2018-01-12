// Code generated by protoc-gen-go. DO NOT EDIT.
// source: distance.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	distance.proto

It has these top-level messages:
	Query
	ResponseElement
	Response
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Query struct {
	Id  string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Top int32  `protobuf:"varint,2,opt,name=top" json:"top,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Query) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Query) GetTop() int32 {
	if m != nil {
		return m.Top
	}
	return 0
}

type ResponseElement struct {
	Id   string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Dist float32 `protobuf:"fixed32,2,opt,name=dist" json:"dist,omitempty"`
}

func (m *ResponseElement) Reset()                    { *m = ResponseElement{} }
func (m *ResponseElement) String() string            { return proto.CompactTextString(m) }
func (*ResponseElement) ProtoMessage()               {}
func (*ResponseElement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ResponseElement) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ResponseElement) GetDist() float32 {
	if m != nil {
		return m.Dist
	}
	return 0
}

type Response struct {
	Responses []*ResponseElement `protobuf:"bytes,1,rep,name=responses" json:"responses,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetResponses() []*ResponseElement {
	if m != nil {
		return m.Responses
	}
	return nil
}

func init() {
	proto.RegisterType((*Query)(nil), "api.Query")
	proto.RegisterType((*ResponseElement)(nil), "api.ResponseElement")
	proto.RegisterType((*Response)(nil), "api.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Distance service

type DistanceClient interface {
	Dist(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
}

type distanceClient struct {
	cc *grpc.ClientConn
}

func NewDistanceClient(cc *grpc.ClientConn) DistanceClient {
	return &distanceClient{cc}
}

func (c *distanceClient) Dist(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/api.Distance/dist", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Distance service

type DistanceServer interface {
	Dist(context.Context, *Query) (*Response, error)
}

func RegisterDistanceServer(s *grpc.Server, srv DistanceServer) {
	s.RegisterService(&_Distance_serviceDesc, srv)
}

func _Distance_Dist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistanceServer).Dist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Distance/Dist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistanceServer).Dist(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

var _Distance_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Distance",
	HandlerType: (*DistanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "dist",
			Handler:    _Distance_Dist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "distance.proto",
}

func init() { proto.RegisterFile("distance.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xc9, 0x2c, 0x2e,
	0x49, 0xcc, 0x4b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54,
	0xd2, 0xe4, 0x62, 0x0d, 0x2c, 0x4d, 0x2d, 0xaa, 0x14, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe0, 0x62, 0x2e, 0xc9, 0x2f, 0x90,
	0x60, 0x52, 0x60, 0xd4, 0x60, 0x0d, 0x02, 0x31, 0x95, 0x4c, 0xb9, 0xf8, 0x83, 0x52, 0x8b, 0x0b,
	0xf2, 0xf3, 0x8a, 0x53, 0x5d, 0x73, 0x52, 0x73, 0x53, 0xf3, 0x4a, 0x30, 0x34, 0x09, 0x71, 0xb1,
	0x80, 0x2c, 0x01, 0xeb, 0x62, 0x0a, 0x02, 0xb3, 0x95, 0xec, 0xb8, 0x38, 0x60, 0xda, 0x84, 0x8c,
	0xb8, 0x38, 0x8b, 0xa0, 0xec, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x11, 0xbd, 0xc4,
	0x82, 0x4c, 0x3d, 0x34, 0x83, 0x83, 0x10, 0xca, 0x8c, 0xf4, 0xb9, 0x38, 0x5c, 0xa0, 0x0e, 0x17,
	0x52, 0x86, 0x98, 0x2f, 0xc4, 0x05, 0xd6, 0x04, 0x76, 0xb8, 0x14, 0x2f, 0x8a, 0x01, 0x4a, 0x0c,
	0x49, 0x6c, 0x60, 0xef, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x92, 0x5f, 0x84, 0xf0,
	0x00, 0x00, 0x00,
}
