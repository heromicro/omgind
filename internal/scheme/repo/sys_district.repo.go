package repo

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/google/wire"
	"github.com/gotidy/ptr"
	"github.com/jossef/format"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/entscheme"
	"github.com/heromicro/omgind/internal/gen/entscheme/sysdistrict"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"
	"github.com/heromicro/omgind/pkg/mw/queue"
	"github.com/heromicro/omgind/pkg/types"
)

// SysDistrictSet 注入SysDistrict
var SysDistrictSet = wire.NewSet(wire.Struct(new(SysDistrict), "*"))

// SysDistrict 行政区域存储
type SysDistrict struct {
	EntCli *entscheme.Client
	Queue  queue.Queuer
}

// ToSchemaSysDistrict 转换为
func ToSchemaSysDistrict(et *entscheme.SysDistrict) *schema.SysDistrict {
	item := new(schema.SysDistrict)
	structure.Copy(et, item)
	return item
}

func ToSchemaSysDistricts(ets entscheme.SysDistricts) []*schema.SysDistrict {
	list := make([]*schema.SysDistrict, len(ets))
	for i, item := range ets {
		list[i] = ToSchemaSysDistrict(item)
	}
	return list
}

func ToEntCreateSysDistrictInput(sch *schema.SysDistrict) *entscheme.CreateSysDistrictInput {
	createinput := new(entscheme.CreateSysDistrictInput)
	structure.Copy(sch, &createinput)

	return createinput
}

func ToEntUpdateSysDistrictInput(sch *schema.SysDistrict) *entscheme.UpdateSysDistrictInput {
	updateinput := new(entscheme.UpdateSysDistrictInput)
	structure.Copy(sch, &updateinput)

	return updateinput
}

func (a *SysDistrict) getQueryOption(opts ...schema.SysDistrictQueryOptions) schema.SysDistrictQueryOptions {
	var opt schema.SysDistrictQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *SysDistrict) Query(ctx context.Context, params schema.SysDistrictQueryParam, opts ...schema.SysDistrictQueryOptions) (*schema.SysDistrictQueryResult, error) {
	// opt := a.getQueryOption(opts...)

	query := a.EntCli.SysDistrict.Query()

	query = query.Where(sysdistrict.DeletedAtIsNil())
	// TODO: 查询条件

	if v := params.ParentID; v != nil {
		if *v == "" {
			query = query.Where(sysdistrict.Or(sysdistrict.ParentIDEQ(*v), sysdistrict.ParentIDIsNil()))
		} else {
			query = query.Where(sysdistrict.ParentIDEQ(*v))
		}
	}

	if v := params.Name; v != "" {
		query = query.Where(sysdistrict.Or(sysdistrict.NameEQ(v), sysdistrict.NameEnEQ(v)))
	}

	if v := params.IsMain; v != nil {
		query = query.Where(sysdistrict.IsMainEQ(*v))
	}

	if v := params.IsReal; v != nil {
		query = query.Where(sysdistrict.IsRealEQ(*v))
	}

	if v := params.IsHot; v != nil {
		query = query.Where(sysdistrict.IsHotEQ(*v))
	}

	if v := params.IsDirect; v != nil {
		query = query.Where(sysdistrict.IsDirectEQ(*v))
	}

	if v := params.IsLeaf; v != nil {
		query = query.Where(sysdistrict.IsLeafEQ(*v))
	}

	if v := params.IsActive; v != nil {
		query = query.Where(sysdistrict.IsActiveEQ(*v))
	}

	if v := params.TreeLevel; v != nil {
		query = query.Where(sysdistrict.TreeLevelEQ(*v))
	}
	if v := params.TreeLeft; v != nil {
		query = query.Where(sysdistrict.TreeLeftGTE(*v))
	}

	if v := params.TreeLeft_St; v != nil {
		query = query.Where(sysdistrict.TreeLeftGTE(*v))
	}
	if v := params.TreeLeft_Ed; v != nil {
		query = query.Where(sysdistrict.TreeLeftLTE(*v))
	}

	if v := params.TreeRight; v != nil {
		query = query.Where(sysdistrict.TreeRightLTE(*v))
	}

	if v := params.TreeRight_St; v != nil {
		query = query.Where(sysdistrict.TreeRightGTE(*v))
	}
	if v := params.TreeRight_Ed; v != nil {
		query = query.Where(sysdistrict.TreeRightLTE(*v))
	}

	// 父级area_code
	if v := params.ParentAreaCode; v != nil {
		query = query.Where(sysdistrict.HasParentWith(sysdistrict.AreaCodeEQ(*v)))
	}
	// 父级pid
	if v := params.ParentParentID; v != nil {
		query = query.Where(sysdistrict.HasParentWith(sysdistrict.ParentIDEQ(*v)))
	}

	count, err := query.Count(ctx)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.SysDistrictQueryResult{PageResult: pr}, nil
	}

	has_order := false
	if v := params.CreatedAt_Order; v != "" {
		query = query.Order(sysdistrict.ByCreatedAt(OrderDirection(v)))
		has_order = true
	}

	if v := params.Name_Order; v != "" {
		query = query.Order(sysdistrict.ByName(OrderDirection(v)))
		has_order = true
	}

	if v := params.TreeID_Order; v != "" {
		query = query.Order(sysdistrict.ByTreeID(OrderDirection(v)))
		has_order = true
	}

	if v := params.TreeLevel_Order; v != "" {
		query = query.Order(sysdistrict.ByTreeLevel(OrderDirection(v)))
		has_order = true
	}

	if v := params.TreeLeft_Order; v != "" {
		query = query.Order(sysdistrict.ByTreeLeft(OrderDirection(v)))
		has_order = true
	}

	if !has_order {
		query = query.Order(sysdistrict.ByTreeID(OrderDirection("asc")))
		query = query.Order(sysdistrict.BySort(OrderDirection("asc")))
	}

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()

	// log.Println(" ----- === ==== ==== params.Offset() ", params.Offset())
	// log.Println(" ------- ==== === count 22222 ", count)

	if params.Offset() > count {
		return &schema.SysDistrictQueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)

	// log.Println(" ----- 00 = ==== ==== list ", list)

	if err1 != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.SysDistrictQueryResult{
		PageResult: pr,
		Data:       ToSchemaSysDistricts(list),
	}

	return qr, nil
}

