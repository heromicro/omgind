// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/ent/sysdistrict"
)

// SysDistrictCreate is the builder for creating a SysDistrict entity.
type SysDistrictCreate struct {
	config
	mutation *SysDistrictMutation
	hooks    []Hook
}

// SetIsDel sets the "is_del" field.
func (sdc *SysDistrictCreate) SetIsDel(b bool) *SysDistrictCreate {
	sdc.mutation.SetIsDel(b)
	return sdc
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableIsDel(b *bool) *SysDistrictCreate {
	if b != nil {
		sdc.SetIsDel(*b)
	}
	return sdc
}

// SetSort sets the "sort" field.
func (sdc *SysDistrictCreate) SetSort(i int32) *SysDistrictCreate {
	sdc.mutation.SetSort(i)
	return sdc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableSort(i *int32) *SysDistrictCreate {
	if i != nil {
		sdc.SetSort(*i)
	}
	return sdc
}

// SetCreatedAt sets the "created_at" field.
func (sdc *SysDistrictCreate) SetCreatedAt(t time.Time) *SysDistrictCreate {
	sdc.mutation.SetCreatedAt(t)
	return sdc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableCreatedAt(t *time.Time) *SysDistrictCreate {
	if t != nil {
		sdc.SetCreatedAt(*t)
	}
	return sdc
}

// SetUpdatedAt sets the "updated_at" field.
func (sdc *SysDistrictCreate) SetUpdatedAt(t time.Time) *SysDistrictCreate {
	sdc.mutation.SetUpdatedAt(t)
	return sdc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableUpdatedAt(t *time.Time) *SysDistrictCreate {
	if t != nil {
		sdc.SetUpdatedAt(*t)
	}
	return sdc
}

// SetDeletedAt sets the "deleted_at" field.
func (sdc *SysDistrictCreate) SetDeletedAt(t time.Time) *SysDistrictCreate {
	sdc.mutation.SetDeletedAt(t)
	return sdc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableDeletedAt(t *time.Time) *SysDistrictCreate {
	if t != nil {
		sdc.SetDeletedAt(*t)
	}
	return sdc
}

// SetIsActive sets the "is_active" field.
func (sdc *SysDistrictCreate) SetIsActive(b bool) *SysDistrictCreate {
	sdc.mutation.SetIsActive(b)
	return sdc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableIsActive(b *bool) *SysDistrictCreate {
	if b != nil {
		sdc.SetIsActive(*b)
	}
	return sdc
}

// SetTreeID sets the "tree_id" field.
func (sdc *SysDistrictCreate) SetTreeID(i int64) *SysDistrictCreate {
	sdc.mutation.SetTreeID(i)
	return sdc
}

// SetNillableTreeID sets the "tree_id" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableTreeID(i *int64) *SysDistrictCreate {
	if i != nil {
		sdc.SetTreeID(*i)
	}
	return sdc
}

// SetTreeLevel sets the "tree_level" field.
func (sdc *SysDistrictCreate) SetTreeLevel(i int32) *SysDistrictCreate {
	sdc.mutation.SetTreeLevel(i)
	return sdc
}

// SetNillableTreeLevel sets the "tree_level" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableTreeLevel(i *int32) *SysDistrictCreate {
	if i != nil {
		sdc.SetTreeLevel(*i)
	}
	return sdc
}

// SetTreeLeft sets the "tree_left" field.
func (sdc *SysDistrictCreate) SetTreeLeft(i int64) *SysDistrictCreate {
	sdc.mutation.SetTreeLeft(i)
	return sdc
}

// SetNillableTreeLeft sets the "tree_left" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableTreeLeft(i *int64) *SysDistrictCreate {
	if i != nil {
		sdc.SetTreeLeft(*i)
	}
	return sdc
}

// SetTreeRight sets the "tree_right" field.
func (sdc *SysDistrictCreate) SetTreeRight(i int64) *SysDistrictCreate {
	sdc.mutation.SetTreeRight(i)
	return sdc
}

// SetNillableTreeRight sets the "tree_right" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableTreeRight(i *int64) *SysDistrictCreate {
	if i != nil {
		sdc.SetTreeRight(*i)
	}
	return sdc
}

// SetIsLeaf sets the "is_leaf" field.
func (sdc *SysDistrictCreate) SetIsLeaf(b bool) *SysDistrictCreate {
	sdc.mutation.SetIsLeaf(b)
	return sdc
}

// SetNillableIsLeaf sets the "is_leaf" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableIsLeaf(b *bool) *SysDistrictCreate {
	if b != nil {
		sdc.SetIsLeaf(*b)
	}
	return sdc
}

// SetTreePath sets the "tree_path" field.
func (sdc *SysDistrictCreate) SetTreePath(s string) *SysDistrictCreate {
	sdc.mutation.SetTreePath(s)
	return sdc
}

// SetNillableTreePath sets the "tree_path" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableTreePath(s *string) *SysDistrictCreate {
	if s != nil {
		sdc.SetTreePath(*s)
	}
	return sdc
}

