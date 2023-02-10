// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/heromicro/omgind/internal/gen/ent"
)

// The SysDictFunc type is an adapter to allow the use of ordinary
// function as SysDict mutator.
type SysDictFunc func(context.Context, *ent.SysDictMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SysDictFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SysDictMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SysDictMutation", m)
	}
	return f(ctx, mv)
}

// The SysDictItemFunc type is an adapter to allow the use of ordinary
// function as SysDictItem mutator.
type SysDictItemFunc func(context.Context, *ent.SysDictItemMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SysDictItemFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SysDictItemMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SysDictItemMutation", m)
	}
	return f(ctx, mv)
}

// The SysJwtBlockFunc type is an adapter to allow the use of ordinary
// function as SysJwtBlock mutator.
type SysJwtBlockFunc func(context.Context, *ent.SysJwtBlockMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SysJwtBlockFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SysJwtBlockMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SysJwtBlockMutation", m)
	}
	return f(ctx, mv)
}

// The SysLoggingFunc type is an adapter to allow the use of ordinary
// function as SysLogging mutator.
type SysLoggingFunc func(context.Context, *ent.SysLoggingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SysLoggingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SysLoggingMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SysLoggingMutation", m)
	}
	return f(ctx, mv)
}

// The SysMenuFunc type is an adapter to allow the use of ordinary
// function as SysMenu mutator.
type SysMenuFunc func(context.Context, *ent.SysMenuMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SysMenuFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SysMenuMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SysMenuMutation", m)
	}
	return f(ctx, mv)
}

// The SysMenuActionFunc type is an adapter to allow the use of ordinary
// function as SysMenuAction mutator.
type SysMenuActionFunc func(context.Context, *ent.SysMenuActionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SysMenuActionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SysMenuActionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SysMenuActionMutation", m)
	}
	return f(ctx, mv)
}

// The SysMenuActionResourceFunc type is an adapter to allow the use of ordinary
// function as SysMenuActionResource mutator.
type SysMenuActionResourceFunc func(context.Context, *ent.SysMenuActionResourceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SysMenuActionResourceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SysMenuActionResourceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SysMenuActionResourceMutation", m)
	}
	return f(ctx, mv)
}

// The SysRoleFunc type is an adapter to allow the use of ordinary
// function as SysRole mutator.
type SysRoleFunc func(context.Context, *ent.SysRoleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SysRoleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SysRoleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SysRoleMutation", m)
	}
	return f(ctx, mv)
}

// The SysRoleMenuFunc type is an adapter to allow the use of ordinary
// function as SysRoleMenu mutator.
type SysRoleMenuFunc func(context.Context, *ent.SysRoleMenuMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SysRoleMenuFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SysRoleMenuMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SysRoleMenuMutation", m)
	}
	return f(ctx, mv)
}

// The SysUserFunc type is an adapter to allow the use of ordinary
// function as SysUser mutator.
type SysUserFunc func(context.Context, *ent.SysUserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SysUserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SysUserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SysUserMutation", m)
	}
	return f(ctx, mv)
}

// The SysUserRoleFunc type is an adapter to allow the use of ordinary
// function as SysUserRole mutator.
type SysUserRoleFunc func(context.Context, *ent.SysUserRoleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SysUserRoleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SysUserRoleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SysUserRoleMutation", m)
	}
	return f(ctx, mv)
}

// The XxxDemoFunc type is an adapter to allow the use of ordinary
// function as XxxDemo mutator.
type XxxDemoFunc func(context.Context, *ent.XxxDemoMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f XxxDemoFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.XxxDemoMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.XxxDemoMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
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
	return func(ctx context.Context, m ent.Mutation) bool {
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
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
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
	return func(_ context.Context, m ent.Mutation) bool {
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
	return func(_ context.Context, m ent.Mutation) bool {
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
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
