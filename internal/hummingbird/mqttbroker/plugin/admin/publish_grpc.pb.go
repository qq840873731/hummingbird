// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package admin

import (
	context "context"

	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// PublishServiceClient is the client API for PublishService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublishServiceClient interface {
	// Publish message to broker
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type publishServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPublishServiceClient(cc grpc.ClientConnInterface) PublishServiceClient {
	return &publishServiceClient{cc}
}

func (c *publishServiceClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/gmqtt.admin.api.PublishService/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublishServiceServer is the server API for PublishService service.
// All implementations must embed UnimplementedPublishServiceServer
// for forward compatibility
type PublishServiceServer interface {
	// Publish message to broker
	Publish(context.Context, *PublishRequest) (*empty.Empty, error)
	mustEmbedUnimplementedPublishServiceServer()
}

// UnimplementedPublishServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPublishServiceServer struct {
}

func (UnimplementedPublishServiceServer) Publish(context.Context, *PublishRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedPublishServiceServer) mustEmbedUnimplementedPublishServiceServer() {}

// UnsafePublishServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublishServiceServer will
// result in compilation errors.
type UnsafePublishServiceServer interface {
	mustEmbedUnimplementedPublishServiceServer()
}

func RegisterPublishServiceServer(s grpc.ServiceRegistrar, srv PublishServiceServer) {
	s.RegisterService(&_PublishService_serviceDesc, srv)
}

func _PublishService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gmqtt.admin.api.PublishService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServiceServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PublishService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gmqtt.admin.api.PublishService",
	HandlerType: (*PublishServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _PublishService_Publish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "publish.proto",
}