// SetName sets the "name" field.
func (sdc *SysDistrictCreate) SetName(s string) *SysDistrictCreate {
	sdc.mutation.SetName(s)
	return sdc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableName(s *string) *SysDistrictCreate {
	if s != nil {
		sdc.SetName(*s)
	}
	return sdc
}

// SetSname sets the "sname" field.
func (sdc *SysDistrictCreate) SetSname(s string) *SysDistrictCreate {
	sdc.mutation.SetSname(s)
	return sdc
}

// SetNillableSname sets the "sname" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableSname(s *string) *SysDistrictCreate {
	if s != nil {
		sdc.SetSname(*s)
	}
	return sdc
}

// SetAbbr sets the "abbr" field.
func (sdc *SysDistrictCreate) SetAbbr(s string) *SysDistrictCreate {
	sdc.mutation.SetAbbr(s)
	return sdc
}

// SetNillableAbbr sets the "abbr" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableAbbr(s *string) *SysDistrictCreate {
	if s != nil {
		sdc.SetAbbr(*s)
	}
	return sdc
}

// SetStCode sets the "st_code" field.
func (sdc *SysDistrictCreate) SetStCode(s string) *SysDistrictCreate {
	sdc.mutation.SetStCode(s)
	return sdc
}

// SetNillableStCode sets the "st_code" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableStCode(s *string) *SysDistrictCreate {
	if s != nil {
		sdc.SetStCode(*s)
	}
	return sdc
}

// SetInitials sets the "initials" field.
func (sdc *SysDistrictCreate) SetInitials(s string) *SysDistrictCreate {
	sdc.mutation.SetInitials(s)
	return sdc
}

// SetNillableInitials sets the "initials" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableInitials(s *string) *SysDistrictCreate {
	if s != nil {
		sdc.SetInitials(*s)
	}
	return sdc
}

// SetPinyin sets the "pinyin" field.
func (sdc *SysDistrictCreate) SetPinyin(s string) *SysDistrictCreate {
	sdc.mutation.SetPinyin(s)
	return sdc
}

// SetNillablePinyin sets the "pinyin" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillablePinyin(s *string) *SysDistrictCreate {
	if s != nil {
		sdc.SetPinyin(*s)
	}
	return sdc
}

// SetParentID sets the "parent_id" field.
func (sdc *SysDistrictCreate) SetParentID(s string) *SysDistrictCreate {
	sdc.mutation.SetParentID(s)
	return sdc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableParentID(s *string) *SysDistrictCreate {
	if s != nil {
		sdc.SetParentID(*s)
	}
	return sdc
}

// SetLongitude sets the "longitude" field.
func (sdc *SysDistrictCreate) SetLongitude(f float64) *SysDistrictCreate {
	sdc.mutation.SetLongitude(f)
	return sdc
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableLongitude(f *float64) *SysDistrictCreate {
	if f != nil {
		sdc.SetLongitude(*f)
	}
	return sdc
}

// SetLatitude sets the "latitude" field.
func (sdc *SysDistrictCreate) SetLatitude(f float64) *SysDistrictCreate {
	sdc.mutation.SetLatitude(f)
	return sdc
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableLatitude(f *float64) *SysDistrictCreate {
	if f != nil {
		sdc.SetLatitude(*f)
	}
	return sdc
}

// SetAreaCode sets the "area_code" field.
func (sdc *SysDistrictCreate) SetAreaCode(s string) *SysDistrictCreate {
	sdc.mutation.SetAreaCode(s)
	return sdc
}

// SetNillableAreaCode sets the "area_code" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableAreaCode(s *string) *SysDistrictCreate {
	if s != nil {
		sdc.SetAreaCode(*s)
	}
	return sdc
}

// SetZipCode sets the "zip_code" field.
func (sdc *SysDistrictCreate) SetZipCode(s string) *SysDistrictCreate {
	sdc.mutation.SetZipCode(s)
	return sdc
}

// SetNillableZipCode sets the "zip_code" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableZipCode(s *string) *SysDistrictCreate {
	if s != nil {
		sdc.SetZipCode(*s)
	}
	return sdc
}

// SetMergeName sets the "merge_name" field.
func (sdc *SysDistrictCreate) SetMergeName(s string) *SysDistrictCreate {
	sdc.mutation.SetMergeName(s)
	return sdc
}

// SetNillableMergeName sets the "merge_name" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableMergeName(s *string) *SysDistrictCreate {
	if s != nil {
		sdc.SetMergeName(*s)
	}
	return sdc
}

