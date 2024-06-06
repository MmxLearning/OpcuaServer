package rpc

import (
	"context"
	opcuaProto "github.com/MmxLearning/OpcuaProto"
	"github.com/MmxLearning/OpcuaServer/internal/pkg/rdpTable"
	"github.com/MmxLearning/OpcuaServer/internal/rpc/middlewares"
	"github.com/MmxLearning/OpcuaServer/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
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

func (s *Opcua) RemoteScreen(srv opcuaProto.Opcua_RemoteScreenServer) error {
	firstMsg, err := srv.Recv()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	screenHello := firstMsg.GetHello()
	if screenHello == nil {
		return status.Error(codes.DataLoss, "screen hello required")
	}

	var listeners sync.Map
	var srvLock sync.Mutex
	unregister := rdpTable.RdpRegister(&rdpTable.Info{
		Name:      screenHello.Name,
		Desc:      screenHello.Desc,
		FrameRate: screenHello.FrameRate,
		SetStream: func(stream bool) error {
			srvLock.Lock()
			defer srvLock.Unlock()
			return srv.Send(&opcuaProto.StreamScreen{
				Stream: stream,
			})
		},
		Listener: &listeners,
	})
	defer unregister()

	for {
		msg, err := srv.Recv()
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		frame := msg.GetData()

		listeners.Range(func(_, value any) bool {
			(*value.(*func([]byte)))(frame)
			return true
		})
	}
}
