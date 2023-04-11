package repo

import (
	"context"
	"time"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/sysdictitem"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"
)

// DictItemSet 注入DictItem
var DictItemSet = wire.NewSet(wire.Struct(new(DictItem), "*"))

// DictItem 字典项存储
type DictItem struct {
	EntCli *ent.Client
}

func ToSchemaDictItem(dit *ent.SysDictItem) *schema.DictItem {
	item := new(schema.DictItem)
	structure.Copy(dit, item)
	return item
}

func ToSchemaDictItems(dits ent.SysDictItems) []*schema.DictItem {
	list := make([]*schema.DictItem, len(dits))
	for i, item := range dits {
		list[i] = ToSchemaDictItem(item)
	}
	return list
}

func ToEntCreateSysDictItemInput(sdi *schema.DictItem) *ent.CreateSysDictItemInput {
	createinput := new(ent.CreateSysDictItemInput)
	structure.Copy(sdi, &createinput)

	return createinput
}

func ToEntUpdateSysDictItemInput(sdi *schema.DictItem) *ent.UpdateSysDictItemInput {
	updateinput := new(ent.UpdateSysDictItemInput)
	structure.Copy(sdi, &updateinput)

	return updateinput
}

func (a *DictItem) getQueryOption(opts ...schema.DictItemQueryOptions) schema.DictItemQueryOptions {
	var opt schema.DictItemQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *DictItem) Query(ctx context.Context, params schema.DictItemQueryParam, opts ...schema.DictItemQueryOptions) (*schema.DictItemQueryResult, error) {
	// opt := a.getQueryOption(opts...)

	query := a.EntCli.SysDictItem.Query().Where(sysdictitem.DeletedAtIsNil())

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if v := params.IDs; len(v) > 0 {
		query = query.Where(sysdictitem.IDIn(v...))
	}

	if v := params.DictID; v != "" {
		query = query.Where(sysdictitem.DictIDEQ(v))
	}

	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.DictItemQueryResult{PageResult: pr}, nil
	}

	// order field

	query = query.Order(sysdictitem.ByID(OrderBy("desc")))

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.DictItemQueryResult{PageResult: pr}, nil
	}

	//分页
	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}
	rlist := ent.SysDictItems(list)

	qr := &schema.DictItemQueryResult{
		PageResult: pr,
		Data:       ToSchemaDictItems(rlist),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *DictItem) Get(ctx context.Context, id string, opts ...schema.DictItemQueryOptions) (*schema.DictItem, error) {
	dictItem, err := a.EntCli.SysDictItem.Query().Where(sysdictitem.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return ToSchemaDictItem(dictItem), nil
}

// Create 创建数据
func (a *DictItem) Create(ctx context.Context, item schema.DictItem) (*schema.DictItem, error) {

	itemInput := ToEntCreateSysDictItemInput(&item)

	dictitem, err := a.EntCli.SysDictItem.Create().SetInput(*itemInput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_dictitem := ToSchemaDictItem(dictitem)
	return sch_dictitem, nil
}

// Update 更新数据
func (a *DictItem) Update(ctx context.Context, id string, item schema.DictItem) (*schema.DictItem, error) {

	oitem, err := a.EntCli.SysDictItem.Query().Where(sysdictitem.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	itemInput := ToEntUpdateSysDictItemInput(&item)

	dictitem, err := oitem.Update().SetInput(*itemInput).Save(ctx)
	if err != nil {
		return nil, err
	}
	sch_dictitem := ToSchemaDictItem(dictitem)

	return sch_dictitem, nil
}

// Delete 删除数据
func (a *DictItem) Delete(ctx context.Context, id string) error {
	dictitem, err := a.EntCli.SysDictItem.Query().Where(sysdictitem.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}
	_, err1 := dictitem.Update().SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)
	return errors.WithStack(err1)
}
