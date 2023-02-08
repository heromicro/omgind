package schema

import (
	"time"
)

// DictItem 字典项对象
type DictItem struct {
	ID string `json:"id"` // 唯一标识

	Label  string `json:"label" binding:"required"`   // 显示值
	Value  int    `json:"value" binding:"required"`   // 字典值
	Status int    `json:"status" binding:"required"`  // 状态(1:启用 2:禁用)
	DictID string `json:"dict_id" binding:"required"` // dict.id
	Memo   string `json:"memo"`                       // 备注
	Sort   int    `json:"sort"`                       // 排序

	Creator   string    `json:"creator"`    // 创建者
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间

}

// is queal
func (di *DictItem) Compare(target *DictItem) bool {
	if di.Label != target.Label {
		return false
	} else if di.Value != target.Value {
		return false
	}
	if di.DictID != target.DictID {
		return false
	}
	if di.Sort != target.Sort {
		return false
	}
	if di.Status != target.Status {
		return false
	}

	return true
}

// DictItemQueryParam 查询条件
type DictItemQueryParam struct {
	PaginationParam
	DictID string //字典id
	IDs    []string
}

// DictItemQueryOptions 查询可选参数项
type DictItemQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// DictItemQueryResult 查询结果
type DictItemQueryResult struct {
	Data       DictItems
	PageResult *PaginationResult
}

// DictItems 字典项列表
type DictItems []*DictItem

// ToMap
func (a DictItems) ToMap() map[string]*DictItem {
	m := make(map[string]*DictItem)
	for _, item := range a {
		m[item.ID + "-" + item.Label] = item
	}
	return m
}
