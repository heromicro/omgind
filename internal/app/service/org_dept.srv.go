package service

import (
	"context"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/schema/repo"
	"github.com/heromicro/omgind/pkg/errors"
)

// OrgDeptSet 注入OrgDept
var OrgDeptSet = wire.NewSet(wire.Struct(new(OrgDept), "*"))

// OrgDept 部门管理
type OrgDept struct {
	EntCli *ent.Client

	OrgDeptRepo *repo.OrgDept
}

// Query 查询数据
func (a *OrgDept) Query(ctx context.Context, params schema.OrgDeptQueryParam, opts ...schema.OrgDeptQueryOptions) (*schema.OrgDeptQueryResult, error) {
	return a.OrgDeptRepo.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *OrgDept) Get(ctx context.Context, id string, opts ...schema.OrgDeptQueryOptions) (*schema.OrgDept, error) {
	item, err := a.OrgDeptRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// View 查询指定数据
func (a *OrgDept) View(ctx context.Context, id string, opts ...schema.OrgDeptQueryOptions) (*schema.OrgDept, error) {
	item, err := a.OrgDeptRepo.View(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// View 查询指定数据
func (a *OrgDept) GetAllSubs(ctx context.Context, pid string, params schema.OrgDeptQueryParam, opts ...schema.OrgDeptQueryOptions) (*schema.OrgDeptQueryResult, error) {
	item, err := a.OrgDeptRepo.GetAllSubs(ctx, pid, params, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// Query 查询数据
func (a *OrgDept) GetTree(ctx context.Context, tpid string, params schema.OrgDeptQueryParam, opts ...schema.OrgDeptQueryOptions) (*schema.OrgDeptQueryTreeResult, error) {
	return a.OrgDeptRepo.GetTree(ctx, tpid, params, opts...)
}

// Create 创建数据
func (a *OrgDept) Create(ctx context.Context, item schema.OrgDept) (*schema.OrgDept, error) {
	// TODO: check?

	sch_OrgDept, err := a.OrgDeptRepo.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return sch_OrgDept, nil
}

// Update 更新数据
func (a *OrgDept) Update(ctx context.Context, id string, item schema.OrgDept) (*schema.OrgDept, error) {

	oitem, err := a.OrgDeptRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if oitem == nil {
		return nil, errors.ErrNotFound
	}

	item.ID = oitem.ID

	return a.OrgDeptRepo.Update(ctx, id, item)
}

// Delete 删除数据
func (a *OrgDept) Delete(ctx context.Context, id string) error {
	oldItem, err := a.OrgDeptRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.OrgDeptRepo.Delete(ctx, id)
}

// UpdateActive 更新状态
func (a *OrgDept) UpdateActive(ctx context.Context, id string, active bool) error {
	oldItem, err := a.OrgDeptRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.OrgDeptRepo.UpdateActive(ctx, id, active)
}
