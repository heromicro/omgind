package config

// system redis配置参数
type SystemConfig struct {
	RunMode     string `mapstructure:"runMode"  json:"runMode"`
	WWW         string `mapstructure:"www"  json:"www"`
	Swagger     bool   `mapstructure:"swagger"  json:"swagger"`
	PrintConfig bool   `mapstructure:"printConfig"  json:"printconfig"`
}
