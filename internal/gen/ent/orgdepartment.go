// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/orgdepartment"
)

// OrgDepartment is the model entity for the OrgDepartment schema.
type OrgDepartment struct {
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
	// memo
	Memo *string `json:"memo,omitempty" sql:"memo"`
	// 名称
	Name *string `json:"name,omitempty"`
	// 助记码
	Code *string `json:"code,omitempty"`
	// 企业id
	OrgID *string `json:"org_id,omitempty"`
	// 创建者
	Creator *string `json:"creator,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrgDepartment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orgdepartment.FieldIsDel, orgdepartment.FieldIsActive:
			values[i] = new(sql.NullBool)
		case orgdepartment.FieldSort:
			values[i] = new(sql.NullInt64)
		case orgdepartment.FieldID, orgdepartment.FieldMemo, orgdepartment.FieldName, orgdepartment.FieldCode, orgdepartment.FieldOrgID, orgdepartment.FieldCreator:
			values[i] = new(sql.NullString)
		case orgdepartment.FieldCreatedAt, orgdepartment.FieldUpdatedAt, orgdepartment.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrgDepartment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrgDepartment fields.
func (od *OrgDepartment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orgdepartment.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				od.ID = value.String
			}
		case orgdepartment.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				od.IsDel = value.Bool
			}
		case orgdepartment.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				od.Sort = int32(value.Int64)
			}
		case orgdepartment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				od.CreatedAt = new(time.Time)
				*od.CreatedAt = value.Time
			}
		case orgdepartment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				od.UpdatedAt = new(time.Time)
				*od.UpdatedAt = value.Time
			}
		case orgdepartment.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				od.DeletedAt = new(time.Time)
				*od.DeletedAt = value.Time
			}
		case orgdepartment.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				od.IsActive = value.Bool
			}
		case orgdepartment.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				od.Memo = new(string)
				*od.Memo = value.String
			}
		case orgdepartment.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				od.Name = new(string)
				*od.Name = value.String
			}
		case orgdepartment.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				od.Code = new(string)
				*od.Code = value.String
			}
		case orgdepartment.FieldOrgID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				od.OrgID = new(string)
				*od.OrgID = value.String
			}
		case orgdepartment.FieldCreator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator", values[i])
			} else if value.Valid {
				od.Creator = new(string)
				*od.Creator = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this OrgDepartment.
// Note that you need to call OrgDepartment.Unwrap() before calling this method if this OrgDepartment
// was returned from a transaction, and the transaction was committed or rolled back.
func (od *OrgDepartment) Update() *OrgDepartmentUpdateOne {
	return NewOrgDepartmentClient(od.config).UpdateOne(od)
}

// Unwrap unwraps the OrgDepartment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (od *OrgDepartment) Unwrap() *OrgDepartment {
	_tx, ok := od.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrgDepartment is not a transactional entity")
	}
	od.config.driver = _tx.drv
	return od
}

// String implements the fmt.Stringer.
func (od *OrgDepartment) String() string {
	var builder strings.Builder
	builder.WriteString("OrgDepartment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", od.ID))
	builder.WriteString("is_del=")
	builder.WriteString(fmt.Sprintf("%v", od.IsDel))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", od.Sort))
	builder.WriteString(", ")
	if v := od.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := od.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := od.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", od.IsActive))
	builder.WriteString(", ")
	if v := od.Memo; v != nil {
		builder.WriteString("memo=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := od.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := od.Code; v != nil {
		builder.WriteString("code=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := od.OrgID; v != nil {
		builder.WriteString("org_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := od.Creator; v != nil {
		builder.WriteString("creator=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// OrgDepartments is a parsable slice of OrgDepartment.
type OrgDepartments []*OrgDepartment