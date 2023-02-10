package gorm

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/heromicro/omgind/pkg/logger"
	"github.com/oklog/ulid/v2"
	"github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var tableName string

// Config 配置参数
type Config struct {
	DBType       string
	DSN          string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
	TableName    string
	TablePrefix  string
}

// New 创建基于gorm的钩子实例(需要指定表名)
func New(c *Config) *Hook {
	tableName = c.TableName

	var dialector gorm.Dialector
	switch strings.ToLower(c.DBType) {
	case "mysql":
		cfg, err := mysqlDriver.ParseDSN(c.DSN)
		if err != nil {
			return nil
		}
		err = createDatabaseWithMySQL(cfg)
		if err != nil {
			return nil
		}

		dialector = mysql.Open(c.DSN)
	case "postgres":
		dialector = postgres.Open(c.DSN)
	case "sqlite3":
		dialector = sqlite.Open(c.DSN)
	default:
		panic("missing dbtype")
	}

	gconfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.TablePrefix,
			SingularTable: true,
		},
	}

	db, err := gorm.Open(dialector, gconfig)
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Second)

	if !db.Migrator().HasTable(new(LogItem)) {
		db.AutoMigrate(new(LogItem))
	}
	return &Hook{
		db: db,
	}
}

func createDatabaseWithMySQL(cfg *mysqlDriver.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", cfg.User, cfg.Passwd, cfg.Addr)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET = `utf8mb4`;", cfg.DBName)
	_, err = db.Exec(query)
	return err
}

// Hook gorm日志钩子
type Hook struct {
	db *gorm.DB
}

// Exec 执行日志写入
func (h *Hook) Exec(entry *logrus.Entry) error {

	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	entropy := rand.New(source)

	id := ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	item := &LogItem{
		ID:        id,
		Level:     entry.Level.String(),
		Message:   entry.Message,
		CreatedAt: entry.Time,
	}

	data := entry.Data
	if v, ok := data[logger.TraceIDKey]; ok {
		item.TraceID, _ = v.(string)
		delete(data, logger.TraceIDKey)
	}
	if v, ok := data[logger.UserIDKey]; ok {
		item.UserID, _ = v.(string)
		delete(data, logger.UserIDKey)
	}
	if v, ok := data[logger.TagKey]; ok {
		item.Tag, _ = v.(string)
		delete(data, logger.TagKey)
	}
	if v, ok := data[logger.StackKey]; ok {
		item.ErrorStack, _ = v.(string)
		delete(data, logger.StackKey)
	}
	if v, ok := data[logger.VersionKey]; ok {
		item.Version, _ = v.(string)
		delete(data, logger.VersionKey)
	}

	if len(data) > 0 {
		b, _ := json.Marshal(data)
		item.Data = string(b)
	}

	return h.db.Create(item).Error
}

// Close 关闭钩子
func (h *Hook) Close() error {
	db, err := h.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

/*

	field.String("level").MaxLen(32).Comment("日志级别"),
	field.String("trace_id").MaxLen(128).Comment("跟踪ID"),
	field.String("user_id").MaxLen(128).Comment("用户ID"),
	field.String("tag").MaxLen(128).Comment("Tag"),

	field.String("version").MaxLen(64).Comment("版本号"),
	field.String("message").Comment("消息"),
	field.Text("data").Immutable().Optional().Nillable().Comment("日志数据(json string)"),
	field.Text("error_stack").Comment("日志数据(json string)"),

*/
// LogItem 存储日志项
type LogItem struct {
	// ID         uint      `gorm:"column:id;primary_key;auto_increment;"` // id
	ID         string    `gorm:"column:id;primary_key;size:36;"`           // id
	Level      string    `gorm:"column:level;size:20;index;size:32;"`      // 日志级别
	TraceID    string    `gorm:"column:trace_id;size:128;index;size:128;"` // 跟踪ID
	UserID     string    `gorm:"column:user_id;size:36;index;size:32;"`    // 用户ID
	Tag        string    `gorm:"column:tag;size:128;index;type:text;"`     // Tag
	Version    string    `gorm:"column:version;index;size:64;"`            // 版本号
	Message    string    `gorm:"column:message;size:1024;"`                // 消息
	Data       string    `gorm:"column:data;type:text;"`                   // 日志数据(json)
	ErrorStack string    `gorm:"column:error_stack;type:text;"`            // Error Stack
	CreatedAt  time.Time `gorm:"column:crtd_at;index"`                     // 创建时间
}

// TableName 表名
func (LogItem) TableName() string {
	return tableName
}
