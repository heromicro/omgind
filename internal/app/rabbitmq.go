package app

import (
	"fmt"

	"github.com/streadway/amqp"

	"github.com/heromicro/omgind/pkg/global"
)

func InitRabbitMQ() (*amqp.Connection, func(), error) {
	cfg := global.CFG.RabbitMQ

	conn, err := amqp.Dial(cfg.DSN())
	fmt.Println(" ------- rabbit mq dsn ", cfg.DSN())
	if err != nil {
		return nil, nil, err
	}

	return conn, func() {
		err := conn.Close()
		if err != nil {
			fmt.Println(" ---- rabbitmq close err: ", err)
		}
	}, nil
}
