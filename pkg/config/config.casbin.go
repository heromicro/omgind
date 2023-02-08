package config

// Casbin casbin配置参数
type CasbinConfig struct {
	Enable           bool
	Debug            bool
	Model            string
	AutoLoad         bool
	AutoLoadInternal int
}
