package schema

import "time"

// Dict 字典对象
type Dict struct {
	ID string `json:"id" ` // 唯一标识

	NameCn string `json:"name_cn" binding:"required"` // 字典名（中）
	NameEn string `json:"name_en" binding:"required"` // 字典名（英）
	Status int    `json:"status" binding:"required"`  // 状态
	Memo   string `json:"memo"`                       // 备注
	Sort   int    `json:"sort"`                       // 排序

	Creator   string    `json:"creator"`    // 创建者
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间

	Items DictItems `json:"items" binding:"required,gt=0"` // 字典项列表

}

// DictQueryParam 查询条件
type DictQueryParam struct {
	PaginationParam
	IDs        []string `form:"-"`          // 唯一标识列表
	NameCn     string   `form:"-"`          // 字典名称(中)
	NameEn     string   `form:"-"`          // 字典名称(英)
	QueryValue string   `form:"queryValue"` // 模糊查询
	Status     int16      `form:"status"`     // 状态(1:启用 2:禁用)
	SqlRaw     string
}

// DictQueryOptions 查询可选参数项
type DictQueryOptions struct {
	OrderFields []*OrderField // 排序字段

}

// DictQueryResult 查询结果
type DictQueryResult struct {
	Data       Dicts
	PageResult *PaginationResult
}

// Dicts 字典列表
type Dicts []*Dict
