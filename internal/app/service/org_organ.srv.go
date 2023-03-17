package service

import (
	"context"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/schema/repo"
	"github.com/heromicro/omgind/pkg/errors"
)

// OrgOrganSet 注入OrgOrgan
var OrgOrganSet = wire.NewSet(wire.Struct(new(OrgOrgan), "*"))

// OrgOrgan 组织管理
type OrgOrgan struct {
	OrgOrganRepo *repo.OrgOrgan
}

// Query 查询数据
func (a *OrgOrgan) Query(ctx context.Context, params schema.OrgOrganQueryParam, opts ...schema.OrgOrganQueryOptions) (*schema.OrgOrganQueryResult, error) {
	return a.OrgOrganRepo.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *OrgOrgan) Get(ctx context.Context, id string, opts ...schema.OrgOrganQueryOptions) (*schema.OrgOrgan, error) {
	item, err := a.OrgOrganRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// View 查询指定数据
func (a *OrgOrgan) View(ctx context.Context, id string, opts ...schema.OrgOrganQueryOptions) (*schema.OrgOrgan, error) {
	item, err := a.OrgOrganRepo.View(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// Create 创建数据
func (a *OrgOrgan) Create(ctx context.Context, item schema.OrgOrgan) (*schema.OrgOrgan, error) {
	// TODO: check?

	sch_orgorgan, err := a.OrgOrganRepo.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return sch_orgorgan, nil
}

// Update 更新数据
func (a *OrgOrgan) Update(ctx context.Context, id string, item schema.OrgOrgan) (*schema.OrgOrgan, error) {

	oitem, err := a.OrgOrganRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if oitem == nil {
		return nil, errors.ErrNotFound
	}

	item.ID = oitem.ID

	return a.OrgOrganRepo.Update(ctx, id, item)
}

// Delete 删除数据
func (a *OrgOrgan) Delete(ctx context.Context, id string) error {
	oldItem, err := a.OrgOrganRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.OrgOrganRepo.Delete(ctx, id)
}

// UpdateActive 更新状态
func (a *OrgOrgan) UpdateActive(ctx context.Context, id string, active bool) error {
	oldItem, err := a.OrgOrganRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.OrgOrganRepo.UpdateActive(ctx, id, active)
}
