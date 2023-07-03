package entity

import (
	"github.com/heromicro/omgind/internal/scheme/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type SysTeam struct {
	ent.Schema
}

func (SysTeam) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
		mixin.ActiveMixin{},
		mixin.MemoMixin{},
		mixin.SoftDelMixin{},
	}
}

func (SysTeam) Fields() []ent.Field {

	return []ent.Field{
		field.String("name").MaxLen(64).Nillable().Optional().StorageKey("name").Comment("名称"),
		field.String("code").MaxLen(64).Nillable().Optional().StorageKey("code").Comment("助记码"),
	}
}

func (SysTeam) Edges() []ent.Edge {
	return []ent.Edge{

		edge.To("users", SysUser.Type).Through("team_users", SysTeamUser.Type),
	}
}

func (SysTeam) Indexes() []ent.Index {
	return []ent.Index{}
}

func (SysTeam) Hooks() []ent.Hook {
	return []ent.Hook{}
}
