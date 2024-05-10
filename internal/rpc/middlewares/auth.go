package middlewares

import (
	"context"
	"github.com/MmxLearning/OpcuaServer/internal/global"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func verifyToken(ctx context.Context) bool {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false
	}
	tokenClient := md.Get("Authorization")
	if len(tokenClient) == 0 {
		return false
	}
	return tokenClient[0] == global.Config.ClientToken
}

func UnaryAuth() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if !verifyToken(ctx) {
			return nil, status.Error(codes.Unauthenticated, "verify client token failed")
		}
		return handler(ctx, req)
	}
}

func StreamAuth() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if !verifyToken(ss.Context()) {
			return status.Error(codes.Unauthenticated, "verify client token failed")
		}
		return handler(srv, ss)
	}
}
