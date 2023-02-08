package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type ActiveMixin struct {
	mixin.Schema
}

func (ActiveMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("is_active").Default(true).Comment("是否活跃"),
	}
}

func (a ActiveMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("is_active"),
	}
}

