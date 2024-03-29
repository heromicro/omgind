// Code generated by ent, DO NOT EDIT.

package mainent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heromicro/omgind/internal/gen/mainent/predicate"
	"github.com/heromicro/omgind/internal/gen/mainent/systeam"
	"github.com/heromicro/omgind/internal/gen/mainent/systeamuser"
	"github.com/heromicro/omgind/internal/gen/mainent/sysuser"
)

// SysTeamUpdate is the builder for updating SysTeam entities.
type SysTeamUpdate struct {
	config
	hooks     []Hook
	mutation  *SysTeamMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SysTeamUpdate builder.
func (stu *SysTeamUpdate) Where(ps ...predicate.SysTeam) *SysTeamUpdate {
	stu.mutation.Where(ps...)
	return stu
}

// SetSort sets the "sort" field.
func (stu *SysTeamUpdate) SetSort(i int32) *SysTeamUpdate {
	stu.mutation.ResetSort()
	stu.mutation.SetSort(i)
	return stu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (stu *SysTeamUpdate) SetNillableSort(i *int32) *SysTeamUpdate {
	if i != nil {
		stu.SetSort(*i)
	}
	return stu
}

// AddSort adds i to the "sort" field.
func (stu *SysTeamUpdate) AddSort(i int32) *SysTeamUpdate {
	stu.mutation.AddSort(i)
	return stu
}

// SetUpdatedAt sets the "updated_at" field.
func (stu *SysTeamUpdate) SetUpdatedAt(t time.Time) *SysTeamUpdate {
	stu.mutation.SetUpdatedAt(t)
	return stu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (stu *SysTeamUpdate) ClearUpdatedAt() *SysTeamUpdate {
	stu.mutation.ClearUpdatedAt()
	return stu
}

// SetDeletedAt sets the "deleted_at" field.
func (stu *SysTeamUpdate) SetDeletedAt(t time.Time) *SysTeamUpdate {
	stu.mutation.SetDeletedAt(t)
	return stu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (stu *SysTeamUpdate) SetNillableDeletedAt(t *time.Time) *SysTeamUpdate {
	if t != nil {
		stu.SetDeletedAt(*t)
	}
	return stu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (stu *SysTeamUpdate) ClearDeletedAt() *SysTeamUpdate {
	stu.mutation.ClearDeletedAt()
	return stu
}

// SetIsActive sets the "is_active" field.
func (stu *SysTeamUpdate) SetIsActive(b bool) *SysTeamUpdate {
	stu.mutation.SetIsActive(b)
	return stu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (stu *SysTeamUpdate) SetNillableIsActive(b *bool) *SysTeamUpdate {
	if b != nil {
		stu.SetIsActive(*b)
	}
	return stu
}

// SetMemo sets the "memo" field.
func (stu *SysTeamUpdate) SetMemo(s string) *SysTeamUpdate {
	stu.mutation.SetMemo(s)
	return stu
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (stu *SysTeamUpdate) SetNillableMemo(s *string) *SysTeamUpdate {
	if s != nil {
		stu.SetMemo(*s)
	}
	return stu
}

// ClearMemo clears the value of the "memo" field.
func (stu *SysTeamUpdate) ClearMemo() *SysTeamUpdate {
	stu.mutation.ClearMemo()
	return stu
}

// SetIsDel sets the "is_del" field.
func (stu *SysTeamUpdate) SetIsDel(b bool) *SysTeamUpdate {
	stu.mutation.SetIsDel(b)
	return stu
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (stu *SysTeamUpdate) SetNillableIsDel(b *bool) *SysTeamUpdate {
	if b != nil {
		stu.SetIsDel(*b)
	}
	return stu
}

// SetName sets the "name" field.
func (stu *SysTeamUpdate) SetName(s string) *SysTeamUpdate {
	stu.mutation.SetName(s)
	return stu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (stu *SysTeamUpdate) SetNillableName(s *string) *SysTeamUpdate {
	if s != nil {
		stu.SetName(*s)
	}
	return stu
}

// ClearName clears the value of the "name" field.
func (stu *SysTeamUpdate) ClearName() *SysTeamUpdate {
	stu.mutation.ClearName()
	return stu
}

// SetCode sets the "code" field.
func (stu *SysTeamUpdate) SetCode(s string) *SysTeamUpdate {
	stu.mutation.SetCode(s)
	return stu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (stu *SysTeamUpdate) SetNillableCode(s *string) *SysTeamUpdate {
	if s != nil {
		stu.SetCode(*s)
	}
	return stu
}

// ClearCode clears the value of the "code" field.
func (stu *SysTeamUpdate) ClearCode() *SysTeamUpdate {
	stu.mutation.ClearCode()
	return stu
}

// AddUserIDs adds the "users" edge to the SysUser entity by IDs.
func (stu *SysTeamUpdate) AddUserIDs(ids ...string) *SysTeamUpdate {
	stu.mutation.AddUserIDs(ids...)
	return stu
}

// AddUsers adds the "users" edges to the SysUser entity.
func (stu *SysTeamUpdate) AddUsers(s ...*SysUser) *SysTeamUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stu.AddUserIDs(ids...)
}

// AddTeamUserIDs adds the "team_users" edge to the SysTeamUser entity by IDs.
func (stu *SysTeamUpdate) AddTeamUserIDs(ids ...string) *SysTeamUpdate {
	stu.mutation.AddTeamUserIDs(ids...)
	return stu
}

// AddTeamUsers adds the "team_users" edges to the SysTeamUser entity.
func (stu *SysTeamUpdate) AddTeamUsers(s ...*SysTeamUser) *SysTeamUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stu.AddTeamUserIDs(ids...)
}

// Mutation returns the SysTeamMutation object of the builder.
func (stu *SysTeamUpdate) Mutation() *SysTeamMutation {
	return stu.mutation
}

// ClearUsers clears all "users" edges to the SysUser entity.
func (stu *SysTeamUpdate) ClearUsers() *SysTeamUpdate {
	stu.mutation.ClearUsers()
	return stu
}

// RemoveUserIDs removes the "users" edge to SysUser entities by IDs.
func (stu *SysTeamUpdate) RemoveUserIDs(ids ...string) *SysTeamUpdate {
	stu.mutation.RemoveUserIDs(ids...)
	return stu
}

// RemoveUsers removes "users" edges to SysUser entities.
func (stu *SysTeamUpdate) RemoveUsers(s ...*SysUser) *SysTeamUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stu.RemoveUserIDs(ids...)
}

// ClearTeamUsers clears all "team_users" edges to the SysTeamUser entity.
func (stu *SysTeamUpdate) ClearTeamUsers() *SysTeamUpdate {
	stu.mutation.ClearTeamUsers()
	return stu
}

// RemoveTeamUserIDs removes the "team_users" edge to SysTeamUser entities by IDs.
func (stu *SysTeamUpdate) RemoveTeamUserIDs(ids ...string) *SysTeamUpdate {
	stu.mutation.RemoveTeamUserIDs(ids...)
	return stu
}

// RemoveTeamUsers removes "team_users" edges to SysTeamUser entities.
func (stu *SysTeamUpdate) RemoveTeamUsers(s ...*SysTeamUser) *SysTeamUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stu.RemoveTeamUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (stu *SysTeamUpdate) Save(ctx context.Context) (int, error) {
	stu.defaults()
	return withHooks(ctx, stu.sqlSave, stu.mutation, stu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (stu *SysTeamUpdate) SaveX(ctx context.Context) int {
	affected, err := stu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (stu *SysTeamUpdate) Exec(ctx context.Context) error {
	_, err := stu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stu *SysTeamUpdate) ExecX(ctx context.Context) {
	if err := stu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (stu *SysTeamUpdate) defaults() {
	if _, ok := stu.mutation.UpdatedAt(); !ok && !stu.mutation.UpdatedAtCleared() {
		v := systeam.UpdateDefaultUpdatedAt()
		stu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stu *SysTeamUpdate) check() error {
	if v, ok := stu.mutation.Memo(); ok {
		if err := systeam.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`mainent: validator failed for field "SysTeam.memo": %w`, err)}
		}
	}
	if v, ok := stu.mutation.Name(); ok {
		if err := systeam.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`mainent: validator failed for field "SysTeam.name": %w`, err)}
		}
	}
	if v, ok := stu.mutation.Code(); ok {
		if err := systeam.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`mainent: validator failed for field "SysTeam.code": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (stu *SysTeamUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysTeamUpdate {
	stu.modifiers = append(stu.modifiers, modifiers...)
	return stu
}

func (stu *SysTeamUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := stu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(systeam.Table, systeam.Columns, sqlgraph.NewFieldSpec(systeam.FieldID, field.TypeString))
	if ps := stu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stu.mutation.Sort(); ok {
		_spec.SetField(systeam.FieldSort, field.TypeInt32, value)
	}
	if value, ok := stu.mutation.AddedSort(); ok {
		_spec.AddField(systeam.FieldSort, field.TypeInt32, value)
	}
	if stu.mutation.CreatedAtCleared() {
		_spec.ClearField(systeam.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := stu.mutation.UpdatedAt(); ok {
		_spec.SetField(systeam.FieldUpdatedAt, field.TypeTime, value)
	}
	if stu.mutation.UpdatedAtCleared() {
		_spec.ClearField(systeam.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := stu.mutation.DeletedAt(); ok {
		_spec.SetField(systeam.FieldDeletedAt, field.TypeTime, value)
	}
	if stu.mutation.DeletedAtCleared() {
		_spec.ClearField(systeam.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := stu.mutation.IsActive(); ok {
		_spec.SetField(systeam.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := stu.mutation.Memo(); ok {
		_spec.SetField(systeam.FieldMemo, field.TypeString, value)
	}
	if stu.mutation.MemoCleared() {
		_spec.ClearField(systeam.FieldMemo, field.TypeString)
	}
	if value, ok := stu.mutation.IsDel(); ok {
		_spec.SetField(systeam.FieldIsDel, field.TypeBool, value)
	}
	if value, ok := stu.mutation.Name(); ok {
		_spec.SetField(systeam.FieldName, field.TypeString, value)
	}
	if stu.mutation.NameCleared() {
		_spec.ClearField(systeam.FieldName, field.TypeString)
	}
	if value, ok := stu.mutation.Code(); ok {
		_spec.SetField(systeam.FieldCode, field.TypeString, value)
	}
	if stu.mutation.CodeCleared() {
		_spec.ClearField(systeam.FieldCode, field.TypeString)
	}
	if stu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   systeam.UsersTable,
			Columns: systeam.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeString),
			},
		}
		createE := &SysTeamUserCreate{config: stu.config, mutation: newSysTeamUserMutation(stu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !stu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   systeam.UsersTable,
			Columns: systeam.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &SysTeamUserCreate{config: stu.config, mutation: newSysTeamUserMutation(stu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   systeam.UsersTable,
			Columns: systeam.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &SysTeamUserCreate{config: stu.config, mutation: newSysTeamUserMutation(stu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if stu.mutation.TeamUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   systeam.TeamUsersTable,
			Columns: []string{systeam.TeamUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(systeamuser.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stu.mutation.RemovedTeamUsersIDs(); len(nodes) > 0 && !stu.mutation.TeamUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   systeam.TeamUsersTable,
			Columns: []string{systeam.TeamUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(systeamuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stu.mutation.TeamUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   systeam.TeamUsersTable,
			Columns: []string{systeam.TeamUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(systeamuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(stu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, stu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{systeam.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	stu.mutation.done = true
	return n, nil
}

// SysTeamUpdateOne is the builder for updating a single SysTeam entity.
type SysTeamUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SysTeamMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetSort sets the "sort" field.
func (stuo *SysTeamUpdateOne) SetSort(i int32) *SysTeamUpdateOne {
	stuo.mutation.ResetSort()
	stuo.mutation.SetSort(i)
	return stuo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (stuo *SysTeamUpdateOne) SetNillableSort(i *int32) *SysTeamUpdateOne {
	if i != nil {
		stuo.SetSort(*i)
	}
	return stuo
}

// AddSort adds i to the "sort" field.
func (stuo *SysTeamUpdateOne) AddSort(i int32) *SysTeamUpdateOne {
	stuo.mutation.AddSort(i)
	return stuo
}

// SetUpdatedAt sets the "updated_at" field.
func (stuo *SysTeamUpdateOne) SetUpdatedAt(t time.Time) *SysTeamUpdateOne {
	stuo.mutation.SetUpdatedAt(t)
	return stuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (stuo *SysTeamUpdateOne) ClearUpdatedAt() *SysTeamUpdateOne {
	stuo.mutation.ClearUpdatedAt()
	return stuo
}

// SetDeletedAt sets the "deleted_at" field.
func (stuo *SysTeamUpdateOne) SetDeletedAt(t time.Time) *SysTeamUpdateOne {
	stuo.mutation.SetDeletedAt(t)
	return stuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (stuo *SysTeamUpdateOne) SetNillableDeletedAt(t *time.Time) *SysTeamUpdateOne {
	if t != nil {
		stuo.SetDeletedAt(*t)
	}
	return stuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (stuo *SysTeamUpdateOne) ClearDeletedAt() *SysTeamUpdateOne {
	stuo.mutation.ClearDeletedAt()
	return stuo
}

// SetIsActive sets the "is_active" field.
func (stuo *SysTeamUpdateOne) SetIsActive(b bool) *SysTeamUpdateOne {
	stuo.mutation.SetIsActive(b)
	return stuo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (stuo *SysTeamUpdateOne) SetNillableIsActive(b *bool) *SysTeamUpdateOne {
	if b != nil {
		stuo.SetIsActive(*b)
	}
	return stuo
}

// SetMemo sets the "memo" field.
func (stuo *SysTeamUpdateOne) SetMemo(s string) *SysTeamUpdateOne {
	stuo.mutation.SetMemo(s)
	return stuo
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (stuo *SysTeamUpdateOne) SetNillableMemo(s *string) *SysTeamUpdateOne {
	if s != nil {
		stuo.SetMemo(*s)
	}
	return stuo
}

// ClearMemo clears the value of the "memo" field.
func (stuo *SysTeamUpdateOne) ClearMemo() *SysTeamUpdateOne {
	stuo.mutation.ClearMemo()
	return stuo
}

// SetIsDel sets the "is_del" field.
func (stuo *SysTeamUpdateOne) SetIsDel(b bool) *SysTeamUpdateOne {
	stuo.mutation.SetIsDel(b)
	return stuo
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (stuo *SysTeamUpdateOne) SetNillableIsDel(b *bool) *SysTeamUpdateOne {
	if b != nil {
		stuo.SetIsDel(*b)
	}
	return stuo
}

// SetName sets the "name" field.
func (stuo *SysTeamUpdateOne) SetName(s string) *SysTeamUpdateOne {
	stuo.mutation.SetName(s)
	return stuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (stuo *SysTeamUpdateOne) SetNillableName(s *string) *SysTeamUpdateOne {
	if s != nil {
		stuo.SetName(*s)
	}
	return stuo
}

// ClearName clears the value of the "name" field.
func (stuo *SysTeamUpdateOne) ClearName() *SysTeamUpdateOne {
	stuo.mutation.ClearName()
	return stuo
}

// SetCode sets the "code" field.
func (stuo *SysTeamUpdateOne) SetCode(s string) *SysTeamUpdateOne {
	stuo.mutation.SetCode(s)
	return stuo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (stuo *SysTeamUpdateOne) SetNillableCode(s *string) *SysTeamUpdateOne {
	if s != nil {
		stuo.SetCode(*s)
	}
	return stuo
}

// ClearCode clears the value of the "code" field.
func (stuo *SysTeamUpdateOne) ClearCode() *SysTeamUpdateOne {
	stuo.mutation.ClearCode()
	return stuo
}

// AddUserIDs adds the "users" edge to the SysUser entity by IDs.
func (stuo *SysTeamUpdateOne) AddUserIDs(ids ...string) *SysTeamUpdateOne {
	stuo.mutation.AddUserIDs(ids...)
	return stuo
}

// AddUsers adds the "users" edges to the SysUser entity.
func (stuo *SysTeamUpdateOne) AddUsers(s ...*SysUser) *SysTeamUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stuo.AddUserIDs(ids...)
}

// AddTeamUserIDs adds the "team_users" edge to the SysTeamUser entity by IDs.
func (stuo *SysTeamUpdateOne) AddTeamUserIDs(ids ...string) *SysTeamUpdateOne {
	stuo.mutation.AddTeamUserIDs(ids...)
	return stuo
}

// AddTeamUsers adds the "team_users" edges to the SysTeamUser entity.
func (stuo *SysTeamUpdateOne) AddTeamUsers(s ...*SysTeamUser) *SysTeamUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stuo.AddTeamUserIDs(ids...)
}

// Mutation returns the SysTeamMutation object of the builder.
func (stuo *SysTeamUpdateOne) Mutation() *SysTeamMutation {
	return stuo.mutation
}

// ClearUsers clears all "users" edges to the SysUser entity.
func (stuo *SysTeamUpdateOne) ClearUsers() *SysTeamUpdateOne {
	stuo.mutation.ClearUsers()
	return stuo
}

// RemoveUserIDs removes the "users" edge to SysUser entities by IDs.
func (stuo *SysTeamUpdateOne) RemoveUserIDs(ids ...string) *SysTeamUpdateOne {
	stuo.mutation.RemoveUserIDs(ids...)
	return stuo
}

// RemoveUsers removes "users" edges to SysUser entities.
func (stuo *SysTeamUpdateOne) RemoveUsers(s ...*SysUser) *SysTeamUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stuo.RemoveUserIDs(ids...)
}

// ClearTeamUsers clears all "team_users" edges to the SysTeamUser entity.
func (stuo *SysTeamUpdateOne) ClearTeamUsers() *SysTeamUpdateOne {
	stuo.mutation.ClearTeamUsers()
	return stuo
}

// RemoveTeamUserIDs removes the "team_users" edge to SysTeamUser entities by IDs.
func (stuo *SysTeamUpdateOne) RemoveTeamUserIDs(ids ...string) *SysTeamUpdateOne {
	stuo.mutation.RemoveTeamUserIDs(ids...)
	return stuo
}

// RemoveTeamUsers removes "team_users" edges to SysTeamUser entities.
func (stuo *SysTeamUpdateOne) RemoveTeamUsers(s ...*SysTeamUser) *SysTeamUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stuo.RemoveTeamUserIDs(ids...)
}

// Where appends a list predicates to the SysTeamUpdate builder.
func (stuo *SysTeamUpdateOne) Where(ps ...predicate.SysTeam) *SysTeamUpdateOne {
	stuo.mutation.Where(ps...)
	return stuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (stuo *SysTeamUpdateOne) Select(field string, fields ...string) *SysTeamUpdateOne {
	stuo.fields = append([]string{field}, fields...)
	return stuo
}

// Save executes the query and returns the updated SysTeam entity.
func (stuo *SysTeamUpdateOne) Save(ctx context.Context) (*SysTeam, error) {
	stuo.defaults()
	return withHooks(ctx, stuo.sqlSave, stuo.mutation, stuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (stuo *SysTeamUpdateOne) SaveX(ctx context.Context) *SysTeam {
	node, err := stuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (stuo *SysTeamUpdateOne) Exec(ctx context.Context) error {
	_, err := stuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stuo *SysTeamUpdateOne) ExecX(ctx context.Context) {
	if err := stuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (stuo *SysTeamUpdateOne) defaults() {
	if _, ok := stuo.mutation.UpdatedAt(); !ok && !stuo.mutation.UpdatedAtCleared() {
		v := systeam.UpdateDefaultUpdatedAt()
		stuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stuo *SysTeamUpdateOne) check() error {
	if v, ok := stuo.mutation.Memo(); ok {
		if err := systeam.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`mainent: validator failed for field "SysTeam.memo": %w`, err)}
		}
	}
	if v, ok := stuo.mutation.Name(); ok {
		if err := systeam.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`mainent: validator failed for field "SysTeam.name": %w`, err)}
		}
	}
	if v, ok := stuo.mutation.Code(); ok {
		if err := systeam.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`mainent: validator failed for field "SysTeam.code": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (stuo *SysTeamUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SysTeamUpdateOne {
	stuo.modifiers = append(stuo.modifiers, modifiers...)
	return stuo
}

func (stuo *SysTeamUpdateOne) sqlSave(ctx context.Context) (_node *SysTeam, err error) {
	if err := stuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(systeam.Table, systeam.Columns, sqlgraph.NewFieldSpec(systeam.FieldID, field.TypeString))
	id, ok := stuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`mainent: missing "SysTeam.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := stuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, systeam.FieldID)
		for _, f := range fields {
			if !systeam.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("mainent: invalid field %q for query", f)}
			}
			if f != systeam.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := stuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stuo.mutation.Sort(); ok {
		_spec.SetField(systeam.FieldSort, field.TypeInt32, value)
	}
	if value, ok := stuo.mutation.AddedSort(); ok {
		_spec.AddField(systeam.FieldSort, field.TypeInt32, value)
	}
	if stuo.mutation.CreatedAtCleared() {
		_spec.ClearField(systeam.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := stuo.mutation.UpdatedAt(); ok {
		_spec.SetField(systeam.FieldUpdatedAt, field.TypeTime, value)
	}
	if stuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(systeam.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := stuo.mutation.DeletedAt(); ok {
		_spec.SetField(systeam.FieldDeletedAt, field.TypeTime, value)
	}
	if stuo.mutation.DeletedAtCleared() {
		_spec.ClearField(systeam.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := stuo.mutation.IsActive(); ok {
		_spec.SetField(systeam.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := stuo.mutation.Memo(); ok {
		_spec.SetField(systeam.FieldMemo, field.TypeString, value)
	}
	if stuo.mutation.MemoCleared() {
		_spec.ClearField(systeam.FieldMemo, field.TypeString)
	}
	if value, ok := stuo.mutation.IsDel(); ok {
		_spec.SetField(systeam.FieldIsDel, field.TypeBool, value)
	}
	if value, ok := stuo.mutation.Name(); ok {
		_spec.SetField(systeam.FieldName, field.TypeString, value)
	}
	if stuo.mutation.NameCleared() {
		_spec.ClearField(systeam.FieldName, field.TypeString)
	}
	if value, ok := stuo.mutation.Code(); ok {
		_spec.SetField(systeam.FieldCode, field.TypeString, value)
	}
	if stuo.mutation.CodeCleared() {
		_spec.ClearField(systeam.FieldCode, field.TypeString)
	}
	if stuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   systeam.UsersTable,
			Columns: systeam.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeString),
			},
		}
		createE := &SysTeamUserCreate{config: stuo.config, mutation: newSysTeamUserMutation(stuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stuo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !stuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   systeam.UsersTable,
			Columns: systeam.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &SysTeamUserCreate{config: stuo.config, mutation: newSysTeamUserMutation(stuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   systeam.UsersTable,
			Columns: systeam.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &SysTeamUserCreate{config: stuo.config, mutation: newSysTeamUserMutation(stuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if stuo.mutation.TeamUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   systeam.TeamUsersTable,
			Columns: []string{systeam.TeamUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(systeamuser.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stuo.mutation.RemovedTeamUsersIDs(); len(nodes) > 0 && !stuo.mutation.TeamUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   systeam.TeamUsersTable,
			Columns: []string{systeam.TeamUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(systeamuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stuo.mutation.TeamUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   systeam.TeamUsersTable,
			Columns: []string{systeam.TeamUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(systeamuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(stuo.modifiers...)
	_node = &SysTeam{config: stuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, stuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{systeam.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	stuo.mutation.done = true
	return _node, nil
}
