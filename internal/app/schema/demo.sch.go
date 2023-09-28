package schema

// Demo 示例对象
// swagger:model
type Demo struct {
	TimeMixin `yaml:"-"`

	ID       string `json:"id" `                          // 唯一标识
	Code     string `json:"code" binding:"required"`      // 编号
	Name     string `json:"name" binding:"required"`      // 名称
	Memo     string `json:"memo"`                         // 备注
	IsActive *bool  `json:"is_active" binding:"required"` // 状态
	Sort     int    `json:"sort,omitempty"`
}

// DemoQueryParam 查询条件
// swagger:parameters
type DemoQueryParam struct {
	PaginationParam
	QueryValue string `form:"q" json:"q"`                 // 查询值
	Code       string `form:"code" json:"code"`           // 编号
	Name       string `form:"name" json:"name"`           // 名称
	IsActive   *bool  `form:"is_active" json:"is_active"` // 状态

	// example: "asc"
	// example: "desc"
	Code_Order string `form:"code__order" json:"code__order"` // asc desc
	// example: "asc"
	// example: "desc"
	Name_Order string `form:"name__order" json:"name__order"` // asc desc

	BasicOrderParam
	TimeOrderParam
}

// DemoQueryOptions 示例对象查询可选参数项
type DemoQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// DemoQueryResult 示例对象查询结果
type DemoQueryResult struct {
	Data       Demos
	PageResult *PaginationResult
}

// demos 对象列表
type Demos []*Demo
