package entity

import (
	"github.com/heromicro/omgind/internal/schema/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type OrgStaff struct {
	ent.Schema
}

func (OrgStaff) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.SoftDelMixin{},
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

		field.Enum("gender").NamedValues("Male", "M", "Female", "F").StorageKey("gender").Nillable().Optional().Comment("性别"),
		field.String("birth_date").Nillable().Optional().StorageKey("birth_date").Comment("出生日期"),
		field.String("iden_no").Nillable().Optional().StorageKey("iden_no").Comment("身份证号码"),

		field.String("worker_no").Nillable().Optional().StorageKey("worker_no").Comment("工号"),
		field.String("cubicle").Nillable().Optional().StorageKey("cubicle").Comment("工位"),
		field.String("entry_date").Nillable().Optional().StorageKey("entry_date").Comment("入职日期"),
		field.String("regular_date").Nillable().Optional().StorageKey("regu_date").Comment("转正日期"),
		field.String("resign_date").Nillable().Optional().StorageKey("resign_date").Comment("离职日期"),

		field.String("org_id").MaxLen(36).Nillable().Optional().StorageKey("org_id").Comment("企业id"),

		field.String("creator").Nillable().Optional().StorageKey("creator").Comment("创建者"),
	}
}

func (OrgStaff) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("organ", OrgOrgan.Type).Ref("staffs").Field("org_id").Unique(),
	}
}

func (OrgStaff) Indexes() []ent.Index {
	return []ent.Index{}
}

func (OrgStaff) Hooks() []ent.Hook {
	return []ent.Hook{}
}
