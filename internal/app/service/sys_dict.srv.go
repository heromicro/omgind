package service

import (
	"context"
	"time"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/sysdict"
	"github.com/heromicro/omgind/internal/gen/mainent/sysdictitem"
	"github.com/heromicro/omgind/internal/scheme/repo"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/str"
)

var DictSet = wire.NewSet(wire.Struct(new(Dict), "*"))

// Demo 示例程序
type Dict struct {
	EntCli *mainent.Client

	DictRepo *repo.Dict
}

// Query 查询数据
func (a *Dict) Query(ctx context.Context, params schema.DictQueryParam, opts ...schema.DictQueryOptions) (*schema.DictQueryResult, error) {
	return a.DictRepo.Query(ctx, params, opts...)
}

// QueryOne 查询数据
func (a *Dict) QueryItems(ctx context.Context, id string, params schema.DictQueryParam, opts ...schema.DictQueryOptions) (*schema.DictQueryResult, error) {
	return a.DictRepo.QueryItems(ctx, id, params, opts...)
}

// Get 查询指定数据
func (a *Dict) Get(ctx context.Context, id string, opts ...schema.DictQueryOptions) (*schema.Dict, error) {

	item, err := a.DictRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// View 查询指定数据
func (a *Dict) View(ctx context.Context, id string, opts ...schema.DictQueryOptions) (*schema.Dict, error) {

	query := a.EntCli.SysDict.Query()
	query = query.WithItems(func(sdiq *mainent.SysDictItemQuery) {
		sdiq.Order(mainent.Asc(sysdictitem.FieldValue)).Select(sysdictitem.FieldID, sysdictitem.FieldValue, sysdictitem.FieldLabel, sysdictitem.FieldIsActive, sysdictitem.FieldMemo, sysdictitem.FieldDictID, sysdictitem.FieldSort)
	})

	dict, err := query.Where(sysdict.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return repo.ToSchemaSysDict(dict), nil
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

	item.NameEn = str.RemoveDuplicateBlank(item.NameEn)
	item.NameCn = str.RemoveDuplicateBlank(item.NameCn)

	err := repo.WithTx(ctx, a.DictRepo.EntCli, func(tx *mainent.Tx) error {

		dictInput := repo.ToEntCreateSysDictInput(&item)
		dictInput.CreatedAt = nil
		dictInput.UpdatedAt = nil

		adict, err := tx.SysDict.Create().SetInput(*dictInput).Save(ctx)

		if err != nil {
			return err
		}

		for _, itm := range item.Items {
			itm.DictID = adict.ID
			ipt := repo.ToEntCreateSysDictItemInput(itm)
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
func (a *Dict) Update(ctx context.Context, id string, item schema.Dict) (*schema.Dict, error) {

	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if oldItem == nil {
		return nil, errors.ErrNotFound
	} else if oldItem.NameEn != item.NameEn || oldItem.NameCn != item.NameCn {
		if err := a.checkName(ctx, item); err != nil {
			return nil, err
		}
	}

	item.ID = oldItem.ID

	item.CreatedAt = oldItem.CreatedAt

	addItems, delItems, updateItems := a.compareDictItems(ctx, oldItem.Items, item.Items)

	err1 := repo.WithTx(ctx, a.DictRepo.EntCli, func(tx *mainent.Tx) error {

		// 添加
		for _, itm := range addItems {

			itm.DictID = id
			inpt := repo.ToEntCreateSysDictItemInput(itm)
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
			_, err := tx.SysDictItem.Update().Where(sysdictitem.IDEQ(itm.ID)).SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)
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

			inpt := repo.ToEntUpdateSysDictItemInput(itm)
			inpt.UpdatedAt = nil
			_, err := tx.SysDictItem.UpdateOneID(itm.ID).SetInput(*inpt).Save(ctx)
			if err != nil {
				if err := tx.Rollback(); err != nil {
					return err
				}
				return err
			}
		}

		dict_input := repo.ToEntUpdateSysDictInput(&item)
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

	if err1 != nil {
		return nil, err1
	}

	nitem, err := a.Get(ctx, id)

	return nitem, err
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

func (a *Dict) UpdateActive(ctx context.Context, id string, isActive bool) error {

	oldItem, err := a.DictRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.DictRepo.UpdateActive(ctx, id, isActive)
}
