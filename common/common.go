package common

import (
	"os"
	"time"
	"log"
	"io"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"github.com/Unknwon/goconfig"
)

var cache = map[string]string{}

type Common struct {}

//相对路径是根据可执行文件当前位置
const (
	confPath  =  "../../conf/config.conf"
	dbPath  =  "../../conf/db.conf"
	envPath  =  "../../conf/env.conf"
)


/**
 * 获取环境变量
 */
func (t *Common) GetEnv() string {
	if val, ok := cache["env"]; ok {
		return val
	}
	env := t.getEnv()
	cache["env"] = env
	return env
}
//读配置文件获取env
func (t *Common) getEnv() string {
	c , err := goconfig.LoadConfigFile(envPath)
	if err != nil {
		log.Fatal("load config file error: ", err)
		return ""
	}
	//先获取运行环境
	env , err := c.GetValue("default" , "env")
	if err != nil {
		log.Fatal("load env file error: ", err)
		return ""
	}
	return env
}

/**
 * 根据键获取值
 */
func (t *Common) GetVal(name string) string {
	env := t.GetEnv()
	key := env + "_" + name
	if val, ok := cache[key]; ok {
		return val
	}

	c , err := goconfig.LoadConfigFile(confPath)
	if err != nil {
		log.Fatal("load config file error: ", err)
	}
	val1 , err := c.GetValue(env , name)
	if err != nil {
		log.Fatal("load env file error: ", err)
	}
	cache[key] = val1
	return val1
}

/**
 * 根据键获取值
 */
func (t *Common) GetDb(name string) string {
	env := t.GetEnv()
	key := env + "_" + name
	if val, ok := cache[key]; ok {
		return val
	}

	c , err := goconfig.LoadConfigFile(dbPath)
	if err != nil {
		log.Fatal("load config file error: ", err)
	}
	val1 , err := c.GetValue(env , name)
	if err != nil {
		log.Fatal("load env file error: ", err)
	}
	cache[key] = val1
	return val1
}

/**
 * 日志写入，需先在调用的模块controller统计目录创建runtime目录
 * @param file_name string 要写入的文件
 * @param file_content string 要写入的内容
 */
func (t *Common) LogInfo(file_name string, file_content string) {
	//文件夹路径
	file, err := os.Open("runtime/" + time.Now().Format("2006-01-02"))
	if err != nil {
		os.Mkdir("runtime/"+time.Now().Format("2006-01-02"), os.ModePerm)
	}
	//确定日志路径
	file_name = "runtime/" + time.Now().Format("2006-01-02") + "/" + file_name
	//打开文件,不存在则创建
	file, err = os.OpenFile(file_name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		println(err.Error())
	}
	//确定输出格式
	trace := log.New(io.MultiWriter(file, os.Stdin), "Info:", log.Ldate|log.Ltime)
	//写入日志内容
	trace.Println(file_content)
}

/**
 * 获取redis链接
 */
func (t *Common) GetRedisClient() (redis.Conn, error) {
	var client redis.Conn
	var host , pass string
	host = t.GetDb("redishost")
	pass = t.GetDb("redispass")
	client , err :=  redis.Dial("tcp", host)
	if err != nil {
		return client, err
	}
	if pass != "" {
		_ , err = client.Do("AUTH", pass)
		if err != nil {
			return client, err
		}
	}
	return client, nil
}

/**
 * 服务调用计数
 */
func (t *Common) CallIncr(srv string, rsp_code int32) {
	var status string
	if rsp_code == 0 {
		status = "1"
	}
	client, err := t.GetRedisClient()
	defer client.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	key := "micro_srv_call_count"
	val := srv + "_" + time.Now().Format("200601021504") + "_" + status
	client.Do("LPUSH", key, val)
	return
}


