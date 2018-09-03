package core

import (
	"github.com/nsqio/go-nsq"
	"time"
	"strconv"
)

type NSQConnection struct {
	Source   string
	Config   *nsq.Config
	Producer *nsq.Producer
	Consumer *nsq.Consumer
}

func NewNSQConnection() QueueConnection {
	return &NSQConnection{}
}

// 设置配置文件
func (n *NSQConnection) SetConfig(content FileContent) QueueConnection {
	config := nsq.NewConfig()
	n.Source = content["host"].(string) + ":" + strconv.Itoa(content["port"].(int))

	if interval, ok := content["lookupd-poll-interval"].(time.Duration); ok {
		config.LookupdPollInterval = interval
	}

	if flight, ok := content["max-in-flight"].(int); ok {
		config.MaxInFlight = flight
	}

	n.Config = config

	return n
}

// 初始化生产者
func (n *NSQConnection) initProducer() QueueConnection {
	producer, err := nsq.NewProducer(n.Source, n.Config)
	ok := producer.Ping()

	if nil != err || nil != ok {
		panic("queue producer connection failed")
	}

	n.Producer = producer

	return n
}

// 初始化消费者
func (n *NSQConnection) InitConsumer(group, channel string) QueueConnection {
	consumer, err := nsq.NewConsumer(group, channel, n.Config)

	if nil != err {
		panic(err)
	}

	n.Consumer = consumer

	return n
}

// 添加任务
func (n *NSQConnection) AddHandler(f interface{}) QueueConnection {
	handle, ok := f.(nsq.HandlerFunc)

	if !ok {
		panic("nsq add handler func type error")
	}

	n.Consumer.AddHandler(handle)

	return n
}

// 运行队列
func (n *NSQConnection) Run() {
	err := n.Consumer.ConnectToNSQD(n.Source)

	if nil != err {
		panic(err)
	}
}

// 案例：异步执行，开启10个消费端来消费
// func runConsumerExample() {
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			x.InitConsumer(...).
// 				AddHandler(...).
// 				Run()
// 		}
// 	}()
// }
