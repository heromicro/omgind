package repo

import (
	"context"
	"time"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/orgdepartment"
	"github.com/heromicro/omgind/internal/gen/ent/orgorgan"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"

	"github.com/google/wire"
)

// OrgDepartmentSet 注入OrgDepartment
var OrgDepartmentSet = wire.NewSet(wire.Struct(new(OrgDepartment), "*"))

// OrgDepartment 部门管理存储
type OrgDepartment struct {
	EntCli *ent.Client
}

// ToSchemaOrgDepartment 转换为
func ToSchemaOrgDepartment(et *ent.OrgDepartment) *schema.OrgDepartment {
	item := new(schema.OrgDepartment)
	structure.Copy(et, item)
	if et.Edges.Organ != nil {
		item.Org = ToSchemaOrgOrganShow(et.Edges.Organ)
	}
	return item
}

func ToSchemaOrgDepartments(ets ent.OrgDepartments) []*schema.OrgDepartment {
	list := make([]*schema.OrgDepartment, len(ets))
	for i, item := range ets {
		list[i] = ToSchemaOrgDepartment(item)
	}
	return list
}

func (a *OrgDepartment) ToEntCreateOrgDepartmentInput(sch *schema.OrgDepartment) *ent.CreateOrgDepartmentInput {
	createinput := new(ent.CreateOrgDepartmentInput)
	structure.Copy(sch, &createinput)

	return createinput
}

func (a *OrgDepartment) ToEntUpdateOrgDepartmentInput(sch *schema.OrgDepartment) *ent.UpdateOrgDepartmentInput {
	updateinput := new(ent.UpdateOrgDepartmentInput)
	structure.Copy(sch, &updateinput)

	return updateinput
}

