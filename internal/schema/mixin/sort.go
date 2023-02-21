package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

//////////////////////////////////

type SortMixin struct {
	mixin.Schema
}

func (m SortMixin) Fields() []ent.Field {
	return []ent.Field{
		FieldSort(9999),
	}
}

func (m SortMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sort"),
	}
}

func FieldSort(dvalue int32) ent.Field {

	f := field.Int32("sort").Default(dvalue)

	return f.StorageKey("sort").
		StructTag(`json:"sort,omitempty" sql:"sort"`).Comment("sort")
}

func IndexSort() ent.Index {
	return index.Fields("sort")
}
