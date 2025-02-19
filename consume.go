package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

// Consume
// @Description: 监听消息
// @receiver c
// @param queueName
// @param consumerTag
// @param autoAck
// @param exclusive
// @param noLocal
// @param noWait
// @param args
// @param callback
// @return error
func (c *RabbitMQClient) Consume(queueName string, consumerTag string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args amqp.Table, callback func(delivery amqp.Delivery)) error {
	// 开始消费消息
	deliveries, err := c.ch.Consume(
		queueName,   // 队列名称
		consumerTag, // 消费者标签，用于区分不同的消费者
		autoAck,     // 是否自动确认
		exclusive,   // 排他消费
		noLocal,     // 不接收自己发送的消息
		noWait,      // 不等待服务器确认
		args,        // 额外参数
	)

	if err != nil {
		return fmt.Errorf("Failed to start consuming: %v", err)
	}

	// 启动一个go func来处理消息
	go func() {
		for delivery := range deliveries {
			callback(delivery)
		}
	}()

	return nil
}
