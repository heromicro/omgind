package entity

//
//// SysCasbinRule holds the schema definition for the SysCasbinRule entity.
//type SysCasbinRule struct {
//	ent.Schema
//}
//
//func (scr SysCasbinRule) Mixin() []ent.Mixin {
//	return []ent.Mixin{
//		//mixin.IDMixin{},
//		mixin.SortMixin{},
//		mixin.TimeMixin{},
//	}
//}
//
//// Fields of the SysCasbinRule.
//func (SysCasbinRule) Fields() []ent.Field {
//	return []ent.Field{
//		field.String("PType").MaxLen(128).StorageKey("p_type").StructTag(`json:"p_type,omitempty"`).Nillable().Optional().Comment("规则类型"),
//		field.String("RoleID").MaxLen(128).StorageKey("v0").StructTag(`json:"v0,omitempty"`).Nillable().Optional().Comment("角色ID"),
//		field.String("Path").MaxLen(128).StorageKey("v1").StructTag(`json:"v1,omitempty"`).Nillable().Optional().Comment("api路径"),
//		field.String("Method").MaxLen(128).StorageKey("v2").StructTag(`json:"v2,omitempty"`).Nillable().Optional().Comment("api访问方法"),
//
//		field.String("v3").MaxLen(128).StorageKey("v3").StructTag(`json:"v3,omitempty"`).Nillable().Optional(),
//		field.String("v4").MaxLen(128).StorageKey("v4").StructTag(`json:"v4,omitempty"`).Nillable().Optional(),
//		field.String("v5").MaxLen(128).StorageKey("v5").StructTag(`json:"v5,omitempty"`).Nillable().Optional(),
//	}
//}
//
//// Edges of the SysCasbinRule.
//func (SysCasbinRule) Edges() []ent.Edge {
//	return nil
//}
