package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

// DeclareQueue
// @Description: 生成/连接队列
// @receiver c
// @param queueName	 队列名称
// @param durable	 持久化
// @param autoDelete 自动删除
// @param exclusive
// @param noWait
// @param args		 自定义参数
// @return amqp.Queue
// @return error
func (c *RabbitMQClient) DeclareQueue(queueName string, durable bool, autoDelete bool, exclusive bool, noWait bool, args amqp.Table) (amqp.Queue, error) {
	q, err := c.ch.QueueDeclare(
		queueName,
		durable,
		autoDelete,
		exclusive,
		noWait,
		args,
	)
	if err != nil {
		return amqp.Queue{}, fmt.Errorf("Failed to declare queue: %v", err)
	}
	return q, nil
}
