package service

import (
	"context"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/schema/repo"
	"github.com/heromicro/omgind/pkg/errors"
)

// OrgDepartmentSet 注入OrgDepartment
var OrgDepartmentSet = wire.NewSet(wire.Struct(new(OrgDepartment), "*"))

// OrgDepartment 部门管理
type OrgDepartment struct {
	EntCli *ent.Client

	OrgDepartmentRepo *repo.OrgDepartment
}

// Query 查询数据
func (a *OrgDepartment) Query(ctx context.Context, params schema.OrgDepartmentQueryParam, opts ...schema.OrgDepartmentQueryOptions) (*schema.OrgDepartmentQueryResult, error) {
	return a.OrgDepartmentRepo.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *OrgDepartment) Get(ctx context.Context, id string, opts ...schema.OrgDepartmentQueryOptions) (*schema.OrgDepartment, error) {
	item, err := a.OrgDepartmentRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// View 查询指定数据
func (a *OrgDepartment) View(ctx context.Context, id string, opts ...schema.OrgDepartmentQueryOptions) (*schema.OrgDepartment, error) {
	item, err := a.OrgDepartmentRepo.View(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// View 查询指定数据
func (a *OrgDepartment) GetAllSubDepts(ctx context.Context, pid string, params schema.OrgDepartmentQueryParam, opts ...schema.OrgDepartmentQueryOptions) (*schema.OrgDepartmentQueryResult, error) {
	item, err := a.OrgDepartmentRepo.GetAllSubDepts(ctx, pid, params, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// Query 查询数据
func (a *OrgDepartment) GetTree(ctx context.Context, tpid string, params schema.OrgDepartmentQueryParam, opts ...schema.OrgDepartmentQueryOptions) (*schema.OrgDepartmentQueryTreeResult, error) {
	return a.OrgDepartmentRepo.GetTree(ctx, tpid, params, opts...)
}

// Create 创建数据
func (a *OrgDepartment) Create(ctx context.Context, item schema.OrgDepartment) (*schema.OrgDepartment, error) {
	// TODO: check?

	sch_orgdepartment, err := a.OrgDepartmentRepo.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return sch_orgdepartment, nil
}

// Update 更新数据
func (a *OrgDepartment) Update(ctx context.Context, id string, item schema.OrgDepartment) (*schema.OrgDepartment, error) {

	oitem, err := a.OrgDepartmentRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if oitem == nil {
		return nil, errors.ErrNotFound
	}

	item.ID = oitem.ID

	return a.OrgDepartmentRepo.Update(ctx, id, item)
}

// Delete 删除数据
func (a *OrgDepartment) Delete(ctx context.Context, id string) error {
	oldItem, err := a.OrgDepartmentRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.OrgDepartmentRepo.Delete(ctx, id)
}

// UpdateActive 更新状态
func (a *OrgDepartment) UpdateActive(ctx context.Context, id string, active bool) error {
	oldItem, err := a.OrgDepartmentRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.OrgDepartmentRepo.UpdateActive(ctx, id, active)
}
