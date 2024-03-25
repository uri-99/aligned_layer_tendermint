// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: alignedlayer/verify/tx.proto

package verify

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
	Msg_UpdateParams_FullMethodName  = "/alignedlayer.verify.Msg/UpdateParams"
	Msg_GnarkPlonk_FullMethodName    = "/alignedlayer.verify.Msg/GnarkPlonk"
	Msg_CairoPlatinum_FullMethodName = "/alignedlayer.verify.Msg/CairoPlatinum"
	Msg_Sp1_FullMethodName           = "/alignedlayer.verify.Msg/Sp1"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	GnarkPlonk(ctx context.Context, in *MsgGnarkPlonk, opts ...grpc.CallOption) (*MsgGnarkPlonkResponse, error)
	CairoPlatinum(ctx context.Context, in *MsgCairoPlatinum, opts ...grpc.CallOption) (*MsgCairoPlatinumResponse, error)
	Sp1(ctx context.Context, in *MsgSp1, opts ...grpc.CallOption) (*MsgSp1Response, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GnarkPlonk(ctx context.Context, in *MsgGnarkPlonk, opts ...grpc.CallOption) (*MsgGnarkPlonkResponse, error) {
	out := new(MsgGnarkPlonkResponse)
	err := c.cc.Invoke(ctx, Msg_GnarkPlonk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CairoPlatinum(ctx context.Context, in *MsgCairoPlatinum, opts ...grpc.CallOption) (*MsgCairoPlatinumResponse, error) {
	out := new(MsgCairoPlatinumResponse)
	err := c.cc.Invoke(ctx, Msg_CairoPlatinum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Sp1(ctx context.Context, in *MsgSp1, opts ...grpc.CallOption) (*MsgSp1Response, error) {
	out := new(MsgSp1Response)
	err := c.cc.Invoke(ctx, Msg_Sp1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	GnarkPlonk(context.Context, *MsgGnarkPlonk) (*MsgGnarkPlonkResponse, error)
	CairoPlatinum(context.Context, *MsgCairoPlatinum) (*MsgCairoPlatinumResponse, error)
	Sp1(context.Context, *MsgSp1) (*MsgSp1Response, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) GnarkPlonk(context.Context, *MsgGnarkPlonk) (*MsgGnarkPlonkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GnarkPlonk not implemented")
}
func (UnimplementedMsgServer) CairoPlatinum(context.Context, *MsgCairoPlatinum) (*MsgCairoPlatinumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CairoPlatinum not implemented")
}
func (UnimplementedMsgServer) Sp1(context.Context, *MsgSp1) (*MsgSp1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sp1 not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GnarkPlonk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGnarkPlonk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GnarkPlonk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_GnarkPlonk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GnarkPlonk(ctx, req.(*MsgGnarkPlonk))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CairoPlatinum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCairoPlatinum)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CairoPlatinum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CairoPlatinum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CairoPlatinum(ctx, req.(*MsgCairoPlatinum))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Sp1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSp1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Sp1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Sp1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Sp1(ctx, req.(*MsgSp1))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "alignedlayer.verify.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "GnarkPlonk",
			Handler:    _Msg_GnarkPlonk_Handler,
		},
		{
			MethodName: "CairoPlatinum",
			Handler:    _Msg_CairoPlatinum_Handler,
		},
		{
			MethodName: "Sp1",
			Handler:    _Msg_Sp1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alignedlayer/verify/tx.proto",
}
