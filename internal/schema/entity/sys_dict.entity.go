package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/schema/mixin"
)

type SysDict struct {
	ent.Schema
}

func (sd SysDict) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.MemoMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
		mixin.ActiveMixin{}}
}

func (sd SysDict) Fields() []ent.Field {

	return []ent.Field{

		field.String("name_cn").StorageKey("name_cn").
			MaxLen(128).StructTag(`json:"name_cn,omitempty"`).Comment("字典名（中）"),
		field.String("name_en").StorageKey("name_en").
			MaxLen(128).StructTag(`json:"name_en,omitempty"`).Comment("字典名（英）"),
	}
}

func (sd SysDict) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("SysDictItems", SysDictItem.Type),
	}
}

func (SysDict) Indexes() []ent.Index {
	return []ent.Index{}
}
