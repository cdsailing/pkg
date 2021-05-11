package kafka

import (
	"errors"
	"github.com/Shopify/sarama"
	"github.com/cdsailing/pkg/log"
)

var (
	producer         sarama.SyncProducer
	bootstrapServers []string
	conf             *sarama.Config
)

type Producer struct {
	sarama.SyncProducer
}

func (this *Producer) Write(data []byte, topic string) (bool, error) {
	if len(topic) <= 0 {
		return false, errors.New("topic不能为空")
	}
	// 构建 消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.ByteEncoder(data)
	if producer == nil {
		producer = InitProducer(bootstrapServers, conf)
	}
	// 发送消息
	message, offset, err := producer.SendMessage(msg)
	if err != nil {
		return false, err
	}
	log.Infof("数据发送到kafka %v %v", offset, message)
	return true, nil
}

func InitProducer(servers []string, config *sarama.Config) Producer {
	// 生成 生产者配置文件
	//config := sarama.NewConfig()
	// 设置生产者 消息 回复等级 0 1 all
	//config.Producer.RequiredAcks = sarama.WaitForAll
	//// 设置生产者 成功 发送消息 将在什么 通道返回
	//config.Producer.Return.Successes = true
	//// 设置生产者 发送的分区
	//config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 连接 kafka
	conf = config
	bootstrapServers = servers
	producer, err := sarama.NewSyncProducer(servers, config)
	if err != nil {
		log.Error(err)
		return Producer{nil}
	}
	return Producer{producer}
}
