package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/heromicro/omgind/internal/schema/mixin"
)


type SysLogging struct {
	ent.Schema
}

func (SysLogging) Annotations() []schema.Annotation  {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_logging"},
	}
}

func (SysLogging) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.MemoMixin{},
	}
}

// Fields of the SysLogging.
func (SysLogging) Fields() []ent.Field {
	return []ent.Field{

		field.String("level").MaxLen(32).Comment("日志级别"),
		field.String("trace_id").MaxLen(128).Comment("跟踪ID"),
		field.String("user_id").MaxLen(128).Comment("用户ID"),
		field.String("tag").MaxLen(128).Comment("Tag"),

		field.String("version").MaxLen(64).Comment("版本号"),
		field.String("message").Comment("消息"),
		field.Text("data").Immutable().Optional().Nillable().Comment("日志数据(json string)"),
		field.Text("error_stack").Comment("日志数据(json string)"),

		mixin.FieldCreateAt(),
	}
}

func (SysLogging) Indexes() []ent.Index {

	return []ent.Index{
		index.Fields("level"),
		index.Fields("trace_id"),
		index.Fields("user_id"),
		index.Fields("tag"),
		mixin.IndexCreateAt(),
	}
}


