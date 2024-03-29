package option

// Log 日志配置参数
type LogConfig struct {
	Level         int
	Format        string
	Output        string
	OutputFile    string
	EnableHook    bool
	HookLevels    []string
	Hook          LogHook
	HookMaxThread int
	HookMaxBuffer int
}

// LogHook 日志钩子
type LogHook string

// IsGorm 是否是gorm钩子
func (h LogHook) IsGorm() bool {
	return h == "gorm"
}

// IsMongo 是否是mongo钩子
func (h LogHook) IsMongo() bool {
	return h == "mongo"
}

// IsEntgo 是否是ent钩子
func (h LogHook) IsEntgo() bool {
	return h == "ent" || h == "entgo"
}

// LogGormHook 日志gorm钩子配置
type LogGormHook struct {
	DBType       string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
	Table        string
}

// LogEntHook 日志ent钩子配置
type LogEntHook struct {
	DBType       string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
	Table        string
}

// LogMongoHook 日志mongo钩子配置
type LogMongoHook struct {
	Collection string
}
