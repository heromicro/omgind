package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/heromicro/omgind/internal/scheme/mixin"
)

// SysMenuActionResource holds the schema definition for the SysMenuActionResource entity.
type SysMenuActionResource struct {
	ent.Schema
}

func (smar SysMenuActionResource) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.SoftDelMixin{},
		mixin.SortMixin{},
		mixin.MemoMixin{},
		mixin.TimeMixin{},
		mixin.ActiveMixin{},
	}
}

// Fields of the SysMenuActionResource.
func (SysMenuActionResource) Fields() []ent.Field {
	return []ent.Field{

		field.String("method").MaxLen(128).NotEmpty().Comment("资源HTTP请求方式(支持正则, get, delete, delete, put, patch )"),
		field.String("path").MaxLen(256).NotEmpty().Comment("资源HTTP请求路径（支持/:id匹配）"),
		field.String("action_id").MaxLen(36).NotEmpty().Comment("sys_menu_action.id"),
	}
}

// Edges of the SysMenuActionResource.
func (SysMenuActionResource) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.From("action", SysMenuAction.Type).Field("action_id").Ref("resources").Unique().Required(),
	}
}

func (SysMenuActionResource) Indexes() []ent.Index {
	return []ent.Index{}
}
