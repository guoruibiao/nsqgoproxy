# nsqgoproxy
a proxy for publish and consume data with nsq&amp;golang


## 配置下环境
```
1 拉镜像
docker pull nsqio/nsq

2 运行nsqlookupd
docker run -d --name lookupd -p 4160:4160 -p 4161:4161 nsqio/nsq /nsqlookupd
通过docker ps -a 找到nsqlookupd的containerid， 再通过docker inspect 找到对应的IP-Address

3 运行nsqd
docker run -d --name=nsqd -p 4150:4150 -p 4151:4151 nsqio/nsq /nsqd --broadcast-address=172.17.0.2 --lookupd-tcp-address=172.17.0.2:4160

--broadcast-address指的是本机IP，其会被注册到lookupd上被发现。

4 运行nsqadmin
docker run -d --name=nsqadmin -p 4171:4171 nsqio/nsq /nsqadmin --lookupd-http-address=172.17.0.2:4161
```

## 安装依赖
```
go get github.com/bitly/go-nsq
go get github.com/pkg/errors
go get github.com/guoruibiao/gorequests
```

## 简单测试
```
go test -v ./...
```

## 实际应用
```
// run-nsq-go-proxy.go
package main

import (
    "github.com/guoruibiao/nsqgoproxy"
)

func main() {
    // serve
    nsqgoproxy.Serve()
}
// 然后在命令行里跑起来
go run run-nsq-go-proxy.go
```

## 部署
拷贝一份Nginx配置，如下：

```
server {

    server_name nsqgoproxy.yourdomain.com;
    listen 80;
    access_log /var/log/nginx/nsqgoproxy.access.log;
    error_log /var/log/nginx/nsqgoproxy.error.log;

    location / {
        proxy_pass http://localhost:9999;
    }
}
```
弄完之后重启下Nginx服务，让新的配置生效。这样就可以如下去访问：
```
➜  ~ curl http://nsqgoproxy.yourdomain.com/

    请求格式：
        1. 首页
        http://nsqgoproxy.yourdomain.com/

        2. 添加消费事件
        http://nsqgoproxy.yourdomain.com/publish?classname=DemoService&methodname=say&name=tiger

➜  ~
```