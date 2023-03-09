// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/sysjwtblock"
)

// SysJwtBlock is the model entity for the SysJwtBlock schema.
type SysJwtBlock struct {
	config `json:"-" sql:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// 是否删除
	IsDel bool `json:"is_del,omitempty"`
	// memo
	Memo *string `json:"memo,omitempty" sql:"memo"`
	// create time
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// update time
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// delete time,
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// 是否活跃
	IsActive bool `json:"is_active,omitempty"`
	// jwt
	Jwt string `json:"jwt,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysJwtBlock) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysjwtblock.FieldIsDel, sysjwtblock.FieldIsActive:
			values[i] = new(sql.NullBool)
		case sysjwtblock.FieldID, sysjwtblock.FieldMemo, sysjwtblock.FieldJwt:
			values[i] = new(sql.NullString)
		case sysjwtblock.FieldCreatedAt, sysjwtblock.FieldUpdatedAt, sysjwtblock.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysJwtBlock", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysJwtBlock fields.
func (sjb *SysJwtBlock) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysjwtblock.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				sjb.ID = value.String
			}
		case sysjwtblock.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				sjb.IsDel = value.Bool
			}
		case sysjwtblock.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				sjb.Memo = new(string)
				*sjb.Memo = value.String
			}
		case sysjwtblock.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sjb.CreatedAt = new(time.Time)
				*sjb.CreatedAt = value.Time
			}
		case sysjwtblock.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sjb.UpdatedAt = new(time.Time)
				*sjb.UpdatedAt = value.Time
			}
		case sysjwtblock.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sjb.DeletedAt = new(time.Time)
				*sjb.DeletedAt = value.Time
			}
		case sysjwtblock.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				sjb.IsActive = value.Bool
			}
		case sysjwtblock.FieldJwt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field jwt", values[i])
			} else if value.Valid {
				sjb.Jwt = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SysJwtBlock.
// Note that you need to call SysJwtBlock.Unwrap() before calling this method if this SysJwtBlock
// was returned from a transaction, and the transaction was committed or rolled back.
func (sjb *SysJwtBlock) Update() *SysJwtBlockUpdateOne {
	return NewSysJwtBlockClient(sjb.config).UpdateOne(sjb)
}

// Unwrap unwraps the SysJwtBlock entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sjb *SysJwtBlock) Unwrap() *SysJwtBlock {
	_tx, ok := sjb.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysJwtBlock is not a transactional entity")
	}
	sjb.config.driver = _tx.drv
	return sjb
}

// String implements the fmt.Stringer.
func (sjb *SysJwtBlock) String() string {
	var builder strings.Builder
	builder.WriteString("SysJwtBlock(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sjb.ID))
	builder.WriteString("is_del=")
	builder.WriteString(fmt.Sprintf("%v", sjb.IsDel))
	builder.WriteString(", ")
	if v := sjb.Memo; v != nil {
		builder.WriteString("memo=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sjb.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := sjb.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := sjb.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", sjb.IsActive))
	builder.WriteString(", ")
	builder.WriteString("jwt=")
	builder.WriteString(sjb.Jwt)
	builder.WriteByte(')')
	return builder.String()
}

// SysJwtBlocks is a parsable slice of SysJwtBlock.
type SysJwtBlocks []*SysJwtBlock
