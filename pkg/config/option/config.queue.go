package option

import (
	_ "github.com/kelseyhightower/envconfig"
)

type QueueProvider string

const (
	RedisQueueProvider QueueProvider = "redis"
)

type Queue struct {
	Type  QueueProvider `toml:"type" envconfig:"OMGIND_QUEUE_PROVIDER"`
	Redis RedisConfig   `toml:"redis"`
}
