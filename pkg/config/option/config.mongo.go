package option

// MongoConfig 使用独立连接，与 Ent/GORM 配置互不影响。
type MongoConfig struct {
	URI            string
	Database       string
	ConnectTimeout int // 秒，默认 10
}
