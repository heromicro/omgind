package repo

import (
	"context"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/sysuser"
	"github.com/heromicro/omgind/internal/gen/ent/sysuserrole"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"
)

// UserSet 注入User
var UserSet = wire.NewSet(wire.Struct(new(User), "*"))

// User 用户存储
type User struct {
	EntCli *ent.Client
}

func ToSchemaSysUser(u *ent.SysUser) *schema.User {
	item := new(schema.User)
	structure.Copy(u, item)
	return item
}

func ToSchemaSysUsers(us ent.SysUsers) []*schema.User {
	list := make([]*schema.User, len(us))
	for i, item := range us {
		list[i] = ToSchemaSysUser(item)
	}
	return list
}

func ToEntCreateSysUserInput(user *schema.User) *ent.CreateSysUserInput {
	createinput := new(ent.CreateSysUserInput)
	structure.Copy(user, &createinput)

	return createinput
}

func ToEntUpdateSysUserInput(user *schema.User) *ent.UpdateSysUserInput {
	updateinput := new(ent.UpdateSysUserInput)
	structure.Copy(user, &updateinput)

	return updateinput
}

func (a *User) getQueryOption(opts ...schema.UserQueryOptions) schema.UserQueryOptions {
	var opt schema.UserQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *User) Query(ctx context.Context, params schema.UserQueryParam, opts ...schema.UserQueryOptions) (*schema.UserQueryResult, error) {
	// opt := a.getQueryOption(opts...)

	query := a.EntCli.SysUser.Query().Where(sysuser.DeletedAtIsNil())

	if v := params.UserName; v != "" {
		query = query.Where(sysuser.UserNameEQ(v))
	}
	if v := params.IsActive; v != nil {
		query = query.Where(sysuser.IsActiveEQ(*v))
	}
	if v := params.RoleIDs; len(v) > 0 {
		//log.Printf(" =000000 ---- subquery -- v %+v ", v)
		query = query.Where(func(s *sql.Selector) {
			sur_t := sql.Table(sysuserrole.Table)
			s.Where(sql.In(
				sysuser.FieldID,
				sql.Select(sysuserrole.FieldUserID).
					From(sur_t).
					Where(
						sql.In(sysuserrole.FieldRoleID, strings.Join(v, ",")),
					),
			))
		})
	}

	if v := params.QueryValue; v != "" {
		query = query.Where(sysuser.Or(
			sysuser.UserNameContains(v), sysuser.RealNameContains(v),
			sysuser.MobileContains(v), sysuser.EmailContains(v),
		))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.UserQueryResult{PageResult: pr}, nil
	}

	if v := params.IsActive_Order; v != "" {
		query = query.Order(sysuser.ByIsActive(OrderDirection(v)))
	}

	if v := params.Sort_Order; v != "" {
		query = query.Order(sysuser.BySort(OrderDirection(v)))
	}

	query = query.Order(sysuser.ByID(OrderDirection("desc")))

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.UserQueryResult{PageResult: pr}, nil
	}

	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}
	rlist := ent.SysUsers(list)

	qr := &schema.UserQueryResult{
		PageResult: pr,
		Data:       ToSchemaSysUsers(rlist),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *User) Get(ctx context.Context, id string, opts ...schema.UserQueryOptions) (*schema.User, error) {
	opt := a.getQueryOption(opts...)
	query := a.EntCli.SysUser.Query()

	user, err := query.Where(sysuser.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	sch_user := ToSchemaSysUser(user)
	if v := opt.UserRoles; v != nil {
		if *v {
			ent_userroles, err := a.EntCli.SysUserRole.Query().Where(sysuserrole.UserIDEQ(id)).Where(sysuserrole.IsDel(false)).All(ctx)
			if err != nil {
				if !ent.IsNotFound(err) {
					return nil, err
				}
			}

			sch_user.UserRoles = ToSchemaSysUserRoles(ent_userroles)
		}
	}

	return sch_user, nil
}

// Create 创建数据
func (a *User) Create(ctx context.Context, item schema.User) (*schema.User, error) {

	item.CreatedAt = nil // ptr.Time(time.Now())
	item.UpdatedAt = nil // ptr.Time(time.Now())

	iteminput := ToEntCreateSysUserInput(&item)

	sysuser, err := a.EntCli.SysUser.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_user := ToSchemaSysUser(sysuser)
	return sch_user, nil
}

// Update 更新数据
func (a *User) Update(ctx context.Context, id string, item schema.User) (*schema.User, error) {

	oitem, err := a.EntCli.SysUser.Query().Where(sysuser.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	item.UpdatedAt = nil // time.Now()
	iteminput := ToEntUpdateSysUserInput(&item)
	user, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	if err != nil {
		return nil, err
	}
	sch_user := ToSchemaSysUser(user)
	return sch_user, nil
}

// Delete 删除数据
func (a *User) Delete(ctx context.Context, id string) error {
	role, err := a.EntCli.SysUser.Query().Where(sysuser.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}
	_, err1 := role.Update().SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)
	return errors.WithStack(err1)
}

// UpdateActive 更新状态
func (a *User) UpdateActive(ctx context.Context, id string, isActive bool) error {
	_, err := a.EntCli.SysUser.UpdateOneID(id).SetIsActive(isActive).Save(ctx)
	return errors.WithStack(err)
}

// UpdatePassword 更新密码
func (a *User) UpdatePassword(ctx context.Context, id, password string) error {
	_, err := a.EntCli.SysUser.UpdateOneID(id).SetPassword(password).Save(ctx)
	return errors.WithStack(err)
}
