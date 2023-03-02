package mixin

import (
	"context"
	"fmt"

	"entgo.io/ent"

	gen "github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/hook"
)

func (MpttTreeMixin) Hooks() []ent.Hook {
	return []ent.Hook{

		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.SysDistrictFunc(func(ctx context.Context, m *gen.SysDistrictMutation) (ent.Value, error) {
					isLeaf, ok := m.IsLeaf()
					if !ok {
						return nil, fmt.Errorf("is_leaf is not ok")
					}
					treeRight, ok := m.TreeRight()
					if !ok {
						return nil, fmt.Errorf("tree right is not ok")
					}
					treeLeft, ok := m.TreeLeft()
					if !ok {
						return nil, fmt.Errorf("tree left is not ok")
					}

					if isLeaf && (treeRight-treeLeft) != 1 {
						return nil, fmt.Errorf("tree is out of order ")
					}
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate|ent.OpDelete|ent.OpUpdate,
		),
	}
}
