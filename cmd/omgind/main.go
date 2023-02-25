/*
Package main omgind

Swagger 文档规则请参考：https://github.com/swaggo/swag#declarative-comments-format

使用方式：

	go get -u github.com/swaggo/swag/cmd/swag
	swag init --generalInfo ./cmd/omgind/main.go --output ./internal/app/swagger
*/
package main

import (
	"github.com/heromicro/omgind/cmd/omgind/enter"
	"github.com/heromicro/omgind/pkg/logger"
)

// VERSION 版本号，可以通过编译的方式指定版本号：go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "7.0.0"

//	@title						omgind
//	@version					7.0.0
//	@description				RBAC scaffolding based on GIN + ENT + CASBIN + WIRE.
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@schemes					http https
//	@basePath					/
//	@contact.name				heromicro
//	@contact.email				wky@foal.com
func main() {
	logger.SetVersion(VERSION)
	enter.Execute()
}
