// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/sysdictitem"
)

// SysDictItem is the model entity for the SysDictItem schema.
type SysDictItem struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// 是否删除
	IsDel bool `json:"is_del,omitempty"`
	// 备注
	Memo string `json:"memo,omitempty"`
	// 排序, 在数据库里的排序
	Sort int32 `json:"sort,omitempty"`
	// 创建时间,由程序自动生成
	CreatedAt time.Time `json:"created_at,omitempty"`
	// 更新时间,由程序自动生成
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 删除时间,
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// 显示值
	Label string `json:"label,omitempty"`
	// 字典值
	Value int `json:"value,omitempty"`
	// 启用状态
	Status int16 `json:"status,omitempty"`
	// sys_dict.id
	DictID string `json:"dict_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysDictItem) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysdictitem.FieldIsDel:
			values[i] = new(sql.NullBool)
		case sysdictitem.FieldSort, sysdictitem.FieldValue, sysdictitem.FieldStatus:
			values[i] = new(sql.NullInt64)
		case sysdictitem.FieldID, sysdictitem.FieldMemo, sysdictitem.FieldLabel, sysdictitem.FieldDictID:
			values[i] = new(sql.NullString)
		case sysdictitem.FieldCreatedAt, sysdictitem.FieldUpdatedAt, sysdictitem.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysDictItem", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysDictItem fields.
func (sdi *SysDictItem) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysdictitem.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				sdi.ID = value.String
			}
		case sysdictitem.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				sdi.IsDel = value.Bool
			}
		case sysdictitem.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				sdi.Memo = value.String
			}
		case sysdictitem.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				sdi.Sort = int32(value.Int64)
			}
		case sysdictitem.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sdi.CreatedAt = value.Time
			}
		case sysdictitem.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sdi.UpdatedAt = value.Time
			}
		case sysdictitem.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sdi.DeletedAt = new(time.Time)
				*sdi.DeletedAt = value.Time
			}
		case sysdictitem.FieldLabel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field label", values[i])
			} else if value.Valid {
				sdi.Label = value.String
			}
		case sysdictitem.FieldValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				sdi.Value = int(value.Int64)
			}
		case sysdictitem.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				sdi.Status = int16(value.Int64)
			}
		case sysdictitem.FieldDictID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dict_id", values[i])
			} else if value.Valid {
				sdi.DictID = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SysDictItem.
// Note that you need to call SysDictItem.Unwrap() before calling this method if this SysDictItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (sdi *SysDictItem) Update() *SysDictItemUpdateOne {
	return (&SysDictItemClient{config: sdi.config}).UpdateOne(sdi)
}

// Unwrap unwraps the SysDictItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sdi *SysDictItem) Unwrap() *SysDictItem {
	_tx, ok := sdi.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysDictItem is not a transactional entity")
	}
	sdi.config.driver = _tx.drv
	return sdi
}

// String implements the fmt.Stringer.
func (sdi *SysDictItem) String() string {
	var builder strings.Builder
	builder.WriteString("SysDictItem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sdi.ID))
	builder.WriteString("is_del=")
	builder.WriteString(fmt.Sprintf("%v", sdi.IsDel))
	builder.WriteString(", ")
	builder.WriteString("memo=")
	builder.WriteString(sdi.Memo)
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", sdi.Sort))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(sdi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sdi.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := sdi.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("label=")
	builder.WriteString(sdi.Label)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(fmt.Sprintf("%v", sdi.Value))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", sdi.Status))
	builder.WriteString(", ")
	builder.WriteString("dict_id=")
	builder.WriteString(sdi.DictID)
	builder.WriteByte(')')
	return builder.String()
}

// SysDictItems is a parsable slice of SysDictItem.
type SysDictItems []*SysDictItem

func (sdi SysDictItems) config(cfg config) {
	for _i := range sdi {
		sdi[_i].config = cfg
	}
}
