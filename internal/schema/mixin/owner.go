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
		FieldOwnerID(true, false, false),
	}
}

func (a OwnerMixin) Indexes() []ent.Index {

	return []ent.Index{
		IndexOwnerID(),
	}
}

func FieldOwnerID(nillable bool, immutable bool, notEmpty bool) ent.Field {

	f := field.String("owner_id").MaxLen(36)
	if nillable {
		f = f.Nillable().Optional()
	}
	if immutable {
		f = f.Immutable()
	}
	if notEmpty {
		f = f.NotEmpty()
	}
	return f.StorageKey("owner_id").
		StructTag(`json:"owner_id,omitempty" sql:"owner_id"`).Comment("owner id")
}

func IndexOwnerID() ent.Index {
	return index.Fields("owner_id")
}
