package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	// 设置Kafka集群地址
	brokerList := []string{"20.1.1.253:9092", "20.1.1.253:9093", "20.1.1.253:9094"}

	// 配置生产者
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// 创建生产者
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln("Failed to close Sarama producer:", err)
		}
	}()

	// 定义要发送到Kafka的消息
	msg := &sarama.ProducerMessage{
		Topic: "int_topic",
	}

	// 每隔5秒推送一次数字到Kafka
	go func() {
		i := 0
		for {
			time.Sleep(5 * time.Second)
			i++
			msg.Value = sarama.StringEncoder(fmt.Sprintf("%d", i))
			partition, offset, err := producer.SendMessage(msg)
			if err != nil {
				log.Printf("Error sending message to Kafka: %s\n", err)
			} else {
				log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
			}
		}
	}()

	// 创建消费者
	consumer, err := sarama.NewConsumer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama consumer:", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln("Failed to close Sarama consumer:", err)
		}
	}()
	time.Sleep(10 * time.Second)

	// 订阅int_topic主题的消息
	partitionConsumer, err := consumer.ConsumePartition("int_topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalln("Failed to start consuming partition:", err)
	}
	// 消费int_topic主题的消息
	go func() {
		for message := range partitionConsumer.Messages() {
			log.Printf("Received message with value %s\n", string(message.Value))
		}
	}()

	// 等待中断信号以退出程序
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals
	log.Println("Exiting")
}
