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

		field.Int32("gender").StorageKey("gender").Nillable().Optional().Comment("性别"),
		field.String("gndr_dict_id").MaxLen(36).Nillable().Optional().StorageKey("gndr_dict_id").Comment("gender options"),

		field.Time("birth_date").Nillable().Optional().StorageKey("birth_date").Comment("出生日期"),

		field.String("iden_no").Nillable().Optional().StorageKey("iden_no").Comment("身份证号码"),

		field.String("iden_addr_id").MaxLen(36).Nillable().Optional().StorageKey("idaddr_id").Comment("身份证地址"),
		field.String("resi_addr_id").MaxLen(36).Nillable().Optional().StorageKey("rsaddr_id").Comment("现居地址"),

		field.String("worker_no").MaxLen(16).Nillable().Optional().StorageKey("worker_no").Comment("工号"),
		field.String("cubicle").MaxLen(32).Nillable().Optional().StorageKey("cubicle").Comment("工位"),

		field.Time("entry_date").Nillable().Optional().StorageKey("entry_date").Comment("入职日期"),
		field.Time("regular_date").Nillable().Optional().StorageKey("regu_date").Comment("转正日期"),
		field.Time("resign_date").Nillable().Optional().StorageKey("resign_date").Comment("离职日期"),

		field.String("org_id").MaxLen(36).Nillable().Optional().StorageKey("org_id").Comment("企业id"),

		field.Int32("emp_stat").Default(1).StorageKey("emp_stat").Comment("在职状态"),

		field.String("empst_dict_id").MaxLen(36).Nillable().Optional().StorageKey("empst_dict_id").Comment("empst_stat options"),

		field.String("creator").Nillable().Optional().StorageKey("creator").Comment("创建者"),
	}
}

func (OrgStaff) Edges() []ent.Edge {

	return []ent.Edge{

		edge.From("organ", OrgOrgan.Type).Ref("staffs").Field("org_id").Unique(),

		edge.From("iden_addr", SysAddress.Type).Ref("staff_iden").Unique().Field("iden_addr_id"),
		edge.From("resi_addr", SysAddress.Type).Ref("staff_resi").Unique().Field("resi_addr_id"),

		// edge.From("gender_dict", OrgOrgan.Type).Ref("staff_gender").Field("gndr_dict_id").Unique(),
		// edge.From("emstat_dict", OrgOrgan.Type).Ref("staff_gender").Field("emst_dict_id").Unique(),
	}
}

func (OrgStaff) Indexes() []ent.Index {
	return []ent.Index{}
}

func (OrgStaff) Hooks() []ent.Hook {
	return []ent.Hook{}
}
