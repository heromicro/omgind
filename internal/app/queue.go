package app

import (
	"github.com/heromicro/omgind/pkg/config"
	"github.com/heromicro/omgind/pkg/config/option"
	"github.com/heromicro/omgind/pkg/mw/queue"
	"github.com/heromicro/omgind/pkg/mw/queue/redisqueue"
	"github.com/heromicro/omgind/pkg/mw/rdb"
	"github.com/heromicro/omgind/pkg/types"
)

func InitQueue(conf *config.AppConfig, rdb *rdb.Redis) (queue.Queuer, func(), error) {

	queueNames := map[string]int{
		string(types.DistrictQueue): 1,
		string(types.ScheduleQueue): 1,
		string(types.DefaultQueue):  1,
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
