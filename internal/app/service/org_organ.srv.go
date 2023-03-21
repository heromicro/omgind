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
	// TODO: check?

	addr_iteminput := a.SysAddressRepo.ToEntCreateSysAddressInput(item.Haddr)
	item.Haddr = nil
	organ_iteminput := a.OrgOrganRepo.ToEntCreateOrgOrganInput(&item)

	var rr_orgorgan *ent.OrgOrgan

	err := repo.WithTx(ctx, a.EntCli, func(tx *ent.Tx) error {

		r_addr, err := a.EntCli.SysAddress.Create().SetInput(*addr_iteminput).Save(ctx)
		if err != nil {
			return err
		}
		rr_orgorgan, err = a.EntCli.OrgOrgan.Create().SetInput(*organ_iteminput).SetHaddrID(r_addr.ID).Save(ctx)
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
	sch_orgorgan.Haddr = repo.ToSchemaSysAddress(r_orgorgan.Edges.Haddr)

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

	addr_iteminput := a.SysAddressRepo.ToEntUpdateSysAddressInput(item.Haddr)
	item.Haddr = nil
	organ_iteminput := a.OrgOrganRepo.ToEntUpdateOrgOrganInput(&item)

	err = repo.WithTx(ctx, a.EntCli, func(tx *ent.Tx) error {

		_, err := a.EntCli.SysAddress.Update().Where(sysaddress.IDEQ(*oitem.HaddrID)).SetInput(*addr_iteminput).Save(ctx)
		if err != nil {
			return err
		}

		_, err = a.EntCli.OrgOrgan.Update().Where(orgorgan.IDEQ(id)).SetInput(*organ_iteminput).Save(ctx)
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
	sch_orgorgan.Haddr = repo.ToSchemaSysAddress(r_orgorgan.Edges.Haddr)

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
