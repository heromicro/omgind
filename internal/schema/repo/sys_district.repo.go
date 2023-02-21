package repo

import (
	"context"
	"time"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/sysdistrict"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"

	"github.com/google/wire"
)

// SysDistrictSet 注入SysDistrict
var SysDistrictSet = wire.NewSet(wire.Struct(new(SysDistrict), "*"))

// SysDistrict 行政区域存储
type SysDistrict struct {
	EntCli *ent.Client
}

// ToSchemaSysDistrict 转换为
func (a *SysDistrict) toSchemaSysDistrict(et *ent.SysDistrict) *schema.SysDistrict {
	item := new(schema.SysDistrict)
	structure.Copy(et, item)
	return item
}

func (a *SysDistrict) toSchemaSysDistricts(ets ent.SysDistricts) []*schema.SysDistrict {
	list := make([]*schema.SysDistrict, len(ets))
	for i, item := range ets {
		list[i] = a.toSchemaSysDistrict(item)
	}
	return list
}

func (a *SysDistrict) ToEntCreateSysDistrictInput(sch *schema.SysDistrict) *ent.CreateSysDistrictInput {
	createinput := new(ent.CreateSysDistrictInput)
	structure.Copy(sch, &createinput)

	return createinput
}

func (a *SysDistrict) ToEntUpdateSysDistrictInput(sch *schema.SysDistrict) *ent.UpdateSysDistrictInput {
	updateinput := new(ent.UpdateSysDistrictInput)
	structure.Copy(sch, &updateinput)

	return updateinput
}

func (a *SysDistrict) getQueryOption(opts ...schema.SysDistrictQueryOptions) schema.SysDistrictQueryOptions {
	var opt schema.SysDistrictQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *SysDistrict) Query(ctx context.Context, params schema.SysDistrictQueryParam, opts ...schema.SysDistrictQueryOptions) (*schema.SysDistrictQueryResult, error) {
	opt := a.getQueryOption(opts...)

	query := a.EntCli.SysDistrict.Query()

	query = query.Where(sysdistrict.DeletedAtIsNil())
	// TODO: 查询条件

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.SysDistrictQueryResult{PageResult: pr}, nil
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	query = query.Order(ParseOrder(opt.OrderFields)...)

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.SysDistrictQueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.SysDistrictQueryResult{
		PageResult: pr,
		Data:       a.toSchemaSysDistricts(list),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *SysDistrict) Get(ctx context.Context, id string, opts ...schema.SysDistrictQueryOptions) (*schema.SysDistrict, error) {

	r_sysdistrict, err := a.EntCli.SysDistrict.Query().Where(sysdistrict.IDEQ(id)).Only(ctx)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return a.toSchemaSysDistrict(r_sysdistrict), nil
}

// Create 创建数据
func (a *SysDistrict) Create(ctx context.Context, item schema.SysDistrict) (*schema.SysDistrict, error) {

	iteminput := a.ToEntCreateSysDistrictInput(&item)
	r_sysdistrict, err := a.EntCli.SysDistrict.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_sysdistrict := a.toSchemaSysDistrict(r_sysdistrict)
	return sch_sysdistrict, nil
}

// Update 更新数据
func (a *SysDistrict) Update(ctx context.Context, id string, item schema.SysDistrict) (*schema.SysDistrict, error) {

	oitem, err := a.EntCli.SysDistrict.Query().Where(sysdistrict.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	iteminput := a.ToEntUpdateSysDistrictInput(&item)

	r_sysdistrict, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	sch_sysdistrict := a.toSchemaSysDistrict(r_sysdistrict)

	return sch_sysdistrict, nil
}

// Delete 删除数据
func (a *SysDistrict) Delete(ctx context.Context, id string) error {

	r_sysdistrict, err := a.EntCli.SysDistrict.Query().Where(sysdistrict.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}

	_, err1 := r_sysdistrict.Update().SetDeletedAt(time.Now()).Save(ctx)

	return errors.WithStack(err1)
}

// UpdateActive 更新状态
func (a *SysDistrict) UpdateActive(ctx context.Context, id string, active bool) error {

	_, err1 := a.EntCli.SysDistrict.Update().Where(sysdistrict.IDEQ(id)).SetIsActive(active).Save(ctx)

	return errors.WithStack(err1)
}
