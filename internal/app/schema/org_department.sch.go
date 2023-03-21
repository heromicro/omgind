package schema

import "time"

// OrgDepartment 部门管理对象
type OrgDepartment struct {
	ID   string `json:"id"`                      // 唯一标识
	Name string `json:"name" binding:"required"` // 名称
	Code string `json:"code"`                    // 助记码

	Memo     *string `json:"memo"`                         //
	IsActive *bool   `json:"is_active" binding:"required"` // 状态
	Sort     int     `json:"sort,omitempty"`

	OrgID string        `json:"org_id"` // 企业id
	Org   *OrgOrganShow `json:"org"`    //

	Creator   string     `json:"creator"`    // 创建者
	CreatedAt *time.Time `json:"created_at"` // 创建时间
	UpdatedAt *time.Time `json:"updated_at"` // 更新时间

}

// OrgDepartmentQueryParam 查询条件
type OrgDepartmentQueryParam struct {
	PaginationParam

	QueryValue string `form:"queryValue"` // 模糊查询

	Name     string `form:"name" json:"name"`           //
	Code     string `form:"code" json:"code"`           //
	IsActive *bool  `form:"is_active" json:"is_active"` //
	OrgID    string `form:"org_id" json:"org_id"`       //

	// example: "asc"
	// example: "desc"
	CreatedAt_Order string `form:"created_at__order" json:"created_at__order"` // asc, desc
	// example: "asc"
	// example: "desc"
	IsActive_Order string `form:"is_active__order" json:"is_active__order"` //  asc desc
	// example: "asc"
	// example: "desc"
	Sort_Order string `form:"sort__order" json:"sort__order"` // asc desc

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
