package worker

import (
	"context"
	"log"

	"github.com/heromicro/omgind/pkg/mw/asyncq/worker/taskmeta"
	"github.com/heromicro/omgind/pkg/mw/queue"
	"github.com/heromicro/omgind/pkg/tipes"
	"github.com/hibiken/asynq"
)

type Consumer struct {
	queue queue.Queuer
	mux   *asynq.ServeMux
	srv   *asynq.Server
	// log   log.StdLogger
}

func NewConsumer(q queue.Queuer) *Consumer {

	srv := asynq.NewServer(
		q.Options().RedisClient,
		asynq.Config{
			Concurrency: 100,
			Queues:      q.Options().Names,
			IsFailure: func(err error) bool {
				if _, ok := err.(*taskmeta.RateLimitError); ok {
					return false
				}
				return true
			},
			RetryDelayFunc: taskmeta.GetRetryDelay,
			// Logger:         l,
		},
	)
	mux := asynq.NewServeMux()

	csm := &Consumer{
		queue: q,
		mux:   mux,
		srv:   srv,
	}

	return csm
}

func (c *Consumer) Start() {
	if err := c.srv.Start(c.mux); err != nil {
		log.Fatalln("error starting worker")
	}
}

func (c *Consumer) RegisterHandlers(taskName tipes.TaskName, handler asynq.Handler) {
	c.mux.Handle(string(taskName), handler)
}

func (c *Consumer) RegisterHandlerFuncs(taskName tipes.TaskName, handler func(context.Context, *asynq.Task) error) {
	c.mux.HandleFunc(string(taskName), handler)
}

func (c *Consumer) Stop() {
	c.srv.Stop()
	c.srv.Shutdown()
}
