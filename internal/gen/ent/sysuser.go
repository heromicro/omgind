// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/sysuser"
)

// SysUser is the model entity for the SysUser schema.
type SysUser struct {
	config `json:"-" sql:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// 是否删除
	IsDel bool `json:"is_del,omitempty"`
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
	// 用户名
	UserName string `json:"user_name,omitempty"`
	// RealName holds the value of the "real_name" field.
	RealName *string `json:"real_name,omitempty"`
	// 名
	FirstName *string `json:"first_name,omitempty"`
	// 姓
	LastName *string `json:"last_name,omitempty"`
	// 密码
	Password string `json:"-" sql:"-"`
	// 电子邮箱
	Email string `json:"email,omitempty"`
	// 电话号码
	Mobile string `json:"mobile,omitempty"`
	// 盐
	Salt         string `json:"salt,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysuser.FieldIsDel, sysuser.FieldIsActive:
			values[i] = new(sql.NullBool)
		case sysuser.FieldSort:
			values[i] = new(sql.NullInt64)
		case sysuser.FieldID, sysuser.FieldUserName, sysuser.FieldRealName, sysuser.FieldFirstName, sysuser.FieldLastName, sysuser.FieldPassword, sysuser.FieldEmail, sysuser.FieldMobile, sysuser.FieldSalt:
			values[i] = new(sql.NullString)
		case sysuser.FieldCreatedAt, sysuser.FieldUpdatedAt, sysuser.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysUser fields.
func (su *SysUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysuser.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				su.ID = value.String
			}
		case sysuser.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				su.IsDel = value.Bool
			}
		case sysuser.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				su.Sort = int32(value.Int64)
			}
		case sysuser.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				su.CreatedAt = new(time.Time)
				*su.CreatedAt = value.Time
			}
		case sysuser.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				su.UpdatedAt = new(time.Time)
				*su.UpdatedAt = value.Time
			}
		case sysuser.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				su.DeletedAt = new(time.Time)
				*su.DeletedAt = value.Time
			}
		case sysuser.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				su.IsActive = value.Bool
			}
		case sysuser.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_name", values[i])
			} else if value.Valid {
				su.UserName = value.String
			}
		case sysuser.FieldRealName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field real_name", values[i])
			} else if value.Valid {
				su.RealName = new(string)
				*su.RealName = value.String
			}
		case sysuser.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				su.FirstName = new(string)
				*su.FirstName = value.String
			}
		case sysuser.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				su.LastName = new(string)
				*su.LastName = value.String
			}
		case sysuser.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				su.Password = value.String
			}
		case sysuser.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				su.Email = value.String
			}
		case sysuser.FieldMobile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mobile", values[i])
			} else if value.Valid {
				su.Mobile = value.String
			}
		case sysuser.FieldSalt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field salt", values[i])
			} else if value.Valid {
				su.Salt = value.String
			}
		default:
			su.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SysUser.
// This includes values selected through modifiers, order, etc.
func (su *SysUser) Value(name string) (ent.Value, error) {
	return su.selectValues.Get(name)
}

// Update returns a builder for updating this SysUser.
// Note that you need to call SysUser.Unwrap() before calling this method if this SysUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (su *SysUser) Update() *SysUserUpdateOne {
	return NewSysUserClient(su.config).UpdateOne(su)
}

// Unwrap unwraps the SysUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (su *SysUser) Unwrap() *SysUser {
	_tx, ok := su.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysUser is not a transactional entity")
	}
	su.config.driver = _tx.drv
	return su
}

// String implements the fmt.Stringer.
func (su *SysUser) String() string {
	var builder strings.Builder
	builder.WriteString("SysUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", su.ID))
	builder.WriteString("is_del=")
	builder.WriteString(fmt.Sprintf("%v", su.IsDel))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", su.Sort))
	builder.WriteString(", ")
	if v := su.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := su.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := su.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", su.IsActive))
	builder.WriteString(", ")
	builder.WriteString("user_name=")
	builder.WriteString(su.UserName)
	builder.WriteString(", ")
	if v := su.RealName; v != nil {
		builder.WriteString("real_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := su.FirstName; v != nil {
		builder.WriteString("first_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := su.LastName; v != nil {
		builder.WriteString("last_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(su.Email)
	builder.WriteString(", ")
	builder.WriteString("mobile=")
	builder.WriteString(su.Mobile)
	builder.WriteString(", ")
	builder.WriteString("salt=")
	builder.WriteString(su.Salt)
	builder.WriteByte(')')
	return builder.String()
}

// SysUsers is a parsable slice of SysUser.
type SysUsers []*SysUser
