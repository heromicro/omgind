package repo

import (
	"context"

	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/sysuserrole"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"
)

// UserRoleSet 注入UserRole
var UserRoleSet = wire.NewSet(wire.Struct(new(UserRole), "*"))


// UserRole 用户角色存储
type UserRole struct {
	EntCli *ent.Client
}

func (a *UserRole) toSchemaSysUserRole(dit *ent.SysUserRole) *schema.UserRole {
	item := new(schema.UserRole)
	structure.Copy(dit, item)
	return item
}

func (a *UserRole) toSchemaSysUserRoles(dits ent.SysUserRoles) []*schema.UserRole {
	list := make([]*schema.UserRole, len(dits))
	for i, item := range dits {
		list[i] = a.toSchemaSysUserRole(item)
	}
	return list
}

func (a *UserRole) getQueryOption(opts ...schema.UserRoleQueryOptions) schema.UserRoleQueryOptions {
	var opt schema.UserRoleQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *UserRole) Query(ctx context.Context, params schema.UserRoleQueryParam, opts ...schema.UserRoleQueryOptions) (*schema.UserRoleQueryResult, error) {
	opt := a.getQueryOption(opts...)

	query := a.EntCli.SysUserRole.Query().Where(sysuserrole.DeletedAtIsNil())

	if v := params.UserID; v != "" {
		query = query.Where(sysuserrole.UserIDEQ(v))
	}
	if v := params.UserIDs; len(v) > 0 {
		query = query.Where(sysuserrole.UserIDIn(v...))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.UserRoleQueryResult{PageResult: pr}, nil
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	query = query.Order(ParseOrder(opt.OrderFields)...)

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.UserRoleQueryResult{PageResult: pr}, nil
	}

	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}
	rlist := ent.SysUserRoles(list)

	qr := &schema.UserRoleQueryResult{
		PageResult: pr,
		Data:       a.toSchemaSysUserRoles(rlist),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *UserRole) Get(ctx context.Context, id string, opts ...schema.UserRoleQueryOptions) (*schema.UserRole, error) {

	userrole, err := a.EntCli.SysUserRole.Query().Where(sysuserrole.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return a.toSchemaSysUserRole(userrole), nil
}

// Create 创建数据
func (a *UserRole) Create(ctx context.Context, item schema.UserRole) (*schema.UserRole, error) {

	userrole, err := a.EntCli.SysUserRole.Create().SetUserID(item.UserID).SetRoleID(item.RoleID).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_userrole := a.toSchemaSysUserRole(userrole)
	return sch_userrole, nil
}

// Update 更新数据
func (a *UserRole) Update(ctx context.Context, id string, item schema.UserRole) (*schema.UserRole, error) {

	oitem, err := a.EntCli.SysUserRole.Query().Where(sysuserrole.IDEQ(id)).Only(ctx)

	if err != nil {
		return nil, err
	}
	userrole, err := oitem.Update().SetUserID(item.UserID).SetRoleID(item.RoleID).Save(ctx)
	sch_userrole := a.toSchemaSysUserRole(userrole)
	return sch_userrole, nil
}

// Delete 删除数据
func (a *UserRole) Delete(ctx context.Context, id string) error {

	err := a.EntCli.SysUserRole.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

// DeleteByUserID 根据用户ID删除数据
func (a *UserRole) DeleteByUserID(ctx context.Context, userID string) error {
	_, err := a.EntCli.SysUserRole.Delete().Where(sysuserrole.UserIDEQ(userID)).Exec(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

