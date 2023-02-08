package config

import (
	"fmt"
)

// Redis redis配置参数
type RedisConfig struct {
	Enable   bool
	Addr     string
	Password string
	Database int
}

func (a RedisConfig) DSN() string {
	return fmt.Sprintf("%s@%s", a.Password, a.Addr)
}
