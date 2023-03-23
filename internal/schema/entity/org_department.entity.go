package entity

import (
	"github.com/heromicro/omgind/internal/schema/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type OrgDepartment struct {
	ent.Schema
}

func (OrgDepartment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.SoftDelMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
		mixin.ActiveMixin{},
		mixin.MemoMixin{},
	}
}

func (OrgDepartment) Fields() []ent.Field {

	return []ent.Field{
		field.String("name").MaxLen(64).Nillable().Optional().StorageKey("name").Comment("名称"),
		field.String("code").MaxLen(16).Nillable().Optional().StorageKey("code").Comment("助记码"),

		field.String("org_id").MaxLen(36).Nillable().Optional().StorageKey("org_id").Comment("企业id"),

		field.String("creator").Nillable().Optional().StorageKey("creator").Comment("创建者"),
	}
}

func (OrgDepartment) Edges() []ent.Edge {
	return []ent.Edge{
		// M2O
		edge.From("organ", OrgOrgan.Type).Ref("departments").Field("org_id").Unique(),
		// O2M
		// edge.To("staffs", OrgStaff.Type),
	}
}

func (OrgDepartment) Indexes() []ent.Index {
	return []ent.Index{}
}

func (OrgDepartment) Hooks() []ent.Hook {
	return []ent.Hook{}
}
