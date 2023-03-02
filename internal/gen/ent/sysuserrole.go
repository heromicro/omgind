// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/sysuserrole"
)

// SysUserRole is the model entity for the SysUserRole schema.
type SysUserRole struct {
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
	// 用户ID, sys_user.id
	UserID string `json:"user_id,omitempty"`
	// 角色ID, sys_role.id
	RoleID string `json:"role_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysUserRole) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysuserrole.FieldIsDel:
			values[i] = new(sql.NullBool)
		case sysuserrole.FieldID, sysuserrole.FieldUserID, sysuserrole.FieldRoleID:
			values[i] = new(sql.NullString)
		case sysuserrole.FieldCreatedAt, sysuserrole.FieldUpdatedAt, sysuserrole.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysUserRole", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysUserRole fields.
func (sur *SysUserRole) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysuserrole.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				sur.ID = value.String
			}
		case sysuserrole.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				sur.IsDel = value.Bool
			}
		case sysuserrole.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sur.CreatedAt = value.Time
			}
		case sysuserrole.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sur.UpdatedAt = value.Time
			}
		case sysuserrole.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sur.DeletedAt = new(time.Time)
				*sur.DeletedAt = value.Time
			}
		case sysuserrole.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				sur.UserID = value.String
			}
		case sysuserrole.FieldRoleID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value.Valid {
				sur.RoleID = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SysUserRole.
// Note that you need to call SysUserRole.Unwrap() before calling this method if this SysUserRole
// was returned from a transaction, and the transaction was committed or rolled back.
func (sur *SysUserRole) Update() *SysUserRoleUpdateOne {
	return NewSysUserRoleClient(sur.config).UpdateOne(sur)
}

// Unwrap unwraps the SysUserRole entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sur *SysUserRole) Unwrap() *SysUserRole {
	_tx, ok := sur.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysUserRole is not a transactional entity")
	}
	sur.config.driver = _tx.drv
	return sur
}

// String implements the fmt.Stringer.
func (sur *SysUserRole) String() string {
	var builder strings.Builder
	builder.WriteString("SysUserRole(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sur.ID))
	builder.WriteString("is_del=")
	builder.WriteString(fmt.Sprintf("%v", sur.IsDel))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(sur.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sur.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := sur.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(sur.UserID)
	builder.WriteString(", ")
	builder.WriteString("role_id=")
	builder.WriteString(sur.RoleID)
	builder.WriteByte(')')
	return builder.String()
}

// SysUserRoles is a parsable slice of SysUserRole.
type SysUserRoles []*SysUserRole
