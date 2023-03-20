// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/orgorgan"
	"github.com/heromicro/omgind/internal/gen/ent/sysaddress"
)

// OrgOrgan is the model entity for the OrgOrgan schema.
type OrgOrgan struct {
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
	// 简称
	Sname *string `json:"sname,omitempty"`
	// 助记码
	Code *string `json:"code,omitempty"`
	// 执照号
	IdenNo *string `json:"iden_no,omitempty"`
	// 所有者user.id
	OwnerID *string `json:"owner_id,omitempty"`
	// 总部id
	HaddrID *string `json:"haddr_id,omitempty"`
	// 创建者
	Creator *string `json:"creator,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrgOrganQuery when eager-loading is set.
	Edges OrgOrganEdges `json:"edges"`
}

// OrgOrganEdges holds the relations/edges for other nodes in the graph.
type OrgOrganEdges struct {
	// Haddr holds the value of the haddr edge.
	Haddr *SysAddress `json:"haddr,omitempty"`
	// Departments holds the value of the departments edge.
	Departments []*OrgDepartment `json:"departments,omitempty"`
	// Staffs holds the value of the staffs edge.
	Staffs []*OrgStaff `json:"staffs,omitempty"`
	// Positions holds the value of the positions edge.
	Positions []*OrgPosition `json:"positions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// HaddrOrErr returns the Haddr value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrgOrganEdges) HaddrOrErr() (*SysAddress, error) {
	if e.loadedTypes[0] {
		if e.Haddr == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: sysaddress.Label}
		}
		return e.Haddr, nil
	}
	return nil, &NotLoadedError{edge: "haddr"}
}

// DepartmentsOrErr returns the Departments value or an error if the edge
// was not loaded in eager-loading.
func (e OrgOrganEdges) DepartmentsOrErr() ([]*OrgDepartment, error) {
	if e.loadedTypes[1] {
		return e.Departments, nil
	}
	return nil, &NotLoadedError{edge: "departments"}
}

// StaffsOrErr returns the Staffs value or an error if the edge
// was not loaded in eager-loading.
func (e OrgOrganEdges) StaffsOrErr() ([]*OrgStaff, error) {
	if e.loadedTypes[2] {
		return e.Staffs, nil
	}
	return nil, &NotLoadedError{edge: "staffs"}
}

// PositionsOrErr returns the Positions value or an error if the edge
// was not loaded in eager-loading.
func (e OrgOrganEdges) PositionsOrErr() ([]*OrgPosition, error) {
	if e.loadedTypes[3] {
		return e.Positions, nil
	}
	return nil, &NotLoadedError{edge: "positions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrgOrgan) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orgorgan.FieldIsDel, orgorgan.FieldIsActive:
			values[i] = new(sql.NullBool)
		case orgorgan.FieldSort:
			values[i] = new(sql.NullInt64)
		case orgorgan.FieldID, orgorgan.FieldMemo, orgorgan.FieldName, orgorgan.FieldSname, orgorgan.FieldCode, orgorgan.FieldIdenNo, orgorgan.FieldOwnerID, orgorgan.FieldHaddrID, orgorgan.FieldCreator:
			values[i] = new(sql.NullString)
		case orgorgan.FieldCreatedAt, orgorgan.FieldUpdatedAt, orgorgan.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrgOrgan", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrgOrgan fields.
func (oo *OrgOrgan) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orgorgan.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				oo.ID = value.String
			}
		case orgorgan.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				oo.IsDel = value.Bool
			}
		case orgorgan.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				oo.Sort = int32(value.Int64)
			}
		case orgorgan.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				oo.CreatedAt = new(time.Time)
				*oo.CreatedAt = value.Time
			}
		case orgorgan.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				oo.UpdatedAt = new(time.Time)
				*oo.UpdatedAt = value.Time
			}
		case orgorgan.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				oo.DeletedAt = new(time.Time)
				*oo.DeletedAt = value.Time
			}
		case orgorgan.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				oo.IsActive = value.Bool
			}
		case orgorgan.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				oo.Memo = new(string)
				*oo.Memo = value.String
			}
		case orgorgan.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				oo.Name = new(string)
				*oo.Name = value.String
			}
		case orgorgan.FieldSname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sname", values[i])
			} else if value.Valid {
				oo.Sname = new(string)
				*oo.Sname = value.String
			}
		case orgorgan.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				oo.Code = new(string)
				*oo.Code = value.String
			}
		case orgorgan.FieldIdenNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field iden_no", values[i])
			} else if value.Valid {
				oo.IdenNo = new(string)
				*oo.IdenNo = value.String
			}
		case orgorgan.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				oo.OwnerID = new(string)
				*oo.OwnerID = value.String
			}
		case orgorgan.FieldHaddrID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field haddr_id", values[i])
			} else if value.Valid {
				oo.HaddrID = new(string)
				*oo.HaddrID = value.String
			}
		case orgorgan.FieldCreator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator", values[i])
			} else if value.Valid {
				oo.Creator = new(string)
				*oo.Creator = value.String
			}
		}
	}
	return nil
}

// QueryHaddr queries the "haddr" edge of the OrgOrgan entity.
func (oo *OrgOrgan) QueryHaddr() *SysAddressQuery {
	return NewOrgOrganClient(oo.config).QueryHaddr(oo)
}

// QueryDepartments queries the "departments" edge of the OrgOrgan entity.
func (oo *OrgOrgan) QueryDepartments() *OrgDepartmentQuery {
	return NewOrgOrganClient(oo.config).QueryDepartments(oo)
}

// QueryStaffs queries the "staffs" edge of the OrgOrgan entity.
func (oo *OrgOrgan) QueryStaffs() *OrgStaffQuery {
	return NewOrgOrganClient(oo.config).QueryStaffs(oo)
}

// QueryPositions queries the "positions" edge of the OrgOrgan entity.
func (oo *OrgOrgan) QueryPositions() *OrgPositionQuery {
	return NewOrgOrganClient(oo.config).QueryPositions(oo)
}

// Update returns a builder for updating this OrgOrgan.
// Note that you need to call OrgOrgan.Unwrap() before calling this method if this OrgOrgan
// was returned from a transaction, and the transaction was committed or rolled back.
func (oo *OrgOrgan) Update() *OrgOrganUpdateOne {
	return NewOrgOrganClient(oo.config).UpdateOne(oo)
}

// Unwrap unwraps the OrgOrgan entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oo *OrgOrgan) Unwrap() *OrgOrgan {
	_tx, ok := oo.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrgOrgan is not a transactional entity")
	}
	oo.config.driver = _tx.drv
	return oo
}

// String implements the fmt.Stringer.
func (oo *OrgOrgan) String() string {
	var builder strings.Builder
	builder.WriteString("OrgOrgan(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oo.ID))
	builder.WriteString("is_del=")
	builder.WriteString(fmt.Sprintf("%v", oo.IsDel))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", oo.Sort))
	builder.WriteString(", ")
	if v := oo.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := oo.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := oo.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", oo.IsActive))
	builder.WriteString(", ")
	if v := oo.Memo; v != nil {
		builder.WriteString("memo=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := oo.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := oo.Sname; v != nil {
		builder.WriteString("sname=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := oo.Code; v != nil {
		builder.WriteString("code=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := oo.IdenNo; v != nil {
		builder.WriteString("iden_no=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := oo.OwnerID; v != nil {
		builder.WriteString("owner_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := oo.HaddrID; v != nil {
		builder.WriteString("haddr_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := oo.Creator; v != nil {
		builder.WriteString("creator=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// OrgOrgans is a parsable slice of OrgOrgan.
type OrgOrgans []*OrgOrgan
