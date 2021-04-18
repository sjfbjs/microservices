package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	//"github.com/micro/go-plugins/registry/consul/v2"
	"hello-client/proto/hello"
)

func main() {
	reg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))

	service := micro.NewService(micro.Registry(reg), micro.Name("sjfbjs.com.api.hello-client"))
	service.Init()
	// Use the generated client stub
	cl := hello.NewHelloService("sjfbjs.com.api.hello", service.Client())

	// Make request
	rsp, err := cl.Call(context.Background(), &hello.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Msg)
}
