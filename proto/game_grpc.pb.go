// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: proto/game.proto

package go_tictactoe_grpc

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

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameClient interface {
	FindGame(ctx context.Context, in *FindGameRequest, opts ...grpc.CallOption) (*FindGameResponse, error)
	JoinGame(ctx context.Context, opts ...grpc.CallOption) (Game_JoinGameClient, error)
}

type gameClient struct {
	cc grpc.ClientConnInterface
}

func NewGameClient(cc grpc.ClientConnInterface) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) FindGame(ctx context.Context, in *FindGameRequest, opts ...grpc.CallOption) (*FindGameResponse, error) {
	out := new(FindGameResponse)
	err := c.cc.Invoke(ctx, "/game.Game/FindGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) JoinGame(ctx context.Context, opts ...grpc.CallOption) (Game_JoinGameClient, error) {
	stream, err := c.cc.NewStream(ctx, &Game_ServiceDesc.Streams[0], "/game.Game/JoinGame", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameJoinGameClient{stream}
	return x, nil
}

type Game_JoinGameClient interface {
	Send(*ClientStream) error
	Recv() (*ServerStream, error)
	grpc.ClientStream
}

type gameJoinGameClient struct {
	grpc.ClientStream
}

func (x *gameJoinGameClient) Send(m *ClientStream) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gameJoinGameClient) Recv() (*ServerStream, error) {
	m := new(ServerStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GameServer is the server API for Game service.
// All implementations must embed UnimplementedGameServer
// for forward compatibility
type GameServer interface {
	FindGame(context.Context, *FindGameRequest) (*FindGameResponse, error)
	JoinGame(Game_JoinGameServer) error
	mustEmbedUnimplementedGameServer()
}

// UnimplementedGameServer must be embedded to have forward compatible implementations.
type UnimplementedGameServer struct {
}

func (UnimplementedGameServer) FindGame(context.Context, *FindGameRequest) (*FindGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindGame not implemented")
}
func (UnimplementedGameServer) JoinGame(Game_JoinGameServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinGame not implemented")
}
func (UnimplementedGameServer) mustEmbedUnimplementedGameServer() {}

// UnsafeGameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServer will
// result in compilation errors.
type UnsafeGameServer interface {
	mustEmbedUnimplementedGameServer()
}

func RegisterGameServer(s grpc.ServiceRegistrar, srv GameServer) {
	s.RegisterService(&Game_ServiceDesc, srv)
}

func _Game_FindGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).FindGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.Game/FindGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).FindGame(ctx, req.(*FindGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_JoinGame_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServer).JoinGame(&gameJoinGameServer{stream})
}

type Game_JoinGameServer interface {
	Send(*ServerStream) error
	Recv() (*ClientStream, error)
	grpc.ServerStream
}

type gameJoinGameServer struct {
	grpc.ServerStream
}

func (x *gameJoinGameServer) Send(m *ServerStream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gameJoinGameServer) Recv() (*ClientStream, error) {
	m := new(ClientStream)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Game_ServiceDesc is the grpc.ServiceDesc for Game service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Game_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "game.Game",
	HandlerType: (*GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindGame",
			Handler:    _Game_FindGame_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinGame",
			Handler:       _Game_JoinGame_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/game.proto",
}