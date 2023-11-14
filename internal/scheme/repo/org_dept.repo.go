package repo

import (
	"context"
	"time"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/orgdept"
	"github.com/heromicro/omgind/internal/gen/mainent/orgorgan"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"

	"github.com/google/wire"
)

// OrgDeptSet 注入OrgDept
var OrgDeptSet = wire.NewSet(wire.Struct(new(OrgDept), "*"))

// OrgDept 部门管理存储
type OrgDept struct {
	EntCli *mainent.Client
}

// ToSchemaOrgDept 转换为
func ToSchemaOrgDept(et *mainent.OrgDept) *schema.OrgDept {
	item := new(schema.OrgDept)
	structure.Copy(et, item)
	if et.Edges.Organ != nil {
		item.Org = ToSchemaOrgOrganShow(et.Edges.Organ)
	}
	if et.Edges.Parent != nil {
		item.Parent = ToSchemaOrgDept(et.Edges.Parent)
	}
	return item
}

func ToSchemaOrgDepts(ets mainent.OrgDepts) []*schema.OrgDept {
	list := make([]*schema.OrgDept, len(ets))
	for i, item := range ets {
		list[i] = ToSchemaOrgDept(item)
	}
	return list
}

func ToEntCreateOrgDeptInput(sch *schema.OrgDept) *mainent.CreateOrgDeptInput {
	createinput := new(mainent.CreateOrgDeptInput)
	structure.Copy(sch, &createinput)

	return createinput
}

func ToEntUpdateOrgDeptInput(sch *schema.OrgDept) *mainent.UpdateOrgDeptInput {
	updateinput := new(mainent.UpdateOrgDeptInput)
	structure.Copy(sch, &updateinput)

	return updateinput
}

