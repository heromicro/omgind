package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/schema/mixin"
)

// SysRoleMenu holds the schema definition for the SysRoleMenu entity.
type SysRoleMenu struct {
	ent.Schema
}

func (rm SysRoleMenu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.SoftDelMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the SysRoleMenu.
func (SysRoleMenu) Fields() []ent.Field {
	return []ent.Field{

		field.String("role_id").MaxLen(36).NotEmpty().Comment("角色ID, sys_role.id"),
		field.String("menu_id").MaxLen(36).NotEmpty().Comment("菜单ID"),
		field.String("action_id").MaxLen(36).Nillable().Optional().Comment("菜单ID, sys_menu_action.id"),
	}
}

//// Edges of the SysRoleMenu.
//func (SysRoleMenu) Edges() []ent.Edge {
//	return []ent.Edge{
//		//edge.To(),
//	}
//}

func (SysRoleMenu) Indexes() []ent.Index {
	return []ent.Index{}
}
