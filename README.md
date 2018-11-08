## 微服务相关文档

### 一.微服务编写

#### 1.定义proto文件

```$xslt

syntax = "proto3";

//service name is kibana
service Kibana {
    //service method is write
    rpc Write(WriteRequest) returns (WriteResponse) {}
}

//writer request struct
message WriteRequest {
    string tag = 1;
    string info = 2;
    string level = 3;
}

//writer response struct
message WriteResponse {
    int32 code = 1;
    string message = 2;
    map<string, string> data = 3;
}


```

#### 2.将定义的proto编译为go文件
```$xslt
protoc --proto_path=.:. --micro_out=. --go_out=. kibana.proto
```

#### 3.编写服务

```$xslt
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
	rsp.Code = 0
	rsp.Message = "success"
	rsp.Data = map[string]string{"tag":req.Tag,"info":req.Info,"level":req.Level}
	err = logics.WriteLog(req.Tag, req.Info, req.Level)
	return rsp, err
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
```

### 二.微服务运行

#### 启动consul

`consul agent -dev &`

面板地址：http://127.0.0.1:8500

#### 启动并注册服务
`cd kibana && go run main.go &`

### 三.服务入参和出参

>> 详细见各服务根目录中 [README.md]

### 四.微服务调用

>> 详细见各服务根目录中 [test/client.go]



