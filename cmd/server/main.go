package main

import (
	"context"
	"github.com/MmxLearning/OpcuaServer/internal/router"
	"github.com/MmxLearning/OpcuaServer/internal/rpc"
	"github.com/MmxLearning/OpcuaServer/tools"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Infoln("Sys Boost")

	httpSrv := &http.Server{
		Addr:    ":80",
		Handler: router.Engine(),
	}
	opcuaRpc := rpc.NewOpcua()

	go tools.RunHttpSrv(httpSrv)
	go tools.RunGrpcSrv(tools.MustTcpListen(":81"), opcuaRpc)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-quit
	log.Infoln("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	err := httpSrv.Shutdown(ctx)
	if err != nil {
		log.Errorln("Http Server Shutdown:", err)
	}

	opcuaRpc.GracefulStop()
}
