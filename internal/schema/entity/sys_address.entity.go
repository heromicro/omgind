package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/schema/mixin"
)

type SysAddress struct {
	ent.Schema
}

func (SysAddress) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.OwnerMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
		mixin.ActiveMixin{},
		mixin.MemoMixin{},
	}
}

func (SysAddress) Fields() []ent.Field {

	return []ent.Field{

		field.String("country").Nillable().Optional().StorageKey("country").Comment("国"),
		field.String("provice").Nillable().Optional().StorageKey("provice").Comment("省"),
		field.String("city").Nillable().Optional().StorageKey("city").Comment("区/市"),
		field.String("county").Nillable().Optional().StorageKey("county").Comment("区/县"),

		field.String("country_id").MaxLen(36).Nillable().Optional().StorageKey("country_id").Comment("国ID"),
		field.String("provice_id").MaxLen(36).Nillable().Optional().StorageKey("provice_id").Comment("省/市ID"),
		field.String("city_id").MaxLen(36).Nillable().Optional().StorageKey("city_id").Comment("区/市ID"),
		field.String("county_id").MaxLen(36).Nillable().Optional().StorageKey("county_id").Comment("区/县ID"),

		field.String("zip_code").MaxLen(8).Nillable().Optional().StorageKey("zip_code").Comment("邮政编码"),
		field.String("daddr").MaxLen(256).Nillable().Optional().StorageKey("daddr").Comment("详细地址"),
		field.String("name").MaxLen(64).Nillable().Optional().StorageKey("name").Comment("联系人"),

		field.String("mobile").MaxLen(64).Nillable().Optional().StorageKey("mobile").Comment("电话"),

		field.String("creator").MaxLen(36).Nillable().Optional().StorageKey("creator").Comment("创建者"),
	}
}

func (SysAddress) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (SysAddress) Indexes() []ent.Index {
	return []ent.Index{}
}

func (SysAddress) Hooks() []ent.Hook {
	return []ent.Hook{}
}
