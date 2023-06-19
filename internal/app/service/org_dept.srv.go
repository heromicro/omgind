package service

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/google/wire"
	"github.com/gotidy/ptr"
	"github.com/hibiken/asynq"
	"github.com/jossef/format"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/orgdept"
	"github.com/heromicro/omgind/internal/gen/ent/orgorgan"
	"github.com/heromicro/omgind/internal/scheme/repo"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/mw/asyncq/worker"
	"github.com/heromicro/omgind/pkg/mw/queue"
	"github.com/heromicro/omgind/pkg/types"
)

// OrgDeptSet 注入OrgDept
// var OrgDeptSet = wire.NewSet(wire.Struct(new(OrgDept), "*"))
var OrgDeptSet = wire.NewSet(NewOrgDeptSrv)

// OrgDept 部门管理
type OrgDept struct {
	EntCli *ent.Client

	OrgDeptRepo *repo.OrgDept

	Queue    queue.Queuer
	consumer *worker.Consumer
}

func NewOrgDeptSrv(entCli *ent.Client, deptRepo *repo.OrgDept, q queue.Queuer, c *worker.Consumer) *OrgDept {

	depSrv := &OrgDept{
		EntCli:      entCli,
		OrgDeptRepo: deptRepo,
		Queue:       q,
		consumer:    c,
	}

	// TODO: register tasks
	depSrv.consumer.RegisterHandlers(types.TaskName_REPAIR_DEPT_TREE_PATH, depSrv)

	return depSrv
}

