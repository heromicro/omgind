package scheme

import (
	"errors"
	"github.com/google/wire"
	"github.com/heromicro/omgind/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"path/filepath"
	"time"
)

// NewGormClient 创建 GORM 连接并返回清理函数，不自动执行迁移。
func NewGormClient(cfg *config.AppConfig) (*gorm.DB, func(), error) {
	noop := func() {}
	if cfg == nil {
		return nil, noop, errors.New("app config is required")
	}
	var dsn string
	switch cfg.Gorm.DBType {
	case "mysql":
		dsn = cfg.MySQL.DSN()
	case "postgres":
		dsn = cfg.Postgres.DSN()
	case "sqlite3":
		dsn = cfg.Sqlite3.DSN()
		if err := os.MkdirAll(filepath.Dir(dsn), 0755); err != nil {
			return nil, noop, err
		}
	default:
		return nil, noop, errors.New("unknown gorm db")
	}
	db, err := gorm.Open(cfg.Gorm.DBType, dsn)
	if err != nil {
		return nil, noop, err
	}
	cleanup := func() { _ = db.Close() }
	db.LogMode(cfg.Gorm.Debug)
	db.DB().SetMaxIdleConns(cfg.Gorm.MaxIdleConns)
	db.DB().SetMaxOpenConns(cfg.Gorm.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(cfg.Gorm.MaxLifetime) * time.Second)
	if err := db.DB().Ping(); err != nil {
		cleanup()
		return nil, noop, err
	}
	return db, cleanup, nil
}

var GormProviderSet = wire.NewSet(NewGormClient)
