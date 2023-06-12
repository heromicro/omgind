package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/heromicro/omgind/internal/schema/mixin"
	"github.com/heromicro/omgind/pkg/helper/str"
)

// SysUser holds the schema definition for the SysUser entity.
type SysUser struct {
	ent.Schema
}

func (su SysUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.SoftDelMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
		mixin.ActiveMixin{},
	}
}

// Fields of the SysUser.
func (SysUser) Fields() []ent.Field {
	return []ent.Field{

		field.String("user_name").StorageKey("user_name").
			MinLen(5).MaxLen(128).NotEmpty().Immutable().Comment("用户名"),
		field.String("real_name").StorageKey("real_name").
			MaxLen(64).Nillable().Optional().Comment(""),
		field.String("first_name").StorageKey("first_name").
			MaxLen(31).Nillable().Optional().Comment("名"),
		field.String("last_name").StorageKey("last_name").
			MaxLen(31).Nillable().Optional().Comment("姓"),
		field.String("password").StorageKey("passwd").
			MaxLen(256).Sensitive().Comment("密码"),
		field.String("email").StorageKey("email").
			MaxLen(256).StructTag(`json:"email,omitempty"`).Comment("电子邮箱"),
		field.String("mobile").StorageKey("mobile").
			MaxLen(20).StructTag(`json:"mobile,omitempty"`).Comment("电话号码"),
		field.String("salt").DefaultFunc(func() string {
			return str.RandomBetween(10, 16)
		}).Comment("盐"),

		// field.Enum("gender").NamedValues("Male", "M", "Female", "F").StorageKey("gender").Nillable().Optional().Comment("性别"),
	}
}

func (su SysUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_name").Unique(),
	}
}

// Edges of the SysUser.
func (SysUser) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.To("organ", OrgOrgan.Type),
		// edge.To("userRoles", SysUserRole.Type).Comment("userroles: "),
	}
}
