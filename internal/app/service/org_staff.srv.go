package service

import (
	"context"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/orgstaff"
	"github.com/heromicro/omgind/internal/schema/repo"
	"github.com/heromicro/omgind/pkg/errors"
)

// OrgStaffSet 注入OrgStaff
var OrgStaffSet = wire.NewSet(wire.Struct(new(OrgStaff), "*"))

// OrgStaff 员工管理
type OrgStaff struct {
	EntCli *ent.Client

	OrgStaffRepo   *repo.OrgStaff
	SysAddressRepo *repo.SysAddress
}

// Query 查询数据
func (a *OrgStaff) Query(ctx context.Context, params schema.OrgStaffQueryParam, opts ...schema.OrgStaffQueryOptions) (*schema.OrgStaffQueryResult, error) {
	return a.OrgStaffRepo.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *OrgStaff) Get(ctx context.Context, id string, opts ...schema.OrgStaffQueryOptions) (*schema.OrgStaff, error) {
	item, err := a.OrgStaffRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// View 查询指定数据
func (a *OrgStaff) View(ctx context.Context, id string, opts ...schema.OrgStaffQueryOptions) (*schema.OrgStaff, error) {
	item, err := a.OrgStaffRepo.View(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// Create 创建数据
func (a *OrgStaff) Create(ctx context.Context, item schema.OrgStaff) (*schema.OrgStaff, error) {

	iden_addr_iteminput := a.SysAddressRepo.ToEntCreateSysAddressInput(item.IdenAddr)
	item.IdenAddr = nil

	resi_addr_iteminput := a.SysAddressRepo.ToEntCreateSysAddressInput(item.ResiAddr)
	item.ResiAddr = nil

	staff_input := a.OrgStaffRepo.ToEntCreateOrgStaffInput(&item)

	var rr_orgstaff *ent.OrgStaff

	err := repo.WithTx(ctx, a.EntCli, func(tx *ent.Tx) error {

		iden_addr, err := a.EntCli.SysAddress.Create().SetInput(*iden_addr_iteminput).Save(ctx)
		if err != nil {
			return err
		}
		resi_addr, err := a.EntCli.SysAddress.Create().SetInput(*resi_addr_iteminput).Save(ctx)
		if err != nil {
			return err
		}

		rr_orgstaff, err = a.EntCli.OrgStaff.Create().SetInput(*staff_input).SetIdenAddrID(iden_addr.ID).SetResiAddrID(resi_addr.ID).Save(ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	r_orgstaff, err := a.EntCli.OrgStaff.Query().Where(orgstaff.IDEQ(rr_orgstaff.ID)).WithIdenAddr().WithResiAddr().First(ctx)

	if err != nil {
		return nil, err
	}

	sch_orgstaff := repo.ToSchemaOrgStaff(r_orgstaff)

	return sch_orgstaff, nil
}

// Update 更新数据
func (a *OrgStaff) Update(ctx context.Context, id string, item schema.OrgStaff) (*schema.OrgStaff, error) {

	oitem, err := a.OrgStaffRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if oitem == nil {
		return nil, errors.ErrNotFound
	}

	item.ID = oitem.ID

	return a.OrgStaffRepo.Update(ctx, id, item)
}

// Delete 删除数据
func (a *OrgStaff) Delete(ctx context.Context, id string) error {
	oldItem, err := a.OrgStaffRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.OrgStaffRepo.Delete(ctx, id)
}

// UpdateActive 更新状态
func (a *OrgStaff) UpdateActive(ctx context.Context, id string, active bool) error {
	oldItem, err := a.OrgStaffRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.OrgStaffRepo.UpdateActive(ctx, id, active)
}
