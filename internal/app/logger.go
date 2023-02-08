package app

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/heromicro/omgind/pkg/global"
	"github.com/heromicro/omgind/pkg/logger"
	loggerhook "github.com/heromicro/omgind/pkg/logger/hook"
	loggergormhook "github.com/heromicro/omgind/pkg/logger/hook/gorm"

	"github.com/sirupsen/logrus"
)

// InitLogger 初始化日志模块
func InitLogger() (func(), error) {
	c := global.CFG.Log
	logger.SetLevel(c.Level)
	logger.SetFormatter(c.Format)

	// 设定日志输出
	var file *os.File
	if c.Output != "" {
		switch c.Output {
		case "stdout":
			logger.SetOutput(os.Stdout)
		case "stderr":
			logger.SetOutput(os.Stderr)
		case "file":
			if name := c.OutputFile; name != "" {
				_ = os.MkdirAll(filepath.Dir(name), 0777)

				f, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					return nil, err
				}
				logger.SetOutput(f)
				file = f
			}
		}
	}

	var hook *loggerhook.Hook
	if c.EnableHook {
		var hookLevels []logrus.Level
		for _, lvl := range c.HookLevels {
			plvl, err := logrus.ParseLevel(lvl)
			if err != nil {
				return nil, err
			}
			hookLevels = append(hookLevels, plvl)
		}

		switch {
		case c.Hook.IsGorm():
			hc := global.CFG.LogGormHook

			var dsn string
			switch hc.DBType {
			case "mysql":
				dsn = global.CFG.MySQL.DSN()
			case "sqlite3":
				dsn = global.CFG.Sqlite3.DSN()
			case "postgres":
				dsn = global.CFG.Postgres.DSN()
			default:
				return nil, errors.New("unknown logger db")
			}

			h := loggerhook.New(loggergormhook.New(&loggergormhook.Config{
				DBType:       hc.DBType,
				DSN:          dsn,
				MaxLifetime:  hc.MaxLifetime,
				MaxOpenConns: hc.MaxOpenConns,
				MaxIdleConns: hc.MaxIdleConns,
				TableName:    hc.Table,
			}),
				loggerhook.SetMaxWorkers(c.HookMaxThread),
				loggerhook.SetMaxQueues(c.HookMaxBuffer),
				loggerhook.SetLevels(hookLevels...),
			)
			logger.AddHook(h)
			hook = h
		}
	}

	return func() {
		if file != nil {
			file.Close()
		}

		if hook != nil {
			hook.Flush()
		}
	}, nil
}