func (a *OrgDepartment) getQueryOption(opts ...schema.OrgDepartmentQueryOptions) schema.OrgDepartmentQueryOptions {
	var opt schema.OrgDepartmentQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *OrgDepartment) Query(ctx context.Context, params schema.OrgDepartmentQueryParam, opts ...schema.OrgDepartmentQueryOptions) (*schema.OrgDepartmentQueryResult, error) {
	opt := a.getQueryOption(opts...)

	query := a.EntCli.OrgDepartment.Query().WithOrgan(func(ooq *ent.OrgOrganQuery) {
		ooq.Select(orgorgan.FieldID, orgorgan.FieldName, orgorgan.FieldSname)
	})

	query = query.Where(orgdepartment.DeletedAtIsNil())
	// TODO: 查询条件

	if v := params.Name; v != "" {
		query = query.Where(orgdepartment.NameContains(v))
	}

	if v := params.Code; v != "" {
		query = query.Where(orgdepartment.CodeContains(v))
	}

	if v := params.OrgID; v != "" {
		query = query.Where(orgdepartment.OrgIDEQ(v))
	}

	if v := params.IsActive; v != nil {
		query = query.Where(orgdepartment.IsActiveEQ(*v))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.OrgDepartmentQueryResult{PageResult: pr}, nil
	}

	if v := params.IsActive_Order; v != "" {
		of := MakeUpOrderField(orgdepartment.FieldIsActive, v)
		opt.OrderFields = append(opt.OrderFields, of)
	}

	if v := params.Sort_Order; v != "" {
		of := MakeUpOrderField(orgdepartment.FieldSort, v)
		opt.OrderFields = append(opt.OrderFields, of)
	}

	if len(opt.OrderFields) == 0 {
		opt.OrderFields = append(opt.OrderFields, schema.NewOrderField(orgdepartment.FieldID, schema.OrderByDESC))
	}

	query = query.Order(ParseOrder(opt.OrderFields)...)

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.OrgDepartmentQueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.OrgDepartmentQueryResult{
		PageResult: pr,
		Data:       ToSchemaOrgDepartments(list),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *OrgDepartment) Get(ctx context.Context, id string, opts ...schema.OrgDepartmentQueryOptions) (*schema.OrgDepartment, error) {

	query := a.EntCli.OrgDepartment.Query().WithOrgan(func(ooq *ent.OrgOrganQuery) {
		ooq.Select(orgorgan.FieldID, orgorgan.FieldName, orgorgan.FieldSname, orgorgan.FieldIdenNo)
	})
	r_orgdepartment, err := query.Where(orgdepartment.IDEQ(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return ToSchemaOrgDepartment(r_orgdepartment), nil
}

// View 查询指定数据
func (a *OrgDepartment) View(ctx context.Context, id string, opts ...schema.OrgDepartmentQueryOptions) (*schema.OrgDepartment, error) {

	query := a.EntCli.OrgDepartment.Query().WithOrgan(func(ooq *ent.OrgOrganQuery) {
		ooq.Select(orgorgan.FieldID, orgorgan.FieldName, orgorgan.FieldSname, orgorgan.FieldIdenNo).WithHaddr()
	})
	r_orgdepartment, err := query.Where(orgdepartment.IDEQ(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return ToSchemaOrgDepartment(r_orgdepartment), nil
}

func (a *OrgDepartment) GetAllSubs(ctx context.Context, pid string, params schema.OrgDepartmentQueryParam, opts ...schema.OrgDepartmentQueryOptions) (*schema.OrgDepartmentQueryResult, error) {

	opt := a.getQueryOption(opts...)

	query := a.EntCli.OrgDepartment.Query()

	query = query.Where(orgdepartment.DeletedAtIsNil(), orgdepartment.IsDelEQ(false))
	// TODO: 查询条件

	if pid == "-" {
		query = query.Where(orgdepartment.Or(orgdepartment.ParentIDEQ(pid), orgdepartment.ParentIDIsNil()))
	} else {
		query = query.Where(orgdepartment.ParentIDEQ(pid))
	}

	if v := params.Name; v != "" {
		query = query.Where(orgdepartment.NameEQ(v))
	}

	if v := params.IsReal; v != nil {
		query = query.Where(orgdepartment.IsRealEQ(*v))
	}

	if v := params.IsLeaf; v != nil {
		query = query.Where(orgdepartment.IsLeafEQ(*v))
	}

	if v := params.IsActive; v != nil {
		query = query.Where(orgdepartment.IsActiveEQ(*v))
	}

	if v := params.TreeLevel; v != nil {
		query = query.Where(orgdepartment.TreeLevelEQ(*v))
	}
	if v := params.TreeLeft; v != nil {
		query = query.Where(orgdepartment.TreeLeftGTE(*v))
	}

	if v := params.TreeLeft_St; v != nil {
		query = query.Where(orgdepartment.TreeLeftGTE(*v))
	}
	if v := params.TreeLeft_Ed; v != nil {
		query = query.Where(orgdepartment.TreeLeftLTE(*v))
	}

	if v := params.TreeRight; v != nil {
		query = query.Where(orgdepartment.TreeRightLTE(*v))
	}

	if v := params.TreeRight_St; v != nil {
		query = query.Where(orgdepartment.TreeRightGTE(*v))
	}
	if v := params.TreeRight_Ed; v != nil {
		query = query.Where(orgdepartment.TreeRightLTE(*v))
	}

	// 父级pid
	if v := params.ParentParentID; v != nil {
		query = query.Where(orgdepartment.HasParentWith(orgdepartment.ParentIDEQ(*v)))
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
		return &schema.OrgDepartmentQueryResult{PageResult: pr}, nil
	}

	// opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))

	if v := params.CreatedAt_Order; v != "" {
		of := MakeUpOrderField(orgdepartment.FieldCreatedAt, v)
		opt.OrderFields = append(opt.OrderFields, of)
	}

	if v := params.Name_Order; v != "" {
		of := MakeUpOrderField(orgdepartment.FieldName, v)
		opt.OrderFields = append(opt.OrderFields, of)
	}

	if v := params.TreeID_Order; v != "" {
		of := MakeUpOrderField(orgdepartment.FieldTreeID, v)
		opt.OrderFields = append(opt.OrderFields, of)
	}

	if v := params.TreeLevel_Order; v != "" {
		of := MakeUpOrderField(orgdepartment.FieldTreeLevel, v)
		opt.OrderFields = append(opt.OrderFields, of)
	}

	if v := params.TreeLeft_Order; v != "" {
		of := MakeUpOrderField(orgdepartment.FieldTreeLeft, v)
		opt.OrderFields = append(opt.OrderFields, of)
	}

	if len(opt.OrderFields) == 0 {
		of := MakeUpOrderField(orgdepartment.FieldTreeID, "asc")
		opt.OrderFields = append(opt.OrderFields, of)

		of2 := MakeUpOrderField(orgdepartment.FieldSort, "asc")
		opt.OrderFields = append(opt.OrderFields, of2)
	}

	query = query.Order(ParseOrder(opt.OrderFields)...)

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()

	// log.Println(" ------- ===== params.Offset() ", params.Offset())
	// log.Println(" ---- === ==== count 22222 ", count)

	if params.Offset() > count {
		return &schema.OrgDepartmentQueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	//
	squery := query.Select(orgdepartment.FieldID, orgdepartment.FieldName, orgdepartment.FieldParentID, orgdepartment.FieldIsLeaf)

	list, err1 := squery.All(ctx)

	// log.Println(" ------- ===== ===== === list ", list)

	if err1 != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.OrgDepartmentQueryResult{
		PageResult: pr,
		Data:       ToSchemaOrgDepartments(list),
	}

	return qr, nil
}

func (a *OrgDepartment) GetTree(ctx context.Context, tpid string, params schema.OrgDepartmentQueryParam, opts ...schema.OrgDepartmentQueryOptions) (*schema.OrgDepartmentQueryTreeResult, error) {
	opt := a.getQueryOption(opts...)

	parent_query := a.EntCli.OrgDepartment.Query().Where(orgdepartment.DeletedAtIsNil())

	parent_query = parent_query.Where(orgdepartment.IDEQ(tpid))

	parent, err := parent_query.Select(orgdepartment.FieldID, orgdepartment.FieldTreeID).First(ctx)
	if err != nil {
		return nil, err
	}

	tree_query := a.EntCli.OrgDepartment.Query().Where(orgdepartment.IsDelEQ(false), orgdepartment.TreeIDEQ(*parent.TreeID)).Where(orgdepartment.IDNEQ(parent.ID))

	of1 := MakeUpOrderField(orgdepartment.FieldTreeLevel, "asc")
	opt.OrderFields = append(opt.OrderFields, of1)
	of2 := MakeUpOrderField(orgdepartment.FieldSort, "asc")
	opt.OrderFields = append(opt.OrderFields, of2)
	tree_query = tree_query.Order(ParseOrder(opt.OrderFields)...)

	select_tree := tree_query.Select(orgdepartment.FieldID, orgdepartment.FieldName, orgdepartment.FieldParentID, orgdepartment.FieldIsLeaf, orgdepartment.FieldTreeID, orgdepartment.FieldTreeLeft, orgdepartment.FieldTreeRight)

	subs, err := select_tree.All(ctx)
	if err != nil {
		return nil, err
	}

	top_query := a.EntCli.OrgDepartment.Query().Where(orgdepartment.IsDel(false), orgdepartment.TreeLevel(1), orgdepartment.ParentIDIsNil())

	top, err := top_query.All(ctx)
	if err != nil {
		return nil, err
	}

	sch_sub := ToSchemaOrgDepartments(subs)
	sch_top := ToSchemaOrgDepartments(top)
	data := &schema.OrgDepartmentQueryTreeResult{
		Top:  sch_top,
		Subs: sch_sub,
	}

	return data, nil
}

// Create 创建数据
func (a *OrgDepartment) Create(ctx context.Context, item schema.OrgDepartment) (*schema.OrgDepartment, error) {

	// TODO: check org_id

	iteminput := a.ToEntCreateOrgDepartmentInput(&item)

	r_orgdepartment, err := a.EntCli.OrgDepartment.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_orgdepartment := ToSchemaOrgDepartment(r_orgdepartment)
	return sch_orgdepartment, nil
}

// Update 更新数据
func (a *OrgDepartment) Update(ctx context.Context, id string, item schema.OrgDepartment) (*schema.OrgDepartment, error) {

	// TODO: check org_id

	oitem, err := a.EntCli.OrgDepartment.Query().Where(orgdepartment.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	iteminput := a.ToEntUpdateOrgDepartmentInput(&item)

	r_orgdepartment, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	sch_orgdepartment := ToSchemaOrgDepartment(r_orgdepartment)

	return sch_orgdepartment, nil
}

// Delete 删除数据
func (a *OrgDepartment) Delete(ctx context.Context, id string) error {

	r_orgdepartment, err := a.EntCli.OrgDepartment.Query().Where(orgdepartment.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}

	_, err1 := r_orgdepartment.Update().SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)

	return errors.WithStack(err1)
}

// UpdateActive 更新状态
func (a *OrgDepartment) UpdateActive(ctx context.Context, id string, active bool) error {

	_, err1 := a.EntCli.OrgDepartment.Update().Where(orgdepartment.IDEQ(id)).SetIsActive(active).Save(ctx)

	return errors.WithStack(err1)
}
