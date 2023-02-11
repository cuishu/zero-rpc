package main

import (
	"fmt"
	"net"

	"{{.Package}}/config"
	"{{.Package}}/proto"
	"{{.Package}}/server"

	"google.golang.org/grpc"
)

func main() {
	config, err := config.NewConfig("config/config.yaml")
	if err != nil {
		panic(err)
	}
	lis, err := net.Listen("tcp", config.Listen)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	proto.Register{{.Service.Name}}Server(grpcServer, server.New{{.Service.Name}}Server(config))
	fmt.Println("listen")
	grpcServer.Serve(lis)
}
