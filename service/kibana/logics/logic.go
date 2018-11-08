package logics

import (
	"encoding/json"
	"log"
	"net"
	"strings"
	"time"
	"micro-srv/common"
	"github.com/garyburd/redigo/redis"
	proto "micro-srv/service/kibana/proto"
	"micro-srv/constant"
)

/**
 * 获取redis链接句柄 写日志比较特殊  专用redis服务器 故单独写在此处
 */
func GetRedisClient() redis.Conn {
	var client redis.Conn
	//此处有删减，涉及到账号密码等信息
	return client
}

/**
 * 写入kibana缓存 生成日志
 * @param tag string 日志标签
 * @param info string 日志信息
 * @param level string 日志级别
 */
func WriteLog(tag string, info string, level string) (rsp *proto.WriteResponse) {
	var prefix string
	baseServ := &common.Common{}
	log_data := make(map[string]interface{})
	if baseServ.GetEnv() == "prod" {
		prefix = ""
	} else {
		prefix = "test_"
	}
	log_data = __getBaseRecord(level,info)
	log_data["tags"] = strings.ToLower(prefix + tag)
	log_data["type"] = strings.ToLower(prefix + tag)
	log_json , _ := json.Marshal(log_data)
	client := GetRedisClient()
	defer client.Close()
	_ , err := client.Do("LPUSH","common_api_access_log", string(log_json))

	reply := proto.WriteResponse{
		Code: 0,
		Message: constant.GetErr(0),
		Data: nil,
	}
	if err != nil {
		reply.Code = 10001
		reply.Message = constant.GetErr(10001)
		reply.Data = nil
	}
	return &reply
}

/**
 * 构造写入日志内容
 * @param level string 日志级别
 * @param info string 日志信息
 * @return map[string]interface{}
 */
func __getBaseRecord(level string,info string) map[string]interface{} {
	log_data := make(map[string]interface{})
	log_data["log_time"] = time.Now().Unix()
	log_data["level"]    = level
	log_data["server"],_ = net.InterfaceAddrs()
	log_data["msg"]      = info
	return log_data
}