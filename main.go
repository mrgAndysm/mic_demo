package main

import (
	"demo/protoorder"
	"demo/service/grpc"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	println("aaa")
	reg := etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"http://192.168.33.11:2379"}

	})

	service := micro.NewService(
		micro.Name("demo.srv.say"),
		micro.Registry(reg),
	)
	service.Init()

	protoorder.RegisterOrderServiceHandler(service.Server(), new(grpc.OrderService))

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
