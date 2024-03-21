package wirex

import (
	"github.com/heromicro/omgind/pkg/config"
	"github.com/heromicro/omgind/pkg/mw/asyncq/worker"
	"github.com/heromicro/omgind/pkg/mw/queue"
)

func InitAsynq(conf *config.AppConfig, q queue.Queuer) (*worker.Consumer, func(), error) {

	consumer := worker.NewConsumer(q)

	go consumer.Start()

	cleanFunc := func() {
		consumer.Stop()
	}
	return consumer, cleanFunc, nil
}
