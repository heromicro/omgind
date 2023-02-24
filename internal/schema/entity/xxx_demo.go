package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/schema/mixin"
)

type XxxDemo struct {
	ent.Schema
}

func (xd XxxDemo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.MemoMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
		mixin.ActiveMixin{},
	}
}

func (xd XxxDemo) Fields() []ent.Field {

	return []ent.Field{

		field.String("code").StorageKey("code").
			MaxLen(128).StructTag(`json:"code,omitempty"`).Comment("编号"),
		field.String("name").StorageKey("name").
			MaxLen(128).StructTag(`json:"name,omitempty"`).Comment("名称"),
	}
}

func (xd XxxDemo) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (xd XxxDemo) Indexes() []ent.Index {
	return []ent.Index{}
}

func (xd XxxDemo) Hooks() []ent.Hook {
	return []ent.Hook{
		//hook.On(func(next ent.Mutator) ent.Mutator {
		//	return hook.XxxDemoFunc(func(ctx context.Context, m *ent.XxxDemoMutation) (ent.Value, error) {
		//
		//		return next.Mutate(ctx, m)
		//	})
		//}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne|ent.OpDelete|ent.OpDeleteOne),
	}
}