func (a *SysDistrict) GetAllSubs(ctx context.Context, pid string, params schema.SysDistrictQueryParam, opts ...schema.SysDistrictQueryOptions) (*schema.SysDistrictQueryResult, error) {

	// opt := a.getQueryOption(opts...)

	query := a.EntCli.SysDistrict.Query()

	query = query.Where(sysdistrict.DeletedAtIsNil(), sysdistrict.IsDelEQ(false))
	// TODO: 查询条件

	if pid == "-" {
		query = query.Where(sysdistrict.Or(sysdistrict.ParentIDEQ(pid), sysdistrict.ParentIDIsNil()))
	} else {
		query = query.Where(sysdistrict.ParentIDEQ(pid))
	}

	if v := params.Name; v != "" {
		query = query.Where(sysdistrict.NameEQ(v))
	}

	if v := params.IsMain; v != nil {
		query = query.Where(sysdistrict.IsMainEQ(*v))
	}

	if v := params.IsReal; v != nil {
		query = query.Where(sysdistrict.IsRealEQ(*v))
	}

	if v := params.IsHot; v != nil {
		query = query.Where(sysdistrict.IsHotEQ(*v))
	}

	if v := params.IsDirect; v != nil {
		query = query.Where(sysdistrict.IsDirectEQ(*v))
	}

	if v := params.IsLeaf; v != nil {
		query = query.Where(sysdistrict.IsLeafEQ(*v))
	}

	if v := params.IsActive; v != nil {
		query = query.Where(sysdistrict.IsActiveEQ(*v))
	}

	if v := params.TreeLevel; v != nil {
		query = query.Where(sysdistrict.TreeLevelEQ(*v))
	}
	if v := params.TreeLeft; v != nil {
		query = query.Where(sysdistrict.TreeLeftGTE(*v))
	}

	if v := params.TreeLeft_St; v != nil {
		query = query.Where(sysdistrict.TreeLeftGTE(*v))
	}
	if v := params.TreeLeft_Ed; v != nil {
		query = query.Where(sysdistrict.TreeLeftLTE(*v))
	}

	if v := params.TreeRight; v != nil {
		query = query.Where(sysdistrict.TreeRightLTE(*v))
	}

	if v := params.TreeRight_St; v != nil {
		query = query.Where(sysdistrict.TreeRightGTE(*v))
	}
	if v := params.TreeRight_Ed; v != nil {
		query = query.Where(sysdistrict.TreeRightLTE(*v))
	}

	// 父级area_code
	if v := params.ParentAreaCode; v != nil {
		query = query.Where(sysdistrict.HasParentWith(sysdistrict.AreaCodeEQ(*v)))
	}
	// 父级pid
	if v := params.ParentParentID; v != nil {
		query = query.Where(sysdistrict.HasParentWith(sysdistrict.ParentIDEQ(*v)))
	}

	count, err := query.Count(ctx)
	// log.Println(" ------- == === === ===== count 11111 ", count)
	// log.Println(" ------- ===== ==== ====  err ", err)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.SysDistrictQueryResult{PageResult: pr}, nil
	}

	has_order := false
	if v := params.CreatedAt_Order; v != "" {
		query = query.Order(sysdistrict.ByCreatedAt(OrderDirection(v)))
		has_order = true
	}

	if v := params.Name_Order; v != "" {
		query = query.Order(sysdistrict.ByName(OrderDirection(v)))
		has_order = true
	}

	if v := params.TreeID_Order; v != "" {
		query = query.Order(sysdistrict.ByTreeID(OrderDirection(v)))
		has_order = true
	}

	if v := params.TreeLevel_Order; v != "" {
		query = query.Order(sysdistrict.ByTreeLevel(OrderDirection(v)))
		has_order = true
	}

	if v := params.TreeLeft_Order; v != "" {
		query = query.Order(sysdistrict.ByTreeLeft(OrderDirection(v)))
		has_order = true
	}

	if !has_order {
		query = query.Order(sysdistrict.ByTreeID(OrderDirection("asc")))
		query = query.Order(sysdistrict.ByTreeLevel(OrderDirection("asc")))
		query = query.Order(sysdistrict.ByTreeLeft(OrderDirection("asc")))
		query = query.Order(sysdistrict.BySort(OrderDirection("asc")))
	}

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()

	// log.Println(" ------- ===== params.Offset() ", params.Offset())
	// log.Println(" ---- === ==== count 22222 ", count)

	if params.Offset() > count {
		return &schema.SysDistrictQueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	//
	squery := query.Select(sysdistrict.FieldID, sysdistrict.FieldName, sysdistrict.FieldSname, sysdistrict.FieldParentID, sysdistrict.FieldAreaCode, sysdistrict.FieldIsLeaf)

	list, err1 := squery.All(ctx)

	// log.Println(" ------- ===== ===== === list ", list)

	if err1 != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.SysDistrictQueryResult{
		PageResult: pr,
		Data:       ToSchemaSysDistricts(list),
	}

	return qr, nil
}

