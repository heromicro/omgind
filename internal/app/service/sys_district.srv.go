package service

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/google/wire"
	"github.com/hibiken/asynq"
	"github.com/jossef/format"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/sysdistrict"
	"github.com/heromicro/omgind/internal/schema/repo"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/mw/asyncq/worker"
	"github.com/heromicro/omgind/pkg/mw/queue"
	"github.com/heromicro/omgind/pkg/types"
)

// SysDistrictSet 注入SysDistrict
// var SysDistrictSet = wire.NewSet(wire.Struct(new(SysDistrict), "*"))
var SysDistrictSet = wire.NewSet(New)

// SysDistrict 行政区域
type SysDistrict struct {
	EntCli *ent.Client

	SysDistrictRepo *repo.SysDistrict
	Queue           queue.Queuer
	consumer        *worker.Consumer
}

func New(entCli *ent.Client, districtRepo *repo.SysDistrict, q queue.Queuer, c *worker.Consumer) *SysDistrict {

	districtSrv := &SysDistrict{
		EntCli:          entCli,
		SysDistrictRepo: districtRepo,
		Queue:           q,
		consumer:        c,
	}

	// TODO: register tasks
	districtSrv.consumer.RegisterHandlers(types.TaskName_REPAIR_DISTRICT_TREE_PATH, districtSrv)

	return districtSrv
}

// repair tree_path, merge_name, merge_sname
func (s *SysDistrict) ProcessTask(ctx context.Context, t *asynq.Task) error {

	// log.Println(" --- ---- === = ", t.Type())
	id := string(t.Payload())
	// log.Println(" -- ----- ==== ", id)

	parent, err := s.EntCli.SysDistrict.Query().Where(sysdistrict.IDEQ(id)).WithChildren(func(sdq *ent.SysDistrictQuery) {

		sdq.Where(sysdistrict.IsDel(false)).Select(sysdistrict.FieldID, sysdistrict.FieldMergeName, sysdistrict.FieldMergeSname, sysdistrict.FieldName, sysdistrict.FieldSname, sysdistrict.FieldIsLeaf, sysdistrict.FieldTreeLeft, sysdistrict.FieldTreeRight, sysdistrict.FieldParentID)

	}).First(ctx)

	// log.Println(" -- ----- ==== err ", err)

	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}
	// var ids []string = []string{}

	for _, child := range parent.Edges.Children {

		var tree_path string
		var merge_name string
		var merge_sname string
		if parent.TreePath != nil && *parent.TreePath != "" {
			tree_path = strings.Join([]string{*parent.TreePath, parent.ID}, "/")
		} else {
			tree_path = parent.ID
		}

		if parent.MergeName != nil && *parent.MergeName != "" {
			merge_name = strings.Join([]string{*parent.MergeName, *child.Name}, ",")
		} else {
			merge_name = strings.Join([]string{*parent.Name, *child.Name}, ",")
		}

		if parent.MergeSname != nil && *parent.MergeSname != "" {
			merge_sname = strings.Join([]string{*parent.MergeSname, *child.Sname}, ",")
		} else {
			merge_sname = strings.Join([]string{*parent.Sname, *child.Sname}, ",")
		}

		update_district := s.EntCli.SysDistrict.Update().Where(sysdistrict.IDEQ(child.ID))
		if tree_path != "" {
			update_district = update_district.SetTreePath(tree_path)
		}
		if merge_name != "" {
			update_district = update_district.SetMergeName(merge_name)
		}
		if merge_sname != "" {
			update_district = update_district.SetMergeSname(merge_sname)
		}

		d := *child.TreeRight - *child.TreeLeft
		if d > 1 {
			update_district = update_district.SetIsLeaf(false)
		} else if d == 1 {
			update_district = update_district.SetIsLeaf(true)
		}

		update_district = update_district.SetTreeLevel(*parent.TreeLevel + 1)

		_, err := update_district.Save(ctx)

		if err != nil {
			log.Println(" --------- ===== ------ ", err)
		} else {
			// ids = append(ids, child.ID)

			if d > 1 {
				jobid := format.String(`{id}-{ml}`, format.Items{"id": child.ID, "ml": time.Now().UnixMilli()})
				job := &queue.Job{
					ID:      jobid,
					Payload: json.RawMessage(child.ID),
					Delay:   200 * time.Millisecond,
				}

				s.Queue.Write(types.TaskName_REPAIR_DISTRICT_TREE_PATH, types.DistrictQueue, job)
			}
		}
	}

	return nil
}

// Query 查询数据
func (a *SysDistrict) Query(ctx context.Context, params schema.SysDistrictQueryParam, opts ...schema.SysDistrictQueryOptions) (*schema.SysDistrictQueryResult, error) {
	return a.SysDistrictRepo.Query(ctx, params, opts...)
}

// Query 查询数据
func (a *SysDistrict) GetAllSubDistricts(ctx context.Context, pid string, params schema.SysDistrictQueryParam, opts ...schema.SysDistrictQueryOptions) (*schema.SysDistrictQueryResult, error) {
	return a.SysDistrictRepo.GetAllSubDistricts(ctx, pid, params, opts...)
}

// Query 查询数据
func (a *SysDistrict) GetTree(ctx context.Context, tpid string, params schema.SysDistrictQueryParam, opts ...schema.SysDistrictQueryOptions) (*schema.SysDistrictQueryTreeResult, error) {
	return a.SysDistrictRepo.GetTree(ctx, tpid, params, opts...)
}

// Get 查询指定数据
func (a *SysDistrict) Get(ctx context.Context, id string, opts ...schema.SysDistrictQueryOptions) (*schema.SysDistrict, error) {
	item, err := a.SysDistrictRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// View 查询指定数据
func (a *SysDistrict) View(ctx context.Context, id string, opts ...schema.SysDistrictQueryOptions) (*schema.SysDistrict, error) {

	item, err := a.SysDistrictRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	// item.Parent = a.SysDistrictRepo.ToSc
	// item.Children =

	return item, nil
}

// Create 创建数据
func (a *SysDistrict) Create(ctx context.Context, item schema.SysDistrict) (*schema.SysDistrict, error) {
	// TODO: check?

	sch_sysdistrict, err := a.SysDistrictRepo.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return sch_sysdistrict, nil
}

// Update 更新数据
func (a *SysDistrict) Update(ctx context.Context, id string, item schema.SysDistrict) (*schema.SysDistrict, error) {

	oitem, err := a.SysDistrictRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if oitem == nil {
		return nil, errors.ErrNotFound
	}

	item.ID = oitem.ID

	return a.SysDistrictRepo.Update(ctx, id, item)
}

// Delete 删除数据
func (a *SysDistrict) Delete(ctx context.Context, id string) error {
	oldItem, err := a.SysDistrictRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.SysDistrictRepo.Delete(ctx, id)
}

// UpdateActive 更新状态
func (a *SysDistrict) UpdateActive(ctx context.Context, id string, active bool) error {
	oldItem, err := a.SysDistrictRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.SysDistrictRepo.UpdateActive(ctx, id, active)
}
