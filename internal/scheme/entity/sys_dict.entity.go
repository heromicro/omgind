package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/heromicro/omgind/internal/scheme/enumtipe"
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

		field.String("dict_key").StorageKey("dict_key").
			MaxLen(64).StructTag(`json:"dict_key,omitempty"`).Comment("字典键"),

		field.Enum("val_tipe").GoType(enumtipe.DictValueTipe("")).Default(enumtipe.DictValueTipeInt.String()).StorageKey("val_tipe").StructTag(`json:"tipe,omitempty"`).Comment("值类型"),
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
	return []ent.Index{
		index.Fields("dict_key").Unique(),
	}
}

func makeupDictFrom(name string, refname string, fieldname string) []ent.Edge {

	return []ent.Edge{
		edge.From(name, SysDict.Type).Ref(refname).Field(fieldname).Unique(),
	}
}
