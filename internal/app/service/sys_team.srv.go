package service

import (
	"context"
	"log"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/systeam"
	"github.com/heromicro/omgind/internal/gen/mainent/sysuser"
	"github.com/heromicro/omgind/internal/scheme/repo"
	"github.com/heromicro/omgind/pkg/errors"
)

// SysTeamSet 注入SysTeam
var SysTeamSet = wire.NewSet(wire.Struct(new(SysTeam), "*"))

// SysTeam 地址管理
type SysTeam struct {
	EntCli *mainent.Client

	SysTeamRepo *repo.SysTeam
}

// Query 查询数据
func (a *SysTeam) Query(ctx context.Context, params schema.SysTeamQueryParam, opts ...schema.SysTeamQueryOptions) (*schema.SysTeamQueryResult, error) {
	return a.SysTeamRepo.Query(ctx, params, opts...)
}

// QuerySelectPage 查询显示项数据
func (a *SysTeam) QuerySelectPage(ctx context.Context, params schema.SysTeamQueryParam,
	opts ...schema.SysTeamQueryOptions) (*schema.SysTeamQueryResult, error) {

	// opt := a.TeamRepo.GetQueryOption(opts...)

	query := a.EntCli.SysTeam.Query().Where(systeam.DeletedAtIsNil())
	if v := params.UserIDs; len(v) > 0 {
		query = query.WithUsers(func(suq *mainent.SysUserQuery) {
			suq.Where(sysuser.IDIn(params.UserIDs...)).Select(sysuser.FieldID, sysuser.FieldIsActive, sysuser.FieldUserName)
		})
	}

	if v := params.QueryValue; v != "" {
		query = query.Where(systeam.Or(
			systeam.NameContains(v),
		))
	}

	if v := params.After; v != "" {
		query = query.Where(systeam.IDGT(v))
	}

	log.Println(" ------ --0 --- ", params)

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	pr := &schema.PaginationResult{Total: count}

	// query = query.Order(ParseOrder(opt.OrderFields)...)

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	pr.After = params.After
	pr.Before = params.Before

	if params.Offset() > count {
		return &schema.SysTeamQueryResult{PageResult: pr}, nil
	}

	query = query.Limit(params.Limit()).Offset(params.Offset())
	qs := query.Select(systeam.FieldID, systeam.FieldName, systeam.FieldIsActive, systeam.FieldCode)

	list, err1 := qs.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.SysTeamQueryResult{
		PageResult: pr,
		Data:       repo.ToSchemaSysTeams(list),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *SysTeam) Get(ctx context.Context, id string, opts ...schema.SysTeamQueryOptions) (*schema.SysTeam, error) {
	item, err := a.SysTeamRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// View 查询指定数据
func (a *SysTeam) View(ctx context.Context, id string, opts ...schema.SysTeamQueryOptions) (*schema.SysTeam, error) {
	item, err := a.SysTeamRepo.View(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// Create 创建数据
func (a *SysTeam) Create(ctx context.Context, item schema.SysTeam) (*schema.SysTeam, error) {
	// TODO: check?

	sch_systeam, err := a.SysTeamRepo.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return sch_systeam, nil
}

// Update 更新数据
func (a *SysTeam) Update(ctx context.Context, id string, item schema.SysTeam) (*schema.SysTeam, error) {

	oitem, err := a.SysTeamRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if oitem == nil {
		return nil, errors.ErrNotFound
	}

	item.ID = oitem.ID

	return a.SysTeamRepo.Update(ctx, id, item)
}

// Delete 删除数据
func (a *SysTeam) Delete(ctx context.Context, id string) error {
	oldItem, err := a.SysTeamRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.SysTeamRepo.Delete(ctx, id)
}

// UpdateActive 更新状态
func (a *SysTeam) UpdateActive(ctx context.Context, id string, active bool) error {
	oldItem, err := a.SysTeamRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.SysTeamRepo.UpdateActive(ctx, id, active)
}
