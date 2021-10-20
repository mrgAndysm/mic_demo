package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main(){
	reg := etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"http://192.168.33.11:8001"}

	})

	re,_ := reg.GetService("")

	//fmt.Println(re[0].Name)
	//fmt.Println(re[0].Metadata)
	//fmt.Println(re[0].Nodes)
	//fmt.Println(re[0].Endpoints)


	fmt.Printf("%V",re[0].Nodes)


}

