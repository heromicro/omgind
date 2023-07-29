package repo

import (
	"context"
	"time"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/sysdict"
	"github.com/heromicro/omgind/internal/gen/mainent/sysdictitem"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"
)

// DictSet 注入Dict
var DictSet = wire.NewSet(wire.Struct(new(Dict), "*"))

// Dict 字典存储
type Dict struct {
	EntCli *mainent.Client
}

func ToSchemaSysDict(dit *mainent.SysDict) *schema.Dict {
	item := new(schema.Dict)
	structure.Copy(dit, item)
	if dit.Edges.Items != nil {
		ditems := ToSchemaDictItems(dit.Edges.Items)
		item.Items = ditems
	}
	return item
}

func ToSchemaSysDicts(dits mainent.SysDicts) []*schema.Dict {
	list := make([]*schema.Dict, len(dits))
	for i, item := range dits {
		list[i] = ToSchemaSysDict(item)
	}
	return list
}

func ToEntCreateSysDictInput(sdi *schema.Dict) *mainent.CreateSysDictInput {
	createinput := new(mainent.CreateSysDictInput)
	structure.Copy(sdi, &createinput)

	return createinput
}

func ToEntUpdateSysDictInput(sdi *schema.Dict) *mainent.UpdateSysDictInput {
	updateinput := new(mainent.UpdateSysDictInput)
	structure.Copy(sdi, &updateinput)

	return updateinput
}

func (a *Dict) getQueryOption(opts ...schema.DictQueryOptions) schema.DictQueryOptions {
	var opt schema.DictQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Dict) Query(ctx context.Context, params schema.DictQueryParam, opts ...schema.DictQueryOptions) (*schema.DictQueryResult, error) {

	// opt := a.getQueryOption(opts...)

	query := a.EntCli.SysDict.Query().Where(sysdict.DeletedAtIsNil(), sysdict.IsDelEQ(false))
	if v := params.WithItem; v != nil && *v {
		query = query.WithItems(func(sdiq *mainent.SysDictItemQuery) {
			sdiq.Where(sysdictitem.IsDelEQ(false)).Select(sysdictitem.FieldID, sysdictitem.FieldValue, sysdictitem.FieldLabel, sysdictitem.FieldIsActive)
		})
	}
	if v := params.IDs; len(v) > 0 {
		query = query.Where(sysdict.IDIn(v...))
	}

	if v := params.NameCn; v != "" {
		query = query.Where(sysdict.NameCnEQ(v))
	}

	if v := params.NameEn; v != "" {
		query = query.Where(sysdict.NameEnEQ(v))
	}

	if v := params.DictKey; v != "" {
		query = query.Where(sysdict.DictKeyEQ(v))
	}

	if v := params.IsActive; v != nil {
		query = query.Where(sysdict.IsActiveEQ(*v))
	}

	if v := params.QueryValue; v != "" {
		query = query.Where(sysdict.Or(sysdict.NameCnContains(v), sysdict.NameEnContains(v), sysdict.MemoContains(v)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.DictQueryResult{PageResult: pr}, nil
	}

	if v := params.Sort_Order; v != "" {
		query = query.Order(sysdict.BySort(OrderDirection(v)))
	}

	if v := params.NameCn_Order; v != "" {
		query = query.Order(sysdict.ByNameCn(OrderDirection(v)))
	}

	if v := params.NameEn_Order; v != "" {
		query = query.Order(sysdict.ByNameEn(OrderDirection(v)))
	}

	if v := params.IsActive_Order; v != "" {
		query = query.Order(sysdict.ByIsActive(OrderDirection(v)))
	}

	query = query.Order(sysdict.ByID(OrderDirection("desc")))

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.DictQueryResult{PageResult: pr}, nil
	}

	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err := query.All(ctx)
	// log.Println(" -------  err ===== === ", err)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	rlist := mainent.SysDicts(list)

	qr := &schema.DictQueryResult{
		PageResult: pr,
		Data:       ToSchemaSysDicts(rlist),
	}

	return qr, nil
}

// Query 查询数据
func (a *Dict) QueryItems(ctx context.Context, id string, params schema.DictQueryParam, opts ...schema.DictQueryOptions) (*schema.DictQueryResult, error) {

	// log.Println(" ------ ======= --- id ", id)
	// log.Println(" ------ ======= -- -params  ", params)

	// opt := a.getQueryOption(opts...)

	query := a.EntCli.SysDict.Query()

	if id != "-" {
		query = query.Where(sysdict.IDEQ(id))
	}

	query = query.WithItems(func(sdiq *mainent.SysDictItemQuery) {
		sdiq.Order(mainent.Asc(sysdictitem.FieldValue)).Select(sysdictitem.FieldID, sysdictitem.FieldValue, sysdictitem.FieldLabel, sysdictitem.FieldIsActive, sysdictitem.FieldMemo, sysdictitem.FieldDictID)
	})

	if id == "-" {
		if v := params.DictKey; v != "" {
			query = query.Where(sysdict.DictKeyEQ(v))
		}
	}

	query = query.Order(sysdict.ByID(OrderDirection("desc")))

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.DictQueryResult{PageResult: pr}, nil
	}

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.DictQueryResult{PageResult: pr}, nil
	}

	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err := query.All(ctx)

	// log.Println(" -------  err ===== === ", err)

	if err != nil {
		return nil, errors.WithStack(err)
	}
	rlist := mainent.SysDicts(list)

	qr := &schema.DictQueryResult{
		PageResult: pr,
		Data:       ToSchemaSysDicts(rlist),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *Dict) Get(ctx context.Context, id string, opts ...schema.DictQueryOptions) (*schema.Dict, error) {

	query := a.EntCli.SysDict.Query()
	query = query.WithItems(func(sdiq *mainent.SysDictItemQuery) {
		sdiq.Order(mainent.Asc(sysdictitem.FieldValue)).Select(sysdictitem.FieldID, sysdictitem.FieldValue, sysdictitem.FieldLabel, sysdictitem.FieldIsActive, sysdictitem.FieldMemo, sysdictitem.FieldDictID)
	})

	dict, err := query.Where(sysdict.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return ToSchemaSysDict(dict), nil
}

// Update 更新数据
func (a *Dict) Update(ctx context.Context, id string, item schema.Dict) (*schema.Dict, error) {
	oitem, err := a.EntCli.SysDict.Query().Where(sysdict.IDEQ(id)).Only(ctx)

	if err != nil {
		return nil, err
	}

	iteminput := ToEntUpdateSysDictInput(&item)
	iteminput.UpdatedAt = nil

	dict, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	if err != nil {
		return nil, err
	}

	sch_dict := ToSchemaSysDict(dict)

	return sch_dict, nil
}

// Delete 删除数据
func (a *Dict) Delete(ctx context.Context, id string) error {
	dict, err := a.EntCli.SysDict.Query().Where(sysdict.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}

	_, err1 := dict.Update().SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)

	return errors.WithStack(err1)
}

// UpdateActive 更新状态
func (a *Dict) UpdateActive(ctx context.Context, id string, isActive bool) error {

	_, err1 := a.EntCli.SysDict.UpdateOneID(id).SetIsActive(isActive).Save(ctx)

	return errors.WithStack(err1)
}
