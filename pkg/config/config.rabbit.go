package config

import "fmt"

// RabbitMQ 配置参数
type RabbitMQConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

// DSN 数据库连接串
func (a RabbitMQConfig) DSN() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d", a.User, a.Password, a.Host, a.Port)
}
