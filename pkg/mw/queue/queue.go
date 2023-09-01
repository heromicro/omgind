package queue

import (
	"encoding/json"
	"time"

	"github.com/heromicro/omgind/pkg/mw/rdb"
	"github.com/heromicro/omgind/pkg/tipes"
)

type Queuer interface {
	Write(tipes.TaskName, tipes.QueueName, *Job) error
	Options() *QueueOptions
}

type Job struct {
	ID      string          `json:"id"`
	Payload json.RawMessage `json:"payload"`
	Delay   time.Duration   `json:"delay"`
}

type QueueOptions struct {
	Names        map[string]int
	Type         string
	RedisClient  *rdb.Redis
	RedisAddress string
	// PrometheusAddress string
}
