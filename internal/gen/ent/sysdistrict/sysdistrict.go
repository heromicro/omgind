// Code generated by ent, DO NOT EDIT.

package sysdistrict

import (
	"time"
)

const (
	// Label holds the string label denoting the sysdistrict type in the database.
	Label = "sys_district"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "crtd_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "uptd_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "dltd_at"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// FieldTreeID holds the string denoting the tree_id field in the database.
	FieldTreeID = "tree_id"
	// FieldTreeLevel holds the string denoting the tree_level field in the database.
	FieldTreeLevel = "tree_level"
	// FieldTreeLeft holds the string denoting the tree_left field in the database.
	FieldTreeLeft = "tree_left"
	// FieldTreeRight holds the string denoting the tree_right field in the database.
	FieldTreeRight = "tree_right"
	// FieldIsLeaf holds the string denoting the is_leaf field in the database.
	FieldIsLeaf = "is_leaf"
	// FieldTreePath holds the string denoting the tree_path field in the database.
	FieldTreePath = "t_path"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSname holds the string denoting the sname field in the database.
	FieldSname = "sname"
	// FieldAbbr holds the string denoting the abbr field in the database.
	FieldAbbr = "abbr"
	// FieldStCode holds the string denoting the st_code field in the database.
	FieldStCode = "stcode"
	// FieldInitials holds the string denoting the initials field in the database.
	FieldInitials = "initials"
	// FieldPinyin holds the string denoting the pinyin field in the database.
	FieldPinyin = "pinyin"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "pid"
	// FieldLongitude holds the string denoting the longitude field in the database.
	FieldLongitude = "longitude"
	// FieldLatitude holds the string denoting the latitude field in the database.
	FieldLatitude = "latitude"
	// FieldAreaCode holds the string denoting the area_code field in the database.
	FieldAreaCode = "area_code"
	// FieldZipCode holds the string denoting the zip_code field in the database.
	FieldZipCode = "zip_code"
	// FieldMergeName holds the string denoting the merge_name field in the database.
	FieldMergeName = "mname"
	// FieldMergeSname holds the string denoting the merge_sname field in the database.
	FieldMergeSname = "msname"
	// FieldExtra holds the string denoting the extra field in the database.
	FieldExtra = "extra"
	// FieldSuffix holds the string denoting the suffix field in the database.
	FieldSuffix = "suffix"
	// FieldIsHot holds the string denoting the is_hot field in the database.
	FieldIsHot = "is_hot"
	// FieldIsReal holds the string denoting the is_real field in the database.
	FieldIsReal = "is_r"
	// FieldIsMain holds the string denoting the is_main field in the database.
	FieldIsMain = "is_m"
	// FieldIsDirect holds the string denoting the is_direct field in the database.
	FieldIsDirect = "is_d"
	// FieldCreator holds the string denoting the creator field in the database.
	FieldCreator = "creator"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// Table holds the table name of the sysdistrict in the database.
	Table = "sys_districts"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "sys_districts"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "pid"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "sys_districts"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "pid"
)

// Columns holds all SQL columns for sysdistrict fields.
var Columns = []string{
	FieldID,
	FieldSort,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldIsActive,
	FieldTreeID,
	FieldTreeLevel,
	FieldTreeLeft,
	FieldTreeRight,
	FieldIsLeaf,
	FieldTreePath,
	FieldName,
	FieldSname,
	FieldAbbr,
	FieldStCode,
	FieldInitials,
	FieldPinyin,
	FieldParentID,
	FieldLongitude,
	FieldLatitude,
	FieldAreaCode,
	FieldZipCode,
	FieldMergeName,
	FieldMergeSname,
	FieldExtra,
	FieldSuffix,
	FieldIsHot,
	FieldIsReal,
	FieldIsMain,
	FieldIsDirect,
	FieldCreator,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultSort holds the default value on creation for the "sort" field.
	DefaultSort int32
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
	// DefaultIsLeaf holds the default value on creation for the "is_leaf" field.
	DefaultIsLeaf bool
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// SnameValidator is a validator for the "sname" field. It is called by the builders before save.
	SnameValidator func(string) error
	// AbbrValidator is a validator for the "abbr" field. It is called by the builders before save.
	AbbrValidator func(string) error
	// StCodeValidator is a validator for the "st_code" field. It is called by the builders before save.
	StCodeValidator func(string) error
	// InitialsValidator is a validator for the "initials" field. It is called by the builders before save.
	InitialsValidator func(string) error
	// PinyinValidator is a validator for the "pinyin" field. It is called by the builders before save.
	PinyinValidator func(string) error
	// ParentIDValidator is a validator for the "parent_id" field. It is called by the builders before save.
	ParentIDValidator func(string) error
	// AreaCodeValidator is a validator for the "area_code" field. It is called by the builders before save.
	AreaCodeValidator func(string) error
	// ZipCodeValidator is a validator for the "zip_code" field. It is called by the builders before save.
	ZipCodeValidator func(string) error
	// MergeNameValidator is a validator for the "merge_name" field. It is called by the builders before save.
	MergeNameValidator func(string) error
	// MergeSnameValidator is a validator for the "merge_sname" field. It is called by the builders before save.
	MergeSnameValidator func(string) error
	// ExtraValidator is a validator for the "extra" field. It is called by the builders before save.
	ExtraValidator func(string) error
	// SuffixValidator is a validator for the "suffix" field. It is called by the builders before save.
	SuffixValidator func(string) error
	// DefaultIsHot holds the default value on creation for the "is_hot" field.
	DefaultIsHot bool
	// DefaultIsReal holds the default value on creation for the "is_real" field.
	DefaultIsReal bool
	// DefaultIsMain holds the default value on creation for the "is_main" field.
	DefaultIsMain bool
	// DefaultIsDirect holds the default value on creation for the "is_direct" field.
	DefaultIsDirect bool
	// CreatorValidator is a validator for the "creator" field. It is called by the builders before save.
	CreatorValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
