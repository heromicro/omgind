package config

import "C"

// Config 配置参数
type Config struct {
	System       SystemConfig `mapstructure:"System"  json:"system"`
	HTTP         HTTPConfig   `mapstructure:"HTTP" json:"http"`
	Menu         MenuConfig   `mapstructure:"menu" json:"menu"`
	Casbin       CasbinConfig `mapstructure:"Casbin" json:"casbin"`
	Log          LogConfig    `mapstructure:"Log" json:"log"`
	LogGormHook  LogGormHook
	LogMongoHook LogMongoHook
	Root         RootConfig        `mapstructure:"Root" json:"root"`
	JWTAuth      JWTAuthConfig     `mapstructure:"JWTAuth" json:"jwtAuth"`
	Monitor      MonitorConfig     `mapstructure:"Monitor" json:"monitor"`
	Captcha      CaptchaConfig     `mapstructure:"Captcha" json:"captcha"`
	RateLimiter  RateLimiterConfig `mapstructure:"RateLimiter" json:"rateLimiter"`
	CORS         CORSConfig        `mapstructure:"CORS" json:"CORS"`
	GZIP         GZIPConfig        `mapstructure:"GZIP" json:"GZIP"`
	Redis        RedisConfig       `mapstructure:"Redis" json:"redis"`
	Ent          EntConfig         `mapstructure:"Ent" json:"ent"`
	MySQL        MySQLConfig       `mapstructure:"MySQL" json:"mysql"`
	Postgres     PostgresConfig    `mapstructure:"Postgres" json:"postgres"`
	Sqlite3      Sqlite3Config     `mapstructure:"Sqlite3" json:"sqlite3"`
	RabbitMQ     RabbitMQConfig    `mapstructure:"RabbitMQ" json:"rabbitmq"`
	InfluxDB     InfluxDBConfig    `mapstructure:"InfluxDB" json:"influxdb"`
}

// IsDebugMode 是否是debug模式
func (c *Config) IsDebugMode() bool {
	return c.System.RunMode == "debug"
}
