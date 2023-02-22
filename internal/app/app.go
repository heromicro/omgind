package app

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/gops/agent"

	"github.com/heromicro/omgind/pkg/global"
	"github.com/heromicro/omgind/pkg/logger"

	// 引入swagger
	_ "github.com/heromicro/omgind/internal/app/swagger"
)

type options struct {
	ConfigFile string
	ModelFile  string
	MenuFile   string
	WWWDir     string
	Version    string
}

// Option 定义配置项
type Option func(*options)

// SetConfigFile 设定配置文件
func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

// SetModelFile 设定casbin模型配置文件
func SetModelFile(s string) Option {
	return func(o *options) {
		o.ModelFile = s
	}
}

// SetWWWDir 设定静态站点目录
func SetWWWDir(s string) Option {
	return func(o *options) {
		o.WWWDir = s
	}
}

// SetMenuFile 设定菜单数据文件
func SetMenuFile(s string) Option {
	return func(o *options) {
		o.MenuFile = s
	}
}

// SetVersion 设定版本号
func SetVersion(s string) Option {
	return func(o *options) {
		o.Version = s
	}
}

// Init 应用初始化
func Init(ctx context.Context, opts ...Option) (func(), error) {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	global.MustLoad(o.ConfigFile)
	if v := o.ModelFile; v != "" {
		global.CFG.Casbin.Model = v
	}
	if v := o.WWWDir; v != "" {
		global.CFG.System.WWW = v
	}
	if v := o.MenuFile; v != "" {
		global.CFG.Menu.Data = v
	}
	global.PrintWithJSON()

	logger.WithContext(ctx).Printf("服务启动，运行模式：%s，版本号：%s，进程号：%d", global.CFG.System.RunMode, o.Version, os.Getpid())

	// 初始化日志模块
	loggerCleanFunc, err := InitLogger()
	if err != nil {
		return nil, err
	}

	// 初始化服务运行监控
	monitorCleanFunc := InitMonitor(ctx)

	// 初始化图形验证码
	redisCli, redisCleanFunc, _ := InitRedisCli()
	global.RdsCli = redisCli

	// 初始化依赖注入器
	injector, injectorCleanFunc, err := BuildInjector(global.CFG)
	if err != nil {
		return nil, err
	}

	// 初始化菜单数据

	fmt.Println(" ---- init data from file global.CFG.Menu.Data == ", global.CFG.Menu.Data)

	if global.CFG.Menu.Enable && global.CFG.Menu.Data != "" {
		err = injector.MenuSrv.InitData(ctx, global.CFG.Menu.Data)

		if err != nil {
			return nil, err
		}
	}

	// 初始化HTTP服务
	httpServerCleanFunc := InitHTTPServer(ctx, injector.Engine)

	return func() {
		redisCleanFunc()
		httpServerCleanFunc()
		injectorCleanFunc()
		monitorCleanFunc()
		loggerCleanFunc()
	}, nil
}

// InitMonitor 初始化服务监控
func InitMonitor(ctx context.Context) func() {
	if c := global.CFG.Monitor; c.Enable {
		// ShutdownCleanup set false to prevent automatically closes on os.Interrupt
		// and close agent manually before service shutting down
		err := agent.Listen(agent.Options{Addr: c.Addr, ConfigDir: c.ConfigDir, ShutdownCleanup: false})
		if err != nil {
			logger.WithContext(ctx).Errorf("Agent monitor error: %s", err.Error())
		}
		return func() {
			agent.Close()
		}
	}
	return func() {}
}

// InitHTTPServer 初始化http服务
func InitHTTPServer(ctx context.Context, handler http.Handler) func() {
	cfg := global.CFG.HTTP
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		logger.WithContext(ctx).Printf("HTTP server is running at %s.", addr)

		var err error
		if cfg.CertFile != "" && cfg.KeyFile != "" {
			srv.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
			err = srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile)
		} else {
			err = srv.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}

	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.ShutdownTimeout))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.WithContext(ctx).Errorf(err.Error())
		}
	}

}

// Run 运行服务
func Run(ctx context.Context, opts ...Option) error {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := Init(ctx, opts...)
	if err != nil {
		return err
	}

EXIT:
	for {
		sig := <-sc
		logger.WithContext(ctx).Infof("接收到信号[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	cleanFunc()
	logger.WithContext(ctx).Infof("服务退出")
	time.Sleep(time.Second)
	os.Exit(state)
	return nil
}
