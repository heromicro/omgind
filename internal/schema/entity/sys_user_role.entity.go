package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/schema/mixin"
)

// SysUserRole holds the schema definition for the SysUserRole entity.
type SysUserRole struct {
	ent.Schema
}

func (syr SysUserRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.SoftDelMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the SysUserRole.
func (SysUserRole) Fields() []ent.Field {
	return []ent.Field{
		//mixin.IdField_ulid(),

		field.String("user_id").MaxLen(36).NotEmpty().Comment("用户ID, sys_user.id"),
		field.String("role_id").MaxLen(36).NotEmpty().Comment("角色ID, sys_role.id"),
	}
}

func (SysUserRole) Indexes() []ent.Index {
	return []ent.Index{
		//index.Fields("id").Unique(),
	}
}

// Edges of the SysUserRole.
func (SysUserRole) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.From("user", SysUser.Type).Field("user_id").Ref("userRoles").Unique().Comment("用户ID: sys_user.id").Required(),
		// edge.From("role", SysRole.Type).Field("role_id").Ref("userRoles").Unique().Comment("角色ID: sys_role.id").Required(),
	}
}
