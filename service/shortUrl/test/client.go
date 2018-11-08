package main

import (
	consulapi "github.com/hashicorp/consul/api"
	"fmt"
	"os"
	"google.golang.org/grpc"
	pb "micro-srv/service/shortUrl/proto"
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

	services1, _, err := client.Catalog().Service("shortUrl", "", nil)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	sendData1(services1)
}


func sendData1(service []*consulapi.CatalogService) {
	if len(service) > 0 {
		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", service[0].ServiceAddress, service[0].ServicePort), grpc.WithInsecure())

		c := pb.NewShortUrlClient(conn)

		//转短
		request := &pb.CreateRequest{
			Url: "http://127.0.0.1:8500/ui/#/dc1/services/kibana",
		}
		r, err := c.Create(context.Background(), request)
		fmt.Println(r, err)

		//转长
		request1 := &pb.QueryRequest{
			Url: "http://mf1.mobi/P7Cw",
		}
		r1, err := c.Query(context.Background(), request1)
		fmt.Println(r1, err)
	}
}