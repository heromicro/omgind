package schema

import "time"

// OrgPosition 职位管理对象
type OrgPosition struct {
	ID        string     `json:"id"`                      // 唯一标识
	Name      string     `json:"name" binding:"required"` // 名称
	code      string     `json:"code" binding:"required"` // 助记码
	OrgID     string     `json:"org_id"`                  // 企业id
	Creator   string     `json:"creator"`                 // 创建者
	CreatedAt *time.Time `json:"created_at"`              // 创建时间
	UpdatedAt *time.Time `json:"updated_at"`              // 更新时间

}

// OrgPositionQueryParam 查询条件
type OrgPositionQueryParam struct {
	PaginationParam
	QueryValue string `form:"queryValue"` // 模糊查询
}

// OrgPositionQueryOptions 查询可选参数项
type OrgPositionQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// OrgPositionQueryResult 查询结果
type OrgPositionQueryResult struct {
	Data       OrgPositions
	PageResult *PaginationResult
}

// OrgPositions 职位管理列表
type OrgPositions []*OrgPosition
