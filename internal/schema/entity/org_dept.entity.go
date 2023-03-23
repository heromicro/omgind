package entity

import (
	"github.com/heromicro/omgind/internal/schema/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type OrgDept struct {
	ent.Schema
}

func (OrgDept) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.SoftDelMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
		mixin.ActiveMixin{},
		mixin.MemoMixin{},
		mixin.MpttTreeMixin{},
	}
}

func (OrgDept) Fields() []ent.Field {

	return []ent.Field{
		field.String("name").MaxLen(64).Nillable().Optional().StorageKey("name").Comment("名称"),
		field.String("code").MaxLen(16).Nillable().Optional().StorageKey("code").Comment("助记码"),

		field.String("org_id").MaxLen(36).Nillable().Optional().StorageKey("org_id").Comment("企业id"),

		field.String("parent_id").MaxLen(36).Nillable().Optional().StorageKey("pid").Comment("父级id"),
		field.Bool("is_real").Default(true).Nillable().Optional().StorageKey("is_r").Comment("是否虚拟部门"),

		field.String("creator").Nillable().Optional().StorageKey("creator").Comment("创建者"),
	}
}

func (OrgDept) Edges() []ent.Edge {
	return []ent.Edge{
		// same O2M
		edge.To("children", OrgDept.Type).From("parent").Field("parent_id").Unique(),

		// M2O
		edge.From("organ", OrgOrgan.Type).Ref("depts").Field("org_id").Unique(),
		// O2M
		// edge.To("staffs", OrgStaff.Type),
	}
}

func (OrgDept) Indexes() []ent.Index {
	return []ent.Index{}
}

func (OrgDept) Hooks() []ent.Hook {
	return []ent.Hook{}
}
