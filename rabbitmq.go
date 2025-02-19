package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

// RabbitMQConfig
// @Description: 连接配置
type RabbitMQConfig struct {
	Host           string // 主机地址
	Port           int    // 端口号，默认是 5672
	User           string // 用户名
	Password       string // 密码
	VHost          string // 虚拟主机名称，默认是 "/"
	TLS            bool   // 是否启用 TLS
	Heartbeat      int    // 心跳间隔（秒）
	ConnectionName string // 连接名称
}

// RabbitMQClient
// @Description: 客户端
type RabbitMQClient struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

// NewRabbitMQClient
// @Description: 实例化
// @param url
// @return *RabbitMQClient
// @return error
func NewRabbitMQClient(config RabbitMQConfig) (*RabbitMQClient, error) {
	if config.Port == 0 {
		if config.TLS {
			config.Port = 5671 // 默认 AMQPS 端口
		} else {
			config.Port = 5672 // 默认 AMQP 端口
		}
	}

	if config.VHost == "" {
		// 默认虚拟主机
		config.VHost = "/"
	}

	// 构建 AMQP URL
	scheme := "amqp"
	if config.TLS {
		scheme = "amqps"
	}

	url := fmt.Sprintf("%s://%s:%s@%s:%d/%s?heartbeat=%d&connection_name=%s",
		scheme, config.User, config.Password, config.Host, config.Port, config.VHost, config.Heartbeat, config.ConnectionName)

	// 连接到 RabbitMQ 服务器
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to RabbitMQ: %v", err)
	}

	// 打开通道
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("Failed to open a channel: %v", err)
	}

	// 返回 RabbitMQClient 实例
	return &RabbitMQClient{
		conn: conn,
		ch:   ch,
	}, nil
}

// Close
// @Description: 关闭连接
// @receiver c
func (c *RabbitMQClient) Close() {
	if c.ch != nil {
		c.ch.Close()
	}
	if c.conn != nil {
		c.conn.Close()
	}
}