func (a *OrgDept) getQueryOption(opts ...schema.OrgDeptQueryOptions) schema.OrgDeptQueryOptions {
	var opt schema.OrgDeptQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *OrgDept) Query(ctx context.Context, params schema.OrgDeptQueryParam, opts ...schema.OrgDeptQueryOptions) (*schema.OrgDeptQueryResult, error) {
	// opt := a.getQueryOption(opts...)

	query := a.EntCli.OrgDept.Query().WithOrgan(func(ooq *mainent.OrgOrganQuery) {
		ooq.Select(orgorgan.FieldID, orgorgan.FieldName, orgorgan.FieldSname)
	})

	query = query.Where(orgdept.DeletedAtIsNil())
	// TODO: 查询条件

	query = query.Where(orgdept.IsShowEQ(*params.IsShow))

	if v := params.Name; v != "" {
		query = query.Where(orgdept.NameContains(v))
	}

	if v := params.Code; v != "" {
		query = query.Where(orgdept.CodeContains(v))
	}

	if v := params.OrgID; v != "" {
		query = query.Where(orgdept.OrgIDEQ(v))
	}

	if v := params.IsActive; v != nil {
		query = query.Where(orgdept.IsActiveEQ(*v))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.OrgDeptQueryResult{PageResult: pr}, nil
	}

	default_order := true
	if v := params.IsActive_Order; v != "" {
		query = query.Order(orgdept.ByIsActive(OrderDirection(v)))
		default_order = false
	}

	if v := params.Sort_Order; v != "" {
		query = query.Order(orgdept.BySort(OrderDirection(v)))
		default_order = false
	}

	if default_order {
		query = query.Order(orgdept.ByTreeID(OrderDirection("asc")))
		query = query.Order(orgdept.ByTreeLevel(OrderDirection("asc")))
		query = query.Order(orgdept.ByID(OrderDirection("asc")))
	}

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.OrgDeptQueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.OrgDeptQueryResult{
		PageResult: pr,
		Data:       ToSchemaOrgDepts(list),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *OrgDept) Get(ctx context.Context, id string, opts ...schema.OrgDeptQueryOptions) (*schema.OrgDept, error) {

	query := a.EntCli.OrgDept.Query().WithOrgan(func(ooq *mainent.OrgOrganQuery) {
		ooq.Select(orgorgan.FieldID, orgorgan.FieldName, orgorgan.FieldSname, orgorgan.FieldIdenNo)
	}).WithParent(func(odq *mainent.OrgDeptQuery) {
		odq.Select(orgdept.FieldID, orgdept.FieldName, orgdept.FieldMergeName, orgdept.FieldTreeID, orgdept.FieldTreePath)
	})
	r_orgdept, err := query.Where(orgdept.IDEQ(id)).Only(ctx)
	if err != nil {
		if mainent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return ToSchemaOrgDept(r_orgdept), nil
}

// View 查询指定数据
func (a *OrgDept) View(ctx context.Context, id string, opts ...schema.OrgDeptQueryOptions) (*schema.OrgDept, error) {

	query := a.EntCli.OrgDept.Query().WithOrgan(func(ooq *mainent.OrgOrganQuery) {
		ooq.Select(orgorgan.FieldID, orgorgan.FieldName, orgorgan.FieldSname, orgorgan.FieldIdenNo).WithHaddr()
	}).WithParent(func(odq *mainent.OrgDeptQuery) {
		odq.Select(orgdept.FieldID, orgdept.FieldName)
	})
	r_orgdept, err := query.Where(orgdept.IDEQ(id)).Only(ctx)
	if err != nil {
		if mainent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return ToSchemaOrgDept(r_orgdept), nil
}

func (a *OrgDept) GetAllSubs(ctx context.Context, pid string, params schema.OrgDeptQueryParam, opts ...schema.OrgDeptQueryOptions) (*schema.OrgDeptQueryResult, error) {

	// opt := a.getQueryOption(opts...)

	query := a.EntCli.OrgDept.Query()

	query = query.Where(orgdept.DeletedAtIsNil(), orgdept.IsDelEQ(false))
	// TODO: 查询条件

	if v := params.OrgID; v == "" {
		return nil, errors.New("no org_id")
	} else {
		query = query.Where(orgdept.OrgID(v))
	}

	if pid == "-" {
		query = query.Where(orgdept.ParentIDIsNil())
	} else {
		query = query.Where(orgdept.ParentIDEQ(pid))
	}

	if v := params.Name; v != "" {
		query = query.Where(orgdept.NameEQ(v))
	}

	if v := params.IsReal; v != nil {
		query = query.Where(orgdept.IsRealEQ(*v))
	}

	if v := params.IsLeaf; v != nil {
		query = query.Where(orgdept.IsLeafEQ(*v))
	}

	if v := params.IsActive; v != nil {
		query = query.Where(orgdept.IsActiveEQ(*v))
	}

	if v := params.TreeLevel; v != nil {
		query = query.Where(orgdept.TreeLevelEQ(*v))
	}
	if v := params.TreeLeft; v != nil {
		query = query.Where(orgdept.TreeLeftGTE(*v))
	}

	if v := params.TreeLeft_St; v != nil {
		query = query.Where(orgdept.TreeLeftGTE(*v))
	}
	if v := params.TreeLeft_Ed; v != nil {
		query = query.Where(orgdept.TreeLeftLTE(*v))
	}

	if v := params.TreeRight; v != nil {
		query = query.Where(orgdept.TreeRightLTE(*v))
	}

	if v := params.TreeRight_St; v != nil {
		query = query.Where(orgdept.TreeRightGTE(*v))
	}
	if v := params.TreeRight_Ed; v != nil {
		query = query.Where(orgdept.TreeRightLTE(*v))
	}

	// 父级pid
	if v := params.ParentParentID; v != nil {
		query = query.Where(orgdept.HasParentWith(orgdept.ParentIDEQ(*v)))
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
		return &schema.OrgDeptQueryResult{PageResult: pr}, nil
	}

	has_order := false
	if v := params.CreatedAt_Order; v != nil {
		query = query.Order(orgdept.ByCreatedAt(OrderDirection(*v)))
		has_order = true
	}

	if v := params.Name_Order; v != "" {
		query = query.Order(orgdept.ByName(OrderDirection(v)))
		has_order = true
	}

	if v := params.TreeID_Order; v != "" {
		query = query.Order(orgdept.ByTreeID(OrderDirection(v)))
		has_order = true
	}

	if v := params.TreeLevel_Order; v != "" {
		query = query.Order(orgdept.ByTreeLevel(OrderDirection(v)))
		has_order = true
	}

	if v := params.TreeLeft_Order; v != "" {
		query = query.Order(orgdept.ByTreeLeft(OrderDirection(v)))
		has_order = true
	}

	if !has_order {
		query = query.Order(orgdept.ByTreeID(OrderDirection("asc")))
		query = query.Order(orgdept.ByTreeLevel(OrderDirection("asc")))
		query = query.Order(orgdept.ByTreeLeft(OrderDirection("asc")))
		query = query.Order(orgdept.BySort(OrderDirection("asc")))
		query = query.Order(orgdept.ByID(OrderDirection("asc")))
	}

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()

	// log.Println(" ------- ===== params.Offset() ", params.Offset())
	// log.Println(" ---- === ==== count 22222 ", count)

	if params.Offset() > count {
		return &schema.OrgDeptQueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	//
	squery := query.Select(orgdept.FieldID, orgdept.FieldName, orgdept.FieldParentID, orgdept.FieldIsLeaf)

	list, err1 := squery.All(ctx)

	// log.Println(" ------- ===== ===== === list ", list)

	if err1 != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.OrgDeptQueryResult{
		PageResult: pr,
		Data:       ToSchemaOrgDepts(list),
	}

	return qr, nil
}

func (a *OrgDept) GetTree(ctx context.Context, tpid string, params schema.OrgDeptQueryParam, opts ...schema.OrgDeptQueryOptions) (*schema.OrgDeptQueryTreeResult, error) {
	// opt := a.getQueryOption(opts...)

	parent_query := a.EntCli.OrgDept.Query().Where(orgdept.DeletedAtIsNil())

	parent_query = parent_query.Where(orgdept.IDEQ(tpid))

	parent, err := parent_query.Select(orgdept.FieldID, orgdept.FieldTreeID).First(ctx)
	if err != nil {
		return nil, err
	}

	tree_query := a.EntCli.OrgDept.Query().Where(orgdept.IsDelEQ(false), orgdept.TreeIDEQ(*parent.TreeID)).Where(orgdept.IDNEQ(parent.ID))

	tree_query = tree_query.Order(orgdept.ByTreeLevel(OrderDirection("asc")))
	tree_query = tree_query.Order(orgdept.ByTreeLeft(OrderDirection("asc")))
	tree_query = tree_query.Order(orgdept.BySort(OrderDirection("asc")))

	select_tree := tree_query.Select(orgdept.FieldID, orgdept.FieldName, orgdept.FieldParentID, orgdept.FieldIsLeaf, orgdept.FieldTreeID, orgdept.FieldTreeLeft, orgdept.FieldTreeRight)

	subs, err := select_tree.All(ctx)
	if err != nil {
		return nil, err
	}

	top_query := a.EntCli.OrgDept.Query().Where(orgdept.IsDel(false), orgdept.TreeLevel(1), orgdept.ParentIDIsNil())

	top, err := top_query.All(ctx)
	if err != nil {
		return nil, err
	}

	sch_sub := ToSchemaOrgDepts(subs)
	sch_top := ToSchemaOrgDepts(top)
	data := &schema.OrgDeptQueryTreeResult{
		Top:  sch_top,
		Subs: sch_sub,
	}

	return data, nil
}

// Delete 删除数据
func (a *OrgDept) Delete(ctx context.Context, id string) error {

	r_orgdept, err := a.EntCli.OrgDept.Query().Where(orgdept.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}

	_, err1 := r_orgdept.Update().SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)

	return errors.WithStack(err1)
}

// UpdateActive 更新状态
func (a *OrgDept) UpdateActive(ctx context.Context, id string, active bool) error {

	_, err1 := a.EntCli.OrgDept.Update().Where(orgdept.IDEQ(id)).SetIsActive(active).Save(ctx)

	return errors.WithStack(err1)
}

func (a *OrgDept) GetLatestTreeID(ctx context.Context) (int64, error) {

	most, err := a.EntCli.OrgDept.Query().Order(orgdept.ByTreeID(OrderDirection("desc"))).Select(orgdept.FieldID, orgdept.FieldTreeID).First(ctx)

	if err != nil {
		if mainent.IsNotFound(err) {
			return 1, nil
		}
		return -1, err
	}

	return *most.TreeID + 1, nil
}
