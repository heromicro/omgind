// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/orgdept"
	"github.com/heromicro/omgind/internal/gen/ent/orgorgan"
)

// OrgDept is the model entity for the OrgDept schema.
type OrgDept struct {
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
	// tree id
	TreeID *int64 `json:"tree_id"`
	// level in tree, toppest level is 1
	TreeLevel *int32 `json:"level"`
	// mptt's left
	TreeLeft *int64 `json:"tree_left"`
	// mptt's right
	TreeRight *int64 `json:"tree_right"`
	// is leaf node
	IsLeaf *bool `json:"is_leaf"`
	// tree path,topest is null or zero length string, subber has fathers ids join by slash(/), eg: pid1/pid2
	TreePath *string `json:"tree_path"`
	// 名称
	Name *string `json:"name,omitempty"`
	// 助记码
	Code *string `json:"code,omitempty"`
	// 企业id
	OrgID *string `json:"org_id,omitempty"`
	// 父级id
	ParentID *string `json:"parent_id,omitempty"`
	// 是否虚拟部门
	IsReal *bool `json:"is_real,omitempty"`
	// 是否显示
	IsShow *bool `json:"is_show,omitempty"`
	// 创建者
	Creator *string `json:"creator,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrgDeptQuery when eager-loading is set.
	Edges OrgDeptEdges `json:"edges"`
}

// OrgDeptEdges holds the relations/edges for other nodes in the graph.
type OrgDeptEdges struct {
	// Parent holds the value of the parent edge.
	Parent *OrgDept `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*OrgDept `json:"children,omitempty"`
	// Organ holds the value of the organ edge.
	Organ *OrgOrgan `json:"organ,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrgDeptEdges) ParentOrErr() (*OrgDept, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: orgdept.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e OrgDeptEdges) ChildrenOrErr() ([]*OrgDept, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// OrganOrErr returns the Organ value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrgDeptEdges) OrganOrErr() (*OrgOrgan, error) {
	if e.loadedTypes[2] {
		if e.Organ == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: orgorgan.Label}
		}
		return e.Organ, nil
	}
	return nil, &NotLoadedError{edge: "organ"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrgDept) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orgdept.FieldIsDel, orgdept.FieldIsActive, orgdept.FieldIsLeaf, orgdept.FieldIsReal, orgdept.FieldIsShow:
			values[i] = new(sql.NullBool)
		case orgdept.FieldSort, orgdept.FieldTreeID, orgdept.FieldTreeLevel, orgdept.FieldTreeLeft, orgdept.FieldTreeRight:
			values[i] = new(sql.NullInt64)
		case orgdept.FieldID, orgdept.FieldMemo, orgdept.FieldTreePath, orgdept.FieldName, orgdept.FieldCode, orgdept.FieldOrgID, orgdept.FieldParentID, orgdept.FieldCreator:
			values[i] = new(sql.NullString)
		case orgdept.FieldCreatedAt, orgdept.FieldUpdatedAt, orgdept.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrgDept", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrgDept fields.
func (od *OrgDept) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orgdept.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				od.ID = value.String
			}
		case orgdept.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				od.IsDel = value.Bool
			}
		case orgdept.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				od.Sort = int32(value.Int64)
			}
		case orgdept.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				od.CreatedAt = new(time.Time)
				*od.CreatedAt = value.Time
			}
		case orgdept.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				od.UpdatedAt = new(time.Time)
				*od.UpdatedAt = value.Time
			}
		case orgdept.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				od.DeletedAt = new(time.Time)
				*od.DeletedAt = value.Time
			}
		case orgdept.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				od.IsActive = value.Bool
			}
		case orgdept.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				od.Memo = new(string)
				*od.Memo = value.String
			}
		case orgdept.FieldTreeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tree_id", values[i])
			} else if value.Valid {
				od.TreeID = new(int64)
				*od.TreeID = value.Int64
			}
		case orgdept.FieldTreeLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tree_level", values[i])
			} else if value.Valid {
				od.TreeLevel = new(int32)
				*od.TreeLevel = int32(value.Int64)
			}
		case orgdept.FieldTreeLeft:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tree_left", values[i])
			} else if value.Valid {
				od.TreeLeft = new(int64)
				*od.TreeLeft = value.Int64
			}
		case orgdept.FieldTreeRight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tree_right", values[i])
			} else if value.Valid {
				od.TreeRight = new(int64)
				*od.TreeRight = value.Int64
			}
		case orgdept.FieldIsLeaf:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_leaf", values[i])
			} else if value.Valid {
				od.IsLeaf = new(bool)
				*od.IsLeaf = value.Bool
			}
		case orgdept.FieldTreePath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tree_path", values[i])
			} else if value.Valid {
				od.TreePath = new(string)
				*od.TreePath = value.String
			}
		case orgdept.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				od.Name = new(string)
				*od.Name = value.String
			}
		case orgdept.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				od.Code = new(string)
				*od.Code = value.String
			}
		case orgdept.FieldOrgID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				od.OrgID = new(string)
				*od.OrgID = value.String
			}
		case orgdept.FieldParentID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				od.ParentID = new(string)
				*od.ParentID = value.String
			}
		case orgdept.FieldIsReal:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_real", values[i])
			} else if value.Valid {
				od.IsReal = new(bool)
				*od.IsReal = value.Bool
			}
		case orgdept.FieldIsShow:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_show", values[i])
			} else if value.Valid {
				od.IsShow = new(bool)
				*od.IsShow = value.Bool
			}
		case orgdept.FieldCreator:
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

// QueryParent queries the "parent" edge of the OrgDept entity.
func (od *OrgDept) QueryParent() *OrgDeptQuery {
	return NewOrgDeptClient(od.config).QueryParent(od)
}

// QueryChildren queries the "children" edge of the OrgDept entity.
func (od *OrgDept) QueryChildren() *OrgDeptQuery {
	return NewOrgDeptClient(od.config).QueryChildren(od)
}

// QueryOrgan queries the "organ" edge of the OrgDept entity.
func (od *OrgDept) QueryOrgan() *OrgOrganQuery {
	return NewOrgDeptClient(od.config).QueryOrgan(od)
}

// Update returns a builder for updating this OrgDept.
// Note that you need to call OrgDept.Unwrap() before calling this method if this OrgDept
// was returned from a transaction, and the transaction was committed or rolled back.
func (od *OrgDept) Update() *OrgDeptUpdateOne {
	return NewOrgDeptClient(od.config).UpdateOne(od)
}

// Unwrap unwraps the OrgDept entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (od *OrgDept) Unwrap() *OrgDept {
	_tx, ok := od.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrgDept is not a transactional entity")
	}
	od.config.driver = _tx.drv
	return od
}

// String implements the fmt.Stringer.
func (od *OrgDept) String() string {
	var builder strings.Builder
	builder.WriteString("OrgDept(")
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
	if v := od.TreeID; v != nil {
		builder.WriteString("tree_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := od.TreeLevel; v != nil {
		builder.WriteString("tree_level=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := od.TreeLeft; v != nil {
		builder.WriteString("tree_left=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := od.TreeRight; v != nil {
		builder.WriteString("tree_right=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := od.IsLeaf; v != nil {
		builder.WriteString("is_leaf=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := od.TreePath; v != nil {
		builder.WriteString("tree_path=")
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
	if v := od.ParentID; v != nil {
		builder.WriteString("parent_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := od.IsReal; v != nil {
		builder.WriteString("is_real=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := od.IsShow; v != nil {
		builder.WriteString("is_show=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := od.Creator; v != nil {
		builder.WriteString("creator=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// OrgDepts is a parsable slice of OrgDept.
type OrgDepts []*OrgDept
