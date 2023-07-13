package schema

import (
	"github.com/heromicro/omgind/pkg/helper/json"
)

// DictItem 字典项对象
type DictItem struct {
	TimeMixin `yaml:"-"`

	ID       string `json:"id" yaml:"id"`                                   // 唯一标识
	Label    string `json:"label" binding:"required"  yaml:"label"`         // 显示值
	Value    int    `json:"value" binding:"required" yaml:"value"`          // 字典值
	IsActive *bool  `json:"is_active" binding:"required"  yaml:"is_active"` // 状态
	DictID   string `json:"dict_id" binding:"required" yaml:"dict_id"`      // dict.id
	Memo     string `json:"memo" yaml:"memo"`                               // 备注
	Sort     int    `json:"sort" yaml:"sort"`                               // 排序

}

func (a *DictItem) String() string {
	return json.MarshalToString(a)
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
	if *di.IsActive != *target.IsActive {
		return false
	}

	return true
}

// DictItemQueryParam 查询条件
type DictItemQueryParam struct {
	PaginationParam
	DictID string //字典id
	IDs    []string
	// QueryValue string `form:"q"` // 模糊查询
	// IsActive   *bool  `form:"is_active"`
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
		m[item.ID+"-"+item.Label] = item
	}
	return m
}
