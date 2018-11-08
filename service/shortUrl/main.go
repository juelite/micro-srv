package main

import (
	proto "micro-srv/service/shortUrl/proto"
	_ "micro-srv/service/shortUrl/logics"
	"micro-srv/common"
	"fmt"
	"golang.org/x/net/context"
	"micro-srv/service/shortUrl/logics"
)

type ShortUrl struct {}

func (s *ShortUrl) Create(ctx context.Context, req *proto.CreateRequest) (rsp *proto.CreateResponse, err error)  {
	rsp = logics.Short(req)
	comm := &common.Common{}
	comm.CallIncr(SRV_NAME, rsp.Code)
	return
}

func (s *ShortUrl) Query(ctx context.Context, req *proto.QueryRequest) (rsp *proto.QueryResponse, err error) {
	rsp = logics.Query(req)
	comm := &common.Common{}
	comm.CallIncr(SRV_NAME, rsp.Code)
	return
}

const (
	SRV_PORT = 50062
	SRV_NAME = "shortUrl"
)

func main() {
	server, listener, err := common.Register(SRV_NAME, SRV_PORT)
	if err != nil {
		fmt.Println(err.Error())
	}
	proto.RegisterShortUrlServer(server, &ShortUrl{})
	err = server.Serve(listener)
	if err != nil {
		fmt.Println(err.Error())
	}
}