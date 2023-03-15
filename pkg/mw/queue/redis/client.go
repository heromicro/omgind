package redis

import (
	"github.com/heromicro/omgind/pkg/mw/queue"
	"github.com/hibiken/asynq"
)

type RedisQueue struct {
	opts      queue.QueueOptions
	client    *asynq.Client
	inspector *asynq.Inspector
}

func NewQueue(opts queue.QueueOptions) queue.Queuer {
	client := asynq.NewClient(opts.RedisClient)
	inspector := asynq.NewInspector(opts.RedisClient)
	return &RedisQueue{
		client:    client,
		opts:      opts,
		inspector: inspector,
	}
}
