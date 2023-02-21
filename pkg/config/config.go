package config

import (
	"C"

	"github.com/heromicro/omgind/pkg/config/option"
	"github.com/spf13/viper"
)
import "log"

// Config 配置参数
type AppConfig struct {
	System       option.SystemConfig `mapstructure:"System"  json:"system"`
	HTTP         option.HTTPConfig   `mapstructure:"HTTP" json:"http"`
	Menu         option.MenuConfig   `mapstructure:"menu" json:"menu"`
	Casbin       option.CasbinConfig `mapstructure:"Casbin" json:"casbin"`
	Log          option.LogConfig    `mapstructure:"Log" json:"log"`
	LogGormHook  option.LogGormHook
	LogMongoHook option.LogMongoHook
	Root         option.RootConfig        `mapstructure:"Root" json:"root"`
	JWTAuth      option.JWTAuthConfig     `mapstructure:"JWTAuth" json:"jwtAuth"`
	Monitor      option.MonitorConfig     `mapstructure:"Monitor" json:"monitor"`
	Captcha      option.CaptchaConfig     `mapstructure:"Captcha" json:"captcha"`
	RateLimiter  option.RateLimiterConfig `mapstructure:"RateLimiter" json:"rateLimiter"`
	CORS         option.CORSConfig        `mapstructure:"CORS" json:"CORS"`
	GZIP         option.GZIPConfig        `mapstructure:"GZIP" json:"GZIP"`
	Redis        option.RedisConfig       `mapstructure:"Redis" json:"redis"`
	Ent          option.EntConfig         `mapstructure:"Ent" json:"ent"`
	MySQL        option.MySQLConfig       `mapstructure:"MySQL" json:"mysql"`
	Postgres     option.PostgresConfig    `mapstructure:"Postgres" json:"postgres"`
	Sqlite3      option.Sqlite3Config     `mapstructure:"Sqlite3" json:"sqlite3"`
	RabbitMQ     option.RabbitMQConfig    `mapstructure:"RabbitMQ" json:"rabbitmq"`
	InfluxDB     option.InfluxDBConfig    `mapstructure:"InfluxDB" json:"influxdb"`
}

// IsDebugMode 是否是debug模式
func (c *AppConfig) IsDebugMode() bool {
	return c.System.RunMode == "debug"
}

func New(pf string) (*viper.Viper, error) {

	var (
		err error
		v   = viper.New()
	)

	v.SetConfigFile(string(pf))
	v.SetConfigType("yml")
	v.SetConfigType("yaml")

	// FIXME: change `omgind` to real
	v.SetEnvPrefix("omgind")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err == nil {
		log.Printf("use config file -> %s\n", v.ConfigFileUsed())
	} else {
		return nil, err
	}

	return v, err
}
