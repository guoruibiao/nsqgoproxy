package nsqgoproxy

import (
	"fmt"
	"os"
	"testing"
)

var nsqProxy NSQProxy
func init() {
	server, err := NewNSQProxy(3)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	nsqProxy = *server
}

func TestNSQProxy_AddEvent(t *testing.T) {
	e := Entity{
		"topicname",
		"DemoService",
		"say",
		[]string{"泰戈尔🤩"},
	}
	fmt.Println(nsqProxy.AddEvent(e))
}

func TestNSQProxy_GetEvent(t *testing.T) {
	topicName := "topicname"
	nsqProxy.GetEvent(topicName)
}

