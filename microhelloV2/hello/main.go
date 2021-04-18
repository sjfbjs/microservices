package main

import (
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/transport/grpc"
	//"github.com/micro/go-plugins/registry/consul/v2"
	"hello/handler"
	hello "hello/proto/hello"
	"hello/subscriber"
)

func main() {
	reg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))

	service := micro.NewService(
		micro.Name("sjfbjs.com.api.hello"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "hello",
		}),
		micro.Registry(reg),
		micro.Transport(grpc.NewTransport()),

		// Setup some flags. Specify --run_client to run the client

		// Add runtime flags
		// We could do this below too
		micro.Flags(&cli.BoolFlag{
			Name:  "run_client",
			Usage: "Launch the client",
		}),
	)

	// Init will parse the command line flags. Any flags set will
	// override the above settings. Options defined here will
	// override anything set on the command line.
	service.Init(
	// Add runtime action
	// We could actually do this above
	//micro.Action(func(c *cli.Context) error {
	//	if c.Bool("run_client") {
	//		runClient(service)
	//		os.Exit(0)
	//	}
	//	return nil
	//}),
	)

	// By default we'll run the server unless the flags catch us

	// Setup the server

	// Register handler
	_ = hello.RegisterHelloHandler(service.Server(), new(handler.Hello))
	_ = micro.RegisterSubscriber("sjfbjs.com.api.hello", service.Server(), new(subscriber.Hello))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
