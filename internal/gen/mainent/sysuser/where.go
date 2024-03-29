// Code generated by ent, DO NOT EDIT.

package sysuser

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/heromicro/omgind/internal/gen/mainent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.SysUser {
	return predicate.SysUser(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.SysUser {
	return predicate.SysUser(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.SysUser {
	return predicate.SysUser(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.SysUser {
	return predicate.SysUser(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.SysUser {
	return predicate.SysUser(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.SysUser {
	return predicate.SysUser(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.SysUser {
	return predicate.SysUser(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.SysUser {
	return predicate.SysUser(sql.FieldContainsFold(FieldID, id))
}

// IsDel applies equality check predicate on the "is_del" field. It's identical to IsDelEQ.
func IsDel(v bool) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldIsDel, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int32) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldSort, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldDeletedAt, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldIsActive, v))
}

// UserName applies equality check predicate on the "user_name" field. It's identical to UserNameEQ.
func UserName(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldUserName, v))
}

// RealName applies equality check predicate on the "real_name" field. It's identical to RealNameEQ.
func RealName(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldRealName, v))
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldFirstName, v))
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldLastName, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldPassword, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldEmail, v))
}

// Mobile applies equality check predicate on the "mobile" field. It's identical to MobileEQ.
func Mobile(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldMobile, v))
}

// Salt applies equality check predicate on the "salt" field. It's identical to SaltEQ.
func Salt(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldSalt, v))
}

// IsDelEQ applies the EQ predicate on the "is_del" field.
func IsDelEQ(v bool) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldIsDel, v))
}

// IsDelNEQ applies the NEQ predicate on the "is_del" field.
func IsDelNEQ(v bool) predicate.SysUser {
	return predicate.SysUser(sql.FieldNEQ(FieldIsDel, v))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int32) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int32) predicate.SysUser {
	return predicate.SysUser(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int32) predicate.SysUser {
	return predicate.SysUser(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int32) predicate.SysUser {
	return predicate.SysUser(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int32) predicate.SysUser {
	return predicate.SysUser(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int32) predicate.SysUser {
	return predicate.SysUser(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int32) predicate.SysUser {
	return predicate.SysUser(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int32) predicate.SysUser {
	return predicate.SysUser(sql.FieldLTE(FieldSort, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.SysUser {
	return predicate.SysUser(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.SysUser {
	return predicate.SysUser(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.SysUser {
	return predicate.SysUser(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.SysUser {
	return predicate.SysUser(sql.FieldNotNull(FieldUpdatedAt))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.SysUser {
	return predicate.SysUser(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.SysUser {
	return predicate.SysUser(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.SysUser {
	return predicate.SysUser(sql.FieldNotNull(FieldDeletedAt))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.SysUser {
	return predicate.SysUser(sql.FieldNEQ(FieldIsActive, v))
}

// UserNameEQ applies the EQ predicate on the "user_name" field.
func UserNameEQ(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldUserName, v))
}

// UserNameNEQ applies the NEQ predicate on the "user_name" field.
func UserNameNEQ(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldNEQ(FieldUserName, v))
}

// UserNameIn applies the In predicate on the "user_name" field.
func UserNameIn(vs ...string) predicate.SysUser {
	return predicate.SysUser(sql.FieldIn(FieldUserName, vs...))
}

// UserNameNotIn applies the NotIn predicate on the "user_name" field.
func UserNameNotIn(vs ...string) predicate.SysUser {
	return predicate.SysUser(sql.FieldNotIn(FieldUserName, vs...))
}

// UserNameGT applies the GT predicate on the "user_name" field.
func UserNameGT(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldGT(FieldUserName, v))
}

// UserNameGTE applies the GTE predicate on the "user_name" field.
func UserNameGTE(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldGTE(FieldUserName, v))
}

// UserNameLT applies the LT predicate on the "user_name" field.
func UserNameLT(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldLT(FieldUserName, v))
}

// UserNameLTE applies the LTE predicate on the "user_name" field.
func UserNameLTE(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldLTE(FieldUserName, v))
}

// UserNameContains applies the Contains predicate on the "user_name" field.
func UserNameContains(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldContains(FieldUserName, v))
}

// UserNameHasPrefix applies the HasPrefix predicate on the "user_name" field.
func UserNameHasPrefix(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldHasPrefix(FieldUserName, v))
}

// UserNameHasSuffix applies the HasSuffix predicate on the "user_name" field.
func UserNameHasSuffix(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldHasSuffix(FieldUserName, v))
}

// UserNameEqualFold applies the EqualFold predicate on the "user_name" field.
func UserNameEqualFold(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEqualFold(FieldUserName, v))
}

// UserNameContainsFold applies the ContainsFold predicate on the "user_name" field.
func UserNameContainsFold(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldContainsFold(FieldUserName, v))
}

// RealNameEQ applies the EQ predicate on the "real_name" field.
func RealNameEQ(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldRealName, v))
}

// RealNameNEQ applies the NEQ predicate on the "real_name" field.
func RealNameNEQ(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldNEQ(FieldRealName, v))
}

// RealNameIn applies the In predicate on the "real_name" field.
func RealNameIn(vs ...string) predicate.SysUser {
	return predicate.SysUser(sql.FieldIn(FieldRealName, vs...))
}

// RealNameNotIn applies the NotIn predicate on the "real_name" field.
func RealNameNotIn(vs ...string) predicate.SysUser {
	return predicate.SysUser(sql.FieldNotIn(FieldRealName, vs...))
}

// RealNameGT applies the GT predicate on the "real_name" field.
func RealNameGT(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldGT(FieldRealName, v))
}

// RealNameGTE applies the GTE predicate on the "real_name" field.
func RealNameGTE(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldGTE(FieldRealName, v))
}

// RealNameLT applies the LT predicate on the "real_name" field.
func RealNameLT(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldLT(FieldRealName, v))
}

// RealNameLTE applies the LTE predicate on the "real_name" field.
func RealNameLTE(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldLTE(FieldRealName, v))
}

// RealNameContains applies the Contains predicate on the "real_name" field.
func RealNameContains(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldContains(FieldRealName, v))
}

// RealNameHasPrefix applies the HasPrefix predicate on the "real_name" field.
func RealNameHasPrefix(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldHasPrefix(FieldRealName, v))
}

