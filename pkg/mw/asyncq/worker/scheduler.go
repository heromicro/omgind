package worker

import (
	"log"

	"github.com/heromicro/omgind/pkg/mw/queue"
	"github.com/heromicro/omgind/pkg/types"
	"github.com/hibiken/asynq"
)

type Scheduler struct {
	// log   log.StdLogger
	queue queue.Queuer
	inner *asynq.Scheduler
}

func NewScheduler(queue queue.Queuer) *Scheduler {
	scheduler := asynq.NewScheduler(queue.Options().RedisClient, &asynq.SchedulerOpts{
		// Logger: log,
	})

	return &Scheduler{
		// log:   log,
		inner: scheduler,
		queue: queue,
	}
}

func (s *Scheduler) Start() {
	if err := s.inner.Start(); err != nil {
		// s.log.WithError(err).Fatal("Could not start scheduler")
		log.Fatalln("Could not start scheduler")
	}
}

func (s *Scheduler) RegisterTask(cronspec string, queue types.QueueName, taskName types.TaskName) {
	task := asynq.NewTask(string(taskName), nil)
	_, err := s.inner.Register(cronspec, task, asynq.Queue(string(queue)))
	if err != nil {
		// s.log.WithError(err).Fatalf("Failed to register %s scheduler task", taskName)
		log.Fatalf("Failed to register %s scheduler task \n ", taskName)
	}
}

func (s *Scheduler) Stop() {
	s.inner.Shutdown()
}
