package scheme

import (
	"context"
	"errors"
	"log"
	"os"
	"path/filepath"
	"time"

	"database/sql"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/google/wire"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"

	"github.com/heromicro/omgind/internal/gen/mainent"
	_ "github.com/heromicro/omgind/internal/gen/mainent/runtime"
	"github.com/heromicro/omgind/pkg/config"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

// 初始化ent存储
func New(acfg *config.AppConfig) (*mainent.Client, func(), error) {
	cfg := acfg.Ent
	cli, cleanFunc, err := NewEntClient(acfg)
	if err != nil {
		return nil, cleanFunc, err
	}

	if cfg.EnableAutoMigrate {
		err := cli.Schema.Create(
			context.Background(),
			// FIXME::  sql/schema: pq: column "id" of relation "xxxx_table" is not an identity column
			//migrate.WithGlobalUniqueID(true),
		)
		if err != nil {
			return nil, cleanFunc, err
		}
	}
	// add hooks
	cli.Use(func(next mainent.Mutator) mainent.Mutator {
		return mainent.MutateFunc(func(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
			start := time.Now()
			defer func() {
				log.Printf("Op=%s\tType=%s\tTime=%s\tConcreteType=%T\n", m.Op(), m.Type(), time.Since(start), m)
			}()
			return next.Mutate(ctx, m)
		})
	})

	return cli, cleanFunc, nil
}

// 创建ent实例
func NewEntClient(cfg *config.AppConfig) (*mainent.Client, func(), error) {

	var dsn string

	switch cfg.Ent.DBType {
	case "mysql":
		dsn = cfg.MySQL.DSN()
	case "sqlite3":
		dsn = cfg.Sqlite3.DSN()
		_ = os.MkdirAll(filepath.Dir(dsn), 0777)
	case "postgres":
		dsn = cfg.Postgres.DSN()
	default:
		return nil, nil, errors.New("unknown ent db")
	}

	log.Println("======== connecting ", cfg.Ent.DBType, " ", dsn)

	db, err := sql.Open(cfg.Ent.DBType, dsn)
	if err != nil {
		return nil, func() {}, err
	}

	// logging to db
	loggerAdapter := zerologadapter.New(zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false}))

	////
	//logger := logrus.New()
	//logger.Level = logrus.DebugLevel
	//logger.Formatter = &logrus.JSONFormatter{}
	//loggerAdapter := logrusadapter.New(logger)

	db = sqldblogger.OpenDriver(dsn, db.Driver(), loggerAdapter)

	drv := entsql.OpenDB(cfg.Ent.DBType, db)

	cleanFunc := func() {
		drv.Close()
	}

	db.SetMaxIdleConns(cfg.Ent.MaxIdleConns)
	db.SetMaxOpenConns(cfg.Ent.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.Ent.MaxLifetime) * time.Second)

	if err := db.Ping(); err != nil {
		return nil, cleanFunc, err
	}

	client := mainent.NewClient(mainent.Driver(drv))

	return client, cleanFunc, nil
}

var ProviderSet = wire.NewSet(New)
