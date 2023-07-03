package service

import (
	"context"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/scheme/repo"
	"github.com/heromicro/omgind/pkg/errors"
)

// SysTeamSet 注入SysTeam
var SysTeamSet = wire.NewSet(wire.Struct(new(SysTeam), "*"))

// SysTeam 地址管理
type SysTeam struct {
	SysTeamRepo *repo.SysTeam
}

// Query 查询数据
func (a *SysTeam) Query(ctx context.Context, params schema.SysTeamQueryParam, opts ...schema.SysTeamQueryOptions) (*schema.SysTeamQueryResult, error) {
	return a.SysTeamRepo.Query(ctx, params, opts...)
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
