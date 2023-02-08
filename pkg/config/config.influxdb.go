package config

import "fmt"

// Influxdb 配置参数
type InfluxDBConfig struct {
	Schema string
	Host   string
	Port   int
	Token  string
	Bucket string
	Org    string
}

// DSN 数据库连接串
func (a InfluxDBConfig) DSN() string {
	return fmt.Sprintf("%s://%s:%d", a.Schema, a.Host, a.Port)
}
