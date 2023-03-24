package service

import (
	"context"
	"strings"

	"github.com/google/wire"
	"github.com/gotidy/ptr"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/orgdept"
	"github.com/heromicro/omgind/internal/gen/ent/orgorgan"
	"github.com/heromicro/omgind/internal/schema/repo"
	"github.com/heromicro/omgind/pkg/errors"
)

// OrgDeptSet 注入OrgDept
var OrgDeptSet = wire.NewSet(wire.Struct(new(OrgDept), "*"))

// OrgDept 部门管理
type OrgDept struct {
	EntCli *ent.Client

	OrgDeptRepo *repo.OrgDept
}

// Query 查询数据
func (a *OrgDept) Query(ctx context.Context, params schema.OrgDeptQueryParam, opts ...schema.OrgDeptQueryOptions) (*schema.OrgDeptQueryResult, error) {
	return a.OrgDeptRepo.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *OrgDept) Get(ctx context.Context, id string, opts ...schema.OrgDeptQueryOptions) (*schema.OrgDept, error) {
	item, err := a.OrgDeptRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// View 查询指定数据
func (a *OrgDept) View(ctx context.Context, id string, opts ...schema.OrgDeptQueryOptions) (*schema.OrgDept, error) {
	item, err := a.OrgDeptRepo.View(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// View 查询指定数据
func (a *OrgDept) GetAllSubs(ctx context.Context, pid string, params schema.OrgDeptQueryParam, opts ...schema.OrgDeptQueryOptions) (*schema.OrgDeptQueryResult, error) {
	item, err := a.OrgDeptRepo.GetAllSubs(ctx, pid, params, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// Query 查询数据
func (a *OrgDept) GetTree(ctx context.Context, tpid string, params schema.OrgDeptQueryParam, opts ...schema.OrgDeptQueryOptions) (*schema.OrgDeptQueryTreeResult, error) {
	return a.OrgDeptRepo.GetTree(ctx, tpid, params, opts...)
}

// Create 创建数据
func (a *OrgDept) Create(ctx context.Context, item schema.OrgDept) (*schema.OrgDept, error) {

	query_org := a.EntCli.OrgOrgan.Query().Where(orgorgan.IDEQ(item.OrgID))
	organ, err := query_org.First(ctx)
	if err != nil {
		return nil, err
	}

	var parent *ent.OrgDept
	if item.ParentID != nil && *item.ParentID != "" {
		parent, err = a.EntCli.OrgDept.Query().Where(orgdept.IDEQ(*item.ParentID), orgdept.OrgIDEQ(organ.ID)).Only(ctx)
		if err != nil {
			return nil, err
		}
	}

	if parent == nil || item.ParentID == nil || *item.ParentID == "" {
		parent, err = a.EntCli.OrgDept.Query().Where(orgdept.OrgIDEQ(item.OrgID), orgdept.TreeLeftEQ(1), orgdept.TreeLevelEQ(1)).Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				la_tree_id, err := a.OrgDeptRepo.GetLatestTreeID(ctx)
				if err != nil {
					return nil, err
				}
				parent, err = a.EntCli.OrgDept.Create().SetOrgID(organ.ID).SetTreeID(la_tree_id).SetName(*organ.Sname).SetTreeLevel(1).SetTreeLeft(1).SetTreeRight(2).SetSort(0).SetIsShow(false).Save(ctx)
				if err != nil {
					return nil, err
				}
			} else {
				return nil, err
			}
		}
		if parent == nil {
			return nil, errors.New("no parent")
		}
	}

	item.ParentID = ptr.String(parent.ID)
	item.TreeID = parent.TreeID
	item.TreeLeft = parent.TreeRight
	item.TreeRight = ptr.Int64(*parent.TreeRight + 1)
	item.IsLeaf = ptr.Bool(true)
	item.TreeLevel = ptr.Int32(*parent.TreeLevel + 1)

	if parent.TreePath == nil || *parent.TreePath == "" {
		item.TreePath = ptr.String(item.ID)
	} else {
		item.TreePath = ptr.String(strings.Join([]string{*parent.TreePath, parent.ID}, "/"))
	}

	dept_create_input := a.OrgDeptRepo.ToEntCreateOrgDeptInput(&item)

	var res_dept *ent.OrgDept
	err = repo.WithTx(ctx, a.EntCli, func(tx *ent.Tx) error {

		res_dept, err = tx.OrgDept.Create().SetInput(*dept_create_input).Save(ctx)

		if err != nil {
			return err
		}
		update_dept_l := tx.OrgDept.Update()
		update_dept_l = update_dept_l.Where(orgdept.IDNEQ(res_dept.ID))
		update_dept_l = update_dept_l.Where(orgdept.TreeIDEQ(*res_dept.TreeID)).Where(orgdept.TreeLeftGT(*parent.TreeRight))
		_, err = update_dept_l.AddTreeLeft(2).Save(ctx)
		if err != nil {
			return err
		}
		update_dept_r := tx.OrgDept.Update()
		update_dept_r = update_dept_r.Where(orgdept.IDNEQ(res_dept.ID))
		update_dept_r = update_dept_r.Where(orgdept.TreeIDEQ(*res_dept.TreeID)).Where(orgdept.TreeRightGTE(*parent.TreeRight))
		_, err = update_dept_r.AddTreeRight(2).Save(ctx)
		if err != nil {
			return err
		}
		// udpate parent's field is_leaf
		_, err = tx.OrgDept.Update().Where(orgdept.IDEQ(parent.ID)).SetIsLeaf(false).Save(ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	sch_OrgDept := repo.ToSchemaOrgDept(res_dept)

	return sch_OrgDept, nil
}

// Update 更新数据
func (a *OrgDept) Update(ctx context.Context, id string, item schema.OrgDept) (*schema.OrgDept, error) {

	oitem, err := a.OrgDeptRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if oitem == nil {
		return nil, errors.ErrNotFound
	}

	item.ID = oitem.ID

	return a.OrgDeptRepo.Update(ctx, id, item)
}

// Delete 删除数据
func (a *OrgDept) Delete(ctx context.Context, id string) error {
	oldItem, err := a.OrgDeptRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.OrgDeptRepo.Delete(ctx, id)
}

// UpdateActive 更新状态
func (a *OrgDept) UpdateActive(ctx context.Context, id string, active bool) error {
	oldItem, err := a.OrgDeptRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.OrgDeptRepo.UpdateActive(ctx, id, active)
}
