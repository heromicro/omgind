package schema

import "time"

// Dict 字典对象
type Dict struct {
	ID string `json:"id" ` // 唯一标识

	NameCn   string `json:"name_cn" binding:"required"`   // 字典名（中）
	NameEn   string `json:"name_en" binding:"required"`   // 字典名（英）
	IsActive *bool  `json:"is_active" binding:"required"` // 状态
	Memo     string `json:"memo"`                         // 备注
	Sort     int    `json:"sort"`                         // 排序

	Creator   string    `json:"creator"`    // 创建者
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间

	Items DictItems `json:"items" binding:"required,gt=0"` // 字典项列表

}

// DictQueryParam 查询条件
type DictQueryParam struct {
	PaginationParam

	QueryValue string `form:"queryValue" json:"queryValue"` // 模糊查询

	IDs      []string `form:"ids" json:"ids"`         // 唯一标识列表
	NameCn   string   `form:"name_cn" json:"name_cn"` // 字典名称(中)
	NameEn   string   `form:"name_en" json:"name_en"` // 字典名称(英)
	WithItem *bool    `form:"wi" json:"wi"`           //

	IsActive *bool `form:"is_active" json:"is_active"` //

	// example: asc
	// example: desc
	NameEn_Order   string `form:"name_en__order" json:"name_en__order"`     // asc/desc
	NameCn_Order   string `form:"name_cn__order" json:"name_cn__order"`     // asc/desc
	IsActive_Order string `form:"is_active__order" json:"is_active__order"` // asc/desc
	Sort_Order     string `form:"sort__order" json:"sort__order"`           // asc/desc

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
