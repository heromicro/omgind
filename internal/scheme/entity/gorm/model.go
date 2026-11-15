package entity

import (
	"context"
	"github.com/jinzhu/gorm"
)

// GetDBWithModel 为一次操作创建独立会话，不修改共享连接。
func GetDBWithModel(ctx context.Context, db *gorm.DB, model interface{}) *gorm.DB {
	return db.New().Set("context", ctx).Model(model)
}