// RealNameHasSuffix applies the HasSuffix predicate on the "real_name" field.
func RealNameHasSuffix(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldHasSuffix(FieldRealName, v))
}

// RealNameIsNil applies the IsNil predicate on the "real_name" field.
func RealNameIsNil() predicate.SysUser {
	return predicate.SysUser(sql.FieldIsNull(FieldRealName))
}

// RealNameNotNil applies the NotNil predicate on the "real_name" field.
func RealNameNotNil() predicate.SysUser {
	return predicate.SysUser(sql.FieldNotNull(FieldRealName))
}

// RealNameEqualFold applies the EqualFold predicate on the "real_name" field.
func RealNameEqualFold(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEqualFold(FieldRealName, v))
}

// RealNameContainsFold applies the ContainsFold predicate on the "real_name" field.
func RealNameContainsFold(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldContainsFold(FieldRealName, v))
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldFirstName, v))
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldNEQ(FieldFirstName, v))
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.SysUser {
	return predicate.SysUser(sql.FieldIn(FieldFirstName, vs...))
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.SysUser {
	return predicate.SysUser(sql.FieldNotIn(FieldFirstName, vs...))
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldGT(FieldFirstName, v))
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldGTE(FieldFirstName, v))
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldLT(FieldFirstName, v))
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldLTE(FieldFirstName, v))
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldContains(FieldFirstName, v))
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldHasPrefix(FieldFirstName, v))
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldHasSuffix(FieldFirstName, v))
}

// FirstNameIsNil applies the IsNil predicate on the "first_name" field.
func FirstNameIsNil() predicate.SysUser {
	return predicate.SysUser(sql.FieldIsNull(FieldFirstName))
}

// FirstNameNotNil applies the NotNil predicate on the "first_name" field.
func FirstNameNotNil() predicate.SysUser {
	return predicate.SysUser(sql.FieldNotNull(FieldFirstName))
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEqualFold(FieldFirstName, v))
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldContainsFold(FieldFirstName, v))
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldLastName, v))
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldNEQ(FieldLastName, v))
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.SysUser {
	return predicate.SysUser(sql.FieldIn(FieldLastName, vs...))
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.SysUser {
	return predicate.SysUser(sql.FieldNotIn(FieldLastName, vs...))
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldGT(FieldLastName, v))
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldGTE(FieldLastName, v))
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldLT(FieldLastName, v))
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldLTE(FieldLastName, v))
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldContains(FieldLastName, v))
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldHasPrefix(FieldLastName, v))
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldHasSuffix(FieldLastName, v))
}

// LastNameIsNil applies the IsNil predicate on the "last_name" field.
func LastNameIsNil() predicate.SysUser {
	return predicate.SysUser(sql.FieldIsNull(FieldLastName))
}

