## 微服务相关文档

### 结构

```
    ├── README.md               
    ├── common                  公共包
    │   ├── common.go           实现一些基础方法，如配置文件读取，获取redis句柄等
    │   └── register.go         服务注册封装
    ├── conf                    配置文件目录
    │   ├── config.conf         项目配置文件，如consul地址等
    │   ├── db.conf             数据库配置文件
    │   └── env.conf            运行环境
    ├── constant                常量包
    │   └── status_code.go      定义所有返回码
    ├── consul                  consul api封装
    │   ├── register.go         
    │   ├── resolver.go
    │   └── watcher.go
    └── service                 服务模块集合
        ├── cache               缓存服务
        │   ├── README.md
        │   ├── logics          业务逻辑包
        │   │   └── redis.go    具体业务
        │   ├── main.go         服务入口
        │   ├── proto           proto文件
        │   │   ├── cache.pb.go 执行protoc.sh生成的文件
        │   │   ├── cache.proto 定义服务
        │   │   └── protoc.sh   编译proto文件
        │   └── test            测试包
        │       └── client.go   用于测试本服务的客户端
        ├── web                 服务面板
        │   ├── README.md
        │   ├── logics          业务逻辑包
        │   │   ├── base.go     具体业务
        │   │   └── db.go       具体业务
        │   ├── main.go         面板主程序入口
        │   ├── db              文件数据库
        │   │   └── 2018_11_02  数据写入日期目录
        │   └── models          数据模型
        ·
        ·
        ·
```


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

### 详细文档

>> (GO-Grpc微服务开发一 概览)[https://blog.csdn.net/weixin_43183475/article/details/83856650]

>> (GO-Grpc微服务开发二 服务编写)[https://blog.csdn.net/weixin_43183475/article/details/83856312]

>> (GO-Grpc微服务开发三 服务调用for golang)[https://blog.csdn.net/weixin_43183475/article/details/83856367]

>> (GO-Grpc微服务开发四 服务调用for php)[https://blog.csdn.net/weixin_43183475/article/details/83856419]

>> (GO-Grpc微服务开发五 服务调用优化)[https://blog.csdn.net/weixin_43183475/article/details/83856470]

>> (GO-Grpc微服务开发六 网关和http调用)[https://blog.csdn.net/weixin_43183475/article/details/84134928]


