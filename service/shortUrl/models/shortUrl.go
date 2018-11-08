package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type ShortUrl struct {
	Id 					int64 		`json:"id"`
	LongUrl 			string 		`json:"long_url"`
	ShortUrl 			string 		`json:"short_url"`
	Express 			int64 		`json:"express"`
	CreateTime 			int64 		`json:"create_time"`
	UpdateTime 			int64 		`json:"update_time"`
}

func init() {
	orm.RegisterModel(new(ShortUrl))
}

func (s *ShortUrl) Add(url ShortUrl) (id int64, err error) {
	url.UpdateTime = time.Now().Unix()
	url.CreateTime = time.Now().Unix()
	o := orm.NewOrm()
	id, err = o.Insert(&url)
	return
}

func (s *ShortUrl) Update(url ShortUrl) (err error) {
	url.UpdateTime = time.Now().Unix()
	o := orm.NewOrm()
	_, err = o.Update(&url)
	return
}

func (s *ShortUrl) GetByShortUrl(url string) (short ShortUrl, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(s).Filter("short_url", url).One(&short)
	return
}