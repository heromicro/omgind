package repo

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/sysrole"
	"github.com/heromicro/omgind/internal/gen/mainent/sysuserrole"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"
)

// RoleSet 注入Role
var RoleSet = wire.NewSet(wire.Struct(new(Role), "*"))

// Role 角色存储
type Role struct {
	EntCli *mainent.Client
}

func ToSchemaRole(role *mainent.SysRole) *schema.Role {
	item := new(schema.Role)
	structure.Copy(role, item)
	return item
}

func ToSchemaRoles(roles mainent.SysRoles) []*schema.Role {
	list := make([]*schema.Role, len(roles))
	for i, item := range roles {
		list[i] = ToSchemaRole(item)
	}
	return list
}

func ToEntCreateSysRoleInput(schrole *schema.Role) *mainent.CreateSysRoleInput {
	createinput := new(mainent.CreateSysRoleInput)
	structure.Copy(schrole, &createinput)

	return createinput
}

func ToEntUpdateSysRoleInput(schrole *schema.Role) *mainent.UpdateSysRoleInput {
	updateinput := new(mainent.UpdateSysRoleInput)
	structure.Copy(schrole, &updateinput)

	return updateinput
}

func (a *Role) getQueryOption(opts ...schema.RoleQueryOptions) schema.RoleQueryOptions {
	var opt schema.RoleQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Role) Query(ctx context.Context, params schema.RoleQueryParam, opts ...schema.RoleQueryOptions) (*schema.RoleQueryResult, error) {

	query := a.EntCli.SysRole.Query().Where(sysrole.DeletedAtIsNil())

	if v := params.IDs; len(v) > 0 {
		query = query.Where(sysrole.IDIn(v...))
	}
	if v := params.Name; v != "" {
		query = query.Where(sysrole.NameEQ(v))
	}
	if v := params.UserID; v != "" {
		query = query.Where(func(s *sql.Selector) {
			sur_t := sql.Table(sysuserrole.Table)
			s.Where(sql.In(
				sysrole.FieldID,
				sql.Select(sysuserrole.FieldRoleID).
					From(sur_t).
					Where(sql.EQ(sysuserrole.FieldUserID, v)),
			),
			)
		})
	}

	if v := params.QueryValue; v != "" {
		query = query.Where(sysrole.Or(sysrole.NameContains(v), sysrole.MemoContains(v)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// get total
	pr := &schema.PaginationResult{Total: count}

	if params.PaginationParam.OnlyCount {
		return &schema.RoleQueryResult{PageResult: pr}, nil
	}

	if v := params.IsActive_Order; v != "" {
		query = query.Order(sysrole.ByIsActive(OrderDirection(v)))
	}

	if v := params.Sort_Order; v != "" {
		query = query.Order(sysrole.BySort(OrderDirection(v)))
	}

	query = query.Order(sysrole.ByID(OrderDirection("desc")))

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()

	if params.Offset() > count {
		return &schema.RoleQueryResult{PageResult: pr}, nil
	}

	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}
	rlist := mainent.SysRoles(list)

	qr := &schema.RoleQueryResult{
		PageResult: pr,
		Data:       ToSchemaRoles(rlist),
	}

	return qr, nil
}

// Query 查询数据
func (a *Role) QuerySelectPage(ctx context.Context, params schema.RoleQueryParam, opts ...schema.RoleQueryOptions) (*schema.RoleQueryResult, error) {

	query := a.EntCli.SysRole.Query().Where(sysrole.DeletedAtIsNil())

	if v := params.IDs; len(v) > 0 {
		query = query.Where(sysrole.IDIn(v...))
	}

	if v := params.Name; v != "" {
		query = query.Where(sysrole.NameEQ(v))
	}

	if v := params.UserID; v != "" {
		query = query.Where(func(s *sql.Selector) {
			sur_t := sql.Table(sysuserrole.Table)
			s.Where(sql.In(
				sysrole.FieldID,
				sql.Select(sysuserrole.FieldRoleID).
					From(sur_t).
					Where(sql.EQ(sysuserrole.FieldUserID, v)),
			),
			)
		})
	}

	if v := params.MustIncludeIDs; len(v) > 0 {
		query = query.Where(sysrole.IDNotIn(v...))
	}

	if v := params.QueryValue; v != "" {
		query = query.Where(sysrole.Or(sysrole.NameContains(v), sysrole.MemoContains(v)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// get total
	pr := &schema.PaginationResult{Total: count}

	if params.PaginationParam.OnlyCount {
		return &schema.RoleQueryResult{PageResult: pr}, nil
	}

	if v := params.Sort_Order; v != "" {
		query = query.Order(sysrole.ByIsActive(OrderDirection(v)))
	}

	query = query.Order(sysrole.ByID(OrderDirection(OrderByDESC.String())))

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()

	if params.Offset() > count {
		return &schema.RoleQueryResult{PageResult: pr}, nil
	}

	query = query.Limit(params.Limit()).Offset(params.Offset())

	qs := query.Select(sysrole.FieldID, sysrole.FieldName, sysrole.FieldIsActive)

	list, err1 := qs.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}
	var inRoles schema.Roles = make(schema.Roles, 0)

	if v := params.MustIncludeIDs; len(v) > 0 {
		query_in := a.EntCli.SysRole.Query().Where(sysrole.IDIn(v...))
		roles_ent, err := query_in.All(ctx)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		inRoles = append(inRoles, ToSchemaRoles(roles_ent)...)
	}

	rlist := mainent.SysRoles(list)

	// qr := &schema.RoleQueryResult{
	// 	PageResult: pr,
	// 	Data:       ToSchemaRoles(rlist),
	// }

	qr := &schema.RoleQueryResult{
		PageResult: pr,
		// Data:       inRoles + ToSchemaRoles(rlist),
	}
	if (len(inRoles)) > 0 {
		inRoles = append(inRoles, ToSchemaRoles(rlist)...)
		qr.Data = inRoles

	} else {
		qr.Data = ToSchemaRoles(rlist)
	}

	return qr, nil
}

// Get 查询指定数据
func (a *Role) Get(ctx context.Context, id string, opts ...schema.RoleQueryOptions) (*schema.Role, error) {

	role, err := a.EntCli.SysRole.Query().Where(sysrole.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return ToSchemaRole(role), nil
}

// Create 创建数据
func (a *Role) Create(ctx context.Context, item schema.Role) (*schema.Role, error) {
	//item.CreatedAt = time.Now()
	//item.UpdatedAt = time.Now()

	iteminput := ToEntCreateSysRoleInput(&item)
	iteminput.CreatedAt = nil
	iteminput.UpdatedAt = nil

	role, err := a.EntCli.SysRole.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}

	sch_role := ToSchemaRole(role)
	return sch_role, nil

}

// Update 更新数据
func (a *Role) Update(ctx context.Context, id string, item schema.Role) (*schema.Role, error) {

	oitem, err := a.EntCli.SysRole.Query().Where(sysrole.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, errors.ErrNotFound
	}

	//item.UpdatedAt = time.Now()
	itemInput := ToEntUpdateSysRoleInput(&item)
	itemInput.UpdatedAt = nil
	role, err := oitem.Update().SetInput(*itemInput).Save(ctx)
	if err != nil {
		return nil, err
	}

	sch_role := ToSchemaRole(role)
	return sch_role, nil
}

// Delete 删除数据
func (a *Role) Delete(ctx context.Context, id string) error {
	role, err := a.EntCli.SysRole.Query().Where(sysrole.IDEQ(id)).Only(ctx)
	if err != nil {
		return errors.ErrNotFound
	}
	_, err1 := role.Update().SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)
	return errors.WithStack(err1)

}

// UpdateActive 更新状态
func (a *Role) UpdateActive(ctx context.Context, id string, isActive bool) error {
	_, err1 := a.EntCli.SysRole.UpdateOneID(id).SetIsActive(isActive).Save(ctx)
	return errors.WithStack(err1)
}
