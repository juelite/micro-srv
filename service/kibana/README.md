### 50061 kibana日志写入服务文档

#### 服务
- service
    `Kibana`
- method
    `Kibana.Write`
- params

| 参数     | 类型      | 是否必填 | 备注     |
| :----:     | :-----:    | :----:    | :----:  |
| tag      | string   | Y       | 日志标签 |
| info     | string   | Y       | 日志内容 |
| level    | string   | N       | 日志级别 |


- response

| 参数     | 类型 | 是否必填 | 备注   |
| :----:   | :----: | :----: | :----: |
| code | string | Y | 状态码 0成功，其他失败 |
| message | string | Y | 提示信息 |
| data | object | Y | 返回数据 |
