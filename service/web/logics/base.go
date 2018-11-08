package logics

import (
	"micro-srv/common"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"micro-srv/service/web/models"
	"strings"
	"strconv"
)

func CallIncrDb() {
	t := &common.Common{}
	client, err := t.GetRedisClient()
	defer func(client redis.Conn) {
		client.Close()
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}(client)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	key := "micro_srv_call_count"
	r := make(map[string]int64)
	for {
		res, err := client.Do("LPOP", key)
		if res == nil || err != nil {
			break
		}
		sl := fmt.Sprintf("%s", res)
		if _, ok := r[sl]; ok {
			r[sl] += 1
		} else {
			r[sl] = 1
		}
	}
	var infos []models.CallCount
	for k, v := range r {
		//cache_201811020943_1
		ss := strings.Split(k, "_")
		if len(ss) != 3 {
			continue
		}
		status, _ := strconv.Atoi(ss[2])
		tmp := &models.CallCount{
			Service: ss[0],
			Time: ss[1],
			Status: int64(status),
			Count: v,
		}
		infos = append(infos, *tmp)
	}
	if len(infos) <= 0{
		return
	}
	infos[0].AddAll(infos)
	return
}
