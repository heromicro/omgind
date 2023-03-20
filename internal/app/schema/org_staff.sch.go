package schema

import "time"

// OrgStaff 员工对象
type OrgStaff struct {
	ID        string `json:"id"`                            // 唯一标识
	FirstName string `json:"first_name" binding:"required"` // 名
	LastName  string `json:"last_name" binding:"required"`  // 姓
	Mobile    string `json:"mobile" binding:"required"`     // 电话

	BirthDate string `json:"birth_date" binding:"required"` // 出生日期
	Gender    string `json:"gender" binding:"required"`     // 性别
	IdenNo    string `json:"iden_no" `                      // 身份证号

	WorkerNo    string `json:"worker_no" binding:"required"`  // 工号
	Cubicle     string `json:"cubicle"`                       // 工位
	EntryDate   string `json:"entry_date" binding:"required"` // 入职日期
	RegularDate string `json:"regular_date"`                  // 转正日期
	ResignDate  string `json:"resign_date"`                   // 离职日期

	OrgID string `json:"org_id"` // 企业id

	IsActive *bool   `json:"is_active" binding:"required"` // 状态
	Sort     int     `json:"sort,omitempty"`
	Memo     *string `json:"memo"` //

	Creator   string     `json:"creator"`    // 创建者
	CreatedAt *time.Time `json:"created_at"` // 创建时间
	UpdatedAt *time.Time `json:"updated_at"` // 更新时间

}

// OrgStaffQueryParam 查询条件
type OrgStaffQueryParam struct {
	PaginationParam
	QueryValue string `form:"queryValue"` // 模糊查询

	Name     string `form:"name"`      //
	IsActive *bool  `form:"is_active"` //
	OrgId    string `form:"org_id"`    //

	CreatedAt_Order string `form:"created_at__order"` // asc, desc
	IsActive_Order  string `form:"is_active__order"`  //
	Sort_Order      string `form:"sort__order"`       // asc desc

}

// OrgStaffQueryOptions 查询可选参数项
type OrgStaffQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// OrgStaffQueryResult 查询结果
type OrgStaffQueryResult struct {
	Data       OrgStaffs
	PageResult *PaginationResult
}

// OrgStaffs 员工管理列表
type OrgStaffs []*OrgStaff
