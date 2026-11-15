package repo

import (
	"context"
	"time"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/sysannex"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"
)

// SysAnnexSet 注入SysAnnex
var SysAnnexSet = wire.NewSet(wire.Struct(new(SysAnnex), "*"))

// SysAnnex 存储
type SysAnnex struct {
	EntCli *mainent.Client
}

// ToSchemaSysAnnex 转换为
func ToSchemaSysAnnex(et *mainent.SysAnnex) *schema.SysAnnex {
	item := new(schema.SysAnnex)
	structure.Copy(et, item)
	return item
}

func ToSchemaSysAnnexes(ets mainent.SysAnnexes) []*schema.SysAnnex {
	list := make([]*schema.SysAnnex, len(ets))
	for i, item := range ets {
		list[i] = ToSchemaSysAnnex(item)
	}
	return list
}

func ToEntCreateSysAnnexInput(sch *schema.SysAnnex) *mainent.CreateSysAnnexInput {
	createinput := new(mainent.CreateSysAnnexInput)
	structure.Copy(sch, &createinput)

	return createinput
}

func ToEntUpdateSysAnnexInput(sch *schema.SysAnnex) *mainent.UpdateSysAnnexInput {
	updateinput := new(mainent.UpdateSysAnnexInput)
	structure.Copy(sch, &updateinput)

	return updateinput
}

func (a *SysAnnex) getQueryOption(opts ...schema.SysAnnexQueryOptions) schema.SysAnnexQueryOptions {
	var opt schema.SysAnnexQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *SysAnnex) Query(ctx context.Context, params schema.SysAnnexQueryParam, opts ...schema.SysAnnexQueryOptions) (*schema.SysAnnexQueryResult, error) {
	// opt := a.getQueryOption(opts...)

	query := a.EntCli.SysAnnex.Query()

	query = query.Where(sysannex.DeletedAtIsNil())
	// TODO: 查询条件

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.SysAnnexQueryResult{PageResult: pr}, nil
	}

	query = query.Order(mainent.Asc(sysannex.FieldSort))

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.SysAnnexQueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.SysAnnexQueryResult{
		PageResult: pr,
		Data:       ToSchemaSysAnnexes(list),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *SysAnnex) Get(ctx context.Context, id string, opts ...schema.SysAnnexQueryOptions) (*schema.SysAnnex, error) {

	r_sysannex, err := a.EntCli.SysAnnex.Query().Where(sysannex.IDEQ(id)).Only(ctx)
	if err != nil {
		if mainent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return ToSchemaSysAnnex(r_sysannex), nil
}

// View 查询指定数据
func (a *SysAnnex) View(ctx context.Context, id string, opts ...schema.SysAnnexQueryOptions) (*schema.SysAnnex, error) {

	r_sysannex, err := a.EntCli.SysAnnex.Query().Where(sysannex.IDEQ(id)).Only(ctx)
	if err != nil {
		if mainent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return ToSchemaSysAnnex(r_sysannex), nil
}

// Create 创建数据
func (a *SysAnnex) Create(ctx context.Context, item schema.SysAnnex) (*schema.SysAnnex, error) {

	iteminput := ToEntCreateSysAnnexInput(&item)
	r_sysannex, err := a.EntCli.SysAnnex.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_sysannex := ToSchemaSysAnnex(r_sysannex)
	return sch_sysannex, nil
}

// Update 更新数据
func (a *SysAnnex) Update(ctx context.Context, id string, item schema.SysAnnex) (*schema.SysAnnex, error) {

	oitem, err := a.EntCli.SysAnnex.Query().Where(sysannex.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	iteminput := ToEntUpdateSysAnnexInput(&item)

	r_sysannex, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	if err != nil {
		return nil, err
	}
	sch_sysannex := ToSchemaSysAnnex(r_sysannex)

	return sch_sysannex, nil
}

// Delete 删除数据
func (a *SysAnnex) Delete(ctx context.Context, id string) error {

	r_sysannex, err := a.EntCli.SysAnnex.Query().Where(sysannex.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}

	_, err1 := r_sysannex.Update().SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)

	return errors.WithStack(err1)
}

// UpdateActive 更新状态
func (a *SysAnnex) UpdateActive(ctx context.Context, id string, active bool) error {

	_, err1 := a.EntCli.SysAnnex.Update().Where(sysannex.IDEQ(id)).SetIsActive(active).Save(ctx)

	return errors.WithStack(err1)
}
