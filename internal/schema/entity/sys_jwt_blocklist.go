package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/schema/mixin"
)

type SysJwtBlock struct {
	ent.Schema
}

func (sjb SysJwtBlock) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.MemoMixin{},
		mixin.TimeMixin{},
		mixin.StatusMixin{},
	}
}

func (sjb SysJwtBlock) Fields() []ent.Field {
	return []ent.Field{

		field.Text("jwt").StorageKey("jwt").NotEmpty().Comment("jwt"),
	}
}

func (SysJwtBlock) Indexes() []ent.Index {
	return []ent.Index{}
}
