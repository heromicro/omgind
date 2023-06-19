package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TreeMixin struct {
	mixin.Schema
}

func (TreeMixin) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("tree_level").Min(1).Nillable().Optional().StorageKey("tree_level").StructTag(`json:"level"`).Comment("level in tree, toppest level is 1"),
		field.Bool("is_leaf").Default(true).Nillable().Optional().StorageKey("is_leaf").StructTag(`json:"is_leaf"`).Comment("is leaf node"),
		field.Text("tree_path").Nillable().Optional().StorageKey("t_path").StructTag(`json:"tree_path"`).Comment("tree path,topest is null or zero length string, subber has fathers ids join by slash(/), eg: pid1/pid2 "),
	}
}

func (a TreeMixin) Indexes() []ent.Index {
	return []ent.Index{}
}
