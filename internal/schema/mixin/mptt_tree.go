package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type MpttTreeMixin struct {
	mixin.Schema
}

func (MpttTreeMixin) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("tree_id").Nillable().Optional().StorageKey("tree_id").StructTag(`json:"tree_id"`).Comment("tree id"),
		field.Int64("tree_level").Nillable().Optional().StorageKey("tree_level").StructTag(`json:"level"`).Comment("层级"),
		field.Int64("tree_left").Nillable().Optional().StorageKey("tree_left").StructTag(`json:"tree_left"`).Comment("mptt左值"),
		field.Int64("tree_right").Nillable().Optional().StorageKey("tree_right").StructTag(`json:"tree_right"`).Comment("mptt右值"),
		field.Bool("is_leaf").Default(true).Nillable().Optional().StorageKey("is_leaf").StructTag(`json:"is_leaf"`).Comment("是否是子叶"),

		field.Text("tree_path").Nillable().Optional().StorageKey("t_path").StructTag(`json:"tree_path"`).Comment("tree path,topest is null or zero length string, subber has fathers ids join by slash(/), eg: pid1/pid2 "),
	}
}

func (a MpttTreeMixin) Indexes() []ent.Index {

	return []ent.Index{
		index.Fields("tree_id"),
		index.Fields("tree_id", "tree_left", "tree_right"),
	}
}
