package entity

import (
	"github.com/heromicro/omgind/internal/scheme/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type SysDistrict struct {
	ent.Schema
}

func (SysDistrict) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.SoftDelMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
		mixin.ActiveMixin{},

		mixin.MpttTreeMixin{},
	}
}

func (SysDistrict) Fields() []ent.Field {

	return []ent.Field{

		field.String("name").MaxLen(128).Nillable().Optional().StorageKey("name").Comment("名称"),
		field.String("name_en").MaxLen(128).Nillable().Optional().StorageKey("name_en").Comment("English名称"),

		field.String("sname").MaxLen(64).Nillable().Optional().StorageKey("sname").Comment("短名称"),
		field.String("sname_en").MaxLen(64).Nillable().Optional().StorageKey("sname_en").Comment("English短名称"),

		field.String("abbr").MaxLen(16).Nillable().Optional().StorageKey("abbr").Comment("简称Abbreviation"),

		field.String("st_code").MaxLen(16).Nillable().Optional().StorageKey("stcode").Comment("统计局区域编码"),
		field.String("initials").MaxLen(32).Nillable().Optional().StorageKey("initials").Comment("简拼"),
		field.String("pinyin").MaxLen(128).Nillable().Optional().StorageKey("pinyin").Comment("拼音"),

		field.String("parent_id").MaxLen(36).Nillable().Optional().StorageKey("pid").Comment("父级id"),

		field.Float("longitude").Nillable().Optional().StorageKey("longitude").Comment("经度"),
		field.Float("latitude").Nillable().Optional().StorageKey("latitude").Comment("经度"),
		field.String("area_code").MaxLen(8).Nillable().Optional().StorageKey("area_code").Comment("电话区号码"),
		field.String("zip_code").MaxLen(8).Nillable().Optional().StorageKey("zip_code").Comment("邮政编码"),
		field.String("merge_name").MaxLen(1024).Nillable().Optional().StorageKey("mname").Comment("带前缀全名称"),
		field.String("merge_sname").MaxLen(1024).Nillable().Optional().StorageKey("msname").Comment("带前缀简名称"),
		field.String("extra").MaxLen(64).Nillable().Optional().StorageKey("extra").Comment("附加信息"),
		field.String("suffix").MaxLen(32).Nillable().Optional().StorageKey("suffix").Comment("后缀, 如 省, 自治区, 旗, 盟"),

		field.Bool("is_hot").Default(false).Nillable().Optional().Comment("热门城市"),
		field.Bool("is_real").Default(true).Nillable().Optional().StorageKey("is_r").Comment("是否虚拟区域"),
		field.Bool("is_main").Default(true).Nillable().Optional().StorageKey("is_m").Comment("是否虚拟区域"),
		field.Bool("is_direct").Default(true).Nillable().Optional().StorageKey("is_d").Comment("是否是直辖"),

		// field.JSON("tag", []string{}).Default([]string{}).Optional().Comment("tag"),

	}
}

func (SysDistrict) Edges() []ent.Edge {

	return []ent.Edge{
		edge.To("children", SysDistrict.Type).From("parent").Field("parent_id").Unique(),
	}
}

func (SysDistrict) Indexes() []ent.Index {
	return []ent.Index{
		// index.Fields("parent_id"),
		index.Fields("is_hot"),
		index.Fields("is_real"),
		index.Fields("is_main"),
		index.Fields("is_direct"),
	}
}
