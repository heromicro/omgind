package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type OwnerMixin struct {
	mixin.Schema
}

func (OwnerMixin) Fields() []ent.Field {

	return []ent.Field{
		OwnerIDField(),
	}
}

func (a OwnerMixin) Indexes() []ent.Index {

	return []ent.Index{
		index.Fields("owner_id"),
	}
}

func OwnerIDField() ent.Field {
	return field.String("owner_id").MaxLen(36).NotEmpty().Immutable().StorageKey("o_id").Comment("所有者id")
}
