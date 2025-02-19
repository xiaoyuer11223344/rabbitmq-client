package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

// PublishMessage
// @Description: 发布消息
// @receiver c
// @param exchange
// @param routingKey
// @param body
// @param mandatory
// @param immediate
// @return error
func (c *RabbitMQClient) PublishMessage(exchange, routingKey, body string, mandatory, immediate bool) error {
	err := c.ch.Publish(
		exchange,   // 交换机
		routingKey, // 路由键
		mandatory,  // 如果消息无法路由到任何队列，是否返回给生产者
		immediate,  // 如果队列中没有消费者，是否立即返回给生产者
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return fmt.Errorf("Failed to publish message: %v", err)
	}
	return nil
}
