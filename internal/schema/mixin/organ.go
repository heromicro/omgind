package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type OrganMixin struct {
	mixin.Schema
}

func (OrganMixin) Fields() []ent.Field {

	return []ent.Field{
		OrganIDField(),
	}
}

func (a OrganMixin) Indexes() []ent.Index {

	return []ent.Index{
		index.Fields("org_id"),
	}
}

func OrganIDField() ent.Field {
	return field.String("org_id").MaxLen(36).NotEmpty().Immutable().StorageKey("org_id").Comment("组织id")
}
