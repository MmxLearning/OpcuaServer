package rpc

import (
	opcuaProto "github.com/MmxLearning/OpcuaProto"
	"github.com/MmxLearning/OpcuaServer/internal/rpc/middlewares"
	"google.golang.org/grpc"
)

func NewOpcua() *grpc.Server {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middlewares.UnaryLogger(),
			middlewares.UnaryAuth(),
		),
		grpc.ChainStreamInterceptor(
			middlewares.StreamLogger(),
			middlewares.StreamAuth(),
		),
	)
	opcuaProto.RegisterOpcuaServer(grpcServer, &Opcua{})
	return grpcServer
}

type Opcua struct {
	opcuaProto.UnimplementedOpcuaServer
}
