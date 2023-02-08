package app

import (
	"github.com/heromicro/omgind/internal/app/middleware"
	"github.com/heromicro/omgind/internal/app/middleware/healthchek"
	"github.com/heromicro/omgind/internal/router"
	"github.com/heromicro/omgind/pkg/global"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitGinEngine 初始化gin引擎
func InitGinEngine(r router.IRouter) *gin.Engine {
	gin.SetMode(global.CFG.System.RunMode)

	app := gin.New()
	app.NoMethod(middleware.NoMethodHandler())
	app.NoRoute(middleware.NoRouteHandler())

	prefixes := r.Prefixes()

	app.Use(healthchek.Default())

	// Trace ID
	app.Use(middleware.TraceMiddleware(middleware.AllowPathPrefixNoSkipper(prefixes...)))

	// Copy body
	app.Use(middleware.CopyBodyMiddleware(middleware.AllowPathPrefixNoSkipper(prefixes...)))

	// Access logger
	app.Use(middleware.LoggerMiddleware(middleware.AllowPathPrefixNoSkipper(prefixes...)))

	// Recover
	app.Use(middleware.RecoveryMiddleware())

	// CORS
	if global.CFG.CORS.Enable {
		app.Use(middleware.CORSMiddleware())
	}

	// GZIP
	if global.CFG.GZIP.Enable {
		app.Use(gzip.Gzip(gzip.BestCompression,
			gzip.WithExcludedExtensions(global.CFG.GZIP.ExcludedExtentions),
			gzip.WithExcludedPaths(global.CFG.GZIP.ExcludedPaths),
		))
	}

	// Router register
	r.Register(app)

	// Swagger
	if global.CFG.System.Swagger {
		app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Website
	if dir := global.CFG.System.WWW; dir != "" {
		app.Use(middleware.WWWMiddleware(dir, middleware.AllowPathPrefixSkipper(prefixes...)))
	}

	return app
}
