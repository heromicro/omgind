package entity

import (
	"github.com/heromicro/omgind/internal/scheme/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type SysTeamUser struct {
	ent.Schema
}

func (SysTeamUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
		mixin.ActiveMixin{},
		mixin.MemoMixin{},
		mixin.SoftDelMixin{},
	}
}

func (SysTeamUser) Fields() []ent.Field {

	return []ent.Field{

		field.String("team_id").MaxLen(36).StorageKey("team_id").Comment("systeam.id"),
		field.String("user_id").MaxLen(36).StorageKey("user_id").Comment("sysuser.id"),
	}
}

func (SysTeamUser) Edges() []ent.Edge {
	return []ent.Edge{

		edge.To("user", SysUser.Type).Unique().Required().Field("user_id"),

		edge.To("team", SysTeam.Type).Unique().Required().Field("team_id"),
	}
}

func (SysTeamUser) Indexes() []ent.Index {
	return []ent.Index{}
}

func (SysTeamUser) Hooks() []ent.Hook {
	return []ent.Hook{}
}
