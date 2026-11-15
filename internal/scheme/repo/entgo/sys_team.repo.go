package repo

import (
	"context"
	"time"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/systeam"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"

	"github.com/google/wire"
)

// SysTeamSet 注入SysTeam
var SysTeamSet = wire.NewSet(wire.Struct(new(SysTeam), "*"))

// SysTeam 地址管理存储
type SysTeam struct {
	EntCli *mainent.Client
}

// ToSchemaSysTeam 转换为
func ToSchemaSysTeam(et *mainent.SysTeam) *schema.SysTeam {
	item := new(schema.SysTeam)
	structure.Copy(et, item)
	return item
}

func ToSchemaSysTeams(ets mainent.SysTeams) []*schema.SysTeam {
	list := make([]*schema.SysTeam, len(ets))
	for i, item := range ets {
		list[i] = ToSchemaSysTeam(item)
	}
	return list
}

func ToEntCreateSysTeamInput(sch *schema.SysTeam) *mainent.CreateSysTeamInput {
	createinput := new(mainent.CreateSysTeamInput)
	structure.Copy(sch, &createinput)

	return createinput
}

func ToEntUpdateSysTeamInput(sch *schema.SysTeam) *mainent.UpdateSysTeamInput {
	updateinput := new(mainent.UpdateSysTeamInput)
	structure.Copy(sch, &updateinput)

	return updateinput
}

func (a *SysTeam) getQueryOption(opts ...schema.SysTeamQueryOptions) schema.SysTeamQueryOptions {
	var opt schema.SysTeamQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *SysTeam) Query(ctx context.Context, params schema.SysTeamQueryParam, opts ...schema.SysTeamQueryOptions) (*schema.SysTeamQueryResult, error) {
	// opt := a.getQueryOption(opts...)

	query := a.EntCli.SysTeam.Query()

	query = query.Where(systeam.DeletedAtIsNil())
	// TODO: 查询条件

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.SysTeamQueryResult{PageResult: pr}, nil
	}

	query = query.Order(mainent.Asc(systeam.FieldSort))

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.SysTeamQueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.SysTeamQueryResult{
		PageResult: pr,
		Data:       ToSchemaSysTeams(list),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *SysTeam) Get(ctx context.Context, id string, opts ...schema.SysTeamQueryOptions) (*schema.SysTeam, error) {

	r_systeam, err := a.EntCli.SysTeam.Query().Where(systeam.IDEQ(id)).Only(ctx)
	if err != nil {
		if mainent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return ToSchemaSysTeam(r_systeam), nil
}

// View 查询指定数据
func (a *SysTeam) View(ctx context.Context, id string, opts ...schema.SysTeamQueryOptions) (*schema.SysTeam, error) {

	r_systeam, err := a.EntCli.SysTeam.Query().Where(systeam.IDEQ(id)).Only(ctx)
	if err != nil {
		if mainent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return ToSchemaSysTeam(r_systeam), nil
}

// Create 创建数据
func (a *SysTeam) Create(ctx context.Context, item schema.SysTeam) (*schema.SysTeam, error) {

	iteminput := ToEntCreateSysTeamInput(&item)
	r_systeam, err := a.EntCli.SysTeam.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_systeam := ToSchemaSysTeam(r_systeam)
	return sch_systeam, nil
}

// Update 更新数据
func (a *SysTeam) Update(ctx context.Context, id string, item schema.SysTeam) (*schema.SysTeam, error) {

	oitem, err := a.EntCli.SysTeam.Query().Where(systeam.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	iteminput := ToEntUpdateSysTeamInput(&item)

	r_systeam, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	if err != nil {
		return nil, err
	}
	sch_systeam := ToSchemaSysTeam(r_systeam)

	return sch_systeam, nil
}

// Delete 删除数据
func (a *SysTeam) Delete(ctx context.Context, id string) error {

	r_systeam, err := a.EntCli.SysTeam.Query().Where(systeam.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}

	_, err1 := r_systeam.Update().SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)

	return errors.WithStack(err1)
}

// UpdateActive 更新状态
func (a *SysTeam) UpdateActive(ctx context.Context, id string, active bool) error {

	_, err1 := a.EntCli.SysTeam.Update().Where(systeam.IDEQ(id)).SetIsActive(active).Save(ctx)

	return errors.WithStack(err1)
}
