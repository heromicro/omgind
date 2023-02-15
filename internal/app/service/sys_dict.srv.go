package service

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/sysdictitem"
	"github.com/heromicro/omgind/internal/schema/repo"
	"github.com/heromicro/omgind/pkg/errors"
	uid "github.com/heromicro/omgind/pkg/helper/uid/ulid"
)

var DictSet = wire.NewSet(wire.Struct(new(Dict), "*"))

// Demo 示例程序
type Dict struct {
	DictRepo     *repo.Dict
	DictItemRepo *repo.DictItem
}

// Query 查询数据
func (a *Dict) Query(ctx context.Context, params schema.DictQueryParam, opts ...schema.DictQueryOptions) (*schema.DictQueryResult, error) {
	return a.DictRepo.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *Dict) Get(ctx context.Context, id string, opts ...schema.DictQueryOptions) (*schema.Dict, error) {
	item, err := a.DictRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}
	ditems, err := a.QueryItems(ctx, id)
	if err != nil {
		return nil, err
	}
	item.Items = ditems

	return item, nil
}

func (d Dict) QueryItems(ctx context.Context, id string) (schema.DictItems, error) {

	result, err := d.DictItemRepo.Query(ctx, schema.DictItemQueryParam{
		DictID: id,
	})

	if err != nil {
		return nil, err
	} else if len(result.Data) == 0 {
		return nil, nil
	}

	return result.Data, nil
}

func (d *Dict) checkName(ctx context.Context, item schema.Dict) error {
	// TODO:: need optimization
	result1, err1 := d.DictRepo.Query(ctx, schema.DictQueryParam{
		PaginationParam: schema.PaginationParam{
			OnlyCount: true,
		},
		NameEn: item.NameEn,
	})

	if err1 != nil {
		return nil
	} else if result1.PageResult.Total > 0 {
		return errors.New400Response("字典名称" + item.NameEn + "已经存在")
	}

	result2, err2 := d.DictRepo.Query(ctx, schema.DictQueryParam{
		PaginationParam: schema.PaginationParam{
			OnlyCount: true,
		},
		NameCn: item.NameCn,
	})

	if err2 != nil {
		return nil
	} else if result2.PageResult.Total > 0 {
		return errors.New400Response("字典名称" + item.NameCn + "已经存在")
	}

	return nil
}

// Create 创建数据
func (a *Dict) Create(ctx context.Context, item schema.Dict) (*schema.IDResult, error) {
	if err := a.checkName(ctx, item); err != nil {
		return nil, err
	}
	item.ID = uid.MustString()

	err := repo.WithTx(ctx, a.DictRepo.EntCli, func(tx *ent.Tx) error {

		dictInput := a.DictRepo.ToEntCreateSysDictInput(&item)
		dictInput.CreatedAt = nil
		dictInput.UpdatedAt = nil

		adict, err := tx.SysDict.Create().SetInput(*dictInput).Save(ctx)

		if err != nil {
			return err
		}

		for _, itm := range item.Items {
			itm.DictID = adict.ID
			ipt := a.DictItemRepo.ToEntCreateSysDictItemInput(itm)
			ipt.CreatedAt = nil
			ipt.UpdatedAt = nil

			_, err := tx.SysDictItem.Create().SetInput(*ipt).Save(ctx)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

// Update 更新数据
func (a *Dict) Update(ctx context.Context, id string, item schema.Dict) error {

	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	} else if oldItem.NameEn != item.NameEn || oldItem.NameCn != item.NameCn {
		if err := a.checkName(ctx, item); err != nil {
			return err
		}
	}

	item.ID = oldItem.ID
	item.Creator = oldItem.Creator
	item.CreatedAt = oldItem.CreatedAt

	addItems, delItems, updateItems := a.compareDictItems(ctx, oldItem.Items, item.Items)

	err1 := repo.WithTx(ctx, a.DictRepo.EntCli, func(tx *ent.Tx) error {

		// 添加
		for _, itm := range addItems {

			itm.DictID = id
			inpt := a.DictItemRepo.ToEntCreateSysDictItemInput(itm)
			itm.DictID = id

			inpt.CreatedAt = nil
			inpt.UpdatedAt = nil
			_, err := tx.SysDictItem.Create().SetInput(*inpt).Save(ctx)
			if err != nil {
				if err := tx.Rollback(); err != nil {
					return err
				}
				return err
			}
		}

		// 删除
		for _, itm := range delItems {
			_, err := tx.SysDictItem.Update().Where(sysdictitem.IDEQ(itm.ID)).SetDeletedAt(time.Now()).Save(ctx)
			if err != nil {
				if err := tx.Rollback(); err != nil {
					//
					return err
				}
				return err
			}
		}

		// 更新
		for _, itm := range updateItems {

			inpt := a.DictItemRepo.ToEntUpdateSysDictItemInput(itm)
			inpt.UpdatedAt = nil
			_, err := tx.SysDictItem.UpdateOneID(itm.ID).SetInput(*inpt).Save(ctx)
			if err != nil {
				if err := tx.Rollback(); err != nil {
					return err
				}
				return err
			}
		}

		dict_input := a.DictRepo.ToEntUpdateSysDictInput(&item)
		dict_input.UpdatedAt = nil
		_, err := tx.SysDict.UpdateOneID(id).SetInput(*dict_input).Save(ctx)

		if err != nil {
			if err := tx.Rollback(); err != nil {
				return err
			}
			return err
		}
		return nil
	})

	return err1
}

func (a *Dict) compareDictItems(ctx context.Context, oldItems, newItems schema.DictItems) (addList, delList, updateList schema.DictItems) {

	mOldItems := oldItems.ToMap()
	mNewItems := newItems.ToMap()

	//log.Printf(" ------ 000--newItems: %+v ", newItems)
	//log.Printf(" ------ 000--oldItems: %+v ", newItems)

	for k, item := range mNewItems {
		if _, ok := mOldItems[k]; ok {
			updateList = append(updateList, item)
			delete(mOldItems, k)
			continue
		}
		addList = append(addList, item)
	}
	for _, item := range mOldItems {
		delList = append(delList, item)
	}
	return
}

// Delete 删除数据
func (a *Dict) Delete(ctx context.Context, id string) error {

	oldItem, err := a.DictRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	return a.DictRepo.Delete(ctx, id)
}

// Delete 删除数据
func (a *Dict) DeleteS(ctx context.Context, id string) error {
	oldItem, err := a.DictRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	_, err1 := a.DictRepo.Update(ctx, id, *oldItem)
	return err1
}

func (a *Dict) UpdateStatus(ctx context.Context, id string, isActive bool) error {

	oldItem, err := a.DictRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.DictRepo.UpdateStatus(ctx, id, isActive)
}
