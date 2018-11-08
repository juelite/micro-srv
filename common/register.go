package common

import (
	"time"
	"google.golang.org/grpc"
	"net"
	"fmt"
	"micro-srv/consul"
	"os"
)

const (
	CONSUL_PORT = 8500
)

func Register(srv_name string, srv_port int) (server *grpc.Server, listener net.Listener, err error) {
	base := Common{}
	server_host := base.GetVal("serveraddr")
	server_addr := server_host  + ":" + fmt.Sprint(srv_port)
	consul_addr := base.GetVal("consuladdr")
	consul_host := consul_addr + ":" + fmt.Sprint(CONSUL_PORT)
	consul_token := base.GetVal("consultoken")
	ttl := 15

	//注册服务到consul
	server = grpc.NewServer()
	listener , err = net.Listen("tcp", server_addr)
	if err != nil {
		fmt.Sprintf("failed to listen: %v", err)
		return
	}

	err = consul.Register(srv_name, server_host, srv_port, consul_host, time.Second * 10, ttl, consul_token)
	if err != nil {
		fmt.Sprintf("failed to listen: %v", err)
		panic(err)
		return
	}
	fmt.Println("service " + srv_name + " is running on: " + fmt.Sprint(os.Getppid()))
	return
}