func (a *SysDistrict) GetTree(ctx context.Context, tpid string, params schema.SysDistrictQueryParam, opts ...schema.SysDistrictQueryOptions) (*schema.SysDistrictQueryTreeResult, error) {
	// opt := a.getQueryOption(opts...)

	parent_query := a.EntCli.SysDistrict.Query().Where(sysdistrict.DeletedAtIsNil())

	parent_query = parent_query.Where(sysdistrict.IDEQ(tpid))

	parent, err := parent_query.Select(sysdistrict.FieldID, sysdistrict.FieldTreeID).First(ctx)
	if err != nil {
		return nil, err
	}

	tree_query := a.EntCli.SysDistrict.Query().Where(sysdistrict.IsDelEQ(false), sysdistrict.TreeIDEQ(*parent.TreeID)).Where(sysdistrict.IDNEQ(parent.ID))

	tree_query = tree_query.Order(sysdistrict.ByTreeLevel(OrderDirection("asc")))
	tree_query = tree_query.Order(sysdistrict.ByTreeLeft(OrderDirection("asc")))
	tree_query = tree_query.Order(sysdistrict.BySort(OrderDirection("asc")))

	select_tree := tree_query.Select(sysdistrict.FieldID, sysdistrict.FieldName, sysdistrict.FieldSname, sysdistrict.FieldParentID, sysdistrict.FieldAreaCode, sysdistrict.FieldIsLeaf, sysdistrict.FieldTreeID, sysdistrict.FieldTreeLeft, sysdistrict.FieldTreeRight)

	subs, err := select_tree.All(ctx)
	if err != nil {
		return nil, err
	}

	top_query := a.EntCli.SysDistrict.Query().Where(sysdistrict.IsDel(false), sysdistrict.TreeLevel(1), sysdistrict.ParentIDIsNil())

	top, err := top_query.All(ctx)
	if err != nil {
		return nil, err
	}

	sch_sub := ToSchemaSysDistricts(subs)
	sch_top := ToSchemaSysDistricts(top)
	data := &schema.SysDistrictQueryTreeResult{
		Top:  sch_top,
		Subs: sch_sub,
	}

	return data, nil
}

