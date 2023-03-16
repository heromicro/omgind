package service

import (
	"context"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/schema/repo"
	"github.com/heromicro/omgind/pkg/errors"
)

// SysAddressSet 注入SysAddress
var SysAddressSet = wire.NewSet(wire.Struct(new(SysAddress), "*"))

// SysAddress 地址管理
type SysAddress struct {
	EntCli *ent.Client

	SysAddressRepo *repo.SysAddress
}

// Query 查询数据
func (a *SysAddress) Query(ctx context.Context, params schema.SysAddressQueryParam, opts ...schema.SysAddressQueryOptions) (*schema.SysAddressQueryResult, error) {
	return a.SysAddressRepo.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *SysAddress) Get(ctx context.Context, id string, opts ...schema.SysAddressQueryOptions) (*schema.SysAddress, error) {
	item, err := a.SysAddressRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// Create 创建数据
func (a *SysAddress) Create(ctx context.Context, item schema.SysAddress) (*schema.SysAddress, error) {
	// TODO: check?

	sch_sysaddress, err := a.SysAddressRepo.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return sch_sysaddress, nil
}

// Update 更新数据
func (a *SysAddress) Update(ctx context.Context, id string, item schema.SysAddress) (*schema.SysAddress, error) {

	oitem, err := a.SysAddressRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if oitem == nil {
		return nil, errors.ErrNotFound
	}

	item.ID = oitem.ID

	return a.SysAddressRepo.Update(ctx, id, item)
}

// Delete 删除数据
func (a *SysAddress) Delete(ctx context.Context, id string) error {
	oldItem, err := a.SysAddressRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.SysAddressRepo.Delete(ctx, id)
}

// UpdateActive 更新状态
func (a *SysAddress) UpdateActive(ctx context.Context, id string, active bool) error {
	oldItem, err := a.SysAddressRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.SysAddressRepo.UpdateActive(ctx, id, active)
}
