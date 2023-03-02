package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type SoftDelMixin struct {
	mixin.Schema
}

func (i SoftDelMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("is_del").Default(false).StructTag(`json:"is_del,omitempty"`).Comment("是否删除"),
	}
}

func (i SoftDelMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("is_del"),
	}
}
