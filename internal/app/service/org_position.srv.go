package service

import (
	"context"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/entscheme"
	"github.com/heromicro/omgind/internal/scheme/repo"
	"github.com/heromicro/omgind/pkg/errors"
)

// OrgPositionSet 注入OrgPosition
var OrgPositionSet = wire.NewSet(wire.Struct(new(OrgPosition), "*"))

// OrgPosition 职位管理
type OrgPosition struct {
	EntCli *entscheme.Client

	OrgPositionRepo *repo.OrgPosition
}

// Query 查询数据
func (a *OrgPosition) Query(ctx context.Context, params schema.OrgPositionQueryParam, opts ...schema.OrgPositionQueryOptions) (*schema.OrgPositionQueryResult, error) {
	return a.OrgPositionRepo.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *OrgPosition) Get(ctx context.Context, id string, opts ...schema.OrgPositionQueryOptions) (*schema.OrgPosition, error) {
	item, err := a.OrgPositionRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// View 查询指定数据
func (a *OrgPosition) View(ctx context.Context, id string, opts ...schema.OrgPositionQueryOptions) (*schema.OrgPosition, error) {
	item, err := a.OrgPositionRepo.View(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// Create 创建数据
func (a *OrgPosition) Create(ctx context.Context, item schema.OrgPosition) (*schema.OrgPosition, error) {
	// TODO: check?

	sch_orgposition, err := a.OrgPositionRepo.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return sch_orgposition, nil
}

// Update 更新数据
func (a *OrgPosition) Update(ctx context.Context, id string, item schema.OrgPosition) (*schema.OrgPosition, error) {

	oitem, err := a.OrgPositionRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if oitem == nil {
		return nil, errors.ErrNotFound
	}

	item.ID = oitem.ID

	return a.OrgPositionRepo.Update(ctx, id, item)
}

// Delete 删除数据
func (a *OrgPosition) Delete(ctx context.Context, id string) error {
	oldItem, err := a.OrgPositionRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.OrgPositionRepo.Delete(ctx, id)
}

// UpdateActive 更新状态
func (a *OrgPosition) UpdateActive(ctx context.Context, id string, active bool) error {
	oldItem, err := a.OrgPositionRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.OrgPositionRepo.UpdateActive(ctx, id, active)
}
