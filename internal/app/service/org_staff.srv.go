package service

import (
	"context"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/orgstaff"
	"github.com/heromicro/omgind/internal/gen/ent/sysaddress"
	"github.com/heromicro/omgind/internal/schema/repo"
	"github.com/heromicro/omgind/pkg/errors"
)

// OrgStaffSet 注入OrgStaff
var OrgStaffSet = wire.NewSet(wire.Struct(new(OrgStaff), "*"))

// OrgStaff 员工管理
type OrgStaff struct {
	EntCli *ent.Client

	OrgStaffRepo *repo.OrgStaff
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

	iden_addr_iteminput := repo.ToEntCreateSysAddressInput(item.IdenAddr)
	item.IdenAddr = nil

	resi_addr_iteminput := repo.ToEntCreateSysAddressInput(item.ResiAddr)
	item.ResiAddr = nil

	staff_input := repo.ToEntCreateOrgStaffInput(&item)

	var rr_orgstaff *ent.OrgStaff

	err := repo.WithTx(ctx, a.EntCli, func(tx *ent.Tx) error {

		iden_addr, err := tx.SysAddress.Create().SetInput(*iden_addr_iteminput).Save(ctx)
		if err != nil {
			return err
		}
		resi_addr, err := tx.SysAddress.Create().SetInput(*resi_addr_iteminput).Save(ctx)
		if err != nil {
			return err
		}

		rr_orgstaff, err = tx.OrgStaff.Create().SetInput(*staff_input).SetIdenAddrID(iden_addr.ID).SetResiAddrID(resi_addr.ID).Save(ctx)
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

	oitem, err := a.EntCli.OrgStaff.Query().Where(orgstaff.IDEQ(id), orgstaff.IsDel(false)).WithIdenAddr().WithResiAddr().First(ctx)
	if err != nil {
		return nil, err
	} else if oitem == nil {
		return nil, errors.ErrNotFound
	}

	iden_addr_create_input := repo.ToEntCreateSysAddressInput(item.IdenAddr)
	iden_addr_update_input := repo.ToEntUpdateSysAddressInput(item.IdenAddr)
	item.IdenAddr = nil

	resi_addr_create_input := repo.ToEntCreateSysAddressInput(item.ResiAddr)
	resi_addr_update_input := repo.ToEntUpdateSysAddressInput(item.ResiAddr)
	item.ResiAddr = nil

	staff_input := repo.ToEntUpdateOrgStaffInput(&item)

	err = repo.WithTx(ctx, a.EntCli, func(tx *ent.Tx) error {

		var iden_addr *ent.SysAddress
		if v := oitem.IdenAddrID; v != nil {
			_, err = tx.SysAddress.Update().Where(sysaddress.IDEQ(*v)).SetInput(*iden_addr_update_input).Save(ctx)
			if err != nil {
				return nil
			}
		} else {
			iden_addr, err = tx.SysAddress.Create().SetInput(*iden_addr_create_input).Save(ctx)
			if err != nil {
				return err
			}
		}

		var resi_addr *ent.SysAddress
		if v := oitem.ResiAddrID; v != nil {
			_, err = tx.SysAddress.Update().Where(sysaddress.IDEQ(*v)).SetInput(*resi_addr_update_input).Save(ctx)
			if err != nil {
				return nil
			}
		} else {
			resi_addr, err = tx.SysAddress.Create().SetInput(*resi_addr_create_input).Save(ctx)
			if err != nil {
				return err
			}
		}

		udpate_staff := tx.OrgStaff.Update().Where(orgstaff.IDEQ(id)).SetInput(*staff_input)
		if iden_addr != nil {
			udpate_staff = udpate_staff.SetIdenAddrID(iden_addr.ID)
		}

		if resi_addr != nil {
			udpate_staff = udpate_staff.SetResiAddrID(resi_addr.ID)
		}

		_, err = udpate_staff.Save(ctx)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	r_staff, err := a.EntCli.OrgStaff.Query().Where(orgstaff.IDEQ(id)).WithIdenAddr().WithResiAddr().First(ctx)
	if err != nil {
		return nil, err
	}

	sch_orgstaff := repo.ToSchemaOrgStaff(r_staff)

	return sch_orgstaff, nil
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
