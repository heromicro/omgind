package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/schema/mixin"
)

// SysMenuAction holds the schema definition for the SysMenuAction entity.
type SysMenuAction struct {
	ent.Schema
}

func (ma SysMenuAction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.SortMixin{},
		mixin.StatusMixin{},
		mixin.MemoMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the SysMenuAction.
func (SysMenuAction) Fields() []ent.Field {
	return []ent.Field{

		field.String("menu_id").MaxLen(36).NotEmpty().Comment("菜单ID"),
		field.String("code").MaxLen(128).NotEmpty().Comment("动作编号"),
		field.String("name").MaxLen(128).NotEmpty().Comment("动作名称"),
	}
}

//// Edges of the SysMenuAction.
//func (SysMenuAction) Edges() []ent.Edge {
//	return []ent.Edge{
//		//edge.To("resources", SysMenuActionResource.Type),
//		//edge.From("menu", SysMenu.Type).Field("menu_id").Ref("actions").Unique().Required(),
//	}
//}

func (SysMenuAction) Indexes() []ent.Index {
	return []ent.Index{}
}
