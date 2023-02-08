package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

//////////////////////////////////

type SortMixin struct {
	mixin.Schema
}

func (m SortMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("sort").
			Default(9999).
			StructTag(`json:"sort,omitempty"`).
			Comment("排序, 在数据库里的排序"),
	}
}

func (m SortMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sort"),
	}
}
