package nsqgoproxy

import "time"

type (
	Entity struct {
		TopicName string
		Classname string
		MethodName string
		Duration time.Duration
		Parameters []string
	}
)
