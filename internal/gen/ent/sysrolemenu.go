// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/sysrolemenu"
)

// SysRoleMenu is the model entity for the SysRoleMenu schema.
type SysRoleMenu struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// 是否删除
	IsDel bool `json:"is_del,omitempty"`
	// create time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// delete time,
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// 角色ID, sys_role.id
	RoleID string `json:"role_id,omitempty"`
	// 菜单ID
	MenuID string `json:"menu_id,omitempty"`
	// 菜单ID, sys_menu_action.id
	ActionID *string `json:"action_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysRoleMenu) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysrolemenu.FieldIsDel:
			values[i] = new(sql.NullBool)
		case sysrolemenu.FieldID, sysrolemenu.FieldRoleID, sysrolemenu.FieldMenuID, sysrolemenu.FieldActionID:
			values[i] = new(sql.NullString)
		case sysrolemenu.FieldCreatedAt, sysrolemenu.FieldUpdatedAt, sysrolemenu.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysRoleMenu", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysRoleMenu fields.
func (srm *SysRoleMenu) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysrolemenu.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				srm.ID = value.String
			}
		case sysrolemenu.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				srm.IsDel = value.Bool
			}
		case sysrolemenu.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				srm.CreatedAt = value.Time
			}
		case sysrolemenu.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				srm.UpdatedAt = value.Time
			}
		case sysrolemenu.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				srm.DeletedAt = new(time.Time)
				*srm.DeletedAt = value.Time
			}
		case sysrolemenu.FieldRoleID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value.Valid {
				srm.RoleID = value.String
			}
		case sysrolemenu.FieldMenuID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field menu_id", values[i])
			} else if value.Valid {
				srm.MenuID = value.String
			}
		case sysrolemenu.FieldActionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field action_id", values[i])
			} else if value.Valid {
				srm.ActionID = new(string)
				*srm.ActionID = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SysRoleMenu.
// Note that you need to call SysRoleMenu.Unwrap() before calling this method if this SysRoleMenu
// was returned from a transaction, and the transaction was committed or rolled back.
func (srm *SysRoleMenu) Update() *SysRoleMenuUpdateOne {
	return NewSysRoleMenuClient(srm.config).UpdateOne(srm)
}

// Unwrap unwraps the SysRoleMenu entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (srm *SysRoleMenu) Unwrap() *SysRoleMenu {
	_tx, ok := srm.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysRoleMenu is not a transactional entity")
	}
	srm.config.driver = _tx.drv
	return srm
}

// String implements the fmt.Stringer.
func (srm *SysRoleMenu) String() string {
	var builder strings.Builder
	builder.WriteString("SysRoleMenu(")
	builder.WriteString(fmt.Sprintf("id=%v, ", srm.ID))
	builder.WriteString("is_del=")
	builder.WriteString(fmt.Sprintf("%v", srm.IsDel))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(srm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(srm.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := srm.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("role_id=")
	builder.WriteString(srm.RoleID)
	builder.WriteString(", ")
	builder.WriteString("menu_id=")
	builder.WriteString(srm.MenuID)
	builder.WriteString(", ")
	if v := srm.ActionID; v != nil {
		builder.WriteString("action_id=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// SysRoleMenus is a parsable slice of SysRoleMenu.
type SysRoleMenus []*SysRoleMenu
