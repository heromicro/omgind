// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/sysmenu"
)

// SysMenu is the model entity for the SysMenu schema.
type SysMenu struct {
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
	// 菜单名称
	Name string `json:"name,omitempty"`
	// 菜单图标
	Icon string `json:"icon,omitempty"`
	// 前端路由
	Router string `json:"router,omitempty"`
	// 是否显示
	IsShow bool `json:"is_show,omitempty"`
	// 父级id
	ParentID *string `json:"parent_id,omitempty"`
	// 父级路径: 1/2/3
	ParentPath *string `json:"parent_path,omitempty"`
	// 层级
	Level int32 `json:"level"`
	// 是否是子叶
	IsLeaf *bool `json:"is_leaf"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysMenu) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysmenu.FieldIsDel, sysmenu.FieldIsActive, sysmenu.FieldIsShow, sysmenu.FieldIsLeaf:
			values[i] = new(sql.NullBool)
		case sysmenu.FieldSort, sysmenu.FieldLevel:
			values[i] = new(sql.NullInt64)
		case sysmenu.FieldID, sysmenu.FieldMemo, sysmenu.FieldName, sysmenu.FieldIcon, sysmenu.FieldRouter, sysmenu.FieldParentID, sysmenu.FieldParentPath:
			values[i] = new(sql.NullString)
		case sysmenu.FieldCreatedAt, sysmenu.FieldUpdatedAt, sysmenu.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysMenu", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysMenu fields.
func (sm *SysMenu) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysmenu.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				sm.ID = value.String
			}
		case sysmenu.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				sm.IsDel = value.Bool
			}
		case sysmenu.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				sm.Memo = new(string)
				*sm.Memo = value.String
			}
		case sysmenu.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				sm.Sort = int32(value.Int64)
			}
		case sysmenu.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sm.CreatedAt = new(time.Time)
				*sm.CreatedAt = value.Time
			}
		case sysmenu.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sm.UpdatedAt = new(time.Time)
				*sm.UpdatedAt = value.Time
			}
		case sysmenu.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sm.DeletedAt = new(time.Time)
				*sm.DeletedAt = value.Time
			}
		case sysmenu.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				sm.IsActive = value.Bool
			}
		case sysmenu.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sm.Name = value.String
			}
		case sysmenu.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				sm.Icon = value.String
			}
		case sysmenu.FieldRouter:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field router", values[i])
			} else if value.Valid {
				sm.Router = value.String
			}
		case sysmenu.FieldIsShow:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_show", values[i])
			} else if value.Valid {
				sm.IsShow = value.Bool
			}
		case sysmenu.FieldParentID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				sm.ParentID = new(string)
				*sm.ParentID = value.String
			}
		case sysmenu.FieldParentPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parent_path", values[i])
			} else if value.Valid {
				sm.ParentPath = new(string)
				*sm.ParentPath = value.String
			}
		case sysmenu.FieldLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				sm.Level = int32(value.Int64)
			}
		case sysmenu.FieldIsLeaf:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_leaf", values[i])
			} else if value.Valid {
				sm.IsLeaf = new(bool)
				*sm.IsLeaf = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SysMenu.
// Note that you need to call SysMenu.Unwrap() before calling this method if this SysMenu
// was returned from a transaction, and the transaction was committed or rolled back.
func (sm *SysMenu) Update() *SysMenuUpdateOne {
	return NewSysMenuClient(sm.config).UpdateOne(sm)
}

// Unwrap unwraps the SysMenu entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sm *SysMenu) Unwrap() *SysMenu {
	_tx, ok := sm.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysMenu is not a transactional entity")
	}
	sm.config.driver = _tx.drv
	return sm
}

// String implements the fmt.Stringer.
func (sm *SysMenu) String() string {
	var builder strings.Builder
	builder.WriteString("SysMenu(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sm.ID))
	builder.WriteString("is_del=")
	builder.WriteString(fmt.Sprintf("%v", sm.IsDel))
	builder.WriteString(", ")
	if v := sm.Memo; v != nil {
		builder.WriteString("memo=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", sm.Sort))
	builder.WriteString(", ")
	if v := sm.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := sm.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := sm.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", sm.IsActive))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(sm.Name)
	builder.WriteString(", ")
	builder.WriteString("icon=")
	builder.WriteString(sm.Icon)
	builder.WriteString(", ")
	builder.WriteString("router=")
	builder.WriteString(sm.Router)
	builder.WriteString(", ")
	builder.WriteString("is_show=")
	builder.WriteString(fmt.Sprintf("%v", sm.IsShow))
	builder.WriteString(", ")
	if v := sm.ParentID; v != nil {
		builder.WriteString("parent_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sm.ParentPath; v != nil {
		builder.WriteString("parent_path=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("level=")
	builder.WriteString(fmt.Sprintf("%v", sm.Level))
	builder.WriteString(", ")
	if v := sm.IsLeaf; v != nil {
		builder.WriteString("is_leaf=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// SysMenus is a parsable slice of SysMenu.
type SysMenus []*SysMenu
