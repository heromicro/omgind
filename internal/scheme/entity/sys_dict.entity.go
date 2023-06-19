package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/heromicro/omgind/internal/scheme/mixin"
)

type SysDict struct {
	ent.Schema
}

func (sd SysDict) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.SoftDelMixin{},
		mixin.MemoMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
		mixin.ActiveMixin{}}
}

func (sd SysDict) Fields() []ent.Field {

	return []ent.Field{

		field.String("name_cn").StorageKey("name_cn").
			MaxLen(128).StructTag(`json:"name_cn,omitempty"`).Comment("字典名（中）"),
		field.String("name_en").StorageKey("name_en").
			MaxLen(128).StructTag(`json:"name_en,omitempty"`).Comment("字典名（英）"),

		field.Enum("tipe").Values("int", "string").Default("int").StorageKey("tipe").StructTag(`json:"tipe,omitempty"`).Comment("值类型"),
	}
}

func (sd SysDict) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("items", SysDictItem.Type),

		// edge.To("staff_gender", OrgStaff.Type).Unique(),
		// edge.To("staff_empyst", OrgStaff.Type).Unique(),

		//

	}
}

func (SysDict) Indexes() []ent.Index {
	return []ent.Index{}
}

func makeupDictFrom(name string, refname string, fieldname string) []ent.Edge {

	return []ent.Edge{
		edge.From(name, SysDict.Type).Ref(refname).Field(fieldname).Unique(),
	}
}
