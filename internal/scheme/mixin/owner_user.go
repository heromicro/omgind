package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type OwnerUserMixin struct {
	mixin.Schema
}

func (OwnerUserMixin) Fields() []ent.Field {

	return []ent.Field{
		FieldOwnerID(true, false, false),
	}
}

func (a OwnerUserMixin) Indexes() []ent.Index {

	return []ent.Index{
		IndexOwnerID(),
	}
}

func FieldOwnerID(nillable bool, immutable bool, notEmpty bool) ent.Field {

	f := field.String("user_id").MaxLen(36)
	if nillable {
		f = f.Nillable().Optional()
	}
	if immutable {
		f = f.Immutable()
	}
	if notEmpty {
		f = f.NotEmpty()
	}
	return f.StorageKey("user_id").
		StructTag(`json:"user_id,omitempty" sql:"user_id"`).Comment("user id")
}

func IndexOwnerID() ent.Index {
	return index.Fields("user_id")
}
