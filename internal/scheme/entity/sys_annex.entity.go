package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/scheme/mixin"
)

type SysAnnex struct {
	ent.Schema
}

func (SysAnnex) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.MemoMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
		mixin.ActiveMixin{},
		mixin.SoftDelMixin{},
	}
}

func (SysAnnex) Fields() []ent.Field {

	return []ent.Field{

		field.String("name").MaxLen(128).MinLen(4).Nillable().Optional().StorageKey("name").Comment("文件名称"),

		field.String("file_path").Nillable().Optional().StorageKey("file_path").Comment("文件路径"),
	}
}

func (SysAnnex) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (SysAnnex) Indexes() []ent.Index {
	return []ent.Index{}
}
