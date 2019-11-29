package nsqgoproxy

import (
		"net/http"
	"fmt"
	"io/ioutil"
)

// nsqproxy 将订阅到的数据扔给下游consumer


// PHP-FPM handler
type PHPFPMHandler struct {
}

func PHPHandler(data string) {
	payload := DataEncode(data)
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
