# Gmsii - 消息机器人

- 企业微信bot


### Build


linux
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build --ldflags "-extldflags -static" -o ./bin/gmsii_linux_amd64 ./main.go
```

win
```bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build --ldflags "-extldflags -static" -o ./bin/gmsii_windows_amd64.exe ./main.go
```

### Run

cli
```bash
nohub gmsii server --bind=0.0.0.0 --port=8000 -c config.json &
```

docker

参考`Dockerfile`

```bash
docker-compose up -d
```


### Route

### 接口：发送文本消息
* 地址：api/v1/send/text
* 类型：POST
* 状态码：200
* 请求接口格式：

```
└─ content: String (文本内容)

```

* 返回接口格式：

```
├─ code: Number 
├─ data: Object : Object 
└─ message: String 

```


### 接口：发送告警消息
* 地址：api/v1/send/error
* 类型：POST
* 状态码：200
* 请求接口格式：

```
├─ title: String (项目名称)
├─ date_time: String (异常时间)
├─ err_code: String (错误码)
├─ err_message: String (错误信息)
├─ err_clazz: String (异常类)
└─ err_trace: String (异常调用栈)

```

* 返回接口格式：

```
├─ code: Number 
├─ data: Object : Object 
└─ message: String 

```


end
