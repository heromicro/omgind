package repo

import (
	"context"
	"time"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/systeamuser"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"

	"github.com/google/wire"
)

// SysTeamUserSet 注入SysTeamUser
var SysTeamUserSet = wire.NewSet(wire.Struct(new(SysTeamUser), "*"))

// SysTeamUser 地址管理存储
type SysTeamUser struct {
	EntCli *mainent.Client
}

// ToSchemaSysTeamUser 转换为
func ToSchemaSysTeamUser(et *mainent.SysTeamUser) *schema.SysTeamUser {
	item := new(schema.SysTeamUser)
	structure.Copy(et, item)
	return item
}

func ToSchemaSysTeamUsers(ets mainent.SysTeamUsers) []*schema.SysTeamUser {
	list := make([]*schema.SysTeamUser, len(ets))
	for i, item := range ets {
		list[i] = ToSchemaSysTeamUser(item)
	}
	return list
}

func ToEntCreateSysTeamUserInput(sch *schema.SysTeamUser) *mainent.CreateSysTeamUserInput {
	createinput := new(mainent.CreateSysTeamUserInput)
	structure.Copy(sch, &createinput)

	return createinput
}

func ToEntUpdateSysTeamUserInput(sch *schema.SysTeamUser) *mainent.UpdateSysTeamUserInput {
	updateinput := new(mainent.UpdateSysTeamUserInput)
	structure.Copy(sch, &updateinput)

	return updateinput
}

func (a *SysTeamUser) getQueryOption(opts ...schema.SysTeamUserQueryOptions) schema.SysTeamUserQueryOptions {
	var opt schema.SysTeamUserQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *SysTeamUser) Query(ctx context.Context, params schema.SysTeamUserQueryParam, opts ...schema.SysTeamUserQueryOptions) (*schema.SysTeamUserQueryResult, error) {
	// opt := a.getQueryOption(opts...)

	query := a.EntCli.SysTeamUser.Query()

	if v := params.IsDel; v != nil && *v {
		query = query.Where(systeamuser.IsDel(*v))
	} else {
		query = query.Where(systeamuser.IsDel(false))
	}

	query = query.Where(systeamuser.DeletedAtIsNil())
	// TODO: 查询条件

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.SysTeamUserQueryResult{PageResult: pr}, nil
	}

	query = query.Order(mainent.Asc(systeamuser.FieldSort))

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.SysTeamUserQueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.SysTeamUserQueryResult{
		PageResult: pr,
		Data:       ToSchemaSysTeamUsers(list),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *SysTeamUser) Get(ctx context.Context, id string, opts ...schema.SysTeamUserQueryOptions) (*schema.SysTeamUser, error) {

	r_systeamuser, err := a.EntCli.SysTeamUser.Query().Where(systeamuser.IDEQ(id)).Only(ctx)
	if err != nil {
		if mainent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return ToSchemaSysTeamUser(r_systeamuser), nil
}

// View 查询指定数据
func (a *SysTeamUser) View(ctx context.Context, id string, opts ...schema.SysTeamUserQueryOptions) (*schema.SysTeamUser, error) {

	r_systeamuser, err := a.EntCli.SysTeamUser.Query().Where(systeamuser.IDEQ(id)).Only(ctx)
	if err != nil {
		if mainent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return ToSchemaSysTeamUser(r_systeamuser), nil
}

// Create 创建数据
func (a *SysTeamUser) Create(ctx context.Context, item schema.SysTeamUser) (*schema.SysTeamUser, error) {

	iteminput := ToEntCreateSysTeamUserInput(&item)
	r_systeamuser, err := a.EntCli.SysTeamUser.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_systeamuser := ToSchemaSysTeamUser(r_systeamuser)
	return sch_systeamuser, nil
}

// Update 更新数据
func (a *SysTeamUser) Update(ctx context.Context, id string, item schema.SysTeamUser) (*schema.SysTeamUser, error) {

	oitem, err := a.EntCli.SysTeamUser.Query().Where(systeamuser.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	iteminput := ToEntUpdateSysTeamUserInput(&item)

	r_systeamuser, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	if err != nil {
		return nil, err
	}
	sch_systeamuser := ToSchemaSysTeamUser(r_systeamuser)

	return sch_systeamuser, nil
}

// Delete 删除数据
func (a *SysTeamUser) Delete(ctx context.Context, id string) error {

	r_systeamuser, err := a.EntCli.SysTeamUser.Query().Where(systeamuser.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}

	_, err1 := r_systeamuser.Update().SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)

	return errors.WithStack(err1)
}

// UpdateActive 更新状态
func (a *SysTeamUser) UpdateActive(ctx context.Context, id string, active bool) error {

	_, err1 := a.EntCli.SysTeamUser.Update().Where(systeamuser.IDEQ(id)).SetIsActive(active).Save(ctx)

	return errors.WithStack(err1)
}
