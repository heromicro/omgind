package repo

import (
	"context"

	"github.com/pkg/errors"

	"github.com/heromicro/omgind/internal/gen/ent"
)

// TransSet 注入Trans
// var TransSet = wire.NewSet(wire.Struct(new(Trans), "*"))

// Trans 事务管理
// type Trans struct {
// 	EntCli *ent.Client
// }

//
//// Exec 执行事务
//func (a *Trans) Exec(ctx context.Context, fn func(context.Context) error) error {
//
//	if _, ok := contextx.FromTrans(ctx); ok {
//		return fn(ctx)
//	}
//	a.EntCli.Tx()
//	return a.DB.Transaction(func(db *gorm.DB) error {
//		return fn(contextx.NewTrans(ctx, db))
//	})
//}

// WithTx best Practices, reusable function that runs callbacks in a transaction
func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {

	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}
