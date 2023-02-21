package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type CreatorMixin struct {
	mixin.Schema
}

func (CreatorMixin) Fields() []ent.Field {

	return []ent.Field{
		FieldCreator(true, true, false),
	}
}

func (a CreatorMixin) Indexes() []ent.Index {

	return []ent.Index{
		IndexCreator(),
	}
}

func FieldCreator(nillable bool, immutable bool, notEmpty bool) ent.Field {

	f := field.String("creator").MaxLen(36)
	if nillable {
		f = f.Nillable().Optional()
	}
	if immutable {
		f = f.Immutable()
	}
	if notEmpty {
		f = f.NotEmpty()
	}
	return f.StorageKey("creator").
		StructTag(`json:"creator,omitempty" sql:"creator"`).Comment("creator id")
}

func IndexCreator() ent.Index {
	return index.Fields("creator")
}
