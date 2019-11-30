package nsqgoproxy

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// nsqproxy 将订阅到的数据扔给下游consumer

// PHP-FPM handler
type PHPFPMHandler struct {
}

func PHPHandler(data string) {
	payload := DataEncode(data)
	// TODO 消费这块其实可以做类似于 round-robin 的形式去扔给不同的消费者
	// 1. 做一套优雅的重启机制
	// 2. 定期读取配置到内存中，每次从内存（也可以是外部存储，如redis，db等）读取配置
	response, err := http.PostForm(PHP_FPM_ADDR, payload)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	bs, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	// 发给钉钉暂时做测试用
	SendToDIngtalk(string(bs))
	fmt.Println(string(bs))
}
