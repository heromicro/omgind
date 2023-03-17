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
		mixin.OwnerOrgMixin{},
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

		field.Enum("gender").NamedValues("Male", "M", "Female", "F").StorageKey("gender").Nillable().Optional().Comment("性别"),
		field.String("birth_date").Nillable().Optional().StorageKey("birth_date").Comment("出生日期"),
		field.String("entry_date").Nillable().Optional().StorageKey("entry_date").Comment("入职日期"),
		field.String("regular_date").Nillable().Optional().StorageKey("regu_date").Comment("转正日期"),
		field.String("iden_no").Nillable().Optional().StorageKey("iden_no").Comment("身份证号码"),

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
