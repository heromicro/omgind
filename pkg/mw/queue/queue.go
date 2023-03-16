package queue

import (
	"encoding/json"
	"time"

	"github.com/heromicro/omgind/pkg/mw/rdb"
	"github.com/heromicro/omgind/pkg/types"
)

type Queuer interface {
	Write(types.TaskName, types.QueueName, *Job) error
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
