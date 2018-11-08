package main

import (
	"micro-srv/common"
	"net/http"
	"io"
	"os"
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"log"
	"github.com/pkg/errors"
	"time"
	_ "micro-srv/service/web/logics"
	"micro-srv/service/web/logics"
	"micro-srv/service/web/models"
)

const (
	WEB_PORT = "5060"
	CONSUL_PORT = 8500
)

func main() {
	http.HandleFunc("/web", WebServer)
	http.HandleFunc("/web/deregister", DeRegisterServer)
	http.HandleFunc("/web/stat", StatServer)
	comm := &common.Common{}
	ip := comm.GetVal("serveraddr")
	fmt.Println("web is running on pid: " + fmt.Sprint(os.Getppid()))
	fmt.Println("web listen port is: " + WEB_PORT)
	err := http.ListenAndServe(ip + ":" + WEB_PORT, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func WebServer(w http.ResponseWriter, req *http.Request) {
	base := &common.Common{}
	logics.CallIncrDb()
	consul_addr := base.GetVal("consuladdr")
	consul_host := consul_addr + ":" + fmt.Sprint(CONSUL_PORT)
	consul_token := base.GetVal("consultoken")

	conf := &consul.Config{Scheme: "http", Address: consul_host, Token:consul_token}
	client, err := consul.NewClient(conf)
	if err != nil {
		io.WriteString(w,errorReturn(err))
		return
	}
	services, _, err := client.Catalog().Services(nil)
	io.WriteString(w, webView(services))
}

func DeRegisterServer(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	p := req.Form
	if len(p["id"]) < 1 {
		io.WriteString(w,errorReturn(errors.New("参数错误")))
		return
	}
	id := p["id"][0]
	if id == "" {
		io.WriteString(w,errorReturn(errors.New("参数错误")))
		return
	}
	//TODO de_register logic
	http.Redirect(w, req, "/web", 302)
}

func StatServer(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	p := req.Form
	var name string
	if len(p["name"]) == 1 {
		name = p["name"][0]
	}
	m := &models.CallCount{}
	date := time.Now().Format("2006_01_02")
	data := m.ReadServiceData(date, name)
	fmt.Println(data)
	if name == "" {
		name = "总览"
	}
	io.WriteString(w, statView(data, name))
}


func errorReturn(err error) string {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>服务面板</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

    <!-- 可选的 Bootstrap 主题文件（一般不用引入） -->
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">

    <!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
</head>
<body>
<div class="container-fluid" style="background-color:#f0f0f0;">
    <div class="row" style="padding: 10px 0;">
        <div class="col-md-9" style="font-size: 16px; font-weight: 700;">
            <a href="/web">微服务面板</a>
        </div>
        <div class="col-md-3">
            <a href="/web/stat">访问统计</a>
        </div>
    </div>
</div>
<div class="container" style="margin-top:30px;">
    <div style="width:500px;margin:auto;padding:10px;border: solid #f0f0f0">
		<p>ERROR: `+err.Error()+`</p>
		<p>TIME: `+fmt.Sprint(time.Now())+`</p>
		<p><a href="/web">GO HOME</a></p>
	</div>
</div>
</body>
</html>
`
	return html
}


func webView(services map[string][]string) string {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>服务面板</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

    <!-- 可选的 Bootstrap 主题文件（一般不用引入） -->
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">

    <!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
</head>
<body>
<div class="container-fluid" style="background-color:#f0f0f0;">
    <div class="row" style="padding: 10px 0;">
        <div class="col-md-9" style="font-size: 16px; font-weight: 700;">
            <a href="/web">微服务面板</a>
        </div>
        <div class="col-md-3">
            <a href="/web/stat">访问统计</a>
        </div>
    </div>
</div>
<div class="container" style="margin-top:30px;">
    <table class="table table-bordered">
        <tr>
            <td>名称</td>
            <td>ID</td>
            <td>操作</td>
        </tr>
`

	var tr string
	for k, v := range services {
		if len(v) > 0 {
			for _, b := range v {
				tr += `<tr>
<td>`+k+`</td>
<td>`+b+`</td>
<td><a href="/web/deregister?id=`+b+`">注销</a> | <a href="/web/stat?name=`+k+`">统计</a></td>
</tr>`
			}
		}

	}
	html += tr
	html += `</table>
</div>
</body>
</html>
`
	return html
}

func statView(data models.Chart, name string) string {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta http-equiv="refresh" content="5">
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>服务面板</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

    <!-- 可选的 Bootstrap 主题文件（一般不用引入） -->
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">

    <script src="https://cdn.bootcss.com/echarts/2.2.7/echarts-all.js"></script>
</head>
<body>
<div class="container-fluid" style="background-color:#f0f0f0;">
    <div class="row" style="padding: 10px 0;">
        <div class="col-md-9" style="font-size: 16px; font-weight: 700;">
            <a href="/web">微服务面板</a>
        </div>
        <div class="col-md-3">
            <a href="/web/stat">访问统计</a>
        </div>
    </div>
</div>
<div class="container" style="margin-top:30px;">
    <div id="main" style="width: 80%;height:500px; margin: auto">

    </div>
</div>
</body>
<script type="text/javascript">
    // 基于准备好的dom，初始化echarts实例
    var myChart = echarts.init(document.getElementById('main'));

    var colors = ['#5793f3', ];


	option = {
		title: {
            text: '`+name+`'
        },
		xAxis: {
			type: 'category',
			data: [`+data.X+`]
		},
		yAxis: {
			type: 'value'
		},
		tooltip: {
			trigger: 'axis'
		},
		legend: {
			data:['成功','失败']
		},
		series: [
			{
                name:'成功',
				itemStyle:{
					normal:{
						color:'#337ab7'
					}
				},
                type:'line',
                stack: '成功',
				smooth: true,
                data:[`+data.Y["s"]+`]
            },
            {
                name:'失败',
				itemStyle:{
					normal:{
						color:'#b77f33'
					}
				},
                type:'line',
                stack: '失败',
				smooth: true,
                data:[`+data.Y["f"]+`]
            }
		]
	};
    // 使用刚指定的配置项和数据显示图表。
    myChart.setOption(option);


</script>
</html>
`
	return html
}