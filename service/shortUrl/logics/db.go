package logics

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"micro-srv/common"
	"fmt"
	"os"
)

func init()  {
	time.LoadLocation("Local")
	connect_db()
}

func connect_db()  {
	comm := &common.Common{}
	host := comm.GetDb("MFK_DB_HOST")
	name := comm.GetDb("MFK_DB_DATABASE")
	port := comm.GetDb("MFK_DB_PORT")
	user := comm.GetDb("MFK_DB_USERNAME")
	pwd := comm.GetDb("MFK_DB_PASSWORD")
	var debug bool
	if comm.GetEnv() != "prod" {
		debug = true
	}
	connectDbFunc(host, name, port, user, pwd, "default", debug)
}

/**
 * 初始化数据库连接
 * @param database_host string 链接地址
 * @param database_name string 数据库名称
 * @param database_port string 数据库端口
 * @param database_user string 数据库用户
 * @param database_pwd string 数据库密码
 * @param debug bool true 开始调试 false 关闭调试
 * @param conn_name string 链接名
 */
func connectDbFunc(database_host, database_name, database_port string , database_user, database_pwd , conn_name string, debug bool) {
	//选择模式
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.Debug = debug
	orm.DefaultTimeLoc = time.Local
	conn := database_user + ":" + database_pwd + "@tcp(" + database_host + ":" + database_port + ")/" + database_name + "?charset=utf8&parseTime=true&loc=Local"
	//注册数据库连接
	err := orm.RegisterDataBase(conn_name, "mysql", conn)
	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	} else {
		fmt.Printf("db conn succ！%s\n", database_name + "@" + database_host)
	}
}

