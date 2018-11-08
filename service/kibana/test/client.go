package main

import (
	consulapi "github.com/hashicorp/consul/api"
	"fmt"
	"os"
	"google.golang.org/grpc"
	pb "micro-srv/service/kibana/proto"
	"golang.org/x/net/context"
)


func main() {
	config := consulapi.Config{
		Address:"http://127.0.0.1:8500",
	}

	client, err := consulapi.NewClient(&config)//非默认情况下需要设置实际的参数
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	services1, _, err := client.Catalog().Service("kibana", "", nil)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	sendData1(services1)
}


func sendData1(service []*consulapi.CatalogService) {
	if len(service) > 0 {
		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", service[0].ServiceAddress, service[0].ServicePort), grpc.WithInsecure())

		c := pb.NewKibanaClient(conn)

		request := &pb.WriteRequest{
			Tag: "micro_test",
			Info: "hello 5.45",
			Level: "info",
		}
		r, err := c.Write(context.Background(), request)
		fmt.Println(r, err)
	}
}