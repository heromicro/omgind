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

func FieldActive(nillable bool, dvalue *bool) ent.Field {

	f := field.Bool("is_active")
	if nillable {
		f = f.Nillable().Optional()
	}
	if dvalue != nil {
		f = f.Default(*dvalue)
	}
	return f.StorageKey("is_active").
		StructTag(`json:"is_active,omitempty" sql:"is_active"`).Comment("is actvie")
}

func IndexActive() ent.Index {
	return index.Fields("is_active")
}
