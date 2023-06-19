package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type OwnerOrgMixin struct {
	mixin.Schema
}

func (OwnerOrgMixin) Fields() []ent.Field {

	return []ent.Field{
		FieldOrgID(true, true, false),
	}
}

func (a OwnerOrgMixin) Indexes() []ent.Index {

	return []ent.Index{
		IndexOrgID(),
	}
}

func FieldOrgID(nillable bool, immutable bool, notEmpty bool) ent.Field {

	f := field.String("org_id").MaxLen(36)
	if nillable {
		f = f.Nillable().Optional()
	}
	if immutable {
		f = f.Immutable()
	}
	if notEmpty {
		f = f.NotEmpty()
	}

	return f.StorageKey("org_id").
		StructTag(`json:"org_id,omitempty" sql:"org_id"`).Comment("organization id")
}

func IndexOrgID() ent.Index {
	return index.Fields("org_id")
}
