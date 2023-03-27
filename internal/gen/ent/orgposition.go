// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/orgorgan"
	"github.com/heromicro/omgind/internal/gen/ent/orgposition"
)

// OrgPosition is the model entity for the OrgPosition schema.
type OrgPosition struct {
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
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrgPositionQuery when eager-loading is set.
	Edges OrgPositionEdges `json:"edges"`
}

// OrgPositionEdges holds the relations/edges for other nodes in the graph.
type OrgPositionEdges struct {
	// Organ holds the value of the organ edge.
	Organ *OrgOrgan `json:"organ,omitempty"`
	// Staffs holds the value of the staffs edge.
	Staffs []*OrgStaff `json:"staffs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OrganOrErr returns the Organ value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrgPositionEdges) OrganOrErr() (*OrgOrgan, error) {
	if e.loadedTypes[0] {
		if e.Organ == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: orgorgan.Label}
		}
		return e.Organ, nil
	}
	return nil, &NotLoadedError{edge: "organ"}
}

// StaffsOrErr returns the Staffs value or an error if the edge
// was not loaded in eager-loading.
func (e OrgPositionEdges) StaffsOrErr() ([]*OrgStaff, error) {
	if e.loadedTypes[1] {
		return e.Staffs, nil
	}
	return nil, &NotLoadedError{edge: "staffs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrgPosition) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orgposition.FieldIsDel, orgposition.FieldIsActive:
			values[i] = new(sql.NullBool)
		case orgposition.FieldSort:
			values[i] = new(sql.NullInt64)
		case orgposition.FieldID, orgposition.FieldMemo, orgposition.FieldName, orgposition.FieldCode, orgposition.FieldOrgID, orgposition.FieldCreator:
			values[i] = new(sql.NullString)
		case orgposition.FieldCreatedAt, orgposition.FieldUpdatedAt, orgposition.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrgPosition", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrgPosition fields.
func (op *OrgPosition) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orgposition.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				op.ID = value.String
			}
		case orgposition.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				op.IsDel = value.Bool
			}
		case orgposition.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				op.Sort = int32(value.Int64)
			}
		case orgposition.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				op.CreatedAt = new(time.Time)
				*op.CreatedAt = value.Time
			}
		case orgposition.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				op.UpdatedAt = new(time.Time)
				*op.UpdatedAt = value.Time
			}
		case orgposition.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				op.DeletedAt = new(time.Time)
				*op.DeletedAt = value.Time
			}
		case orgposition.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				op.IsActive = value.Bool
			}
		case orgposition.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				op.Memo = new(string)
				*op.Memo = value.String
			}
		case orgposition.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				op.Name = new(string)
				*op.Name = value.String
			}
		case orgposition.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				op.Code = new(string)
				*op.Code = value.String
			}
		case orgposition.FieldOrgID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				op.OrgID = new(string)
				*op.OrgID = value.String
			}
		case orgposition.FieldCreator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator", values[i])
			} else if value.Valid {
				op.Creator = new(string)
				*op.Creator = value.String
			}
		}
	}
	return nil
}

// QueryOrgan queries the "organ" edge of the OrgPosition entity.
func (op *OrgPosition) QueryOrgan() *OrgOrganQuery {
	return NewOrgPositionClient(op.config).QueryOrgan(op)
}

// QueryStaffs queries the "staffs" edge of the OrgPosition entity.
func (op *OrgPosition) QueryStaffs() *OrgStaffQuery {
	return NewOrgPositionClient(op.config).QueryStaffs(op)
}

// Update returns a builder for updating this OrgPosition.
// Note that you need to call OrgPosition.Unwrap() before calling this method if this OrgPosition
// was returned from a transaction, and the transaction was committed or rolled back.
func (op *OrgPosition) Update() *OrgPositionUpdateOne {
	return NewOrgPositionClient(op.config).UpdateOne(op)
}

// Unwrap unwraps the OrgPosition entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (op *OrgPosition) Unwrap() *OrgPosition {
	_tx, ok := op.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrgPosition is not a transactional entity")
	}
	op.config.driver = _tx.drv
	return op
}

// String implements the fmt.Stringer.
func (op *OrgPosition) String() string {
	var builder strings.Builder
	builder.WriteString("OrgPosition(")
	builder.WriteString(fmt.Sprintf("id=%v, ", op.ID))
	builder.WriteString("is_del=")
	builder.WriteString(fmt.Sprintf("%v", op.IsDel))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", op.Sort))
	builder.WriteString(", ")
	if v := op.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := op.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := op.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", op.IsActive))
	builder.WriteString(", ")
	if v := op.Memo; v != nil {
		builder.WriteString("memo=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := op.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := op.Code; v != nil {
		builder.WriteString("code=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := op.OrgID; v != nil {
		builder.WriteString("org_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := op.Creator; v != nil {
		builder.WriteString("creator=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// OrgPositions is a parsable slice of OrgPosition.
type OrgPositions []*OrgPosition
