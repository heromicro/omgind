package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/heromicro/omgind/internal/schema/mixin"
)

type OrgOrgan struct {
	ent.Schema
}

func (OrgOrgan) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.SoftDelMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
		mixin.ActiveMixin{},
		mixin.MemoMixin{},
	}
}

func (OrgOrgan) Fields() []ent.Field {

	return []ent.Field{
		field.String("name").MaxLen(256).Nillable().Optional().StorageKey("name").Comment("名称"),
		field.String("sname").MaxLen(64).Nillable().Optional().StorageKey("sname").Comment("简称"),
		field.String("code").MaxLen(16).Nillable().Optional().StorageKey("code").Comment("助记码"),

		field.String("iden_no").MaxLen(20).Nillable().Optional().StorageKey("iden_no").Comment("执照号"),

		field.String("owner_id").MaxLen(36).Nillable().Optional().StorageKey("owner_id").Comment("所有者user.id"),
		field.String("haddr_id").MaxLen(36).Nillable().Optional().StorageKey("haddr_id").Comment("总部id"),

		field.String("creator").Nillable().Optional().StorageKey("creator").Comment("创建者"),
	}
}

func (OrgOrgan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("haddr", SysAddress.Type).Ref("organ").Unique().Field("haddr_id"),
		// edge.From("owner", SysUser.Type).Ref("organ").Unique().Field("owner_id"),
	}
}

func (OrgOrgan) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_id"),
	}
}

func (OrgOrgan) Hooks() []ent.Hook {
	return []ent.Hook{}
}
