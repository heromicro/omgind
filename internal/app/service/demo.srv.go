package service

import (
	"context"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/schema/repo"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"
)

// DemoSet 注入Demo
var DemoSet = wire.NewSet(wire.Struct(new(Demo), "*"))

// Demo 示例程序
type Demo struct {
	EntCli *ent.Client

	DemoRepo *repo.Demo
}

// ToSchemaDemo 转换为
func toSchemaDemo(a *ent.XxxDemo) *schema.Demo {
	item := new(schema.Demo)
	structure.Copy(a, item)
	return item
}

func toSchemaDemos(us ent.XxxDemos) []*schema.Demo {
	list := make([]*schema.Demo, len(us))
	for i, item := range us {
		list[i] = toSchemaDemo(item)
	}
	return list
}

func toEntCreateDemoInput(a *schema.Demo) *ent.CreateXxxDemoInput {
	createinput := new(ent.CreateXxxDemoInput)
	structure.Copy(a, &createinput)

	return createinput
}

func toEntUpdateDemoInput(a *schema.Demo) *ent.UpdateXxxDemoInput {
	updateinput := new(ent.UpdateXxxDemoInput)
	structure.Copy(a, &updateinput)

	return updateinput
}

func toEntCreateDemoInputs(dms schema.Demos) []*ent.CreateXxxDemoInput {
	list := make([]*ent.CreateXxxDemoInput, len(dms))
	for i, item := range dms {
		list[i] = toEntCreateDemoInput(item)
	}
	return list
}

// Query 查询数据
func (a *Demo) Query(ctx context.Context, params schema.DemoQueryParam, opts ...schema.DemoQueryOptions) (*schema.DemoQueryResult, error) {
	return a.DemoRepo.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *Demo) Get(ctx context.Context, id string, opts ...schema.DemoQueryOptions) (*schema.Demo, error) {
	item, err := a.DemoRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}
	return item, nil
}

func (a *Demo) checkCode(ctx context.Context, code string) error {
	result, err := a.DemoRepo.Query(ctx, schema.DemoQueryParam{
		PaginationParam: schema.PaginationParam{
			OnlyCount: true,
		},
		Code: code,
	})

	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New400Response("编号已经存在")
	}

	return nil
}

// Create 创建数据
func (a *Demo) Create(ctx context.Context, item schema.Demo) (*schema.Demo, error) {

	err := a.checkCode(ctx, item.Code)
	if err != nil {
		return nil, err
	}

	sch_demo, err := a.DemoRepo.Create(ctx, item)

	if err != nil {
		return nil, err
	}

	return sch_demo, nil
}

// Update 更新数据
func (a *Demo) Update(ctx context.Context, id string, item schema.Demo) (*schema.Demo, error) {

	oitem, err := a.DemoRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if oitem == nil {
		return nil, errors.ErrNotFound
	} else if oitem.Code != item.Code {
		if err := a.checkCode(ctx, item.Code); err != nil {
			return nil, err
		}
	}
	item.ID = oitem.ID

	return a.DemoRepo.Update(ctx, id, item)
}

// Delete 删除数据
func (a *Demo) Delete(ctx context.Context, id string) error {

	oldItem, err := a.DemoRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	return a.DemoRepo.Delete(ctx, id)
}

// UpdateStatus 更新状态
func (a *Demo) UpdateStatus(ctx context.Context, id string, isActive bool) error {
	oldItem, err := a.DemoRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.DemoRepo.UpdateStatus(ctx, id, isActive)
}
