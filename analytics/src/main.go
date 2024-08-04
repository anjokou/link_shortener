package main

import (
	"fmt"
	"link_shortener_analytics/actions"
	"link_shortener_analytics/metrics"
	"link_shortener_analytics/routers"
	appRpc "link_shortener_analytics/rpc"
	"link_shortener_analytics/settings"
	"net"
	"net/rpc"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	metricsDA := metrics.NewInMemoryMetricsDA()

	actions.InitMetricActions(metricsDA)

	server := gin.Default()
	routers.ApplyMetricsRoutes(server)

	go startRpcServer()
	server.Run(":" + strconv.Itoa(settings.HttpListeningPort()))
}

func startRpcServer() {
	metricsActions := new(appRpc.MetricsRpcActions)

	err := rpc.Register(metricsActions)
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(settings.RpcListeningPort()))
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Printf("RPC server listening on port %d\n", settings.RpcListeningPort())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Serve the connection
		go rpc.ServeConn(conn)
	}
}
