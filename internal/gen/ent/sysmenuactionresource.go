// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/sysmenuactionresource"
)

// SysMenuActionResource is the model entity for the SysMenuActionResource schema.
type SysMenuActionResource struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// 是否删除
	IsDel bool `json:"is_del,omitempty"`
	// 排序, 在数据库里的排序
	Sort int32 `json:"sort,omitempty"`
	// 备注
	Memo string `json:"memo,omitempty"`
	// 创建时间,由程序自动生成
	CreatedAt time.Time `json:"created_at,omitempty"`
	// 更新时间,由程序自动生成
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 删除时间,
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// 状态,
	Status int16 `json:"status,omitempty"`
	// 资源HTTP请求方式(支持正则, get, delete, delete, put, patch )
	Method string `json:"method,omitempty"`
	// 资源HTTP请求路径（支持/:id匹配）
	Path string `json:"path,omitempty"`
	// sys_menu_action.id
	ActionID string `json:"action_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysMenuActionResource) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysmenuactionresource.FieldIsDel:
			values[i] = new(sql.NullBool)
		case sysmenuactionresource.FieldSort, sysmenuactionresource.FieldStatus:
			values[i] = new(sql.NullInt64)
		case sysmenuactionresource.FieldID, sysmenuactionresource.FieldMemo, sysmenuactionresource.FieldMethod, sysmenuactionresource.FieldPath, sysmenuactionresource.FieldActionID:
			values[i] = new(sql.NullString)
		case sysmenuactionresource.FieldCreatedAt, sysmenuactionresource.FieldUpdatedAt, sysmenuactionresource.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysMenuActionResource", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysMenuActionResource fields.
func (smar *SysMenuActionResource) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysmenuactionresource.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				smar.ID = value.String
			}
		case sysmenuactionresource.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				smar.IsDel = value.Bool
			}
		case sysmenuactionresource.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				smar.Sort = int32(value.Int64)
			}
		case sysmenuactionresource.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				smar.Memo = value.String
			}
		case sysmenuactionresource.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				smar.CreatedAt = value.Time
			}
		case sysmenuactionresource.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				smar.UpdatedAt = value.Time
			}
		case sysmenuactionresource.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				smar.DeletedAt = new(time.Time)
				*smar.DeletedAt = value.Time
			}
		case sysmenuactionresource.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				smar.Status = int16(value.Int64)
			}
		case sysmenuactionresource.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				smar.Method = value.String
			}
		case sysmenuactionresource.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				smar.Path = value.String
			}
		case sysmenuactionresource.FieldActionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field action_id", values[i])
			} else if value.Valid {
				smar.ActionID = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SysMenuActionResource.
// Note that you need to call SysMenuActionResource.Unwrap() before calling this method if this SysMenuActionResource
// was returned from a transaction, and the transaction was committed or rolled back.
func (smar *SysMenuActionResource) Update() *SysMenuActionResourceUpdateOne {
	return (&SysMenuActionResourceClient{config: smar.config}).UpdateOne(smar)
}

// Unwrap unwraps the SysMenuActionResource entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (smar *SysMenuActionResource) Unwrap() *SysMenuActionResource {
	_tx, ok := smar.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysMenuActionResource is not a transactional entity")
	}
	smar.config.driver = _tx.drv
	return smar
}

// String implements the fmt.Stringer.
func (smar *SysMenuActionResource) String() string {
	var builder strings.Builder
	builder.WriteString("SysMenuActionResource(")
	builder.WriteString(fmt.Sprintf("id=%v, ", smar.ID))
	builder.WriteString("is_del=")
	builder.WriteString(fmt.Sprintf("%v", smar.IsDel))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", smar.Sort))
	builder.WriteString(", ")
	builder.WriteString("memo=")
	builder.WriteString(smar.Memo)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(smar.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(smar.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := smar.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", smar.Status))
	builder.WriteString(", ")
	builder.WriteString("method=")
	builder.WriteString(smar.Method)
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(smar.Path)
	builder.WriteString(", ")
	builder.WriteString("action_id=")
	builder.WriteString(smar.ActionID)
	builder.WriteByte(')')
	return builder.String()
}

// SysMenuActionResources is a parsable slice of SysMenuActionResource.
type SysMenuActionResources []*SysMenuActionResource

func (smar SysMenuActionResources) config(cfg config) {
	for _i := range smar {
		smar[_i].config = cfg
	}
}
