// Code generated by ent, DO NOT EDIT.

package entscheme

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/entscheme/systeam"
)

// SysTeam is the model entity for the SysTeam schema.
type SysTeam struct {
	config `json:"-" sql:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
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
	// 是否删除
	IsDel bool `json:"is_del,omitempty"`
	// 名称
	Name *string `json:"name,omitempty"`
	// 助记码
	Code *string `json:"code,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SysTeamQuery when eager-loading is set.
	Edges        SysTeamEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SysTeamEdges holds the relations/edges for other nodes in the graph.
type SysTeamEdges struct {
	// Users holds the value of the users edge.
	Users []*SysUser `json:"users,omitempty"`
	// TeamUsers holds the value of the team_users edge.
	TeamUsers []*SysTeamUser `json:"team_users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e SysTeamEdges) UsersOrErr() ([]*SysUser, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// TeamUsersOrErr returns the TeamUsers value or an error if the edge
// was not loaded in eager-loading.
func (e SysTeamEdges) TeamUsersOrErr() ([]*SysTeamUser, error) {
	if e.loadedTypes[1] {
		return e.TeamUsers, nil
	}
	return nil, &NotLoadedError{edge: "team_users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysTeam) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case systeam.FieldIsActive, systeam.FieldIsDel:
			values[i] = new(sql.NullBool)
		case systeam.FieldSort:
			values[i] = new(sql.NullInt64)
		case systeam.FieldID, systeam.FieldMemo, systeam.FieldName, systeam.FieldCode:
			values[i] = new(sql.NullString)
		case systeam.FieldCreatedAt, systeam.FieldUpdatedAt, systeam.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysTeam fields.
func (st *SysTeam) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case systeam.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				st.ID = value.String
			}
		case systeam.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				st.Sort = int32(value.Int64)
			}
		case systeam.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				st.CreatedAt = new(time.Time)
				*st.CreatedAt = value.Time
			}
		case systeam.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				st.UpdatedAt = new(time.Time)
				*st.UpdatedAt = value.Time
			}
		case systeam.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				st.DeletedAt = new(time.Time)
				*st.DeletedAt = value.Time
			}
		case systeam.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				st.IsActive = value.Bool
			}
		case systeam.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				st.Memo = new(string)
				*st.Memo = value.String
			}
		case systeam.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				st.IsDel = value.Bool
			}
		case systeam.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				st.Name = new(string)
				*st.Name = value.String
			}
		case systeam.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				st.Code = new(string)
				*st.Code = value.String
			}
		default:
			st.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SysTeam.
// This includes values selected through modifiers, order, etc.
func (st *SysTeam) Value(name string) (ent.Value, error) {
	return st.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the SysTeam entity.
func (st *SysTeam) QueryUsers() *SysUserQuery {
	return NewSysTeamClient(st.config).QueryUsers(st)
}

// QueryTeamUsers queries the "team_users" edge of the SysTeam entity.
func (st *SysTeam) QueryTeamUsers() *SysTeamUserQuery {
	return NewSysTeamClient(st.config).QueryTeamUsers(st)
}

// Update returns a builder for updating this SysTeam.
// Note that you need to call SysTeam.Unwrap() before calling this method if this SysTeam
// was returned from a transaction, and the transaction was committed or rolled back.
func (st *SysTeam) Update() *SysTeamUpdateOne {
	return NewSysTeamClient(st.config).UpdateOne(st)
}

// Unwrap unwraps the SysTeam entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (st *SysTeam) Unwrap() *SysTeam {
	_tx, ok := st.config.driver.(*txDriver)
	if !ok {
		panic("entscheme: SysTeam is not a transactional entity")
	}
	st.config.driver = _tx.drv
	return st
}

// String implements the fmt.Stringer.
func (st *SysTeam) String() string {
	var builder strings.Builder
	builder.WriteString("SysTeam(")
	builder.WriteString(fmt.Sprintf("id=%v, ", st.ID))
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", st.Sort))
	builder.WriteString(", ")
	if v := st.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := st.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := st.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", st.IsActive))
	builder.WriteString(", ")
	if v := st.Memo; v != nil {
		builder.WriteString("memo=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("is_del=")
	builder.WriteString(fmt.Sprintf("%v", st.IsDel))
	builder.WriteString(", ")
	if v := st.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := st.Code; v != nil {
		builder.WriteString("code=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// SysTeams is a parsable slice of SysTeam.
type SysTeams []*SysTeam
