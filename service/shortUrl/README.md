### 50062 长短链接转化服务文档 

#### 转短服务
- service
    `ShortUrl`
- method
    `ShortUrl.Create`
- params

| 参数     | 类型      | 是否必填 | 备注     |
| :----:     | :-----:    | :----:    | :----:  |
| url      | string   | Y       | 需转换的长链接 |
| express     | int   | N       | 预设有效期，0或不传永久有效 |


- response

| 参数     | 类型 | 是否必填 | 备注   |
| :----:   | :----: | :----: | :----: |
| code | string | Y | 状态码 0成功，其他失败 |
| message | string | Y | 提示信息 |
| data | object | Y | 返回数据 |

data:

| 参数     | 类型 | 是否必填 | 备注   |
| :----:   | :----: | :----: | :----: |
| url | string | Y | 转换后的短链接 |


#### 转长服务
- service
    `ShortUrl`
- method
    `ShortUrl.Query`
- params

| 参数     | 类型      | 是否必填 | 备注     |
| :----:     | :-----:    | :----:    | :----:  |
| url      | string   | Y       | 需转换的短链接 |


- response

| 参数     | 类型 | 是否必填 | 备注   |
| :----:   | :----: | :----: | :----: |
| code | string | Y | 状态码 0成功，其他失败 |
| message | string | Y | 提示信息 |
| data | object | Y | 返回数据 |

data:

| 参数     | 类型 | 是否必填 | 备注   |
| :----:   | :----: | :----: | :----: |
| url | string | Y | 转换后的长链接 |
