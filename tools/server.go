package tools

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func MustTcpListen(addr string) net.Listener {
	tcpListen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("tcp listen addr %s failed: %v", addr, err)
	}
	return tcpListen
}

func RunHttpSrv(srv *http.Server) {
	err := srv.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return
		}
		log.Fatalln("run api server failed:", err)
	}
}

func RunGrpcSrv(tcpListen net.Listener, srv *grpc.Server) {
	err := srv.Serve(tcpListen)
	if err != nil {
		if errors.Is(err, grpc.ErrServerStopped) {
			return
		}
		log.Fatalln("run rpc server failed:", err)
	}
}
