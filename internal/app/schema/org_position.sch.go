package schema

// OrgPosition 职位管理对象
type OrgPosition struct {
	TimeMixin `yaml:"-"`

	ID   string `json:"id"`                      // 唯一标识
	Name string `json:"name" binding:"required"` // 名称
	Code string `json:"code"`                    // 助记码

	Memo     *string `json:"memo"`                         //
	IsActive *bool   `json:"is_active" binding:"required"` // 状态
	Sort     int     `json:"sort,omitempty"`

	OrgMixin
}

// OrgPositionQueryParam 查询条件
type OrgPositionQueryParam struct {
	PaginationParam
	QueryValue string `form:"q" json:"q"` // 模糊查询

	Name     string `form:"name" json:"name"`           //
	Code     string `form:"code" json:"code"`           //
	IsActive *bool  `form:"is_active" json:"is_active"` //
	OrgID    string `form:"org_id" json:"org_id"`       //

	BasicOrderParam
	TimeOrderParam
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
