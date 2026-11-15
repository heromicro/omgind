package option

// GormConfig 复用 MySQL/Postgres/Sqlite3 的连接参数。
type GormConfig struct {
	DBType       string
	Debug        bool
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}
