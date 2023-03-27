package entity

import (
	"github.com/heromicro/omgind/internal/schema/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type OrgPosition struct {
	ent.Schema
}

func (OrgPosition) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.SoftDelMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
		mixin.ActiveMixin{},
		mixin.MemoMixin{},
	}
}

func (OrgPosition) Fields() []ent.Field {

	return []ent.Field{

		field.String("name").MaxLen(64).Nillable().Optional().StorageKey("name").Comment("名称"),
		field.String("code").MaxLen(16).Nillable().Optional().StorageKey("code").Comment("助记码"),
		field.String("org_id").MaxLen(36).Nillable().Optional().StorageKey("org_id").Comment("企业id"),

		field.String("creator").Nillable().Optional().StorageKey("creator").Comment("创建者"),
	}
}

func (OrgPosition) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("organ", OrgOrgan.Type).Ref("positions").Field("org_id").Unique(),

		// M2O
		edge.To("staffs", OrgStaff.Type),
	}
}

func (OrgPosition) Indexes() []ent.Index {
	return []ent.Index{}
}

func (OrgPosition) Hooks() []ent.Hook {
	return []ent.Hook{}
}
