package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{

		field.Time("created_at").
			StorageKey("crtd_at").
			//StructTag(`json:"created_at,omitempty" sql:"crtd_at"`).
			Immutable().
			Default(time.Now).
			Comment("创建时间,由程序自动生成"),

		field.Time("updated_at").
			StorageKey("uptd_at").
			//StructTag(`json:"updated_at,omitempty" sql:"uptd_at"`).
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间,由程序自动生成"),

		field.Time("deleted_at").
			StorageKey("dltd_at").Nillable().Optional().
			StructTag(`json:"deleted_at,omitempty"`).
			Comment("删除时间,"),
	}
}

func (TimeMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("updated_at"),
		index.Fields("deleted_at"),
	}
}

func FieldCreateAt() ent.Field {
	return field.Time("created_at").
		StorageKey("crtd_at").
		//StructTag(`json:"created_at,omitempty" sql:"crtd_at"`).
		Immutable().
		Default(time.Now).Comment("创建时间,由程序自动生成")
}

func IndexCreateAt() ent.Index {
	return index.Fields("created_at")
}

// TODO:: Hooks
//func (TimeMixin) Hooks() []ent.Hook {
//	return []ent.Hook{
//		hook.On(func(mutator ent.Mutator) ent.Mutator {
//
//		}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne),
//		//func(next ent.Mutator) ent.Mutator {
//		//
//		//	//	fmt.Println(" +++++ ----- ======= 111 ")
//		//
//		//	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
//		//		//		fmt.Println(" +++++ ----- ======= 222 ")
//		//		//
//		//		//		switch op := m.Op(); {
//		//		//		case op.Is(ent.OpCreate):
//		//		//			//m.SetField("created_at", time.Now())
//		//		//
//		//		//		case op.Is(ent.OpUpdate | ent.OpUpdateOne):
//		//		//			//m.SetField("updated_at", time.Now())
//		//		//
//		//		//			//case op.Is(ent.OpDelete | ent.OpDeleteOne):
//		//		//			//	m.SetField("deleted_at", time.Now())
//		//		//
//		//		//		}
//		//		return next.Mutate(ctx, m)
//		//	})
//		//},
//	}
//}
