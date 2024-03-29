// Code generated by ent, DO NOT EDIT.

package mainent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/mainent/orgorgan"
	"github.com/heromicro/omgind/internal/gen/mainent/orgstaff"
	"github.com/heromicro/omgind/internal/gen/mainent/sysaddress"
)

// SysAddress is the model entity for the SysAddress schema.
type SysAddress struct {
	config `json:"-" sql:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// 是否删除
	IsDel bool `json:"is_del,omitempty"`
	// user id
	UserID *string `json:"user_id,omitempty" sql:"user_id"`
	// organization id
	OrgID *string `json:"org_id,omitempty" sql:"org_id"`
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
	// 国
	Country *string `json:"country,omitempty"`
	// 省
	Province *string `json:"province,omitempty"`
	// 区/市
	City *string `json:"city,omitempty"`
	// 区/县
	County *string `json:"county,omitempty"`
	// 国ID
	CountryID *string `json:"country_id,omitempty"`
	// 省/市ID
	ProvinceID *string `json:"province_id,omitempty"`
	// 区/市ID
	CityID *string `json:"city_id,omitempty"`
	// 区/县ID
	CountyID *string `json:"county_id,omitempty"`
	// 邮政编码
	ZipCode *string `json:"zip_code,omitempty"`
	// 详细地址
	Daddr *string `json:"daddr,omitempty"`
	// 联系名
	FirstName *string `json:"first_name,omitempty"`
	// 联系姓
	LastName *string `json:"last_name,omitempty"`
	// 国际区号
	AreaCode *string `json:"area_code,omitempty"`
	// 电话
	Mobile *string `json:"mobile,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SysAddressQuery when eager-loading is set.
	Edges        SysAddressEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SysAddressEdges holds the relations/edges for other nodes in the graph.
type SysAddressEdges struct {
	// Organ holds the value of the organ edge.
	Organ *OrgOrgan `json:"organ,omitempty"`
	// StaffResi holds the value of the staff_resi edge.
	StaffResi *OrgStaff `json:"staff_resi,omitempty"`
	// StaffIden holds the value of the staff_iden edge.
	StaffIden *OrgStaff `json:"staff_iden,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OrganOrErr returns the Organ value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SysAddressEdges) OrganOrErr() (*OrgOrgan, error) {
	if e.Organ != nil {
		return e.Organ, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: orgorgan.Label}
	}
	return nil, &NotLoadedError{edge: "organ"}
}

// StaffResiOrErr returns the StaffResi value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SysAddressEdges) StaffResiOrErr() (*OrgStaff, error) {
	if e.StaffResi != nil {
		return e.StaffResi, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: orgstaff.Label}
	}
	return nil, &NotLoadedError{edge: "staff_resi"}
}

// StaffIdenOrErr returns the StaffIden value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SysAddressEdges) StaffIdenOrErr() (*OrgStaff, error) {
	if e.StaffIden != nil {
		return e.StaffIden, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: orgstaff.Label}
	}
	return nil, &NotLoadedError{edge: "staff_iden"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysAddress) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysaddress.FieldIsDel, sysaddress.FieldIsActive:
			values[i] = new(sql.NullBool)
		case sysaddress.FieldSort:
			values[i] = new(sql.NullInt64)
		case sysaddress.FieldID, sysaddress.FieldUserID, sysaddress.FieldOrgID, sysaddress.FieldMemo, sysaddress.FieldCountry, sysaddress.FieldProvince, sysaddress.FieldCity, sysaddress.FieldCounty, sysaddress.FieldCountryID, sysaddress.FieldProvinceID, sysaddress.FieldCityID, sysaddress.FieldCountyID, sysaddress.FieldZipCode, sysaddress.FieldDaddr, sysaddress.FieldFirstName, sysaddress.FieldLastName, sysaddress.FieldAreaCode, sysaddress.FieldMobile:
			values[i] = new(sql.NullString)
		case sysaddress.FieldCreatedAt, sysaddress.FieldUpdatedAt, sysaddress.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysAddress fields.
func (sa *SysAddress) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysaddress.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				sa.ID = value.String
			}
		case sysaddress.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				sa.IsDel = value.Bool
			}
		case sysaddress.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				sa.UserID = new(string)
				*sa.UserID = value.String
			}
		case sysaddress.FieldOrgID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				sa.OrgID = new(string)
				*sa.OrgID = value.String
			}
		case sysaddress.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				sa.Sort = int32(value.Int64)
			}
		case sysaddress.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sa.CreatedAt = new(time.Time)
				*sa.CreatedAt = value.Time
			}
		case sysaddress.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sa.UpdatedAt = new(time.Time)
				*sa.UpdatedAt = value.Time
			}
		case sysaddress.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sa.DeletedAt = new(time.Time)
				*sa.DeletedAt = value.Time
			}
		case sysaddress.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				sa.IsActive = value.Bool
			}
		case sysaddress.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				sa.Memo = new(string)
				*sa.Memo = value.String
			}
		case sysaddress.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				sa.Country = new(string)
				*sa.Country = value.String
			}
		case sysaddress.FieldProvince:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field province", values[i])
			} else if value.Valid {
				sa.Province = new(string)
				*sa.Province = value.String
			}
		case sysaddress.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				sa.City = new(string)
				*sa.City = value.String
			}
		case sysaddress.FieldCounty:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field county", values[i])
			} else if value.Valid {
				sa.County = new(string)
				*sa.County = value.String
			}
		case sysaddress.FieldCountryID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country_id", values[i])
			} else if value.Valid {
				sa.CountryID = new(string)
				*sa.CountryID = value.String
			}
		case sysaddress.FieldProvinceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field province_id", values[i])
			} else if value.Valid {
				sa.ProvinceID = new(string)
				*sa.ProvinceID = value.String
			}
		case sysaddress.FieldCityID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city_id", values[i])
			} else if value.Valid {
				sa.CityID = new(string)
				*sa.CityID = value.String
			}
		case sysaddress.FieldCountyID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field county_id", values[i])
			} else if value.Valid {
				sa.CountyID = new(string)
				*sa.CountyID = value.String
			}
		case sysaddress.FieldZipCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field zip_code", values[i])
			} else if value.Valid {
				sa.ZipCode = new(string)
				*sa.ZipCode = value.String
			}
		case sysaddress.FieldDaddr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field daddr", values[i])
			} else if value.Valid {
				sa.Daddr = new(string)
				*sa.Daddr = value.String
			}
		case sysaddress.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				sa.FirstName = new(string)
				*sa.FirstName = value.String
			}
		case sysaddress.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				sa.LastName = new(string)
				*sa.LastName = value.String
			}
		case sysaddress.FieldAreaCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field area_code", values[i])
			} else if value.Valid {
				sa.AreaCode = new(string)
				*sa.AreaCode = value.String
			}
		case sysaddress.FieldMobile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mobile", values[i])
			} else if value.Valid {
				sa.Mobile = new(string)
				*sa.Mobile = value.String
			}
		default:
			sa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SysAddress.
// This includes values selected through modifiers, order, etc.
func (sa *SysAddress) Value(name string) (ent.Value, error) {
	return sa.selectValues.Get(name)
}

// QueryOrgan queries the "organ" edge of the SysAddress entity.
func (sa *SysAddress) QueryOrgan() *OrgOrganQuery {
	return NewSysAddressClient(sa.config).QueryOrgan(sa)
}

// QueryStaffResi queries the "staff_resi" edge of the SysAddress entity.
func (sa *SysAddress) QueryStaffResi() *OrgStaffQuery {
	return NewSysAddressClient(sa.config).QueryStaffResi(sa)
}

// QueryStaffIden queries the "staff_iden" edge of the SysAddress entity.
func (sa *SysAddress) QueryStaffIden() *OrgStaffQuery {
	return NewSysAddressClient(sa.config).QueryStaffIden(sa)
}

// Update returns a builder for updating this SysAddress.
// Note that you need to call SysAddress.Unwrap() before calling this method if this SysAddress
// was returned from a transaction, and the transaction was committed or rolled back.
func (sa *SysAddress) Update() *SysAddressUpdateOne {
	return NewSysAddressClient(sa.config).UpdateOne(sa)
}

// Unwrap unwraps the SysAddress entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sa *SysAddress) Unwrap() *SysAddress {
	_tx, ok := sa.config.driver.(*txDriver)
	if !ok {
		panic("mainent: SysAddress is not a transactional entity")
	}
	sa.config.driver = _tx.drv
	return sa
}

// String implements the fmt.Stringer.
func (sa *SysAddress) String() string {
	var builder strings.Builder
	builder.WriteString("SysAddress(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sa.ID))
	builder.WriteString("is_del=")
	builder.WriteString(fmt.Sprintf("%v", sa.IsDel))
	builder.WriteString(", ")
	if v := sa.UserID; v != nil {
		builder.WriteString("user_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sa.OrgID; v != nil {
		builder.WriteString("org_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", sa.Sort))
	builder.WriteString(", ")
	if v := sa.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := sa.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := sa.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", sa.IsActive))
	builder.WriteString(", ")
	if v := sa.Memo; v != nil {
		builder.WriteString("memo=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sa.Country; v != nil {
		builder.WriteString("country=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sa.Province; v != nil {
		builder.WriteString("province=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sa.City; v != nil {
		builder.WriteString("city=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sa.County; v != nil {
		builder.WriteString("county=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sa.CountryID; v != nil {
		builder.WriteString("country_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sa.ProvinceID; v != nil {
		builder.WriteString("province_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sa.CityID; v != nil {
		builder.WriteString("city_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sa.CountyID; v != nil {
		builder.WriteString("county_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sa.ZipCode; v != nil {
		builder.WriteString("zip_code=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sa.Daddr; v != nil {
		builder.WriteString("daddr=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sa.FirstName; v != nil {
		builder.WriteString("first_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sa.LastName; v != nil {
		builder.WriteString("last_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sa.AreaCode; v != nil {
		builder.WriteString("area_code=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sa.Mobile; v != nil {
		builder.WriteString("mobile=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// SysAddresses is a parsable slice of SysAddress.
type SysAddresses []*SysAddress