func (s *OrgDept) ProcessTask(ctx context.Context, t *asynq.Task) error {

	id := string(t.Payload())

	log.Println(" -------- ==== id : ", id)

	parent, err := s.EntCli.OrgDept.Query().Where(orgdept.IDEQ(id)).WithChildren(func(sdq *ent.OrgDeptQuery) {
		sdq.Where(orgdept.IsDel(false)).Select(orgdept.FieldID, orgdept.FieldName, orgdept.FieldIsLeaf, orgdept.FieldTreeLeft, orgdept.FieldTreeRight, orgdept.FieldParentID)

	}).First(ctx)

	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	for _, child := range parent.Edges.Children {

		var tree_path string
		if parent.TreePath != nil && *parent.TreePath != "" {
			tree_path = strings.Join([]string{*parent.TreePath, parent.ID}, "/")
		} else {
			tree_path = parent.ID
		}

		var merge_name string
		if parent.MergeName != nil && *parent.MergeName != "" {
			merge_name = strings.Join([]string{*parent.MergeName, *child.Name}, "/")
		} else {
			merge_name = strings.Join([]string{*parent.Name, *child.Name}, "/")
		}

		update_dept := s.EntCli.OrgDept.Update().Where(orgdept.IDEQ(child.ID))

		if tree_path != "" {
			update_dept = update_dept.SetTreePath(tree_path)
		}

		if merge_name != "" {
			update_dept = update_dept.SetMergeName(merge_name)
		}

		d := *child.TreeRight - *child.TreeLeft
		if d > 1 {
			update_dept = update_dept.SetIsLeaf(false)
		} else if d == 1 {
			update_dept = update_dept.SetIsLeaf(true)
		}
		update_dept = update_dept.SetTreeLevel(*parent.TreeLevel + 1)
		_, err = update_dept.Save(ctx)

		if err == nil {
			// ids = append(ids, child.ID)

			if d > 1 {
				jobid := format.String(`{id}-{ml}`, format.Items{"id": child.ID, "ml": time.Now().UnixMilli()})
				job := &queue.Job{
					ID:      jobid,
					Payload: json.RawMessage(child.ID),
					Delay:   200 * time.Millisecond,
				}

				s.Queue.Write(types.TaskName_REPAIR_DEPT_TREE_PATH, types.RepaireTreeQueue, job)
			}

		}

	}

	return nil
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

	query_org := a.EntCli.OrgOrgan.Query().Where(orgorgan.IDEQ(*item.OrgID))
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
		parent, err = a.EntCli.OrgDept.Query().Where(orgdept.OrgIDEQ(*item.OrgID), orgdept.TreeLeftEQ(1), orgdept.TreeLevelEQ(1)).Only(ctx)
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
		item.TreePath = ptr.String(parent.ID)
	} else {
		item.TreePath = ptr.String(strings.Join([]string{*parent.TreePath, parent.ID}, "/"))
	}

	if parent.MergeName != nil && *parent.MergeName != "" {
		item.MergeName = ptr.String(strings.Join([]string{*parent.MergeName, item.Name}, "/"))
	} else {
		item.MergeName = ptr.String(item.Name)
	}

	dept_create_input := repo.ToEntCreateOrgDeptInput(&item)

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

	oitem, err := a.EntCli.OrgDept.Query().Where(orgdept.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	} else if oitem == nil {
		return nil, errors.ErrNotFound
	}

	// do not modify top item
	if oitem.ParentID == nil {
		return nil, errors.New("do not modify top item")
	}

	if item.ParentID == nil || *item.ParentID == "" {
		return nil, errors.New("no pid ")
	}

	if *item.ParentID == id {
		return nil, errors.New("make self parent")
	}

	if item.OrgID != nil && *item.OrgID != "" && *item.OrgID != *oitem.OrgID {
		return nil, errors.New("do not change org_id")
	}
	item.OrgID = nil

	jobid := format.String(`{id}-{ml}`, format.Items{"id": id, "ml": time.Now().UnixMilli()})
	job := &queue.Job{
		ID:      jobid,
		Payload: json.RawMessage(id),
		Delay:   100 * time.Millisecond,
	}

	// check pid changed or not

	query_oparent := a.EntCli.OrgDept.Query()
	query_oparent = query_oparent.Where(orgdept.IDEQ(*oitem.ParentID), orgdept.TreeIDEQ(*oitem.TreeID))
	oparent, err := query_oparent.First(ctx)

	if err != nil {
		return nil, err
	}

	oparent_has_children, err := a.EntCli.OrgDept.Query().Where(orgdept.IsDelEQ(false), orgdept.ParentIDEQ(oparent.ID)).Count(ctx)
	if err != nil {
		return nil, err
	}

	log.Println(" ------- ===== oparent_has_children ", oparent_has_children)

	query_nparent := a.EntCli.OrgDept.Query()
	query_nparent = query_nparent.Where(orgdept.DeletedAtIsNil()).Where(orgdept.IDEQ(*item.ParentID))
	query_nparent = query_nparent.Where(orgdept.And(orgdept.TreeIDEQ(*oitem.TreeID), orgdept.Or(orgdept.TreeLeftLT(*oitem.TreeLeft), orgdept.TreeRightGT(*oitem.TreeRight))))
	nparent, err := query_nparent.First(ctx)

	if err != nil {
		return nil, err
	}

	if *oitem.ParentID == *item.ParentID {
		// pid no change

		item.TreeID = nil
		item.TreeLeft = nil
		item.TreeRight = nil
		item.TreeLevel = ptr.Int32(*nparent.TreeLevel + 1)
		if (*oitem.TreeRight - *oitem.TreeLeft) > 1 {
			item.IsLeaf = ptr.Bool(false)
		} else {
			item.IsLeaf = ptr.Bool(true)
		}

		if nparent.TreePath == nil || *nparent.TreePath == "" {
			item.TreePath = &nparent.ID
		} else {
			item.TreePath = ptr.String(strings.Join([]string{*nparent.TreePath, nparent.ID}, "/"))
		}

		if nparent.MergeName != nil && *nparent.MergeName != "" {
			item.MergeName = ptr.String(strings.Join([]string{*nparent.MergeName, item.Name}, "/"))
		} else {
			item.MergeName = ptr.String(item.Name)
		}

		iteminput := repo.ToEntUpdateOrgDeptInput(&item)

		_, err := a.EntCli.OrgDept.Update().Where(orgdept.IDEQ(id)).SetInput(*iteminput).Save(ctx)
		if err != nil {
			return nil, err
		}

		_ = a.Queue.Write(types.TaskName_REPAIR_DEPT_TREE_PATH, types.RepaireTreeQueue, job)

	} else {

		ids, err := a.EntCli.OrgDept.Query().Where(orgdept.TreeIDEQ(*oitem.TreeID), orgdept.And(orgdept.TreeLeftGTE(*oitem.TreeLeft), orgdept.TreeRightLTE(*oitem.TreeRight))).IDs(ctx)
		if err != nil {
			return nil, err
		}
		d1 := *oitem.TreeRight - *oitem.TreeLeft + 1
		var d2 int64
		if *nparent.TreeRight > *oparent.TreeRight {
			d2 = *nparent.TreeRight - *oitem.TreeRight - 1
		} else {
			d2 = *nparent.TreeRight - *oitem.TreeLeft
		}

		item.TreeID = nil
		item.TreeLeft = nil
		item.TreeRight = nil
		item.TreeLevel = ptr.Int32(*nparent.TreeLevel + 1)
		if (*oitem.TreeRight - *oitem.TreeLeft) > 1 {
			item.IsLeaf = ptr.Bool(false)
		} else {
			item.IsLeaf = ptr.Bool(true)
		}

		if nparent.TreePath != nil && *nparent.TreePath != "" {
			item.TreePath = ptr.String(strings.Join([]string{*nparent.TreePath, nparent.ID}, "/"))
		} else {
			item.TreePath = nil
		}

		if nparent.MergeName != nil && *nparent.MergeName != "" {
			item.MergeName = ptr.String(strings.Join([]string{*nparent.MergeName, item.Name}, "/"))
		} else {
			item.MergeName = ptr.String(item.Name)
		}

		iteminput := repo.ToEntUpdateOrgDeptInput(&item)

		err = repo.WithTx(ctx, a.EntCli, func(tx *ent.Tx) error {

			// repair old parent's left/right
			_, err := tx.OrgDept.Update().Where(orgdept.TreeIDEQ(*oparent.TreeID), orgdept.TreeLeftGT(*oparent.TreeRight), orgdept.TreeLeftLTE(*nparent.TreeLeft)).AddTreeLeft(-d1).Save(ctx)
			if err != nil {
				return nil
			}
			_, err = tx.OrgDept.Update().Where(orgdept.TreeIDEQ(*oparent.TreeID), orgdept.TreeRightGTE(*oparent.TreeRight), orgdept.TreeRightLT(*nparent.TreeRight)).AddTreeRight(-d1).Save(ctx)
			if err != nil {
				return nil
			}

			// repair new parent's left/rightrepair new parent's left/right
			_, err = tx.OrgDept.Update().Where(orgdept.TreeIDEQ(*nparent.TreeID), orgdept.TreeLeftGT(*nparent.TreeRight), orgdept.TreeLeftLT(*oitem.TreeLeft)).AddTreeLeft(d1).Save(ctx)
			if err != nil {
				return err
			}
			_, err = tx.OrgDept.Update().Where(orgdept.TreeIDEQ(*nparent.TreeID), orgdept.TreeRightGTE(*nparent.TreeRight), orgdept.TreeRightLT(*oitem.TreeLeft)).AddTreeRight(d1).Save(ctx)
			if err != nil {
				return err
			}

			// step 1: repair the sub items left/right
			_, err = tx.OrgDept.Update().Where(orgdept.IDIn(ids...)).AddTreeLeft(d2).AddTreeRight(d2).Save(ctx)
			if err != nil {
				return err
			}

			// step 4:  udpate
			_, err = tx.OrgDept.Update().Where(orgdept.IDEQ(id)).SetInput(*iteminput).Save(ctx)
			if err != nil {
				return err
			}

			if oparent_has_children == 1 {
				_, err := tx.OrgDept.Update().Where(orgdept.IDEQ(oparent.ID)).SetIsLeaf(true).Save(ctx)
				if err != nil {
					return err
				}
			}

			_, err = tx.OrgDept.Update().Where(orgdept.IDEQ(nparent.ID)).SetIsLeaf(false).Save(ctx)
			if err != nil {
				return err
			}

			// step 5: trigger update subs's tree_path
			_ = a.Queue.Write(types.TaskName_REPAIR_DEPT_TREE_PATH, types.RepaireTreeQueue, job)

			return nil
		})

		if err != nil {
			return nil, err
		}
	}

	item.ID = oitem.ID

	res_dept, err := a.EntCli.OrgDept.Query().Where(orgdept.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	sch_OrgDept := repo.ToSchemaOrgDept(res_dept)

	return sch_OrgDept, nil
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
