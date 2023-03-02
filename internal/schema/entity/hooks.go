package entity

import "entgo.io/ent"

func (SysDistrict) Hooks() []ent.Hook {
	return []ent.Hook{
		/*
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
		*/
	}
}

func (SysLogging) Hooks() []ent.Hook {
	return []ent.Hook{

		/*
			hook.On(
				func(next ent.Mutator) ent.Mutator {
					return hook.SysLoggingFunc(func(ctx context.Context, m *gen.SysLoggingMutation) (ent.Value, error) {
						_, ok := m.ID()
						if !ok {
							return nil, fmt.Errorf("id is not ok")
						}
						return next.Mutate(ctx, m)
					})
				},
				ent.OpCreate,
			),
		*/
		//

	}
}