// SetMergeSname sets the "merge_sname" field.
func (sdc *SysDistrictCreate) SetMergeSname(s string) *SysDistrictCreate {
	sdc.mutation.SetMergeSname(s)
	return sdc
}

// SetNillableMergeSname sets the "merge_sname" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableMergeSname(s *string) *SysDistrictCreate {
	if s != nil {
		sdc.SetMergeSname(*s)
	}
	return sdc
}

// SetExtra sets the "extra" field.
func (sdc *SysDistrictCreate) SetExtra(s string) *SysDistrictCreate {
	sdc.mutation.SetExtra(s)
	return sdc
}

// SetNillableExtra sets the "extra" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableExtra(s *string) *SysDistrictCreate {
	if s != nil {
		sdc.SetExtra(*s)
	}
	return sdc
}

// SetSuffix sets the "suffix" field.
func (sdc *SysDistrictCreate) SetSuffix(s string) *SysDistrictCreate {
	sdc.mutation.SetSuffix(s)
	return sdc
}

// SetNillableSuffix sets the "suffix" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableSuffix(s *string) *SysDistrictCreate {
	if s != nil {
		sdc.SetSuffix(*s)
	}
	return sdc
}

// SetIsHot sets the "is_hot" field.
func (sdc *SysDistrictCreate) SetIsHot(b bool) *SysDistrictCreate {
	sdc.mutation.SetIsHot(b)
	return sdc
}

// SetNillableIsHot sets the "is_hot" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableIsHot(b *bool) *SysDistrictCreate {
	if b != nil {
		sdc.SetIsHot(*b)
	}
	return sdc
}

// SetIsReal sets the "is_real" field.
func (sdc *SysDistrictCreate) SetIsReal(b bool) *SysDistrictCreate {
	sdc.mutation.SetIsReal(b)
	return sdc
}

// SetNillableIsReal sets the "is_real" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableIsReal(b *bool) *SysDistrictCreate {
	if b != nil {
		sdc.SetIsReal(*b)
	}
	return sdc
}

// SetIsMain sets the "is_main" field.
func (sdc *SysDistrictCreate) SetIsMain(b bool) *SysDistrictCreate {
	sdc.mutation.SetIsMain(b)
	return sdc
}

// SetNillableIsMain sets the "is_main" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableIsMain(b *bool) *SysDistrictCreate {
	if b != nil {
		sdc.SetIsMain(*b)
	}
	return sdc
}

// SetIsDirect sets the "is_direct" field.
func (sdc *SysDistrictCreate) SetIsDirect(b bool) *SysDistrictCreate {
	sdc.mutation.SetIsDirect(b)
	return sdc
}

// SetNillableIsDirect sets the "is_direct" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableIsDirect(b *bool) *SysDistrictCreate {
	if b != nil {
		sdc.SetIsDirect(*b)
	}
	return sdc
}

// SetCreator sets the "creator" field.
func (sdc *SysDistrictCreate) SetCreator(s string) *SysDistrictCreate {
	sdc.mutation.SetCreator(s)
	return sdc
}

// SetID sets the "id" field.
func (sdc *SysDistrictCreate) SetID(s string) *SysDistrictCreate {
	sdc.mutation.SetID(s)
	return sdc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sdc *SysDistrictCreate) SetNillableID(s *string) *SysDistrictCreate {
	if s != nil {
		sdc.SetID(*s)
	}
	return sdc
}

// SetParent sets the "parent" edge to the SysDistrict entity.
func (sdc *SysDistrictCreate) SetParent(s *SysDistrict) *SysDistrictCreate {
	return sdc.SetParentID(s.ID)
}

