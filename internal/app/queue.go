package app

import (
	"github.com/heromicro/omgind/pkg/config"
	"github.com/heromicro/omgind/pkg/config/option"
	"github.com/heromicro/omgind/pkg/mw/queue"
	"github.com/heromicro/omgind/pkg/mw/queue/redisqueue"
	"github.com/heromicro/omgind/pkg/mw/rdb"
	"github.com/heromicro/omgind/pkg/tipes"
)

func InitQueue(conf *config.AppConfig, rdb *rdb.Redis) (queue.Queuer, func(), error) {

	queueNames := map[string]int{
		string(tipes.RepaireTreeQueue): 1,
		string(tipes.ScheduleQueue):    1,
		string(tipes.DefaultQueue):     1,
	}

	opts := &queue.QueueOptions{
		Names:       queueNames,
		RedisClient: rdb,
		Type:        string(option.RedisQueueProvider),
		// PrometheusAddress: cfg.Prometheus.Dsn,
	}

	q := redisqueue.NewQueue(opts)

	cleanFunc := func() {

	}

	return q, cleanFunc, nil
}
