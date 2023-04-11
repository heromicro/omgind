// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/sysmenuaction"
)

// SysMenuAction is the model entity for the SysMenuAction schema.
type SysMenuAction struct {
	config `json:"-" sql:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// 是否删除
	IsDel bool `json:"is_del,omitempty"`
	// sort
	Sort int32 `json:"sort,omitempty" sql:"sort"`
	// 是否活跃
	IsActive bool `json:"is_active,omitempty"`
	// memo
	Memo *string `json:"memo,omitempty" sql:"memo"`
	// create time
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// update time
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// delete time,
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// 菜单ID
	MenuID string `json:"menu_id,omitempty"`
	// 动作编号
	Code string `json:"code,omitempty"`
	// 动作名称
	Name         string `json:"name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysMenuAction) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysmenuaction.FieldIsDel, sysmenuaction.FieldIsActive:
			values[i] = new(sql.NullBool)
		case sysmenuaction.FieldSort:
			values[i] = new(sql.NullInt64)
		case sysmenuaction.FieldID, sysmenuaction.FieldMemo, sysmenuaction.FieldMenuID, sysmenuaction.FieldCode, sysmenuaction.FieldName:
			values[i] = new(sql.NullString)
		case sysmenuaction.FieldCreatedAt, sysmenuaction.FieldUpdatedAt, sysmenuaction.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysMenuAction fields.
func (sma *SysMenuAction) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysmenuaction.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				sma.ID = value.String
			}
		case sysmenuaction.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				sma.IsDel = value.Bool
			}
		case sysmenuaction.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				sma.Sort = int32(value.Int64)
			}
		case sysmenuaction.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				sma.IsActive = value.Bool
			}
		case sysmenuaction.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				sma.Memo = new(string)
				*sma.Memo = value.String
			}
		case sysmenuaction.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sma.CreatedAt = new(time.Time)
				*sma.CreatedAt = value.Time
			}
		case sysmenuaction.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sma.UpdatedAt = new(time.Time)
				*sma.UpdatedAt = value.Time
			}
		case sysmenuaction.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sma.DeletedAt = new(time.Time)
				*sma.DeletedAt = value.Time
			}
		case sysmenuaction.FieldMenuID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field menu_id", values[i])
			} else if value.Valid {
				sma.MenuID = value.String
			}
		case sysmenuaction.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				sma.Code = value.String
			}
		case sysmenuaction.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sma.Name = value.String
			}
		default:
			sma.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SysMenuAction.
// This includes values selected through modifiers, order, etc.
func (sma *SysMenuAction) Value(name string) (ent.Value, error) {
	return sma.selectValues.Get(name)
}

// Update returns a builder for updating this SysMenuAction.
// Note that you need to call SysMenuAction.Unwrap() before calling this method if this SysMenuAction
// was returned from a transaction, and the transaction was committed or rolled back.
func (sma *SysMenuAction) Update() *SysMenuActionUpdateOne {
	return NewSysMenuActionClient(sma.config).UpdateOne(sma)
}

// Unwrap unwraps the SysMenuAction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sma *SysMenuAction) Unwrap() *SysMenuAction {
	_tx, ok := sma.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysMenuAction is not a transactional entity")
	}
	sma.config.driver = _tx.drv
	return sma
}

// String implements the fmt.Stringer.
func (sma *SysMenuAction) String() string {
	var builder strings.Builder
	builder.WriteString("SysMenuAction(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sma.ID))
	builder.WriteString("is_del=")
	builder.WriteString(fmt.Sprintf("%v", sma.IsDel))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", sma.Sort))
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", sma.IsActive))
	builder.WriteString(", ")
	if v := sma.Memo; v != nil {
		builder.WriteString("memo=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sma.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := sma.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := sma.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("menu_id=")
	builder.WriteString(sma.MenuID)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(sma.Code)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(sma.Name)
	builder.WriteByte(')')
	return builder.String()
}

// SysMenuActions is a parsable slice of SysMenuAction.
type SysMenuActions []*SysMenuAction
