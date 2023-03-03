// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/syslogging"
)

// SysLogging is the model entity for the SysLogging schema.
type SysLogging struct {
	config `json:"-" sql:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// create time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// delete time,
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// 是否删除
	IsDel bool `json:"is_del,omitempty"`
	// memo
	Memo *string `json:"memo,omitempty" sql:"memo"`
	// 日志级别
	Level *string `json:"level,omitempty"`
	// 跟踪ID
	TraceID *string `json:"trace_id,omitempty"`
	// 用户ID
	UserID *string `json:"user_id,omitempty"`
	// Tag
	Tag *string `json:"tag,omitempty"`
	// 版本号
	Version *string `json:"version,omitempty"`
	// 消息
	Message *string `json:"message,omitempty"`
	// 日志数据(json string)
	Data *string `json:"data,omitempty"`
	// 日志数据(json string)
	ErrorStack *string `json:"error_stack,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysLogging) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case syslogging.FieldIsDel:
			values[i] = new(sql.NullBool)
		case syslogging.FieldID, syslogging.FieldMemo, syslogging.FieldLevel, syslogging.FieldTraceID, syslogging.FieldUserID, syslogging.FieldTag, syslogging.FieldVersion, syslogging.FieldMessage, syslogging.FieldData, syslogging.FieldErrorStack:
			values[i] = new(sql.NullString)
		case syslogging.FieldCreatedAt, syslogging.FieldUpdatedAt, syslogging.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysLogging", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysLogging fields.
func (sl *SysLogging) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case syslogging.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				sl.ID = value.String
			}
		case syslogging.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sl.CreatedAt = value.Time
			}
		case syslogging.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sl.UpdatedAt = value.Time
			}
		case syslogging.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sl.DeletedAt = new(time.Time)
				*sl.DeletedAt = value.Time
			}
		case syslogging.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				sl.IsDel = value.Bool
			}
		case syslogging.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				sl.Memo = new(string)
				*sl.Memo = value.String
			}
		case syslogging.FieldLevel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				sl.Level = new(string)
				*sl.Level = value.String
			}
		case syslogging.FieldTraceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field trace_id", values[i])
			} else if value.Valid {
				sl.TraceID = new(string)
				*sl.TraceID = value.String
			}
		case syslogging.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				sl.UserID = new(string)
				*sl.UserID = value.String
			}
		case syslogging.FieldTag:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tag", values[i])
			} else if value.Valid {
				sl.Tag = new(string)
				*sl.Tag = value.String
			}
		case syslogging.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				sl.Version = new(string)
				*sl.Version = value.String
			}
		case syslogging.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				sl.Message = new(string)
				*sl.Message = value.String
			}
		case syslogging.FieldData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value.Valid {
				sl.Data = new(string)
				*sl.Data = value.String
			}
		case syslogging.FieldErrorStack:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field error_stack", values[i])
			} else if value.Valid {
				sl.ErrorStack = new(string)
				*sl.ErrorStack = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SysLogging.
// Note that you need to call SysLogging.Unwrap() before calling this method if this SysLogging
// was returned from a transaction, and the transaction was committed or rolled back.
func (sl *SysLogging) Update() *SysLoggingUpdateOne {
	return NewSysLoggingClient(sl.config).UpdateOne(sl)
}

// Unwrap unwraps the SysLogging entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sl *SysLogging) Unwrap() *SysLogging {
	_tx, ok := sl.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysLogging is not a transactional entity")
	}
	sl.config.driver = _tx.drv
	return sl
}

// String implements the fmt.Stringer.
func (sl *SysLogging) String() string {
	var builder strings.Builder
	builder.WriteString("SysLogging(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sl.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sl.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sl.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := sl.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("is_del=")
	builder.WriteString(fmt.Sprintf("%v", sl.IsDel))
	builder.WriteString(", ")
	if v := sl.Memo; v != nil {
		builder.WriteString("memo=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sl.Level; v != nil {
		builder.WriteString("level=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sl.TraceID; v != nil {
		builder.WriteString("trace_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sl.UserID; v != nil {
		builder.WriteString("user_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sl.Tag; v != nil {
		builder.WriteString("tag=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sl.Version; v != nil {
		builder.WriteString("version=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sl.Message; v != nil {
		builder.WriteString("message=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sl.Data; v != nil {
		builder.WriteString("data=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sl.ErrorStack; v != nil {
		builder.WriteString("error_stack=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// SysLoggings is a parsable slice of SysLogging.
type SysLoggings []*SysLogging
