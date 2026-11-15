package entgo

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/heromicro/omgind/internal/scheme/mixin"
)

// SysRole holds the schema definition for the SysRole entity.
type SysRole struct {
	ent.Schema
}

func (sr SysRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.SoftDelMixin{},
		mixin.ActiveMixin{},
		mixin.SortMixin{},
		mixin.MemoMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the SysRole.
func (SysRole) Fields() []ent.Field {
	return []ent.Field{
 field.String("org_id").Default("").MaxLen(36).Comment("Tenant organization ID; empty for platform roles"),
		field.String("name").MaxLen(64).MinLen(2).NotEmpty().Comment("角色名称")}
}

// Edges of the SysRole.
func (SysRole) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.To("userRoles", SysUserRole.Type).Comment("userroles"),
	}
}

func (SysRole) Indexes() []ent.Index {
	return []ent.Index{}
}
