package logics

import (
	proto "micro-srv/service/shortUrl/proto"
	"strings"
	"micro-srv/service/shortUrl/models"
	"time"
	"micro-srv/constant"
)

const (
	SHOR_URL_DOMAIN = "http://mf1.mobi/"
)

func Short(req *proto.CreateRequest) (*proto.CreateResponse) {
	var resp proto.CreateResponse
	model := &models.ShortUrl{}
	model.LongUrl = req.Url
	model.Express = req.Express
	id, err := model.Add(*model)

	if err != nil {
		resp.Code = 20001
		resp.Message = constant.GetErr(20001)
		resp.Data = map[string]string{}
		return &resp
	}
	model.ShortUrl = base10ToBase62(id)
	model.Id = id
	model.CreateTime = time.Now().Unix()
	err = model.Update(*model)
	if err != nil {
		resp.Code = 20001
		resp.Message = constant.GetErr(20001)
		resp.Data = map[string]string{}
		return &resp
	}
	resp.Code = 0
	resp.Message = constant.GetErr(0)
	resp.Data = map[string]string{"url": SHOR_URL_DOMAIN + model.ShortUrl}
	return &resp
}


func Query(req *proto.QueryRequest) (*proto.QueryResponse) {
	var resp proto.QueryResponse
	model := &models.ShortUrl{}
	url := strings.Replace(req.Url, SHOR_URL_DOMAIN, "", 1)
	short, err := model.GetByShortUrl(url)
	if err != nil {
		resp.Code = 20002
		resp.Message = constant.GetErr(20002)
		return &resp
	}
	if short.Express > 0 && time.Now().Unix() > short.CreateTime + short.Express {
		resp.Code = 20003
		resp.Message = constant.GetErr(20003)
		return &resp
	}
	resp.Code = 0
	resp.Message = constant.GetErr(0)
	resp.Data = map[string]string{"url": short.LongUrl}
	return &resp
}


func base10ToBase62(id int64) string {
	var chars = strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", "")
	tempE := []int64{}
	for id > 0 {
		tempE = append(tempE, id%62)
		id /= 62
	}
	res := ""
	for _, val := range tempE {
		res += chars[val]
	}
	runes := []rune(res)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
