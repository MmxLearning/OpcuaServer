package middlewares

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"time"
)

func UnaryLogger() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		startAt := time.Now()
		remote, _ := peer.FromContext(ctx)
		remoteAddr := remote.Addr.String()

		resp, err = handler(ctx, req)

		logger := log.WithFields(log.Fields{
			"component": "rpc",
			"method":    info.FullMethod,
			"ip":        remoteAddr,
			"cost":      time.Now().Sub(startAt),
		})
		if err == nil {
			logger.Info("success")
		} else {
			logger.Errorln("error:", err)
		}
		return resp, err
	}
}

func StreamLogger() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		startAt := time.Now()
		remote, _ := peer.FromContext(ss.Context())
		remoteAddr := remote.Addr.String()
		logger := log.WithFields(log.Fields{
			"component": "rpc",
			"method":    info.FullMethod,
			"ip":        remoteAddr,
		})
		log.Infof("STREAM start")

		err := handler(srv, ss)

		logger = logger.WithField("cost", time.Now().Sub(startAt))
		if err != nil {
			logger.Warnln("STREAM end with error:", err)
		} else {
			logger.Infof("STREAM complete")
		}
		return err
	}
}
