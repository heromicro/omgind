package mixin

import (
	"entgo.io/ent"

	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type MemoMixin struct {
	mixin.Schema
}

func (MemoMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("memo").MaxLen(1024).Default("").Comment("备注"),
	}
}
