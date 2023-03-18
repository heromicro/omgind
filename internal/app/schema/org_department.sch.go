package schema

import "time"

// OrgDepartment 部门管理对象
type OrgDepartment struct {
	ID    string `json:"id"`                      // 唯一标识
	Name  string `json:"name" binding:"required"` // 名称
	Code  string `json:"code"`                    // 助记码
	OrgID string `json:"org_id"`                  // 企业id

	Memo     *string `json:"memo"`                         //
	IsActive *bool   `json:"is_active" binding:"required"` // 状态
	Sort     int     `json:"sort,omitempty"`

	Creator   string     `json:"creator"`    // 创建者
	CreatedAt *time.Time `json:"created_at"` // 创建时间
	UpdatedAt *time.Time `json:"updated_at"` // 更新时间

}

// OrgDepartmentQueryParam 查询条件
type OrgDepartmentQueryParam struct {
	PaginationParam
	QueryValue string  `form:"queryValue"` // 模糊查询
	Name       string  `form:"name"`       //
	Code       *string `form:"code"`       //
	IsActive   *bool   `form:"is_active"`  //
	OrgId      string  `form:"org_id"`     //

	IsActive_Order *bool `form:"is_active__order"` //

}

// OrgDepartmentQueryOptions 查询可选参数项
type OrgDepartmentQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// OrgDepartmentQueryResult 查询结果
type OrgDepartmentQueryResult struct {
	Data       OrgDepartments
	PageResult *PaginationResult
}

// OrgDepartments 部门管理列表
type OrgDepartments []*OrgDepartment
