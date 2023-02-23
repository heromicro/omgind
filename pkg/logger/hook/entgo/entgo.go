package entgo

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/pkg/logger"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
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

func New(c *Config) *Hook {

	eclient, _, err := NewEntClient(c)
	if err != nil {
		panic(err)
	}

	return &Hook{
		eclient: eclient,
	}
}

// NewGormDB 创建DB实例
func NewEntClient(cfg *Config) (*ent.Client, func(), error) {

	db, err := sql.Open(cfg.DBType, cfg.DSN)
	if err != nil {
		return nil, func() {}, err
	}

	// logging to db
	loggerAdapter := zerologadapter.New(zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false}))

	db = sqldblogger.OpenDriver(cfg.DSN, db.Driver(), loggerAdapter)

	drv := entsql.OpenDB(cfg.DBType, db)

	cleanFunc := func() {
		drv.Close()
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Second)

	if err := db.Ping(); err != nil {
		return nil, cleanFunc, err
	}

	client := ent.NewClient(ent.Driver(drv))

	return client, cleanFunc, nil
}

// Hook gorm日志钩子
type Hook struct {
	eclient *ent.Client
}

// Exec 执行日志写入
func (h *Hook) Exec(entry *logrus.Entry) error {

	create_logger := h.eclient.SysLogging.Create()

	create_logger = create_logger.SetLevel(entry.Level.String())
	create_logger = create_logger.SetMessage(entry.Message)
	create_logger = create_logger.SetMessage(entry.Message)

	data := entry.Data
	if v, ok := data[logger.TraceIDKey]; ok {
		traceID, _ := v.(string)
		create_logger = create_logger.SetTraceID(traceID)
		delete(data, logger.TraceIDKey)
	}
	if v, ok := data[logger.UserIDKey]; ok {
		userID, _ := v.(string)
		create_logger = create_logger.SetUserID(userID)
		delete(data, logger.UserIDKey)
	}
	if v, ok := data[logger.TagKey]; ok {
		tag, _ := v.(string)
		create_logger = create_logger.SetTag(tag)
		delete(data, logger.TagKey)
	}
	if v, ok := data[logger.StackKey]; ok {
		errorStack, _ := v.(string)
		create_logger = create_logger.SetErrorStack(errorStack)
		delete(data, logger.StackKey)
	}
	if v, ok := data[logger.VersionKey]; ok {
		version, _ := v.(string)
		create_logger = create_logger.SetVersion(version)
		delete(data, logger.VersionKey)
	}

	if len(data) > 0 {
		b, _ := json.Marshal(data)
		data := string(b)
		create_logger = create_logger.SetData(data)
	}

	_, err := create_logger.Save(context.Background())

	return err
}

// Close 关闭钩子
func (h *Hook) Close() error {
	return h.eclient.Close()
}