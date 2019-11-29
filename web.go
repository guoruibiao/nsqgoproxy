package nsqgoproxy

import (
	"net/http"
	"log"
	"fmt"
	"strconv"
	"os"
)

var nsqProxy *NSQProxy
func Serve() {
	proxy, err := NewNSQProxy(3)
	nsqProxy = proxy
	// 一直监听 跑起来
	go func() {
		nsqProxy.GetEvent("topicname")
	}()

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/publish", publish)

	err = http.ListenAndServe(":" + strconv.Itoa(SERVER_PORT), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


func publish(writer http.ResponseWriter, request *http.Request) {
	queryMap := request.URL.Query()
	e := &Entity{
		"topicname",
		queryMap.Get("classname"),
		queryMap.Get("methodname"),
		[]string{queryMap.Get("name")},
	}
	if ok, err := nsqProxy.AddEvent(*e); ok == false {
		fmt.Println(err)
		fmt.Fprintln(writer, err.Error())
	}else{
		fmt.Fprintln(writer, "添加事件成功，等待被消费")
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	indexhtml := `
    请求格式：
        1. 首页
        http://nsqgoproxy.vps90.vps.changbaops.com/

        2. 添加消费事件
        http://nsqgoproxy.vps90.vps.changbaops.com/publish?classname=DemoService&methodname=say&name=tiger
`
	fmt.Fprintln(writer, indexhtml)
}
