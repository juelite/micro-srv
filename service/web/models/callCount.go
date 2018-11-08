package models

import (
	"micro-srv/common"
	"encoding/json"
	"time"
	"fmt"
	"strings"
)

type CallCount struct {
	Service 			string			`json:"service"`		//服务名称
	Time 				string 			`json:"time"`
	Count 				int64			`json:"count"`			//调用次数
	Status 				int64			`json:"status"`			//调用状态 1成功 0失败
}

const (
	CALL_STATUS_SUCCESS = 1		//调用成功常量
	CALL_STATUS_FAIL = 0		//调用失败常量

	DB_NAME = "call_count"
)



func (c *CallCount) AddAll(infos []CallCount) {
	date := time.Now().Format("2006_01_02")
	info_old := c.ReadData(date)
	new_info := make(map[string]CallCount)
	if len(info_old) > 0 {
		for _, b := range infos {
			key := b.Service+"_"+b.Time+"_"+fmt.Sprint(b.Status)
			if _, ok := info_old[key]; ok {
				old := info_old[key]
				old.Count += b.Count
				new_info[key] = old
			} else {
				new_info[key] = b
			}
		}
		for k, v := range info_old {
			if _, ok := new_info[k]; !ok {
				new_info[k] = v
			}
		}
	} else {
		for _, b := range infos {
			key := b.Service+"_"+b.Time+"_"+fmt.Sprint(b.Status)
			new_info[key] = b
		}
	}
	fd := &common.FileDb{}
	fd.Insert(DB_NAME, new_info)
}

func (c *CallCount) ReadData(date string) (infos map[string]CallCount) {
	fd := &common.FileDb{}
	res := fd.Read(DB_NAME, date)
	json.Unmarshal(res, &infos)
	return
}

func (c *CallCount) ReadServiceData(date, name string) (Chart) {
	fd := &common.FileDb{}
	res := fd.Read(DB_NAME, date)
	var infos map[string]CallCount
	json.Unmarshal(res, &infos)
	//
	// [
	//		"x" => "["16:08","16:09","16:10","16:11","16:12","16:13","16:14","16:15"...]"
	//		"y" => [
	//			"s" => "[12,4356,78,96,64,78,57,89...]"
	//			"f" => "[122,436,718,926,634,784,557,869...]"
	// 		]
	// ]

	var (
		x, s, f string
		xs []string
	)

	now := time.Now().Add(-5 * time.Minute)

	for i := -30; i < 0; i++ {
		t := now.Add(time.Duration(i) * time.Minute).Format("15:04")
		x += `"` + t + `",`
		xs = append(xs, now.Add(time.Duration(i) * time.Minute).Format("200601021504"))
	}

	ss := map[string]int64{}
	fs := map[string]int64{}

	if name != "" {
		for _, v := range xs {
			ss[v] = 0
			fs[v] = 0
			k0 := name + "_" + v + "_0"
			k1 := name + "_" + v + "_1"
			for a, b := range infos {
				if k0 == a {
					fs[v] += b.Count
				}
				if k1 == a {
					ss[v] += b.Count
				}
			}
		}
	} else {
		for _, v := range xs {
			ss[v] = 0
			fs[v] = 0
			for a, b := range infos {
				ar := strings.Split(a, "_")
				if ar[1] == v && ar[2] == "0" {
					fs[v] += b.Count
				}

				if ar[1] == v && ar[2] == "1" {
					ss[v] += b.Count
				}
			}
		}
	}

	for _, v := range ss {
		s += fmt.Sprint(v) + ","
	}
	for _, v := range fs {
		f += fmt.Sprint(v) + ","
	}
	x = strings.TrimRight(x, ",")
	s = strings.TrimRight(s, ",")
	f = strings.TrimRight(f, ",")
	y := make(map[string]string)
	y["s"] = s
	y["f"] = f
	d := &Chart{}
	d.X = x
	d.Y = y
	return *d
}

type Chart struct {
	X 		string 					`json:"x"`
	Y 		map[string]string 		`json:"x"`
}