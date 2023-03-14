package rabbit

import (
	"log"

	"github.com/google/wire"
	"github.com/heromicro/omgind/pkg/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func New(conf *config.AppConfig) (*amqp.Connection, func(), error) {
	cfg := conf.RabbitMQ

	conn, err := amqp.Dial(cfg.DSN())
	log.Println(" ------- rabbit mq dsn ", cfg.DSN())
	if err != nil {
		return nil, nil, err
	}

	//ch, err := conn.Channel()
	//ch.Publish()
	//amqp.Authentication()

	return conn, func() {
		err := conn.Close()
		if err != nil {
			log.Println(" ---- rabbitmq close err: ", err)
		}
	}, nil
}

var ProviderSet = wire.NewSet(New)