// Get 查询指定数据
func (a *SysDistrict) Get(ctx context.Context, id string, opts ...schema.SysDistrictQueryOptions) (*schema.SysDistrict, error) {
	query_district := a.EntCli.SysDistrict.Query()

	query_district = query_district.WithParent()

	r_sysdistrict, err := query_district.Where(sysdistrict.IDEQ(id)).Only(ctx)
	if err != nil {
		if entscheme.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	r_schema := ToSchemaSysDistrict(r_sysdistrict)
	if r_sysdistrict.Edges.Parent != nil {
		r_schema.Parent = ToSchemaSysDistrict(r_sysdistrict.Edges.Parent)
	}

	return r_schema, nil
}

// Create 创建数据
func (a *SysDistrict) Create(ctx context.Context, item schema.SysDistrict) (*schema.SysDistrict, error) {

	var res_sysdistrict *entscheme.SysDistrict
	var err error

	// check pid
	if item.ParentID == nil || *item.ParentID == "" {
		// no pid, top level
		var opt schema.SysDistrictQueryOptions
		opt.OrderFields = append(opt.OrderFields, schema.NewOrderField(sysdistrict.FieldTreeID, schema.OrderByDESC))

		most, err := a.GetLatestTreeID(ctx)
		if err != nil {
			return nil, err
		}

		item.TreeID = ptr.Int64(most)

		item.TreeLeft = ptr.Int64(1)
		item.TreeRight = ptr.Int64(2)
		item.IsLeaf = ptr.Bool(true)
		item.TreeLevel = ptr.Int32(1)
		item.TreePath = nil

		if item.ParentID != nil && *item.ParentID == "" {
			item.ParentID = nil
		}

		if item.MergeName == nil || *item.MergeName == "" {
			item.MergeName = &item.Name
		}

		if item.MergeSname == nil || *item.MergeSname == "" {
			if item.Sname != nil && *item.Sname != "" {
				item.MergeSname = item.Sname
			}
		}

		iteminput := ToEntCreateSysDistrictInput(&item)

		res_sysdistrict, err = a.EntCli.SysDistrict.Create().SetInput(*iteminput).Save(ctx)
		if err != nil {
			return nil, err
		}

	} else {

		parent, err := a.EntCli.SysDistrict.Query().Where(sysdistrict.IDEQ(*item.ParentID)).First(ctx)

		if err != nil {
			return nil, err
		}
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

		if item.MergeName == nil || *item.MergeName == "" {

			if parent.MergeName != nil && *parent.MergeName != "" {
				item.MergeName = ptr.String(strings.Join([]string{*parent.MergeName, item.Name}, ","))
			} else {
				item.MergeName = ptr.String(item.Name)
			}
		}

		if item.MergeSname == nil || *item.MergeSname == "" {

			if item.Sname != nil && *item.Sname != "" {
				if parent.MergeSname != nil && *parent.MergeSname != "" {
					item.MergeSname = ptr.String(strings.Join([]string{*parent.MergeSname, *item.Sname}, ","))
				} else {
					item.MergeSname = item.Sname
				}
			}
		}

		iteminput := ToEntCreateSysDistrictInput(&item)

		err = WithTx(ctx, a.EntCli, func(tx *entscheme.Tx) error {

			res_sysdistrict, err = tx.SysDistrict.Create().SetInput(*iteminput).Save(ctx)

			if err != nil {
				return err
			}
			update_district_l := tx.SysDistrict.Update()
			update_district_l = update_district_l.Where(sysdistrict.IDNEQ(res_sysdistrict.ID))
			update_district_l = update_district_l.Where(sysdistrict.TreeIDEQ(*parent.TreeID))
			update_district_l = update_district_l.Where(sysdistrict.TreeLeftGT(*parent.TreeRight))
			_, err := update_district_l.AddTreeLeft(2).Save(ctx)
			if err != nil {
				return err
			}
			update_district_r := tx.SysDistrict.Update()
			update_district_r = update_district_r.Where(sysdistrict.IDNEQ(res_sysdistrict.ID))
			update_district_r = update_district_r.Where(sysdistrict.TreeIDEQ(*parent.TreeID))
			update_district_r = update_district_r.Where(sysdistrict.TreeRightGTE(*parent.TreeRight))
			_, err = update_district_r.AddTreeRight(2).Save(ctx)
			if err != nil {
				return nil
			}
			// udpate parent's field is_leaf
			_, err = tx.SysDistrict.Update().Where(sysdistrict.IDEQ(parent.ID)).SetIsLeaf(false).Save(ctx)

			if err != nil {
				return err
			}
			return nil
		})
	}

	if err != nil {
		return nil, err
	}
	sch_sysdistrict := ToSchemaSysDistrict(res_sysdistrict)
	return sch_sysdistrict, nil
}

// Update 更新数据
func (a *SysDistrict) Update(ctx context.Context, id string, item schema.SysDistrict) (*schema.SysDistrict, error) {

	oitem, err := a.EntCli.SysDistrict.Query().Where(sysdistrict.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	// iteminput := ToEntUpdateSysDistrictInput(&item)

	jobid := format.String(`{id}-{ml}`, format.Items{"id": id, "ml": time.Now().UnixMilli()})
	job := &queue.Job{
		ID:      jobid,
		Payload: json.RawMessage(id),
		Delay:   100 * time.Millisecond,
	}

	log.Println(" ------- ===== ---- oitem.ParentID ", oitem.ParentID)
	log.Println(" ------- ===== ---- nitem.ParentID  ", item.ParentID)
	if *item.ParentID == "" {
		log.Println(" ------- ===== ---- nitem.ParentID  [", *item.ParentID, "]")
	}

	if *item.ParentID == id {
		return nil, errors.New("make self parent")
	}

	// check pid changed or not
	if oitem.ParentID == nil {
		// old data.pid is nil
		if item.ParentID == nil || *item.ParentID == "" {
			// pid no change

			item.TreeID = nil
			item.TreeLeft = nil
			item.TreeRight = nil
			item.TreeLevel = nil
			item.TreePath = nil

			if item.ParentID != nil && *item.ParentID == "" {
				item.ParentID = nil
			}

			if item.MergeName == nil || *item.MergeName == "" {
				item.MergeName = &item.Name
			}

			if item.MergeSname == nil || *item.MergeSname == "" {
				item.MergeSname = item.Sname
			}

			iteminput := ToEntUpdateSysDistrictInput(&item)
			log.Println(" ------ ====== name ", iteminput.Name)

			_, err := a.EntCli.SysDistrict.Update().Where(sysdistrict.IDEQ(id)).SetInput(*iteminput).Save(ctx)
			log.Println(" ------ ====== err ", err)
			if err != nil {
				return nil, err
			}

		} else {
			// pid changed
			// bring top item to sub

			query_nparent := a.EntCli.SysDistrict.Query()
			query_nparent = query_nparent.Where(sysdistrict.DeletedAtIsNil()).Where(sysdistrict.IDEQ(*item.ParentID))
			query_nparent = query_nparent.Where(sysdistrict.Or(sysdistrict.TreeIDNEQ(*oitem.TreeID), sysdistrict.And(sysdistrict.TreeIDEQ(*oitem.TreeID), sysdistrict.Or(sysdistrict.TreeLeftLT(*oitem.TreeLeft), sysdistrict.TreeRightGT(*oitem.TreeRight)))))

			nparent, err := query_nparent.First(ctx)
			if err != nil {
				return nil, err
			}

			item.TreeID = ptr.Int64(*nparent.TreeID)
			item.TreeLeft = nil
			item.TreeRight = nil
			item.TreeLevel = ptr.Int32(*nparent.TreeLevel + 1)

			if nparent.TreePath == nil || *nparent.TreePath == "" {
				item.TreePath = ptr.String(nparent.ID)
			} else {
				item.TreePath = ptr.String(strings.Join([]string{*nparent.TreePath, nparent.ID}, "/"))
			}

			if nparent.MergeName != nil && *nparent.MergeName != "" {
				if nparent.MergeName != nil && *nparent.MergeName != "" {
					item.MergeName = ptr.String(strings.Join([]string{*nparent.MergeName, item.Name}, ","))
				} else {
					item.MergeName = ptr.String(item.Name)
				}
			}

			if nparent.MergeSname != nil && *nparent.MergeSname != "" {
				item.MergeSname = ptr.String(strings.Join([]string{*nparent.MergeSname, *item.Sname}, ","))
			} else {
				item.MergeSname = item.Sname
			}

			iteminput := ToEntUpdateSysDistrictInput(&item)

			err = WithTx(ctx, a.EntCli, func(tx *entscheme.Tx) error {
				// step 1: update old item's and it's sub's tree_left and tree_right
				count, err := tx.SysDistrict.Update().Where(sysdistrict.TreeID(*oitem.TreeID)).AddTreeLeft(*nparent.TreeRight - 1).AddTreeRight(*nparent.TreeRight - 1).Save(ctx)
				if err != nil {
					return err
				}
				if count != int(*oitem.TreeRight)/2 {
					return errors.New("the count of update old top item and it's subs tree_left/tree_right mismatch the real  ")
				}
				// step 3: update new parent's and it's upper's tree_left
				_, err = tx.SysDistrict.Update().Where(sysdistrict.TreeID(*nparent.TreeID)).Where(sysdistrict.TreeLeftGT(*nparent.TreeRight)).AddTreeLeft(*oitem.TreeRight).Save(ctx)
				if err != nil {
					return err
				}
				// step 4: update new parent's and it's upper's tree_right
				_, err = tx.SysDistrict.Update().Where(sysdistrict.TreeIDEQ(*nparent.TreeID)).Where(sysdistrict.TreeRightGTE(*nparent.TreeRight)).AddTreeRight(*oitem.TreeRight).Save(ctx)
				if err != nil {
					return err
				}

				// step 6: update old item's and it's sub's tree_id
				count, err = tx.SysDistrict.Update().Where(sysdistrict.TreeIDEQ(*oitem.TreeID)).SetTreeID(*nparent.TreeID).Save(ctx)
				if err != nil {
					return err
				}

				_, err = tx.SysDistrict.Update().Where(sysdistrict.IDEQ(nparent.ID)).SetIsLeaf(false).Save(ctx)
				if err != nil {
					return err
				}

				if count != int(*oitem.TreeRight)/2 {
					return errors.New("the count of update old item's and it's subs mismatched.")
				}

				// step 7:
				_, err = tx.SysDistrict.Update().Where(sysdistrict.IDEQ(id)).SetInput(*iteminput).Save(ctx)
				if err != nil {
					return err
				}

				// step 8: update item's and it's sub's tree_path, merge_name, merge_sname
				// TODO: fix tree_path of id's and id's subs
				// TODO: maybe can trigger celery to do this
				job.ID = format.String(`{id}-{ml}`, format.Items{"id": nparent.ID, "ml": time.Now().UnixMilli()})
				job.Payload = json.RawMessage(nparent.ID)
				err = a.Queue.Write(types.TaskName_REPAIR_DISTRICT_TREE_PATH, types.RepaireTreeQueue, job)

				log.Println(" -- -- --- 0 -- --- -- 0000 enqueue err : ", err)

				return nil
			})

			if err != nil {
				return nil, err
			}
		}
	} else {
		// old data.pid is not nil

		query_oparent := a.EntCli.SysDistrict.Query()
		query_oparent = query_oparent.Where(sysdistrict.IDEQ(*oitem.ParentID))
		oparent, err := query_oparent.First(ctx)

		if err != nil {
			return nil, err
		}
		oparent_has_children, err := a.EntCli.SysDistrict.Query().Where(sysdistrict.IsDelEQ(false), sysdistrict.ParentIDEQ(oparent.ID)).Count(ctx)
		if err != nil {
			return nil, err
		}

		log.Println(" -------- ----- oparent.ID- ", oparent.ID)
		// new item.id is empty
		if item.ParentID == nil || *item.ParentID == "" {
			// bring  to top level,
			//

			most, err := a.GetLatestTreeID(ctx)
			if err != nil {
				return nil, err
			}

			log.Println(" -------- ----- most.tree_id- ", most)
			item.TreeID = ptr.Int64(most)
			item.TreeLeft = nil
			item.TreeRight = nil
			item.TreeLevel = ptr.Int32(1)
			item.IsLeaf = nil
			item.TreePath = nil

			if item.ParentID != nil && *item.ParentID == "" {
				item.ParentID = nil
			}

			item.MergeName = &item.Name
			item.MergeSname = item.Sname

			iteminput := ToEntUpdateSysDistrictInput(&item)
			iteminput.ClearTreePath = true
			iteminput.ClearParentID = true

			distance := *oitem.TreeRight - *oitem.TreeLeft + 1

			log.Println(" ---- ==== yyyyyyyy ")

			err = WithTx(ctx, a.EntCli, func(tx *entscheme.Tx) error {
				// step:1 update tree_id with most newest,

				count1, err := tx.SysDistrict.Update().Where(sysdistrict.TreeIDEQ(*oitem.TreeID), sysdistrict.TreeLeftGTE(*oitem.TreeLeft), sysdistrict.TreeRightLTE(*oitem.TreeRight)).SetTreeID(most).Save(ctx)
				if err != nil {
					return err
				}
				if count1 <= 0 {
					return errors.New("failed to update tree_id")
				}

				// step 2: new parent is oitem
				count2, err := tx.SysDistrict.Update().Where(sysdistrict.TreeIDEQ(most), sysdistrict.TreeLeftGTE(*oitem.TreeLeft), sysdistrict.TreeRightLTE(*oitem.TreeRight)).AddTreeLeft(-*oitem.TreeLeft + 1).AddTreeRight(-*oitem.TreeLeft + 1).Save(ctx)
				if err != nil {
					return err
				}
				if count1 != count2 {
					return errors.New("the result of update tree_id and tree_left/tree_right mismatched.")
				}

				// step 3: oparent minus tree_left
				count3, err := tx.SysDistrict.Update().Where(sysdistrict.TreeID(*oparent.TreeID), sysdistrict.TreeLeftGTE(*oitem.TreeRight)).AddTreeLeft(-distance).Save(ctx)
				if err != nil {
					return err
				}
				log.Println(" --- --- === === ", count3)

				// step 4: oparent minus tree_right
				count4, err := tx.SysDistrict.Update().Where(sysdistrict.TreeID(*oparent.TreeID), sysdistrict.TreeRightGT(*oitem.TreeRight)).AddTreeRight(-distance).Save(ctx)
				if err != nil {
					return err
				}
				log.Println(" --- --- === === ", count4)
				if oparent_has_children == 1 {
					_, err := tx.SysDistrict.Update().Where(sysdistrict.IDEQ(oparent.ID)).SetIsLeaf(true).Save(ctx)
					if err != nil {
						return err
					}
				}
				// step 5: update other fileds
				_, err = tx.SysDistrict.Update().Where(sysdistrict.IDEQ(id)).SetInput(*iteminput).Save(ctx)
				if err != nil {
					return err
				}

				// step 6: trigger udpate tree_path and merge_name, merge_sname,

				err = a.Queue.Write(types.TaskName_REPAIR_DISTRICT_TREE_PATH, types.RepaireTreeQueue, job)

				log.Println(" -- -- --- 0 -- --- -- 1111 enqueue err : ", err)

				return nil
			})

			if err != nil {
				return nil, err
			}

		} else {

			query_nparent := a.EntCli.SysDistrict.Query()
			query_nparent = query_nparent.Where(sysdistrict.DeletedAtIsNil()).Where(sysdistrict.IDEQ(*item.ParentID))
			query_nparent = query_nparent.Where(sysdistrict.Or(sysdistrict.TreeIDNEQ(*oitem.TreeID), sysdistrict.And(sysdistrict.TreeIDEQ(*oitem.TreeID), sysdistrict.Or(sysdistrict.TreeLeftLT(*oitem.TreeLeft), sysdistrict.TreeRightGT(*oitem.TreeRight)))))

			nparent, err := query_nparent.First(ctx)

			if err != nil {
				return nil, err
			}
			log.Println(" -------- - nparent.ID----- ", nparent.ID)

			if *oitem.ParentID == *item.ParentID {
				// new data.pid no changed

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
					item.MergeName = ptr.String(strings.Join([]string{*nparent.MergeName, item.Name}, ","))
				} else {
					item.MergeName = ptr.String(item.Name)
				}

				if nparent.MergeSname != nil && *nparent.MergeSname != "" {
					item.MergeSname = ptr.String(strings.Join([]string{*nparent.MergeSname, *item.Sname}, ","))
				} else {
					item.MergeSname = item.Sname

				}

				iteminput := ToEntUpdateSysDistrictInput(&item)

				_, err := a.EntCli.SysDistrict.Update().Where(sysdistrict.IDEQ(id)).SetInput(*iteminput).Save(ctx)
				if err != nil {
					return nil, err
				}

			} else {
				// new data.pid changed

				// query_nparent := a.EntCli.SysDistrict.Query()
				// query_nparent = query_nparent.Where(sysdistrict.DeletedAtIsNil()).Where(sysdistrict.IDEQ(*item.ParentID))
				// query_nparent = query_nparent.Where(sysdistrict.Or(sysdistrict.TreeIDNEQ(*oitem.TreeID), sysdistrict.And(sysdistrict.TreeIDEQ(*oitem.TreeID), sysdistrict.Or(sysdistrict.TreeLeftLT(*oitem.TreeLeft), sysdistrict.TreeRightGT(*oitem.TreeRight)))))
				// nparent, err := query_nparent.First(ctx)
				// if err != nil {
				// 	return nil, err
				// }

				item.TreeID = nparent.TreeID
				item.TreeLeft = nil
				item.TreeRight = nil
				item.TreeLevel = ptr.Int32(*nparent.TreeLevel + 1)
				if nparent.TreePath != nil && *nparent.TreePath != "" {
					item.TreePath = ptr.String(strings.Join([]string{*nparent.TreePath, nparent.ID}, "/"))
				} else {
					item.TreePath = ptr.String(*&nparent.ID)
				}

				if nparent.MergeName != nil && *nparent.MergeName != "" {
					item.MergeName = ptr.String(strings.Join([]string{*nparent.MergeName, item.Name}, ","))
				} else {
					item.MergeName = ptr.String(item.Name)
				}

				if item.Sname != nil && *item.Sname != "" {
					if nparent.MergeSname != nil && *nparent.MergeSname != "" {
						item.MergeSname = ptr.String(strings.Join([]string{*nparent.MergeSname, *item.Sname}, ","))
					} else {
						item.MergeSname = item.Sname
					}
				}

				iteminput := ToEntUpdateSysDistrictInput(&item)
				// old parent need to minus d1, new parent need to add d1
				d1 := *oitem.TreeRight - *oitem.TreeLeft + 1
				// oitem and subs need to add d2
				d2 := *nparent.TreeRight - *oitem.TreeLeft

				log.Println(" ------ ==---- == *nparent.TreeID ", *nparent.TreeID)
				log.Println(" ------ ==---- ==  *oparent.TreeID ", *oparent.TreeID)
				log.Println(" ------ ==---- ==  *oparent.TreeID == *nparent.TreeID ", *oparent.TreeID == *nparent.TreeID)

				if *nparent.TreeID == *oparent.TreeID { // same tree move
					//

					ids, err := a.EntCli.SysDistrict.Query().Where(sysdistrict.TreeIDEQ(*oitem.TreeID), sysdistrict.And(sysdistrict.TreeLeftGTE(*oitem.TreeLeft), sysdistrict.TreeRightLTE(*oitem.TreeRight))).IDs(ctx)
					if err != nil {
						return nil, err
					}

					var d2 int64
					if *nparent.TreeRight > *oparent.TreeRight {
						d2 = *nparent.TreeRight - *oitem.TreeRight - 1
					} else {
						d2 = *nparent.TreeRight - *oitem.TreeLeft
					}

					err = WithTx(ctx, a.EntCli, func(tx *entscheme.Tx) error {

						// step 2: repair old parent's left/right
						_, err := tx.SysDistrict.Update().Where(sysdistrict.TreeIDEQ(*oparent.TreeID), sysdistrict.And(sysdistrict.TreeLeftGT(*oparent.TreeRight), sysdistrict.TreeLeftLTE(*nparent.TreeLeft))).AddTreeLeft(-d1).Save(ctx)
						if err != nil {
							return err
						}

						_, err = tx.SysDistrict.Update().Where(sysdistrict.TreeIDEQ(*oparent.TreeID), sysdistrict.And(sysdistrict.TreeRightGTE(*oparent.TreeRight), sysdistrict.TreeRightLT(*nparent.TreeRight))).AddTreeRight(-d1).Save(ctx)
						if err != nil {
							return err
						}

						// step 3: repair new parent's left/right
						_, err = tx.SysDistrict.Update().Where(sysdistrict.TreeIDEQ(*nparent.TreeID), sysdistrict.And(sysdistrict.TreeLeftGT(*nparent.TreeRight), sysdistrict.TreeLeftLT(*oitem.TreeLeft))).AddTreeLeft(d1).Save(ctx)
						if err != nil {
							return err
						}
						_, err = tx.SysDistrict.Update().Where(sysdistrict.TreeIDEQ(*nparent.TreeID), sysdistrict.And(sysdistrict.TreeRightGTE(*nparent.TreeRight), sysdistrict.TreeRightLT(*oitem.TreeLeft))).AddTreeRight(d1).Save(ctx)
						if err != nil {
							return err
						}

						//
						// step 1: repair the sub items left/right
						_, err = tx.SysDistrict.Update().Where(sysdistrict.IDIn(ids...)).AddTreeLeft(d2).AddTreeRight(d2).Save(ctx)

						if err != nil {
							return err
						}

						// step 4:  udpate
						_, err = tx.SysDistrict.Update().Where(sysdistrict.IDEQ(id)).SetInput(*iteminput).Save(ctx)
						if err != nil {
							return err
						}

						if oparent_has_children == 1 {
							_, err := tx.SysDistrict.Update().Where(sysdistrict.IDEQ(oparent.ID)).SetIsLeaf(true).Save(ctx)
							if err != nil {
								return err
							}
						}
						// step 5: trigger update subs's tree_path, merge_name, merge_sname
						// job.ID =
						err = a.Queue.Write(types.TaskName_REPAIR_DISTRICT_TREE_PATH, types.RepaireTreeQueue, job)

						log.Println(" -- -- --- 0 -- --- -- 2222 enqueue err : ", err)
						return nil
					})

					if err != nil {
						return nil, err
					}

				} else {

					err = WithTx(ctx, a.EntCli, func(tx *entscheme.Tx) error {

						// step 1: repair old parent's left/right
						_, err := tx.SysDistrict.Update().Where(sysdistrict.TreeIDEQ(*oparent.TreeID), sysdistrict.TreeLeftGT(*oitem.TreeRight)).AddTreeLeft(-d1).Save(ctx)
						if err != nil {
							return err
						}

						_, err = tx.SysDistrict.Update().Where(sysdistrict.TreeIDEQ(*oparent.TreeID), sysdistrict.TreeRightGT(*oitem.TreeRight)).AddTreeRight(-d1).Save(ctx)
						if err != nil {
							return err
						}

						// step 2: repair new parent's left/right
						_, err = tx.SysDistrict.Update().Where(sysdistrict.TreeIDEQ(*nparent.TreeID), sysdistrict.TreeLeftGT(*nparent.TreeRight)).AddTreeLeft(d1).Save(ctx)
						if err != nil {
							return err
						}
						_, err = tx.SysDistrict.Update().Where(sysdistrict.TreeID(*nparent.TreeID), sysdistrict.TreeRightGT(*nparent.TreeRight)).AddTreeRight(d1).Save(ctx)
						if err != nil {
							return err
						}

						// step 3: repair the sub items left/right, switch tree_id
						_, err = tx.SysDistrict.Update().Where(sysdistrict.TreeIDEQ(*oparent.TreeID), sysdistrict.TreeLeftGTE(*oitem.TreeLeft), sysdistrict.TreeRightLTE(*oitem.TreeRight)).AddTreeLeft(d2).AddTreeRight(d2).SetTreeID(*nparent.TreeID).Save(ctx)
						if err != nil {
							return err
						}

						// step 3:  udpate
						_, err = tx.SysDistrict.Update().Where(sysdistrict.IDEQ(id)).SetInput(*iteminput).Save(ctx)
						if err != nil {
							return err
						}

						// step 5: trigger update subs's tree_path, merge_name, merge_sname
						// job.ID =
						err = a.Queue.Write(types.TaskName_REPAIR_DISTRICT_TREE_PATH, types.RepaireTreeQueue, job)

						log.Println(" -- -- --- 0 -- --- -- 3333 enqueue err : ", err)

						return nil
					})

					if err != nil {
						return nil, err
					}
				}
			}
		}
	}

	res_sysdistrict, err := a.EntCli.SysDistrict.Query().Where(sysdistrict.IDEQ(id)).First(ctx)

	if err != nil {
		return nil, err
	}
	sch_sysdistrict := ToSchemaSysDistrict(res_sysdistrict)

	return sch_sysdistrict, nil
}

func (a *SysDistrict) GetLatestTreeID(ctx context.Context) (int64, error) {

	var opt schema.SysDistrictQueryOptions
	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField(sysdistrict.FieldTreeID, schema.OrderByDESC))

	most, err := a.EntCli.SysDistrict.Query().Order(sysdistrict.ByTreeID(OrderDirection("desc"))).Select(sysdistrict.FieldID, sysdistrict.FieldTreeID).First(ctx)

	if err != nil {
		if entscheme.IsNotFound(err) {
			return 1, nil
		}
		return -1, err
	}

	return *most.TreeID + 1, nil
}

