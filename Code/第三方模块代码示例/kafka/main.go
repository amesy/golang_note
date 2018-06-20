package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()

	// Ack确认消息落盘, 依消息重要性。
	config.Producer.RequiredAcks = sarama.WaitForAll

	// 一个Topic就是一个消息队列，消息数据会按照一定的算法落到各Partitioner的机器上，提高并发性。此处采用随机分配。
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}

	// 指定topic。
	msg.Topic = "nginx_log"

	// 具体的日志数据。
	// 日志信息。
	msg.Value = sarama.StringEncoder("this is a good test, my message is good")

	client, err := sarama.NewSyncProducer([]string{"10.68.7.20:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}

	defer client.Close()

	// 同步和异步的区别：同步的话消息发送完就已存在kafka中，异步的话，先发给后台的channel,再发往kafka。

	// offset偏移量。
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}

	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
