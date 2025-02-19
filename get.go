package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

// GetMessage
// @Description: 获取消息
// @receiver c
// @param queueName
// @param autoAck
// @return *amqp.Delivery
// @return bool
// @return error
func (c *RabbitMQClient) GetMessage(queueName string, autoAck bool) (*amqp.Delivery, bool, error) {
	msg, ok, err := c.ch.Get(queueName, autoAck)
	if err != nil {
		return nil, false, fmt.Errorf("Failed to get message: %v", err)
	}
	return &msg, ok, nil
}
