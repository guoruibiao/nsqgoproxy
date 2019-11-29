package nsqgoproxy

import (
	"sync"
	"log"
	"github.com/bitly/go-nsq"
	"github.com/pkg/errors"
	"bytes"
	"encoding/json"
		"os"
	"fmt"
	)

// 提供:
// 1. 上层承接并扔到nsq
// 2. 从nsq订阅，并扔给下层consumer


type NSQProxy struct {
	waitgroup sync.WaitGroup
	poolSize int
	jobsPool chan int
	producer *nsq.Producer
}

func NewNSQProxy(poolcount int) (*NSQProxy, error) {
	config := nsq.NewConfig()

	producer, err := nsq.NewProducer(NSQ_ADDR, config)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("can not create nsq producer")
	}

	return &NSQProxy{
		sync.WaitGroup{},
		poolcount,
		make(chan int, poolcount),
		producer,
	}, nil
}

func (p *NSQProxy) AddEvent(e Entity) (bool, error) {
	// TODO 校验 消费方是否有此消费能力
	writer := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(writer)
	encoder.Encode(e)
	p.producer.Publish(e.TopicName, writer.Bytes())
	return true, nil
}

func (p *NSQProxy) GetEvent(topicName string) (e *Entity, err error) {
	consumer, err := nsq.NewConsumer(topicName, "channel", nsq.NewConfig())
	if err != nil {
		return nil, err
	}
	consumer.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		fmt.Println("message: ", string((msg.Body)))
		PHPHandler(string(msg.Body))
		// TODO 在此处做消费处理
		return nil
	}))
	err = consumer.ConnectToNSQD(NSQ_ADDR)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	// 消费者等待上方handler处理
	<- make(chan bool)
	fmt.Println("consumer consumed.")
	return
}