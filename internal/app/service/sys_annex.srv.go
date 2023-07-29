package service

import (
	"context"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/scheme/repo"
	"github.com/heromicro/omgind/pkg/errors"
)

// SysAnnexSet 注入SysAnnex
var SysAnnexSet = wire.NewSet(wire.Struct(new(SysAnnex), "*"))

// SysAnnex
type SysAnnex struct {
	SysAnnexRepo *repo.SysAnnex
}

// Query 查询数据
func (a *SysAnnex) Query(ctx context.Context, params schema.SysAnnexQueryParam, opts ...schema.SysAnnexQueryOptions) (*schema.SysAnnexQueryResult, error) {
	return a.SysAnnexRepo.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *SysAnnex) Get(ctx context.Context, id string, opts ...schema.SysAnnexQueryOptions) (*schema.SysAnnex, error) {
	item, err := a.SysAnnexRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// View 查询指定数据
func (a *SysAnnex) View(ctx context.Context, id string, opts ...schema.SysAnnexQueryOptions) (*schema.SysAnnex, error) {
	item, err := a.SysAnnexRepo.View(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// Create 创建数据
func (a *SysAnnex) Create(ctx context.Context, item schema.SysAnnex) (*schema.SysAnnex, error) {
	// TODO: check?

	sch_sysannex, err := a.SysAnnexRepo.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return sch_sysannex, nil
}

// Update 更新数据
func (a *SysAnnex) Update(ctx context.Context, id string, item schema.SysAnnex) (*schema.SysAnnex, error) {

	oitem, err := a.SysAnnexRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if oitem == nil {
		return nil, errors.ErrNotFound
	}

	item.ID = oitem.ID

	return a.SysAnnexRepo.Update(ctx, id, item)
}

// Delete 删除数据
func (a *SysAnnex) Delete(ctx context.Context, id string) error {
	oldItem, err := a.SysAnnexRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.SysAnnexRepo.Delete(ctx, id)
}

// UpdateActive 更新状态
func (a *SysAnnex) UpdateActive(ctx context.Context, id string, active bool) error {
	oldItem, err := a.SysAnnexRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.SysAnnexRepo.UpdateActive(ctx, id, active)
}
