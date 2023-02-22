package app

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/heromicro/omgind/pkg/global"
)

func InitRabbitMQ() (*amqp.Connection, func(), error) {
	cfg := global.CFG.RabbitMQ

	conn, err := amqp.Dial(cfg.DSN())
	log.Println(" ------- rabbit mq dsn ", cfg.DSN())
	if err != nil {
		return nil, nil, err
	}

	return conn, func() {
		err := conn.Close()
		if err != nil {
			log.Println(" ---- rabbitmq close err: ", err)
		}
	}, nil
}
