// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// UmurClient is the client API for Umur service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UmurClient interface {
	HitungUmur(ctx context.Context, in *RequestHitungUmur, opts ...grpc.CallOption) (*ResponseHitungUmur, error)
}

type umurClient struct {
	cc grpc.ClientConnInterface
}

func NewUmurClient(cc grpc.ClientConnInterface) UmurClient {
	return &umurClient{cc}
}

func (c *umurClient) HitungUmur(ctx context.Context, in *RequestHitungUmur, opts ...grpc.CallOption) (*ResponseHitungUmur, error) {
	out := new(ResponseHitungUmur)
	err := c.cc.Invoke(ctx, "/proto.Umur/HitungUmur", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UmurServer is the server API for Umur service.
// All implementations should embed UnimplementedUmurServer
// for forward compatibility
type UmurServer interface {
	HitungUmur(context.Context, *RequestHitungUmur) (*ResponseHitungUmur, error)
}

// UnimplementedUmurServer should be embedded to have forward compatible implementations.
type UnimplementedUmurServer struct {
}

func (UnimplementedUmurServer) HitungUmur(context.Context, *RequestHitungUmur) (*ResponseHitungUmur, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HitungUmur not implemented")
}

// UnsafeUmurServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UmurServer will
// result in compilation errors.
type UnsafeUmurServer interface {
	mustEmbedUnimplementedUmurServer()
}

func RegisterUmurServer(s *grpc.Server, srv UmurServer) {
	s.RegisterService(&_Umur_serviceDesc, srv)
}

func _Umur_HitungUmur_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestHitungUmur)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmurServer).HitungUmur(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Umur/HitungUmur",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmurServer).HitungUmur(ctx, req.(*RequestHitungUmur))
	}
	return interceptor(ctx, in, info, handler)
}

var _Umur_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Umur",
	HandlerType: (*UmurServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HitungUmur",
			Handler:    _Umur_HitungUmur_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "umur.proto",
}
