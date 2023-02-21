package mixin

import (
	"entgo.io/ent"
	"github.com/gotidy/ptr"

	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type MemoMixin struct {
	mixin.Schema
}

func (MemoMixin) Fields() []ent.Field {
	return []ent.Field{
		FieldMemo(1024, ptr.String(""), true, false),
	}
}

func FieldMemo(maxlen int, dvalue *string, nillable bool, notEmpty bool) ent.Field {

	f := field.String("memo").MaxLen(maxlen)
	if dvalue != nil {
		f = f.Default(*dvalue)
	}
	if nillable {
		f = f.Nillable().Optional()
	}
	if notEmpty {
		f = f.NotEmpty()
	}

	return f.StorageKey("memo").
		StructTag(`json:"memo,omitempty" sql:"memo"`).Comment("memo")
}