// LastNameNotNil applies the NotNil predicate on the "last_name" field.
func LastNameNotNil() predicate.SysUser {
	return predicate.SysUser(sql.FieldNotNull(FieldLastName))
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEqualFold(FieldLastName, v))
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldContainsFold(FieldLastName, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.SysUser {
	return predicate.SysUser(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.SysUser {
	return predicate.SysUser(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldContainsFold(FieldPassword, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.SysUser {
	return predicate.SysUser(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.SysUser {
	return predicate.SysUser(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldContainsFold(FieldEmail, v))
}

// MobileEQ applies the EQ predicate on the "mobile" field.
func MobileEQ(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldMobile, v))
}

// MobileNEQ applies the NEQ predicate on the "mobile" field.
func MobileNEQ(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldNEQ(FieldMobile, v))
}

// MobileIn applies the In predicate on the "mobile" field.
func MobileIn(vs ...string) predicate.SysUser {
	return predicate.SysUser(sql.FieldIn(FieldMobile, vs...))
}

// MobileNotIn applies the NotIn predicate on the "mobile" field.
func MobileNotIn(vs ...string) predicate.SysUser {
	return predicate.SysUser(sql.FieldNotIn(FieldMobile, vs...))
}

// MobileGT applies the GT predicate on the "mobile" field.
func MobileGT(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldGT(FieldMobile, v))
}

// MobileGTE applies the GTE predicate on the "mobile" field.
func MobileGTE(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldGTE(FieldMobile, v))
}

// MobileLT applies the LT predicate on the "mobile" field.
func MobileLT(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldLT(FieldMobile, v))
}

// MobileLTE applies the LTE predicate on the "mobile" field.
func MobileLTE(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldLTE(FieldMobile, v))
}

// MobileContains applies the Contains predicate on the "mobile" field.
func MobileContains(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldContains(FieldMobile, v))
}

// MobileHasPrefix applies the HasPrefix predicate on the "mobile" field.
func MobileHasPrefix(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldHasPrefix(FieldMobile, v))
}

// MobileHasSuffix applies the HasSuffix predicate on the "mobile" field.
func MobileHasSuffix(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldHasSuffix(FieldMobile, v))
}

// MobileEqualFold applies the EqualFold predicate on the "mobile" field.
func MobileEqualFold(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEqualFold(FieldMobile, v))
}

// MobileContainsFold applies the ContainsFold predicate on the "mobile" field.
func MobileContainsFold(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldContainsFold(FieldMobile, v))
}

// SaltEQ applies the EQ predicate on the "salt" field.
func SaltEQ(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEQ(FieldSalt, v))
}

// SaltNEQ applies the NEQ predicate on the "salt" field.
func SaltNEQ(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldNEQ(FieldSalt, v))
}

// SaltIn applies the In predicate on the "salt" field.
func SaltIn(vs ...string) predicate.SysUser {
	return predicate.SysUser(sql.FieldIn(FieldSalt, vs...))
}

// SaltNotIn applies the NotIn predicate on the "salt" field.
func SaltNotIn(vs ...string) predicate.SysUser {
	return predicate.SysUser(sql.FieldNotIn(FieldSalt, vs...))
}

// SaltGT applies the GT predicate on the "salt" field.
func SaltGT(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldGT(FieldSalt, v))
}

// SaltGTE applies the GTE predicate on the "salt" field.
func SaltGTE(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldGTE(FieldSalt, v))
}

// SaltLT applies the LT predicate on the "salt" field.
func SaltLT(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldLT(FieldSalt, v))
}

// SaltLTE applies the LTE predicate on the "salt" field.
func SaltLTE(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldLTE(FieldSalt, v))
}

// SaltContains applies the Contains predicate on the "salt" field.
func SaltContains(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldContains(FieldSalt, v))
}

// SaltHasPrefix applies the HasPrefix predicate on the "salt" field.
func SaltHasPrefix(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldHasPrefix(FieldSalt, v))
}

// SaltHasSuffix applies the HasSuffix predicate on the "salt" field.
func SaltHasSuffix(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldHasSuffix(FieldSalt, v))
}

// SaltEqualFold applies the EqualFold predicate on the "salt" field.
func SaltEqualFold(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldEqualFold(FieldSalt, v))
}

// SaltContainsFold applies the ContainsFold predicate on the "salt" field.
func SaltContainsFold(v string) predicate.SysUser {
	return predicate.SysUser(sql.FieldContainsFold(FieldSalt, v))
}

// HasTeams applies the HasEdge predicate on the "teams" edge.
func HasTeams() predicate.SysUser {
	return predicate.SysUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TeamsTable, TeamsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamsWith applies the HasEdge predicate on the "teams" edge with a given conditions (other predicates).
func HasTeamsWith(preds ...predicate.SysTeam) predicate.SysUser {
	return predicate.SysUser(func(s *sql.Selector) {
		step := newTeamsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTeamUsers applies the HasEdge predicate on the "team_users" edge.
func HasTeamUsers() predicate.SysUser {
	return predicate.SysUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, TeamUsersTable, TeamUsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamUsersWith applies the HasEdge predicate on the "team_users" edge with a given conditions (other predicates).
func HasTeamUsersWith(preds ...predicate.SysTeamUser) predicate.SysUser {
	return predicate.SysUser(func(s *sql.Selector) {
		step := newTeamUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysUser) predicate.SysUser {
	return predicate.SysUser(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysUser) predicate.SysUser {
	return predicate.SysUser(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SysUser) predicate.SysUser {
	return predicate.SysUser(sql.NotPredicates(p))
}
