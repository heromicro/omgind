package entity

import (
	"github.com/heromicro/omgind/internal/schema/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type OrgStaff struct {
	ent.Schema
}

func (OrgStaff) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
		mixin.ActiveMixin{},
		mixin.MemoMixin{},
	}
}

func (OrgStaff) Fields() []ent.Field {

	return []ent.Field{
		field.String("first_name").MaxLen(64).Nillable().Optional().StorageKey("first_name").Comment("名"),
		field.String("last_name").MaxLen(64).Nillable().Optional().StorageKey("last_name").Comment("姓"),
		field.String("mobile").MaxLen(32).Nillable().Optional().StorageKey("mobile").Comment("电话"),
		field.String("creator").Nillable().Optional().StorageKey("creator").Comment("创建者"),
	}
}

func (OrgStaff) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (OrgStaff) Indexes() []ent.Index {
	return []ent.Index{}
}

func (OrgStaff) Hooks() []ent.Hook {
	return []ent.Hook{}
}
