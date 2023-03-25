package schema

import "time"

type OrgStaff0 struct {
	ID        string `json:"id"`                            // 唯一标识
	FirstName string `json:"first_name" binding:"required"` // 名
	LastName  string `json:"last_name" binding:"required"`  // 姓
	Mobile    string `json:"mobile" binding:"required"`     // 电话

	BirthDate    *time.Time `json:"birth_date" binding:"required"`     // 出生日期
	Gender       int32      `json:"gender" binding:"required"`         // 性别
	GenderDictID string     `json:"gender_dict_id" binding:"required"` // 性别dict id
	GenderDict   *Dict      `json:"gender_dict"`                       //
	IdenNo       string     `json:"iden_no" `                          // 身份证号

	OrgMixin

	IdenAddrID *string     `json:"iden_addr_id"` // 身份证地址
	IdenAddr   *SysAddress `json:"iden_addr"`    // 身份证地址
	ResiAddrID *string     `json:"resi_addr_id"` // 现居地址
	ResiAddr   *SysAddress `json:"resi_addr"`    // 现居地址

}

type OrgStaff1 struct {
	WorkerNo     string `json:"worker_no"`                         // 工号
	Cubicle      string `json:"cubicle"`                           // 工位
	EmpyStat     *int32 `json:"empy_stat" binding:"required"`      // 在职状态
	EmpystDictID string `json:"empyst_dict_id" binding:"required"` // 在职状态
	EmpystDict   string `json:"empyst_dict"`                       //

	EntryDate   *time.Time `json:"entry_date" binding:"required"` // 入职日期
	RegularDate *time.Time `json:"regular_date"`                  // 转正日期
	ResignDate  *time.Time `json:"resign_date"`                   // 离职日期

}

// OrgStaff 员工对象
type OrgStaff struct {
	OrgStaff0
	OrgStaff1

	OrgID string        `json:"org_id"` // 企业id
	Org   *OrgOrganShow `json:"org"`    //

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
	QueryValue string `form:"queryValue" json:"queryValue"` // 模糊查询

	FirstName string `form:"first_name" json:"first_name"` //
	LastName  string `form:"last_name" json:"last_name"`   //
	Gender    *int32 `form:"gender" json:"gender"`         //
	IsActive  *bool  `form:"is_active" json:"is_active"`   //
	OrgID     string `form:"org_id" json:"org_id"`         //

	// example: asc
	// example: desc
	CreatedAt_Order string `form:"created_at__order" json:"created_at__order"` // asc, desc
	IsActive_Order  string `form:"is_active__order" json:"is_active__order"`   //
	Sort_Order      string `form:"sort__order" json:"sort__order"`             // asc/desc

	BirthDate_Order string `form:"birth_date__order" json:"birth_date__order"` // asc/desc

	WorkerNo_Order    string `form:"worker_no__order" json:"worker_no__order"`       // asc/desc
	EntryDate_Order   string `form:"entry_date__order" json:"entry_date__order"`     // asc/desc
	RegularDate_Order string `form:"regular_date__order" json:"regular_date__order"` // asc/desc
	ResignDate_Order  string `form:"resign_date__order" json:"resign_date__order"`   // asc/desc

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

type OrgStaff0s []*OrgStaff0

// OrgStaffs 员工管理列表
type OrgStaffs []*OrgStaff
