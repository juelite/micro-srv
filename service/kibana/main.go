package main

import (
	proto "micro-srv/service/kibana/proto"
	"micro-srv/service/kibana/logics"
	"fmt"
	"micro-srv/common"
	"golang.org/x/net/context"
)

type Kibana struct{}

const (
	SRV_PORT = 50061
	SRV_NAME = "kibana"
)

func (k *Kibana) Write(ctx context.Context, req *proto.WriteRequest) (rsp *proto.WriteResponse, err error) {
	rsp = logics.WriteLog(req.Tag, req.Info, req.Level)
	comm := &common.Common{}
	comm.CallIncr(SRV_NAME, rsp.Code)
	return
}

func main() {
	server, listener, err := common.Register(SRV_NAME, SRV_PORT)
	if err != nil {
		fmt.Println(err.Error())
	}
	proto.RegisterKibanaServer(server, &Kibana{})
	err = server.Serve(listener)
	if err != nil {
		fmt.Println(err.Error())
	}
}