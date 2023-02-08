package config

// Sqlite3 sqlite3配置参数
type Sqlite3Config struct {
	Path string
}

// DSN 数据库连接串
func (a Sqlite3Config) DSN() string {
	return a.Path
}
