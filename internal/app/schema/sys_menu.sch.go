package schema

import (
	"strings"

	"github.com/heromicro/omgind/pkg/helper/json"
)

// Menu 菜单对象
type Menu struct {
	TimeMixin `yaml:"-"`

	ID         string  `json:"id" yaml:"id"`                                  // 唯一标识
	Name       string  `json:"name" binding:"required" yaml:"name"`           // 菜单名称
	Sort       int     `json:"sort" yaml:"sort"`                              // 排序值
	Icon       string  `json:"icon" yaml:"icon"`                              // 菜单图标
	Router     string  `json:"router" yaml:"router"`                          // 访问路由
	ParentID   *string `json:"parent_id" yaml:"parent_id"`                    // 父级ID
	ParentPath string  `json:"parent_path" yaml:"parent_path"`                // 父级路径
	IsActive   *bool   `json:"is_active" binding:"required" yaml:"is_active"` // 状态
	Memo       string  `json:"memo" yaml:"memo"`                              // 备注
	IsShow     *bool   `json:"is_show" yaml:"is_show"`                        // 是否显示
	OpenBlank  *bool   `json:"open_blank" yaml:"open_blank"`                  // 是否打开新标签

	Level  int32 `json:"level" yaml:"level"`     // 层级
	IsLeaf bool  `json:"is_leaf" yaml:"is_leaf"` // 是否子叶

	Actions  MenuActions `json:"actions" yaml:"actions"` // 动作列表
	Parent   *Menu       `json:"parent" yaml:"-"`        // 父级
	Children Menus       `json:"children" yaml:"-"`      // 子级
}

func (a *Menu) String() string {
	return json.MarshalToString(a)
}

// MenuQueryParam 查询条件
type MenuQueryParam struct {
	PaginationParam
	IDs              []string `form:"-"`   // 唯一标识列表
	Name             string   `form:"-"`   // 菜单名称
	PrefixParentPath string   `form:"-"`   // 父级路径(前缀模糊查询)
	QueryValue       string   `form:"q"`   // 模糊查询
	ParentID         *string  `form:"pid"` // 父级内码
	IsShow           *bool    `form:"isShow"`
	IsActive         *bool    `form:"is_active"` // 状态
	Level            *int32   `form:"level"`     //

	Level_Order string `form:"level__order" json:"level__order"` // asc/desc

	WithParent   *bool `form:"w__parent" json:"w__parent"`     //
	WithChildren *bool `form:"w__children" json:"w__children"` //

	BasicOrderParam
	TimeOrderParam
}

func (a *MenuQueryParam) String() string {
	return json.MarshalToString(a)
}

// MenuQueryOptions 查询可选参数项
type MenuQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// MenuQueryResult 查询结果
type MenuQueryResult struct {
	Data       Menus
	PageResult *PaginationResult
}

// Menus 菜单列表
type Menus []*Menu

func (a Menus) Len() int {
	return len(a)
}

func (a Menus) Less(i, j int) bool {
	return a[i].Sort > a[j].Sort
}

func (a Menus) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// ToMap 转换为键值映射
func (a Menus) ToMap() map[string]*Menu {
	m := make(map[string]*Menu)
	for _, item := range a {
		m[item.ID] = item
	}
	return m
}

// SplitParentIDs 拆分父级路径的唯一标识列表
func (a Menus) SplitParentIDs() []string {
	idList := make([]string, 0, len(a))
	mIDList := make(map[string]struct{})

	for _, item := range a {
		if _, ok := mIDList[item.ID]; ok || item.ParentPath == "" {
			continue
		}

		for _, pp := range strings.Split(item.ParentPath, "/") {
			if _, ok := mIDList[pp]; ok {
				continue
			}
			idList = append(idList, pp)
			mIDList[pp] = struct{}{}
		}
	}

	return idList
}

// ToTree 转换为菜单树
func (a Menus) ToTree() MenuTrees {

	list := make(MenuTrees, len(a))
	for i, item := range a {
		list[i] = &MenuTree{
			ID:         item.ID,
			Name:       item.Name,
			Icon:       item.Icon,
			Router:     item.Router,
			ParentID:   item.ParentID,
			ParentPath: item.ParentPath,
			Sort:       item.Sort,
			IsShow:     item.IsShow,
			Level:      item.Level,
			IsLeaf:     item.IsLeaf,
			IsActive:   item.IsActive,
			Actions:    item.Actions,
		}
	}

	//fmt.Printf(" == ===== === %t \n ", *list[0].IsShow)

	return list.ToTree()
}

// FillMenuAction 填充菜单动作列表
func (a Menus) FillMenuAction(mActions map[string]MenuActions) Menus {
	for _, item := range a {
		if v, ok := mActions[item.ID]; ok {
			item.Actions = v
		}
	}

	return a
}

// ----------------------------------------MenuTree--------------------------------------

// MenuTree 菜单树
type MenuTree struct {
	ID         string  `yaml:"id" json:"id"`                   // 唯一标识
	Name       string  `yaml:"name" json:"name"`               // 菜单名称
	Icon       string  `yaml:"icon" json:"icon"`               // 菜单图标
	Router     string  `yaml:"router,omitempty" json:"router"` // 访问路由
	ParentID   *string `yaml:"parent_id" json:"parent_id"`     // 父级ID
	ParentPath string  `yaml:"-" json:"parent_path"`           // 父级路径

	Sort      int   `yaml:"sort" json:"sort"` // 排序值
	IsShow    *bool `json:"is_show"`
	IsActive  *bool `yaml:"-" json:"is_active"`           // 状态
	OpenBlank *bool `json:"open_blank" yaml:"open_blank"` // 是否打开新标签

	Level    int32       `json:"level"`                                        // 层级
	IsLeaf   bool        `json:"is_leaf"`                                      // 是否子叶
	Actions  MenuActions `yaml:"actions,omitempty" json:"actions"`             // 动作列表
	Children *MenuTrees  `yaml:"children,omitempty" json:"children,omitempty"` // 子级树

	Memo string `json:"-" yaml:"memo"` // 备注

}

// MenuTrees 菜单树列表
type MenuTrees []*MenuTree

// ToTree 转换为树形结构
func (a MenuTrees) ToTree() MenuTrees {
	mi := make(map[string]*MenuTree)
	for _, item := range a {
		mi[item.ID] = item
	}

	var list MenuTrees
	for _, item := range a {
		if item.ParentID == nil || *item.ParentID == "" {
			list = append(list, item)
			continue
		}
		if pitem, ok := mi[*item.ParentID]; ok {
			if pitem.Children == nil {
				children := MenuTrees{item}
				pitem.Children = &children
				continue
			}
			*pitem.Children = append(*pitem.Children, item)
		}
	}
	return list
}
