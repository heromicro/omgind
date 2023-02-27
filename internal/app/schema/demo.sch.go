package schema

import "time"

// Demo 示例对象
// swagger:model
type Demo struct {
	ID       string `json:"id" `                          // 唯一标识
	Code     string `json:"code" binding:"required"`      // 编号
	Name     string `json:"name" binding:"required"`      // 名称
	Memo     string `json:"memo"`                         // 备注
	IsActive *bool  `json:"is_active" binding:"required"` // 状态
	Sort     int    `json:"sort,omitempty"`

	Creator   string     `json:"creator" `    // 创建者
	CreatedAt *time.Time `json:"created_at" ` // 创建时间
	UpdatedAt *time.Time `json:"updated_at" ` // 更新时间
}

// DemoQueryParam 查询条件
// swagger:parameters
type DemoQueryParam struct {
	PaginationParam
	QueryValue string `form:"queryValue"` // 查询值
	Code       string `form:"code"`       // 编号
	Name       string `form:"name"`       // 名称
	IsActive   *bool  `form:"is_active"`  // 状态

	CreatedAt_St *int64 `form:"created_at__st"` // create_at 开始
	CreatedAt_Et *int64 `form:"created_at__et"` // create_at 结束

	CreatedAt_Order string `form:"created_at__order"` // asc, desc
	Code_Order      string `form:"code__order"`       // asc desc
	Name_Order      string `form:"name__order"`       // asc desc
	Sort_Order      string `form:"sort__order"`       // asc desc

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
