package repo

import (
	"context"
	"time"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/xxxdemo"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"
)

var DemoSet = wire.NewSet(wire.Struct(new(Demo), "*"))

type Demo struct {
	EntCli *mainent.Client
}

// ToSchemaDemo 转换为
func ToSchemaDemo(xd *mainent.XxxDemo) *schema.Demo {
	item := new(schema.Demo)
	structure.Copy(xd, item)
	return item
}

func ToSchemaDemos(xds mainent.XxxDemos) []*schema.Demo {
	list := make([]*schema.Demo, len(xds))
	for i, item := range xds {
		list[i] = ToSchemaDemo(item)
	}
	return list
}

func ToEntCreateDemoInput(sch *schema.Demo) *mainent.CreateXxxDemoInput {
	createinput := new(mainent.CreateXxxDemoInput)
	structure.Copy(sch, &createinput)

	return createinput
}

func ToEntUpdateDemoInput(sch *schema.Demo) *mainent.UpdateXxxDemoInput {
	updateinput := new(mainent.UpdateXxxDemoInput)
	structure.Copy(sch, &updateinput)

	return updateinput
}

func (a *Demo) getQueryOption(opts ...schema.DemoQueryOptions) schema.DemoQueryOptions {
	var opt schema.DemoQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Demo) Query(ctx context.Context, params schema.DemoQueryParam, opts ...schema.DemoQueryOptions) (*schema.DemoQueryResult, error) {

	// opt := a.getQueryOption(opts...)

	query := a.EntCli.XxxDemo.Query().Where(xxxdemo.DeletedAtIsNil())

	if v := params.Code; v != "" {
		query = query.Where(xxxdemo.CodeContains(v))
	}

	if v := params.Name; v != "" {
		query = query.Where(xxxdemo.NameContains(v))
	}

	if v := params.QueryValue; v != "" {
		query = query.Where(xxxdemo.Or(xxxdemo.CodeContains(v), xxxdemo.NameContains(v)))
	}

	if v := params.CreatedAt_St; v != nil {
		st := time.UnixMilli(*v)
		query = query.Where(xxxdemo.CreatedAtGTE(st))
	}

	if v := params.CreatedAt_Ed; v != nil {
		et := time.UnixMilli(*v)
		query = query.Where(xxxdemo.CreatedAtLTE(et))
	}

	if v := params.Sort_St; v != nil {
		query = query.Where(xxxdemo.SortGTE(*v))
	}

	if v := params.Sort_Ed; v != nil {
		query = query.Where(xxxdemo.SortLTE(*v))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.DemoQueryResult{PageResult: pr}, nil
	}

	// order field
	if v := params.Sort_Order; v != "" {
		query = query.Order(xxxdemo.BySort(OrderDirection(v)))
	}

	if v := params.Code_Order; v != "" {
		query = query.Order(xxxdemo.ByCode(OrderDirection(v)))
	}

	if v := params.CreatedAt_Order; v != nil {
		query = query.Order(xxxdemo.ByCreatedAt(OrderDirection(*v)))
	}

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.DemoQueryResult{PageResult: pr}, nil
	}

	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}
	rlist := mainent.XxxDemos(list)

	qr := &schema.DemoQueryResult{
		PageResult: pr,
		Data:       ToSchemaDemos(rlist),
	}
	return qr, nil
}

// Get 查询指定数据
func (a *Demo) Get(ctx context.Context, id string, opts ...schema.DemoQueryOptions) (*schema.Demo, error) {

	xxxdemo, err := a.EntCli.XxxDemo.Query().Where(xxxdemo.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return ToSchemaDemo(xxxdemo), nil
}

// Create 创建数据
func (a *Demo) Create(ctx context.Context, item schema.Demo) (*schema.Demo, error) {

	iteminput := ToEntCreateDemoInput(&item)
	iteminput.CreatedAt = nil
	iteminput.UpdatedAt = nil
	xxxdemo, err := a.EntCli.XxxDemo.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_demo := ToSchemaDemo(xxxdemo)
	return sch_demo, nil
}

// Update 更新数据
func (a *Demo) Update(ctx context.Context, id string, item schema.Demo) (*schema.Demo, error) {

	oitem, err := a.EntCli.XxxDemo.Query().Where(xxxdemo.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	iteminput := ToEntUpdateDemoInput(&item)
	iteminput.UpdatedAt = nil

	xxxdemo, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	if err != nil {
		return nil, err
	}
	sch_demo := ToSchemaDemo(xxxdemo)

	return sch_demo, nil
}

// Delete 删除数据
func (a *Demo) Delete(ctx context.Context, id string) error {

	xxxdemo, err := a.EntCli.XxxDemo.Query().Where(xxxdemo.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}

	_, err1 := xxxdemo.Update().SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)

	return errors.WithStack(err1)
}

// UpdateActive 更新状态
func (a *Demo) UpdateActive(ctx context.Context, id string, isActive bool) error {

	_, err1 := a.EntCli.XxxDemo.Update().Where(xxxdemo.IDEQ(id)).SetIsActive(isActive).Save(ctx)

	return errors.WithStack(err1)
}
