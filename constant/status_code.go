package constant

func GetErr(code int64) string {
	err := map[int64]string{
			0: "成功",
		10001: "写入日志出错",
		20001: "短连接生成失败",
		20002: "短链接解析失败",
		20003: "短链已失效",
		30001: "缓存服务器连接异常",
		30002: "缓存写入失败",
		30003: "缓存读取异常",
		30004: "没有找到键对应的值或已失效",
		30005: "缓存删除失败",
	}

	for k, v := range err {
		if k == code {
			return v
		}
	}
	return "未定义错误码"
}