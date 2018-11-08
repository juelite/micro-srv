package common

import (
	"os"
	"time"
	"encoding/json"
	"io/ioutil"
)

type FileDb struct {

}


/**
 * 文件数据
 * @param db_name string 要写入的文件名
 * @param data string 要写入的内容
 */
func (f *FileDb) Insert(db_name string, data interface{}) {
	//文件夹路径
	file, err := os.Open("db" + "/" + time.Now().Format("2006_01_02"))
	if err != nil {
		os.Mkdir("db" + "/" + time.Now().Format("2006_01_02"), os.ModePerm)
	}
	//确定日志路径
	db_path := "db/" + time.Now().Format("2006_01_02") + "/" + db_name
	//打开文件,不存在则创建
	file, err = os.OpenFile(db_path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		println(err.Error())
	}
	bt, _ := json.Marshal(data)
	file.WriteString(string(bt))
}

/**
 * 文件数据
 * @param db_name string 要写入的文件名
 * @param data string 要写入的内容
 */
func (f *FileDb) Update(db_name string, data interface{}) {
	//文件夹路径
	file, err := os.Open("db" + "/" + time.Now().Format("2006_01_02"))
	if err != nil {
		os.Mkdir("db" + "/" + time.Now().Format("2006_01_02"), os.ModePerm)
	}
	//确定日志路径
	db_path := "db/" + time.Now().Format("2006_01_02") + "/" + db_name
	//打开文件,不存在则创建
	file, err = os.OpenFile(db_path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		println(err.Error())
	}
	bt, _ := json.Marshal(data)
	file.WriteString(string(bt))
}

/**
 * 文件数据
 * @param db_name string 要写入的文件名
 * @param date string 如 2018_11_02
 */
func (f *FileDb) Read(db_name string, date string) []byte {
	//确定日志路径
	db_path := "db/" + date + "/" + db_name
	//打开文件,不存在则创建
	file, err := os.OpenFile(db_path, os.O_RDONLY, 0666)
	if err != nil {
		println(err.Error())
		return nil
	}
	defer file.Close()
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		println(err.Error())
		return nil
	}
	return contents
}


