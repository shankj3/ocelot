// Code generated by protoc-gen-go. DO NOT EDIT.
// source: werkerserver.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Request struct {
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Request) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type Response struct {
	OutputLine string `protobuf:"bytes,1,opt,name=outputLine" json:"outputLine,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *Response) GetOutputLine() string {
	if m != nil {
		return m.OutputLine
	}
	return ""
}

type Info struct {
	WerkerFacts  *WerkerFacts `protobuf:"bytes,1,opt,name=werkerFacts" json:"werkerFacts,omitempty"`
	ActiveHashes []string     `protobuf:"bytes,6,rep,name=activeHashes" json:"activeHashes,omitempty"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (m *Info) String() string            { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *Info) GetWerkerFacts() *WerkerFacts {
	if m != nil {
		return m.WerkerFacts
	}
	return nil
}

func (m *Info) GetActiveHashes() []string {
	if m != nil {
		return m.ActiveHashes
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "models.Request")
	proto.RegisterType((*Response)(nil), "models.Response")
	proto.RegisterType((*Info)(nil), "models.Info")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Build service

type BuildClient interface {
	BuildInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (Build_BuildInfoClient, error)
	KillHash(ctx context.Context, in *Request, opts ...grpc.CallOption) (Build_KillHashClient, error)
	GetInfo(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*Info, error)
}

type buildClient struct {
	cc *grpc.ClientConn
}

func NewBuildClient(cc *grpc.ClientConn) BuildClient {
	return &buildClient{cc}
}

func (c *buildClient) BuildInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (Build_BuildInfoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Build_serviceDesc.Streams[0], c.cc, "/models.Build/BuildInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &buildBuildInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Build_BuildInfoClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type buildBuildInfoClient struct {
	grpc.ClientStream
}

func (x *buildBuildInfoClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *buildClient) KillHash(ctx context.Context, in *Request, opts ...grpc.CallOption) (Build_KillHashClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Build_serviceDesc.Streams[1], c.cc, "/models.Build/KillHash", opts...)
	if err != nil {
		return nil, err
	}
	x := &buildKillHashClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Build_KillHashClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type buildKillHashClient struct {
	grpc.ClientStream
}

func (x *buildKillHashClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *buildClient) GetInfo(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*Info, error) {
	out := new(Info)
	err := grpc.Invoke(ctx, "/models.Build/GetInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Build service

type BuildServer interface {
	BuildInfo(*Request, Build_BuildInfoServer) error
	KillHash(*Request, Build_KillHashServer) error
	GetInfo(context.Context, *google_protobuf.Empty) (*Info, error)
}

func RegisterBuildServer(s *grpc.Server, srv BuildServer) {
	s.RegisterService(&_Build_serviceDesc, srv)
}

func _Build_BuildInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BuildServer).BuildInfo(m, &buildBuildInfoServer{stream})
}

type Build_BuildInfoServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type buildBuildInfoServer struct {
	grpc.ServerStream
}

func (x *buildBuildInfoServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _Build_KillHash_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BuildServer).KillHash(m, &buildKillHashServer{stream})
}

type Build_KillHashServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type buildKillHashServer struct {
	grpc.ServerStream
}

func (x *buildKillHashServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _Build_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Build/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildServer).GetInfo(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Build_serviceDesc = grpc.ServiceDesc{
	ServiceName: "models.Build",
	HandlerType: (*BuildServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _Build_GetInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BuildInfo",
			Handler:       _Build_BuildInfo_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "KillHash",
			Handler:       _Build_KillHash_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "werkerserver.proto",
}

func init() { proto.RegisterFile("werkerserver.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x1b, 0x8d, 0x69, 0x73, 0x1b, 0x50, 0x46, 0x90, 0x12, 0x51, 0xca, 0xac, 0x8a, 0x8b,
	0xa9, 0x46, 0x7c, 0x81, 0x82, 0x7f, 0xe8, 0x2a, 0x1b, 0xc1, 0x5d, 0xd2, 0xde, 0x36, 0xc1, 0x69,
	0x66, 0x9c, 0x9f, 0x8a, 0x8f, 0xe3, 0x9b, 0x4a, 0x66, 0x1a, 0x5a, 0x77, 0xee, 0x6e, 0x4e, 0xbe,
	0x73, 0xcf, 0x9d, 0x03, 0xe4, 0x0b, 0xd5, 0x07, 0x2a, 0x8d, 0x6a, 0x83, 0x8a, 0x49, 0x25, 0x8c,
	0x20, 0xd1, 0x5a, 0x2c, 0x90, 0xeb, 0xf4, 0x7c, 0x25, 0xc4, 0x8a, 0xe3, 0xd4, 0xa9, 0xa5, 0x5d,
	0x4e, 0x71, 0x2d, 0xcd, 0xb7, 0x87, 0xd2, 0xc4, 0x1b, 0xfd, 0x17, 0xbd, 0x80, 0x7e, 0x8e, 0x9f,
	0x16, 0xb5, 0x21, 0x04, 0xc2, 0xaa, 0xd0, 0xd5, 0x28, 0x18, 0x07, 0x93, 0x38, 0x77, 0x33, 0xbd,
	0x82, 0x41, 0x8e, 0x5a, 0x8a, 0x46, 0x23, 0xb9, 0x04, 0x10, 0xd6, 0x48, 0x6b, 0x5e, 0xeb, 0x06,
	0xb7, 0xd4, 0x9e, 0x42, 0x0b, 0x08, 0x9f, 0x9b, 0xa5, 0x20, 0x77, 0x30, 0xf4, 0x11, 0x0f, 0xc5,
	0xdc, 0x68, 0x07, 0x0e, 0xb3, 0x53, 0xe6, 0x6f, 0x63, 0x6f, 0xbb, 0x5f, 0xf9, 0x3e, 0x47, 0x28,
	0x24, 0xc5, 0xdc, 0xd4, 0x1b, 0x7c, 0x2a, 0x74, 0x85, 0x7a, 0x14, 0x8d, 0x0f, 0x27, 0x71, 0xfe,
	0x47, 0xcb, 0x7e, 0x02, 0x38, 0x9a, 0xd9, 0x9a, 0x2f, 0x48, 0x06, 0xb1, 0x1b, 0x5c, 0xe2, 0x71,
	0xb7, 0x7c, 0xfb, 0x94, 0xf4, 0x64, 0x27, 0xf8, 0xe3, 0x69, 0xef, 0x3a, 0x20, 0x37, 0x30, 0x78,
	0xa9, 0x39, 0x6f, 0x77, 0xfd, 0xdf, 0xd2, 0x7f, 0x44, 0xe3, 0x42, 0xce, 0x98, 0x6f, 0x95, 0x75,
	0xad, 0xb2, 0xfb, 0xb6, 0xd5, 0x34, 0xe9, 0x8c, 0x2d, 0x45, 0x7b, 0xb3, 0xf0, 0xfd, 0x40, 0x96,
	0x65, 0xe4, 0xa8, 0xdb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0x9d, 0xa0, 0x62, 0xa7, 0x01,
	0x00, 0x00,
}
