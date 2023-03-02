// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/gen/ent/sysaddress"
)

// SysAddress is the model entity for the SysAddress schema.
type SysAddress struct {
	config `json:"-" sql:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// 是否删除
	IsDel bool `json:"is_del,omitempty"`
	// owner id
	OwnerID *string `json:"owner_id,omitempty" sql:"owner_id"`
	// sort
	Sort int32 `json:"sort,omitempty" sql:"sort"`
	// create time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// delete time,
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// 是否活跃
	IsActive bool `json:"is_active,omitempty"`
	// memo
	Memo *string `json:"memo,omitempty" sql:"memo"`
	// 国
	Country *string `json:"country,omitempty"`
	// 省
	Provice *string `json:"provice,omitempty"`
	// 区/市
	City *string `json:"city,omitempty"`
	// 区/县
	County *string `json:"county,omitempty"`
	// 国ID
	CountryID *string `json:"country_id,omitempty"`
	// 省/市ID
	ProviceID *string `json:"provice_id,omitempty"`
	// 区/市ID
	CityID *string `json:"city_id,omitempty"`
	// 区/县ID
	CountyID *string `json:"county_id,omitempty"`
	// 邮政编码
	ZipCode *string `json:"zip_code,omitempty"`
	// 详细地址
	Daddr *string `json:"daddr,omitempty"`
	// 联系人
	Name *string `json:"name,omitempty"`
	// 电话
	Mobile *string `json:"mobile,omitempty"`
	// 创建者
	Creator *string `json:"creator,omitempty"`
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
		case sysaddress.FieldID, sysaddress.FieldOwnerID, sysaddress.FieldMemo, sysaddress.FieldCountry, sysaddress.FieldProvice, sysaddress.FieldCity, sysaddress.FieldCounty, sysaddress.FieldCountryID, sysaddress.FieldProviceID, sysaddress.FieldCityID, sysaddress.FieldCountyID, sysaddress.FieldZipCode, sysaddress.FieldDaddr, sysaddress.FieldName, sysaddress.FieldMobile, sysaddress.FieldCreator:
			values[i] = new(sql.NullString)
		case sysaddress.FieldCreatedAt, sysaddress.FieldUpdatedAt, sysaddress.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysAddress", columns[i])
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
		case sysaddress.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				sa.OwnerID = new(string)
				*sa.OwnerID = value.String
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
				sa.CreatedAt = value.Time
			}
		case sysaddress.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sa.UpdatedAt = value.Time
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
		case sysaddress.FieldProvice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provice", values[i])
			} else if value.Valid {
				sa.Provice = new(string)
				*sa.Provice = value.String
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
		case sysaddress.FieldProviceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provice_id", values[i])
			} else if value.Valid {
				sa.ProviceID = new(string)
				*sa.ProviceID = value.String
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
		case sysaddress.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sa.Name = new(string)
				*sa.Name = value.String
			}
		case sysaddress.FieldMobile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mobile", values[i])
			} else if value.Valid {
				sa.Mobile = new(string)
				*sa.Mobile = value.String
			}
		case sysaddress.FieldCreator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator", values[i])
			} else if value.Valid {
				sa.Creator = new(string)
				*sa.Creator = value.String
			}
		}
	}
	return nil
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
		panic("ent: SysAddress is not a transactional entity")
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
	if v := sa.OwnerID; v != nil {
		builder.WriteString("owner_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", sa.Sort))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(sa.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sa.UpdatedAt.Format(time.ANSIC))
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
	if v := sa.Provice; v != nil {
		builder.WriteString("provice=")
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
	if v := sa.ProviceID; v != nil {
		builder.WriteString("provice_id=")
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
	if v := sa.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sa.Mobile; v != nil {
		builder.WriteString("mobile=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sa.Creator; v != nil {
		builder.WriteString("creator=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// SysAddresses is a parsable slice of SysAddress.
type SysAddresses []*SysAddress
