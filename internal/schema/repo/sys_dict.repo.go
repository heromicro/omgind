package repo

import (
	"context"
	"log"
	"time"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/sysdict"
	"github.com/heromicro/omgind/internal/gen/ent/sysdictitem"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"
)

// DictSet 注入Dict
var DictSet = wire.NewSet(wire.Struct(new(Dict), "*"))

// Dict 字典存储
type Dict struct {
	EntCli *ent.Client
}

func ToSchemaSysDict(dit *ent.SysDict) *schema.Dict {
	item := new(schema.Dict)
	structure.Copy(dit, item)
	if dit.Edges.Items != nil {
		ditems := ToSchemaDictItems(dit.Edges.Items)
		item.Items = ditems
	}
	return item
}

func ToSchemaSysDicts(dits ent.SysDicts) []*schema.Dict {
	list := make([]*schema.Dict, len(dits))
	for i, item := range dits {
		list[i] = ToSchemaSysDict(item)
	}
	return list
}

func (a *Dict) ToEntCreateSysDictInput(sdi *schema.Dict) *ent.CreateSysDictInput {
	createinput := new(ent.CreateSysDictInput)
	structure.Copy(sdi, &createinput)

	return createinput
}

func (a *Dict) ToEntUpdateSysDictInput(sdi *schema.Dict) *ent.UpdateSysDictInput {
	updateinput := new(ent.UpdateSysDictInput)
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

	opt := a.getQueryOption(opts...)

	query := a.EntCli.SysDict.Query().Where(sysdict.DeletedAtIsNil(), sysdict.IsDelEQ(false))
	if v := params.WithItem; v != nil && *v {
		query = query.WithItems(func(sdiq *ent.SysDictItemQuery) {
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
		of := MakeUpOrderField(sysdict.FieldSort, v)
		opt.OrderFields = append(opt.OrderFields, of)
	}

	if v := params.NameCn_Order; v != "" {
		of := MakeUpOrderField(sysdict.FieldNameCn, v)
		opt.OrderFields = append(opt.OrderFields, of)
	}

	if v := params.NameEn_Order; v != "" {
		of := MakeUpOrderField(sysdict.FieldNameEn, v)
		opt.OrderFields = append(opt.OrderFields, of)
	}

	if v := params.IsActive_Order; v != "" {
		of := MakeUpOrderField(sysdict.FieldIsActive, v)
		opt.OrderFields = append(opt.OrderFields, of)
	}

	if len(opt.OrderFields) == 0 {
		opt.OrderFields = append(opt.OrderFields, schema.NewOrderField(sysdict.FieldID, schema.OrderByDESC))
	}

	query = query.Order(ParseOrder(opt.OrderFields)...)

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
	rlist := ent.SysDicts(list)

	qr := &schema.DictQueryResult{
		PageResult: pr,
		Data:       ToSchemaSysDicts(rlist),
	}

	return qr, nil
}

// Query 查询数据
func (a *Dict) QueryItems(ctx context.Context, id string, params schema.DictQueryParam, opts ...schema.DictQueryOptions) (*schema.DictQueryResult, error) {

	log.Println(" ------ ======= --- id ", id)
	log.Println(" ------ ======= -- -params  ", params)

	opt := a.getQueryOption(opts...)

	query := a.EntCli.SysDict.Query()

	if id != "-" {
		query = query.Where(sysdict.IDEQ(id))
	}

	query = query.WithItems(func(sdiq *ent.SysDictItemQuery) {
		sdiq.Order(ent.Asc(sysdictitem.FieldValue)).Select(sysdictitem.FieldID, sysdictitem.FieldValue, sysdictitem.FieldLabel, sysdictitem.FieldIsActive, sysdictitem.FieldMemo, sysdictitem.FieldDictID)
	})

	if id == "-" {
		if v := params.NameCn; v != "" {
			query = query.Where(sysdict.NameCnEQ(v))
		} else if v := params.NameEn; v != "" {
			query = query.Where(sysdict.NameEnEQ(v))
		}
	}

	if len(opt.OrderFields) == 0 {
		opt.OrderFields = append(opt.OrderFields, schema.NewOrderField(sysdict.FieldID, schema.OrderByDESC))
	}

	query = query.Order(ParseOrder(opt.OrderFields)...)

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

	log.Println(" -------  err ===== === ", err)

	if err != nil {
		return nil, errors.WithStack(err)
	}
	rlist := ent.SysDicts(list)

	qr := &schema.DictQueryResult{
		PageResult: pr,
		Data:       ToSchemaSysDicts(rlist),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *Dict) Get(ctx context.Context, id string, opts ...schema.DictQueryOptions) (*schema.Dict, error) {

	query := a.EntCli.SysDict.Query()
	query = query.WithItems(func(sdiq *ent.SysDictItemQuery) {
		sdiq.Order(ent.Asc(sysdictitem.FieldValue)).Select(sysdictitem.FieldID, sysdictitem.FieldValue, sysdictitem.FieldLabel, sysdictitem.FieldIsActive, sysdictitem.FieldMemo)
	})

	dict, err := query.Where(sysdict.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return ToSchemaSysDict(dict), nil
}

// Create 创建数据
func (a *Dict) Create(ctx context.Context, item schema.Dict) (*schema.Dict, error) {
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	iteminput := a.ToEntCreateSysDictInput(&item)
	sysdict, err := a.EntCli.SysDict.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_dict := ToSchemaSysDict(sysdict)
	return sch_dict, nil
}

// Update 更新数据
func (a *Dict) Update(ctx context.Context, id string, item schema.Dict) (*schema.Dict, error) {
	oitem, err := a.EntCli.SysDict.Query().Where(sysdict.IDEQ(id)).Only(ctx)

	if err != nil {
		return nil, err
	}

	item.UpdatedAt = time.Now()
	iteminput := a.ToEntUpdateSysDictInput(&item)
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

// UpdateStatus 更新状态
func (a *Dict) UpdateStatus(ctx context.Context, id string, isActive bool) error {

	_, err1 := a.EntCli.SysDict.UpdateOneID(id).SetIsActive(isActive).Save(ctx)

	return errors.WithStack(err1)
}
