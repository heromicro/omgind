package entity

// all entity hooks in the file

///*
import (
	"context"
	"fmt"

	"entgo.io/ent"

	gen "github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/hook"
)

// doc https://entgo.io/docs/hooks/#hooks-registration

func (SysAddress) Hooks() []ent.Hook {
	return []ent.Hook{
		//

	}
}

func (SysDistrict) Hooks() []ent.Hook {
	return []ent.Hook{
		//

	}
}

func (SysDictItem) Hooks() []ent.Hook {
	return []ent.Hook{
		//

	}
}

func (SysDict) Hooks() []ent.Hook {
	return []ent.Hook{
		//

	}
}

func (SysLogging) Hooks() []ent.Hook {
	return []ent.Hook{

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

		//

	}
}

func (SysMenuActionResource) Hooks() []ent.Hook {
	return []ent.Hook{
		//

	}
}

func (SysMenuAction) Hooks() []ent.Hook {
	return []ent.Hook{
		//

	}
}

func (SysMenu) Hooks() []ent.Hook {
	return []ent.Hook{
		//

	}
}

func (SysRole) Hooks() []ent.Hook {
	return []ent.Hook{
		//

	}
}

func (SysUser) Hooks() []ent.Hook {
	return []ent.Hook{
		//

	}
}

//*/ ///
