// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BallingServiceClient is the client API for BallingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BallingServiceClient interface {
	CalculateScore(ctx context.Context, in *CalculateScoreRequest, opts ...grpc.CallOption) (*CalculateScoreResponse, error)
}

type ballingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBallingServiceClient(cc grpc.ClientConnInterface) BallingServiceClient {
	return &ballingServiceClient{cc}
}

func (c *ballingServiceClient) CalculateScore(ctx context.Context, in *CalculateScoreRequest, opts ...grpc.CallOption) (*CalculateScoreResponse, error) {
	out := new(CalculateScoreResponse)
	err := c.cc.Invoke(ctx, "/BallingService/CalculateScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BallingServiceServer is the server API for BallingService service.
// All implementations must embed UnimplementedBallingServiceServer
// for forward compatibility
type BallingServiceServer interface {
	CalculateScore(context.Context, *CalculateScoreRequest) (*CalculateScoreResponse, error)
	mustEmbedUnimplementedBallingServiceServer()
}

// UnimplementedBallingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBallingServiceServer struct {
}

func (UnimplementedBallingServiceServer) CalculateScore(context.Context, *CalculateScoreRequest) (*CalculateScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateScore not implemented")
}
func (UnimplementedBallingServiceServer) mustEmbedUnimplementedBallingServiceServer() {}

// UnsafeBallingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BallingServiceServer will
// result in compilation errors.
type UnsafeBallingServiceServer interface {
	mustEmbedUnimplementedBallingServiceServer()
}

func RegisterBallingServiceServer(s grpc.ServiceRegistrar, srv BallingServiceServer) {
	s.RegisterService(&BallingService_ServiceDesc, srv)
}

func _BallingService_CalculateScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BallingServiceServer).CalculateScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BallingService/CalculateScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BallingServiceServer).CalculateScore(ctx, req.(*CalculateScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BallingService_ServiceDesc is the grpc.ServiceDesc for BallingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BallingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BallingService",
	HandlerType: (*BallingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculateScore",
			Handler:    _BallingService_CalculateScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ballingService.proto",
}
