package redisqueue

import (
	"errors"

	"github.com/google/wire"
	"github.com/heromicro/omgind/pkg/config"
	"github.com/heromicro/omgind/pkg/config/option"
	"github.com/heromicro/omgind/pkg/mw/queue"
	"github.com/heromicro/omgind/pkg/mw/rdb"
	"github.com/heromicro/omgind/pkg/tipes"
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
	"github.com/oklog/ulid/v2"
)

var ProviderSet = wire.NewSet(NewClient)

type RedisQueue struct {
	opts      *queue.QueueOptions
	client    *asynq.Client
	inspector *asynq.Inspector
}

func NewClient(cfg *config.AppConfig, rdb *rdb.Redis) (*asynq.Client, error) {
	// https://github.com/frain-dev/convoy.git
	if cfg.Queue.Type != option.RedisQueueProvider {
		return nil, errors.New("please select the redis driver in your config")
	}

	client := asynq.NewClient(rdb)

	return client, nil
}

func NewQueue(opts *queue.QueueOptions) queue.Queuer {

	client := asynq.NewClient(opts.RedisClient)
	inspector := asynq.NewInspector(opts.RedisClient)
	return &RedisQueue{
		client:    client,
		opts:      opts,
		inspector: inspector,
	}
}

func (q *RedisQueue) Write(taskName tipes.TaskName, queueName tipes.QueueName, job *queue.Job) error {
	if job.ID == "" {
		job.ID = ulid.Make().String()
	}
	t := asynq.NewTask(string(taskName), job.Payload, asynq.Queue(string(queueName)), asynq.TaskID(job.ID), asynq.ProcessIn(job.Delay))
	_, err := q.client.Enqueue(t)
	return err
}

func (q *RedisQueue) Options() *queue.QueueOptions {
	return q.opts
}

func (q *RedisQueue) Monitor() *asynqmon.HTTPHandler {
	h := asynqmon.New(asynqmon.Options{
		RootPath:     "/queue/monitoring",
		RedisConnOpt: q.opts.RedisClient,
		// PrometheusAddress: q.opts.PrometheusAddress,
	})
	return h
}

func (q *RedisQueue) Inspector() *asynq.Inspector {
	return q.inspector
}

func (q *RedisQueue) DeleteEventDeliveriesfromQueue(queuename tipes.QueueName, ids []string) error {
	for _, id := range ids {
		taskInfo, err := q.inspector.GetTaskInfo(string(queuename), id)
		if err != nil {
			return err
		}
		if taskInfo.State == asynq.TaskStateActive {
			err = q.inspector.CancelProcessing(id)
			if err != nil {
				return err
			}
		}
		err = q.inspector.DeleteTask(string(queuename), id)
		if err != nil {
			return err
		}
	}
	return nil
}
