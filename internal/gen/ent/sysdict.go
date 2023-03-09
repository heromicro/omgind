// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/sysdict"
)

// SysDict is the model entity for the SysDict schema.
type SysDict struct {
	config `json:"-" sql:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// 是否删除
	IsDel bool `json:"is_del,omitempty"`
	// memo
	Memo *string `json:"memo,omitempty" sql:"memo"`
	// sort
	Sort int32 `json:"sort,omitempty" sql:"sort"`
	// create time
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// update time
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// delete time,
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// 是否活跃
	IsActive bool `json:"is_active,omitempty"`
	// 字典名（中）
	NameCn string `json:"name_cn,omitempty"`
	// 字典名（英）
	NameEn string `json:"name_en,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysDict) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysdict.FieldIsDel, sysdict.FieldIsActive:
			values[i] = new(sql.NullBool)
		case sysdict.FieldSort:
			values[i] = new(sql.NullInt64)
		case sysdict.FieldID, sysdict.FieldMemo, sysdict.FieldNameCn, sysdict.FieldNameEn:
			values[i] = new(sql.NullString)
		case sysdict.FieldCreatedAt, sysdict.FieldUpdatedAt, sysdict.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysDict", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysDict fields.
func (sd *SysDict) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysdict.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				sd.ID = value.String
			}
		case sysdict.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				sd.IsDel = value.Bool
			}
		case sysdict.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				sd.Memo = new(string)
				*sd.Memo = value.String
			}
		case sysdict.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				sd.Sort = int32(value.Int64)
			}
		case sysdict.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sd.CreatedAt = new(time.Time)
				*sd.CreatedAt = value.Time
			}
		case sysdict.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sd.UpdatedAt = new(time.Time)
				*sd.UpdatedAt = value.Time
			}
		case sysdict.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sd.DeletedAt = new(time.Time)
				*sd.DeletedAt = value.Time
			}
		case sysdict.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				sd.IsActive = value.Bool
			}
		case sysdict.FieldNameCn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_cn", values[i])
			} else if value.Valid {
				sd.NameCn = value.String
			}
		case sysdict.FieldNameEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_en", values[i])
			} else if value.Valid {
				sd.NameEn = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SysDict.
// Note that you need to call SysDict.Unwrap() before calling this method if this SysDict
// was returned from a transaction, and the transaction was committed or rolled back.
func (sd *SysDict) Update() *SysDictUpdateOne {
	return NewSysDictClient(sd.config).UpdateOne(sd)
}

// Unwrap unwraps the SysDict entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sd *SysDict) Unwrap() *SysDict {
	_tx, ok := sd.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysDict is not a transactional entity")
	}
	sd.config.driver = _tx.drv
	return sd
}

// String implements the fmt.Stringer.
func (sd *SysDict) String() string {
	var builder strings.Builder
	builder.WriteString("SysDict(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sd.ID))
	builder.WriteString("is_del=")
	builder.WriteString(fmt.Sprintf("%v", sd.IsDel))
	builder.WriteString(", ")
	if v := sd.Memo; v != nil {
		builder.WriteString("memo=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", sd.Sort))
	builder.WriteString(", ")
	if v := sd.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := sd.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := sd.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", sd.IsActive))
	builder.WriteString(", ")
	builder.WriteString("name_cn=")
	builder.WriteString(sd.NameCn)
	builder.WriteString(", ")
	builder.WriteString("name_en=")
	builder.WriteString(sd.NameEn)
	builder.WriteByte(')')
	return builder.String()
}

// SysDicts is a parsable slice of SysDict.
type SysDicts []*SysDict
