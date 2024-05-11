package rpc

import (
	"context"
	opcuaProto "github.com/MmxLearning/OpcuaProto"
	"github.com/MmxLearning/OpcuaServer/internal/rpc/middlewares"
	"github.com/MmxLearning/OpcuaServer/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"unsafe"
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

func (s *Opcua) ReportOpcua(_ context.Context, msg *opcuaProto.OpcuaMessage) (*opcuaProto.OpcuaResult, error) {
	model, err := service.Opcua.Store(
		msg.Name, msg.NodeId,
		unsafe.String(unsafe.SliceData(msg.Data), len(msg.Data)),
		msg.Timestamp,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &opcuaProto.OpcuaResult{
		Id: uint64(model.ID),
	}, nil
}
