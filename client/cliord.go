package main

import (
	"context"
	"fmt"

	"demo/protoorder"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	reg := etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"http://192.168.33.11:8001"}

	})

	service := micro.NewService(
		micro.Name("demo.srv.say"),
		micro.Registry(reg),
	)

	service.Init()

	order := protoorder.NewOrderService("demo.srv.say", service.Client())

	res, err := order.CreateOrder(context.TODO(), &protoorder.QOrderInfo{
		GoodId: 232323,
		Price:  6,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.OrderId)
}