// Delete 删除数据
func (a *SysDistrict) Delete(ctx context.Context, id string) error {

	tobeDel, err := a.EntCli.SysDistrict.Query().Where(sysdistrict.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}
	distance := *tobeDel.TreeRight - *tobeDel.TreeLeft + 1

	err = WithTx(ctx, a.EntCli, func(tx *entscheme.Tx) error {

		_, err = tx.SysDistrict.Update().Where(sysdistrict.TreeIDEQ(*tobeDel.TreeID), sysdistrict.TreeLeftGTE(*tobeDel.TreeLeft), sysdistrict.TreeRightLTE(*tobeDel.TreeRight)).SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)
		if err != nil {
			return err
		}

		_, err := tx.SysDistrict.Update().Where(sysdistrict.TreeIDEQ(*tobeDel.TreeID), sysdistrict.TreeLeftGT(*tobeDel.TreeRight)).AddTreeLeft(-distance).Save(ctx)
		if err != nil {
			return err
		}

		_, err = tx.SysDistrict.Update().Where(sysdistrict.TreeIDEQ(*tobeDel.TreeID), sysdistrict.TreeRightGT(*tobeDel.TreeRight)).AddTreeRight(-distance).Save(ctx)
		if err != nil {
			return err
		}

		if tobeDel.ParentID != nil {
			// _, err = tx.SysDistrict.Update().Where(sysdistrict.IDEQ(*tobeDel.ParentID)).RemoveChildIDs(id).Save(ctx)
			_, err = tx.SysDistrict.Update().Where(sysdistrict.IDEQ(id)).ClearParentID().Save(ctx)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(err)
}

// UpdateActive 更新状态
func (a *SysDistrict) UpdateActive(ctx context.Context, id string, active bool) error {

	_, err1 := a.EntCli.SysDistrict.Update().Where(sysdistrict.IDEQ(id)).SetIsActive(active).Save(ctx)

	return errors.WithStack(err1)
}
