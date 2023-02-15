// Code generated by zero-rpc. DO NOT EDIT.

package main

import (
	"flag"
	"fmt"
	"os"

	"{{.Module}}/config"
	"{{.Module}}/proto"
	"{{.Module}}/server"
	"{{.Module}}/svc"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	_ "go.uber.org/automaxprocs"
	"google.golang.org/grpc"
)

var (
	BuildTime string
	Version string
)

func init() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "")
	flag.Parse()
	if showVersion {
		fmt.Println("version: ", Version)
		fmt.Println("build at:", BuildTime)
		os.Exit(0)
	}
}

func main() {
	config, err := config.NewConfig("config/config.yaml")
	if err != nil {
		panic(err)
	}
	srv, err := zrpc.NewServer(zrpc.RpcServerConf{
		ListenOn: config.Listen,
		Etcd: discov.EtcdConf{
			Hosts: config.Etcd.Hosts,
			Key: "{{.Package}}.rpc",
		},
	}, func(s *grpc.Server) {
		proto.Register{{.Service.Name}}Server(s, server.New{{.Service.Name}}Server(svc.NewSvc(config)))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("listen:", config.Listen)
	srv.Start()
}
