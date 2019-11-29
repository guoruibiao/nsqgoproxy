package nsqgoproxy

import (
	"net/url"
	"github.com/guoruibiao/gorequests"
		"fmt"
	)

func DataEncode(data string) url.Values {
	return url.Values{"data": {data}}
}

func SendToDIngtalk(message string) {
	payload := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": message,
		},
		"at":  map[string]interface{}{
			"atMobiles": nil,
			"isAtAll": true,
		},
	}
	response, err := gorequests.NewRequest("POST", DINGTALK_WEBHOOK).Body(payload).DoRequest()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.Content())
}