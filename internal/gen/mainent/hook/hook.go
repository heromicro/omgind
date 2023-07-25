// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/heromicro/omgind/internal/gen/mainent"
)

// The OrgDeptFunc type is an adapter to allow the use of ordinary
// function as OrgDept mutator.
type OrgDeptFunc func(context.Context, *mainent.OrgDeptMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f OrgDeptFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.OrgDeptMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.OrgDeptMutation", m)
}

// The OrgOrganFunc type is an adapter to allow the use of ordinary
// function as OrgOrgan mutator.
type OrgOrganFunc func(context.Context, *mainent.OrgOrganMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f OrgOrganFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.OrgOrganMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.OrgOrganMutation", m)
}

// The OrgPositionFunc type is an adapter to allow the use of ordinary
// function as OrgPosition mutator.
type OrgPositionFunc func(context.Context, *mainent.OrgPositionMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f OrgPositionFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.OrgPositionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.OrgPositionMutation", m)
}

// The OrgStaffFunc type is an adapter to allow the use of ordinary
// function as OrgStaff mutator.
type OrgStaffFunc func(context.Context, *mainent.OrgStaffMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f OrgStaffFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.OrgStaffMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.OrgStaffMutation", m)
}

// The SysAddressFunc type is an adapter to allow the use of ordinary
// function as SysAddress mutator.
type SysAddressFunc func(context.Context, *mainent.SysAddressMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f SysAddressFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.SysAddressMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.SysAddressMutation", m)
}

// The SysDictFunc type is an adapter to allow the use of ordinary
// function as SysDict mutator.
type SysDictFunc func(context.Context, *mainent.SysDictMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f SysDictFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.SysDictMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.SysDictMutation", m)
}

// The SysDictItemFunc type is an adapter to allow the use of ordinary
// function as SysDictItem mutator.
type SysDictItemFunc func(context.Context, *mainent.SysDictItemMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f SysDictItemFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.SysDictItemMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.SysDictItemMutation", m)
}

// The SysDistrictFunc type is an adapter to allow the use of ordinary
// function as SysDistrict mutator.
type SysDistrictFunc func(context.Context, *mainent.SysDistrictMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f SysDistrictFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.SysDistrictMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.SysDistrictMutation", m)
}

// The SysJwtBlockFunc type is an adapter to allow the use of ordinary
// function as SysJwtBlock mutator.
type SysJwtBlockFunc func(context.Context, *mainent.SysJwtBlockMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f SysJwtBlockFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.SysJwtBlockMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.SysJwtBlockMutation", m)
}

// The SysLoggingFunc type is an adapter to allow the use of ordinary
// function as SysLogging mutator.
type SysLoggingFunc func(context.Context, *mainent.SysLoggingMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f SysLoggingFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.SysLoggingMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.SysLoggingMutation", m)
}

// The SysMenuFunc type is an adapter to allow the use of ordinary
// function as SysMenu mutator.
type SysMenuFunc func(context.Context, *mainent.SysMenuMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f SysMenuFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.SysMenuMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.SysMenuMutation", m)
}

// The SysMenuActionFunc type is an adapter to allow the use of ordinary
// function as SysMenuAction mutator.
type SysMenuActionFunc func(context.Context, *mainent.SysMenuActionMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f SysMenuActionFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.SysMenuActionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.SysMenuActionMutation", m)
}

// The SysMenuActionResourceFunc type is an adapter to allow the use of ordinary
// function as SysMenuActionResource mutator.
type SysMenuActionResourceFunc func(context.Context, *mainent.SysMenuActionResourceMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f SysMenuActionResourceFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.SysMenuActionResourceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.SysMenuActionResourceMutation", m)
}

// The SysRoleFunc type is an adapter to allow the use of ordinary
// function as SysRole mutator.
type SysRoleFunc func(context.Context, *mainent.SysRoleMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f SysRoleFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.SysRoleMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.SysRoleMutation", m)
}

// The SysRoleMenuFunc type is an adapter to allow the use of ordinary
// function as SysRoleMenu mutator.
type SysRoleMenuFunc func(context.Context, *mainent.SysRoleMenuMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f SysRoleMenuFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.SysRoleMenuMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.SysRoleMenuMutation", m)
}

// The SysTeamFunc type is an adapter to allow the use of ordinary
// function as SysTeam mutator.
type SysTeamFunc func(context.Context, *mainent.SysTeamMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f SysTeamFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.SysTeamMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.SysTeamMutation", m)
}

// The SysTeamUserFunc type is an adapter to allow the use of ordinary
// function as SysTeamUser mutator.
type SysTeamUserFunc func(context.Context, *mainent.SysTeamUserMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f SysTeamUserFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.SysTeamUserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.SysTeamUserMutation", m)
}

// The SysUserFunc type is an adapter to allow the use of ordinary
// function as SysUser mutator.
type SysUserFunc func(context.Context, *mainent.SysUserMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f SysUserFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.SysUserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.SysUserMutation", m)
}

// The SysUserRoleFunc type is an adapter to allow the use of ordinary
// function as SysUserRole mutator.
type SysUserRoleFunc func(context.Context, *mainent.SysUserRoleMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f SysUserRoleFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.SysUserRoleMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.SysUserRoleMutation", m)
}

// The XxxDemoFunc type is an adapter to allow the use of ordinary
// function as XxxDemo mutator.
type XxxDemoFunc func(context.Context, *mainent.XxxDemoMutation) (mainent.Value, error)

// Mutate calls f(ctx, m).
func (f XxxDemoFunc) Mutate(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
	if mv, ok := m.(*mainent.XxxDemoMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *mainent.XxxDemoMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, mainent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m mainent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m mainent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m mainent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op mainent.Op) Condition {
	return func(_ context.Context, m mainent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m mainent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m mainent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m mainent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk mainent.Hook, cond Condition) mainent.Hook {
	return func(next mainent.Mutator) mainent.Mutator {
		return mainent.MutateFunc(func(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, mainent.Delete|mainent.Create)
func On(hk mainent.Hook, op mainent.Op) mainent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, mainent.Update|mainent.UpdateOne)
func Unless(hk mainent.Hook, op mainent.Op) mainent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) mainent.Hook {
	return func(mainent.Mutator) mainent.Mutator {
		return mainent.MutateFunc(func(context.Context, mainent.Mutation) (mainent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []mainent.Hook {
//		return []mainent.Hook{
//			Reject(mainent.Delete|mainent.Update),
//		}
//	}
func Reject(op mainent.Op) mainent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []mainent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...mainent.Hook) Chain {
	return Chain{append([]mainent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() mainent.Hook {
	return func(mutator mainent.Mutator) mainent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...mainent.Hook) Chain {
	newHooks := make([]mainent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
