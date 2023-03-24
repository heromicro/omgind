package service

import (
	"context"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/orgorgan"
	"github.com/heromicro/omgind/internal/gen/ent/sysaddress"
	"github.com/heromicro/omgind/internal/schema/repo"
	"github.com/heromicro/omgind/pkg/errors"
)

// OrgOrganSet 注入OrgOrgan
var OrgOrganSet = wire.NewSet(wire.Struct(new(OrgOrgan), "*"))

// OrgOrgan 组织管理
type OrgOrgan struct {
	EntCli *ent.Client

	OrgOrganRepo   *repo.OrgOrgan
	SysAddressRepo *repo.SysAddress
	OrgDeptRepo    *repo.OrgDept
}

// Query 查询数据
func (a *OrgOrgan) Query(ctx context.Context, params schema.OrgOrganQueryParam, opts ...schema.OrgOrganQueryOptions) (*schema.OrgOrganQueryResult, error) {
	return a.OrgOrganRepo.Query(ctx, params, opts...)
}

// Query 查询数据
func (a *OrgOrgan) QuerySelect(ctx context.Context, params schema.OrgOrganQueryParam, opts ...schema.OrgOrganQueryOptions) (*schema.OrgOrganQueryResult, error) {
	return a.OrgOrganRepo.QuerySelect(ctx, params, opts...)
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

	addr_iteminput := a.SysAddressRepo.ToEntCreateSysAddressInput(item.Haddr)
	item.Haddr = nil
	organ_iteminput := a.OrgOrganRepo.ToEntCreateOrgOrganInput(&item)

	var rr_orgorgan *ent.OrgOrgan

	la_tree_id, err := a.OrgDeptRepo.GetLatestTreeID(ctx)
	if err != nil {
		return nil, err
	}

	err = repo.WithTx(ctx, a.EntCli, func(tx *ent.Tx) error {

		r_addr, err := tx.SysAddress.Create().SetInput(*addr_iteminput).Save(ctx)
		if err != nil {
			return err
		}
		rr_orgorgan, err = tx.OrgOrgan.Create().SetInput(*organ_iteminput).SetHaddrID(r_addr.ID).Save(ctx)
		if err != nil {
			return err
		}

		// create an not show dept record
		_, err = tx.OrgDept.Create().SetOrgID(rr_orgorgan.ID).SetTreeID(la_tree_id).SetName(item.Sname).SetTreeLevel(1).SetTreeLeft(1).SetTreeRight(2).SetSort(0).SetIsShow(false).SetIsLeaf(false).Save(ctx)

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	r_orgorgan, err := a.EntCli.OrgOrgan.Query().Where(orgorgan.IDEQ(rr_orgorgan.ID)).WithHaddr().First(ctx)
	if err != nil {
		return nil, err
	}

	sch_orgorgan := repo.ToSchemaOrgOrgan(r_orgorgan)
	// sch_orgorgan.Haddr = repo.ToSchemaSysAddress(r_orgorgan.Edges.Haddr)

	return sch_orgorgan, nil
}

// Update 更新数据
func (a *OrgOrgan) Update(ctx context.Context, id string, item schema.OrgOrgan) (*schema.OrgOrgan, error) {

	oitem, err := a.EntCli.OrgOrgan.Query().Where(orgorgan.IDEQ(id), orgorgan.IsDel(false)).WithHaddr().First(ctx)
	if err != nil {
		return nil, err
	} else if oitem == nil {
		return nil, errors.ErrNotFound
	}

	addr_create_input := a.SysAddressRepo.ToEntCreateSysAddressInput(item.Haddr)
	addr_update_input := a.SysAddressRepo.ToEntUpdateSysAddressInput(item.Haddr)
	item.Haddr = nil
	organ_iteminput := a.OrgOrganRepo.ToEntUpdateOrgOrganInput(&item)

	err = repo.WithTx(ctx, a.EntCli, func(tx *ent.Tx) error {
		var haddr *ent.SysAddress
		if v := oitem.HaddrID; v != nil {
			_, err := tx.SysAddress.Update().Where(sysaddress.IDEQ(*v)).SetInput(*addr_update_input).Save(ctx)
			if err != nil {
				return err
			}
		} else {
			haddr, err = tx.SysAddress.Create().SetInput(*addr_create_input).Save(ctx)
			if err != nil {
				return err
			}
		}

		update_orga := tx.OrgOrgan.Update().Where(orgorgan.IDEQ(id)).SetInput(*organ_iteminput)
		if haddr != nil {
			update_orga = update_orga.SetHaddrID(haddr.ID)
		}

		_, err = update_orga.Save(ctx)
		if err != nil {
			return nil
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	r_orgorgan, err := a.EntCli.OrgOrgan.Query().Where(orgorgan.IDEQ(id)).WithHaddr().First(ctx)
	if err != nil {
		return nil, err
	}

	sch_orgorgan := repo.ToSchemaOrgOrgan(r_orgorgan)

	return sch_orgorgan, nil
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
