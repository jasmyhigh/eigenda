// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: node/v2/node_v2.proto

package v2

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

const (
	Dispersal_StoreChunks_FullMethodName = "/node.v2.Dispersal/StoreChunks"
	Dispersal_NodeInfo_FullMethodName    = "/node.v2.Dispersal/NodeInfo"
)

// DispersalClient is the client API for Dispersal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DispersalClient interface {
	StoreChunks(ctx context.Context, in *StoreChunksRequest, opts ...grpc.CallOption) (*StoreChunksReply, error)
	NodeInfo(ctx context.Context, in *NodeInfoRequest, opts ...grpc.CallOption) (*NodeInfoReply, error)
}

type dispersalClient struct {
	cc grpc.ClientConnInterface
}

func NewDispersalClient(cc grpc.ClientConnInterface) DispersalClient {
	return &dispersalClient{cc}
}

func (c *dispersalClient) StoreChunks(ctx context.Context, in *StoreChunksRequest, opts ...grpc.CallOption) (*StoreChunksReply, error) {
	out := new(StoreChunksReply)
	err := c.cc.Invoke(ctx, Dispersal_StoreChunks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispersalClient) NodeInfo(ctx context.Context, in *NodeInfoRequest, opts ...grpc.CallOption) (*NodeInfoReply, error) {
	out := new(NodeInfoReply)
	err := c.cc.Invoke(ctx, Dispersal_NodeInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispersalServer is the server API for Dispersal service.
// All implementations must embed UnimplementedDispersalServer
// for forward compatibility
type DispersalServer interface {
	StoreChunks(context.Context, *StoreChunksRequest) (*StoreChunksReply, error)
	NodeInfo(context.Context, *NodeInfoRequest) (*NodeInfoReply, error)
	mustEmbedUnimplementedDispersalServer()
}

// UnimplementedDispersalServer must be embedded to have forward compatible implementations.
type UnimplementedDispersalServer struct {
}

func (UnimplementedDispersalServer) StoreChunks(context.Context, *StoreChunksRequest) (*StoreChunksReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreChunks not implemented")
}
func (UnimplementedDispersalServer) NodeInfo(context.Context, *NodeInfoRequest) (*NodeInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeInfo not implemented")
}
func (UnimplementedDispersalServer) mustEmbedUnimplementedDispersalServer() {}

// UnsafeDispersalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DispersalServer will
// result in compilation errors.
type UnsafeDispersalServer interface {
	mustEmbedUnimplementedDispersalServer()
}

func RegisterDispersalServer(s grpc.ServiceRegistrar, srv DispersalServer) {
	s.RegisterService(&Dispersal_ServiceDesc, srv)
}

func _Dispersal_StoreChunks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreChunksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispersalServer).StoreChunks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dispersal_StoreChunks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispersalServer).StoreChunks(ctx, req.(*StoreChunksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dispersal_NodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispersalServer).NodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dispersal_NodeInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispersalServer).NodeInfo(ctx, req.(*NodeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dispersal_ServiceDesc is the grpc.ServiceDesc for Dispersal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dispersal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "node.v2.Dispersal",
	HandlerType: (*DispersalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreChunks",
			Handler:    _Dispersal_StoreChunks_Handler,
		},
		{
			MethodName: "NodeInfo",
			Handler:    _Dispersal_NodeInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node/v2/node_v2.proto",
}

const (
	Retrieval_GetChunks_FullMethodName = "/node.v2.Retrieval/GetChunks"
	Retrieval_NodeInfo_FullMethodName  = "/node.v2.Retrieval/NodeInfo"
)

// RetrievalClient is the client API for Retrieval service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RetrievalClient interface {
	// GetChunks retrieves the chunks for a blob custodied at the Node.
	GetChunks(ctx context.Context, in *GetChunksRequest, opts ...grpc.CallOption) (*GetChunksReply, error)
	// Retrieve node info metadata
	NodeInfo(ctx context.Context, in *NodeInfoRequest, opts ...grpc.CallOption) (*NodeInfoReply, error)
}

type retrievalClient struct {
	cc grpc.ClientConnInterface
}

func NewRetrievalClient(cc grpc.ClientConnInterface) RetrievalClient {
	return &retrievalClient{cc}
}

func (c *retrievalClient) GetChunks(ctx context.Context, in *GetChunksRequest, opts ...grpc.CallOption) (*GetChunksReply, error) {
	out := new(GetChunksReply)
	err := c.cc.Invoke(ctx, Retrieval_GetChunks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retrievalClient) NodeInfo(ctx context.Context, in *NodeInfoRequest, opts ...grpc.CallOption) (*NodeInfoReply, error) {
	out := new(NodeInfoReply)
	err := c.cc.Invoke(ctx, Retrieval_NodeInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RetrievalServer is the server API for Retrieval service.
// All implementations must embed UnimplementedRetrievalServer
// for forward compatibility
type RetrievalServer interface {
	// GetChunks retrieves the chunks for a blob custodied at the Node.
	GetChunks(context.Context, *GetChunksRequest) (*GetChunksReply, error)
	// Retrieve node info metadata
	NodeInfo(context.Context, *NodeInfoRequest) (*NodeInfoReply, error)
	mustEmbedUnimplementedRetrievalServer()
}

// UnimplementedRetrievalServer must be embedded to have forward compatible implementations.
type UnimplementedRetrievalServer struct {
}

func (UnimplementedRetrievalServer) GetChunks(context.Context, *GetChunksRequest) (*GetChunksReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChunks not implemented")
}
func (UnimplementedRetrievalServer) NodeInfo(context.Context, *NodeInfoRequest) (*NodeInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeInfo not implemented")
}
func (UnimplementedRetrievalServer) mustEmbedUnimplementedRetrievalServer() {}

// UnsafeRetrievalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RetrievalServer will
// result in compilation errors.
type UnsafeRetrievalServer interface {
	mustEmbedUnimplementedRetrievalServer()
}

func RegisterRetrievalServer(s grpc.ServiceRegistrar, srv RetrievalServer) {
	s.RegisterService(&Retrieval_ServiceDesc, srv)
}

func _Retrieval_GetChunks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChunksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrievalServer).GetChunks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Retrieval_GetChunks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrievalServer).GetChunks(ctx, req.(*GetChunksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Retrieval_NodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrievalServer).NodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Retrieval_NodeInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrievalServer).NodeInfo(ctx, req.(*NodeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Retrieval_ServiceDesc is the grpc.ServiceDesc for Retrieval service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Retrieval_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "node.v2.Retrieval",
	HandlerType: (*RetrievalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChunks",
			Handler:    _Retrieval_GetChunks_Handler,
		},
		{
			MethodName: "NodeInfo",
			Handler:    _Retrieval_NodeInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node/v2/node_v2.proto",
}
