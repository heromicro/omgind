package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/schema/mixin"
)

type SysDictItem struct {
	ent.Schema
}

func (sdd SysDictItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.SoftDelMixin{},
		mixin.MemoMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
		mixin.ActiveMixin{},
	}
}

func (sdd SysDictItem) Fields() []ent.Field {
	return []ent.Field{

		field.String("label").StorageKey("label").
			MaxLen(128).StructTag(`json:"label,omitempty"`).Comment("显示值"),
		field.Int("value").StorageKey("val").StructTag(`json:"value,omitempty"`).Comment("字典值"),

		field.String("dict_id").MaxLen(36).Nillable().Optional().Comment("sys_dict.id"),
	}
}

func (sdd SysDictItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("dict", SysDict.Type).Field("dict_id").Ref("items").Unique(),
	}
}

func (SysDictItem) Indexes() []ent.Index {
	return []ent.Index{}
}