// AddChildIDs adds the "children" edge to the SysDistrict entity by IDs.
func (sdc *SysDistrictCreate) AddChildIDs(ids ...string) *SysDistrictCreate {
	sdc.mutation.AddChildIDs(ids...)
	return sdc
}

// AddChildren adds the "children" edges to the SysDistrict entity.
func (sdc *SysDistrictCreate) AddChildren(s ...*SysDistrict) *SysDistrictCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sdc.AddChildIDs(ids...)
}

// Mutation returns the SysDistrictMutation object of the builder.
func (sdc *SysDistrictCreate) Mutation() *SysDistrictMutation {
	return sdc.mutation
}

// Save creates the SysDistrict in the database.
func (sdc *SysDistrictCreate) Save(ctx context.Context) (*SysDistrict, error) {
	var (
		err  error
		node *SysDistrict
	)
	sdc.defaults()
	if len(sdc.hooks) == 0 {
		if err = sdc.check(); err != nil {
			return nil, err
		}
		node, err = sdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysDistrictMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sdc.check(); err != nil {
				return nil, err
			}
			sdc.mutation = mutation
			if node, err = sdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sdc.hooks) - 1; i >= 0; i-- {
			if sdc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sdc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sdc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SysDistrict)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SysDistrictMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sdc *SysDistrictCreate) SaveX(ctx context.Context) *SysDistrict {
	v, err := sdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdc *SysDistrictCreate) Exec(ctx context.Context) error {
	_, err := sdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdc *SysDistrictCreate) ExecX(ctx context.Context) {
	if err := sdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sdc *SysDistrictCreate) defaults() {
	if _, ok := sdc.mutation.IsDel(); !ok {
		v := sysdistrict.DefaultIsDel
		sdc.mutation.SetIsDel(v)
	}
	if _, ok := sdc.mutation.Sort(); !ok {
		v := sysdistrict.DefaultSort
		sdc.mutation.SetSort(v)
	}
	if _, ok := sdc.mutation.CreatedAt(); !ok {
		v := sysdistrict.DefaultCreatedAt()
		sdc.mutation.SetCreatedAt(v)
	}
	if _, ok := sdc.mutation.UpdatedAt(); !ok {
		v := sysdistrict.DefaultUpdatedAt()
		sdc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sdc.mutation.IsActive(); !ok {
		v := sysdistrict.DefaultIsActive
		sdc.mutation.SetIsActive(v)
	}
	if _, ok := sdc.mutation.IsLeaf(); !ok {
		v := sysdistrict.DefaultIsLeaf
		sdc.mutation.SetIsLeaf(v)
	}
	if _, ok := sdc.mutation.IsHot(); !ok {
		v := sysdistrict.DefaultIsHot
		sdc.mutation.SetIsHot(v)
	}
	if _, ok := sdc.mutation.IsReal(); !ok {
		v := sysdistrict.DefaultIsReal
		sdc.mutation.SetIsReal(v)
	}
	if _, ok := sdc.mutation.IsMain(); !ok {
		v := sysdistrict.DefaultIsMain
		sdc.mutation.SetIsMain(v)
	}
	if _, ok := sdc.mutation.IsDirect(); !ok {
		v := sysdistrict.DefaultIsDirect
		sdc.mutation.SetIsDirect(v)
	}
	if _, ok := sdc.mutation.ID(); !ok {
		v := sysdistrict.DefaultID()
		sdc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdc *SysDistrictCreate) check() error {
	if _, ok := sdc.mutation.IsDel(); !ok {
		return &ValidationError{Name: "is_del", err: errors.New(`ent: missing required field "SysDistrict.is_del"`)}
	}
	if _, ok := sdc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "SysDistrict.sort"`)}
	}
	if _, ok := sdc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SysDistrict.created_at"`)}
	}
	if _, ok := sdc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SysDistrict.updated_at"`)}
	}
	if _, ok := sdc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "SysDistrict.is_active"`)}
	}
	if v, ok := sdc.mutation.Name(); ok {
		if err := sysdistrict.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "SysDistrict.name": %w`, err)}
		}
	}
	if v, ok := sdc.mutation.Sname(); ok {
		if err := sysdistrict.SnameValidator(v); err != nil {
			return &ValidationError{Name: "sname", err: fmt.Errorf(`ent: validator failed for field "SysDistrict.sname": %w`, err)}
		}
	}
	if v, ok := sdc.mutation.Abbr(); ok {
		if err := sysdistrict.AbbrValidator(v); err != nil {
			return &ValidationError{Name: "abbr", err: fmt.Errorf(`ent: validator failed for field "SysDistrict.abbr": %w`, err)}
		}
	}
	if v, ok := sdc.mutation.StCode(); ok {
		if err := sysdistrict.StCodeValidator(v); err != nil {
			return &ValidationError{Name: "st_code", err: fmt.Errorf(`ent: validator failed for field "SysDistrict.st_code": %w`, err)}
		}
	}
	if v, ok := sdc.mutation.Initials(); ok {
		if err := sysdistrict.InitialsValidator(v); err != nil {
			return &ValidationError{Name: "initials", err: fmt.Errorf(`ent: validator failed for field "SysDistrict.initials": %w`, err)}
		}
	}
	if v, ok := sdc.mutation.Pinyin(); ok {
		if err := sysdistrict.PinyinValidator(v); err != nil {
			return &ValidationError{Name: "pinyin", err: fmt.Errorf(`ent: validator failed for field "SysDistrict.pinyin": %w`, err)}
		}
	}
	if v, ok := sdc.mutation.ParentID(); ok {
		if err := sysdistrict.ParentIDValidator(v); err != nil {
			return &ValidationError{Name: "parent_id", err: fmt.Errorf(`ent: validator failed for field "SysDistrict.parent_id": %w`, err)}
		}
	}
	if v, ok := sdc.mutation.AreaCode(); ok {
		if err := sysdistrict.AreaCodeValidator(v); err != nil {
			return &ValidationError{Name: "area_code", err: fmt.Errorf(`ent: validator failed for field "SysDistrict.area_code": %w`, err)}
		}
	}
	if v, ok := sdc.mutation.ZipCode(); ok {
		if err := sysdistrict.ZipCodeValidator(v); err != nil {
			return &ValidationError{Name: "zip_code", err: fmt.Errorf(`ent: validator failed for field "SysDistrict.zip_code": %w`, err)}
		}
	}
	if v, ok := sdc.mutation.MergeName(); ok {
		if err := sysdistrict.MergeNameValidator(v); err != nil {
			return &ValidationError{Name: "merge_name", err: fmt.Errorf(`ent: validator failed for field "SysDistrict.merge_name": %w`, err)}
		}
	}
	if v, ok := sdc.mutation.MergeSname(); ok {
		if err := sysdistrict.MergeSnameValidator(v); err != nil {
			return &ValidationError{Name: "merge_sname", err: fmt.Errorf(`ent: validator failed for field "SysDistrict.merge_sname": %w`, err)}
		}
	}
	if v, ok := sdc.mutation.Extra(); ok {
		if err := sysdistrict.ExtraValidator(v); err != nil {
			return &ValidationError{Name: "extra", err: fmt.Errorf(`ent: validator failed for field "SysDistrict.extra": %w`, err)}
		}
	}
	if v, ok := sdc.mutation.Suffix(); ok {
		if err := sysdistrict.SuffixValidator(v); err != nil {
			return &ValidationError{Name: "suffix", err: fmt.Errorf(`ent: validator failed for field "SysDistrict.suffix": %w`, err)}
		}
	}
	if _, ok := sdc.mutation.Creator(); !ok {
		return &ValidationError{Name: "creator", err: errors.New(`ent: missing required field "SysDistrict.creator"`)}
	}
	if v, ok := sdc.mutation.Creator(); ok {
		if err := sysdistrict.CreatorValidator(v); err != nil {
			return &ValidationError{Name: "creator", err: fmt.Errorf(`ent: validator failed for field "SysDistrict.creator": %w`, err)}
		}
	}
	if v, ok := sdc.mutation.ID(); ok {
		if err := sysdistrict.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "SysDistrict.id": %w`, err)}
		}
	}
	return nil
}

func (sdc *SysDistrictCreate) sqlSave(ctx context.Context) (*SysDistrict, error) {
	_node, _spec := sdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected SysDistrict.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (sdc *SysDistrictCreate) createSpec() (*SysDistrict, *sqlgraph.CreateSpec) {
	var (
		_node = &SysDistrict{config: sdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysdistrict.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysdistrict.FieldID,
			},
		}
	)
	if id, ok := sdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sdc.mutation.IsDel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysdistrict.FieldIsDel,
		})
		_node.IsDel = value
	}
	if value, ok := sdc.mutation.Sort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdistrict.FieldSort,
		})
		_node.Sort = value
	}
	if value, ok := sdc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdistrict.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sdc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdistrict.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sdc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdistrict.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := sdc.mutation.IsActive(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysdistrict.FieldIsActive,
		})
		_node.IsActive = value
	}
	if value, ok := sdc.mutation.TreeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdistrict.FieldTreeID,
		})
		_node.TreeID = &value
	}
	if value, ok := sdc.mutation.TreeLevel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdistrict.FieldTreeLevel,
		})
		_node.TreeLevel = &value
	}
	if value, ok := sdc.mutation.TreeLeft(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdistrict.FieldTreeLeft,
		})
		_node.TreeLeft = &value
	}
	if value, ok := sdc.mutation.TreeRight(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdistrict.FieldTreeRight,
		})
		_node.TreeRight = &value
	}
	if value, ok := sdc.mutation.IsLeaf(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysdistrict.FieldIsLeaf,
		})
		_node.IsLeaf = &value
	}
	if value, ok := sdc.mutation.TreePath(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdistrict.FieldTreePath,
		})
		_node.TreePath = &value
	}
	if value, ok := sdc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdistrict.FieldName,
		})
		_node.Name = &value
	}
	if value, ok := sdc.mutation.Sname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdistrict.FieldSname,
		})
		_node.Sname = &value
	}
	if value, ok := sdc.mutation.Abbr(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdistrict.FieldAbbr,
		})
		_node.Abbr = &value
	}
	if value, ok := sdc.mutation.StCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdistrict.FieldStCode,
		})
		_node.StCode = &value
	}
	if value, ok := sdc.mutation.Initials(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdistrict.FieldInitials,
		})
		_node.Initials = &value
	}
	if value, ok := sdc.mutation.Pinyin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdistrict.FieldPinyin,
		})
		_node.Pinyin = &value
	}
	if value, ok := sdc.mutation.Longitude(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sysdistrict.FieldLongitude,
		})
		_node.Longitude = &value
	}
	if value, ok := sdc.mutation.Latitude(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sysdistrict.FieldLatitude,
		})
		_node.Latitude = &value
	}
	if value, ok := sdc.mutation.AreaCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdistrict.FieldAreaCode,
		})
		_node.AreaCode = &value
	}
	if value, ok := sdc.mutation.ZipCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdistrict.FieldZipCode,
		})
		_node.ZipCode = &value
	}
	if value, ok := sdc.mutation.MergeName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdistrict.FieldMergeName,
		})
		_node.MergeName = &value
	}
	if value, ok := sdc.mutation.MergeSname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdistrict.FieldMergeSname,
		})
		_node.MergeSname = &value
	}
	if value, ok := sdc.mutation.Extra(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdistrict.FieldExtra,
		})
		_node.Extra = &value
	}
	if value, ok := sdc.mutation.Suffix(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdistrict.FieldSuffix,
		})
		_node.Suffix = &value
	}
	if value, ok := sdc.mutation.IsHot(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysdistrict.FieldIsHot,
		})
		_node.IsHot = &value
	}
	if value, ok := sdc.mutation.IsReal(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysdistrict.FieldIsReal,
		})
		_node.IsReal = &value
	}
	if value, ok := sdc.mutation.IsMain(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysdistrict.FieldIsMain,
		})
		_node.IsMain = &value
	}
	if value, ok := sdc.mutation.IsDirect(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysdistrict.FieldIsDirect,
		})
		_node.IsDirect = &value
	}
	if value, ok := sdc.mutation.Creator(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdistrict.FieldCreator,
		})
		_node.Creator = value
	}
	if nodes := sdc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sysdistrict.ParentTable,
			Columns: []string{sysdistrict.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: sysdistrict.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sdc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysdistrict.ChildrenTable,
			Columns: []string{sysdistrict.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: sysdistrict.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SysDistrictCreateBulk is the builder for creating many SysDistrict entities in bulk.
type SysDistrictCreateBulk struct {
	config
	builders []*SysDistrictCreate
}

// Save creates the SysDistrict entities in the database.
func (sdcb *SysDistrictCreateBulk) Save(ctx context.Context) ([]*SysDistrict, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sdcb.builders))
	nodes := make([]*SysDistrict, len(sdcb.builders))
	mutators := make([]Mutator, len(sdcb.builders))
	for i := range sdcb.builders {
		func(i int, root context.Context) {
			builder := sdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysDistrictMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sdcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sdcb *SysDistrictCreateBulk) SaveX(ctx context.Context) []*SysDistrict {
	v, err := sdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdcb *SysDistrictCreateBulk) Exec(ctx context.Context) error {
	_, err := sdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdcb *SysDistrictCreateBulk) ExecX(ctx context.Context) {
	if err := sdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
