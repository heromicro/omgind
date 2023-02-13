package schema

// ----------------------------------------MenuAction--------------------------------------

// MenuAction 菜单动作对象
type MenuAction struct {
	ID        string              `yaml:"-" json:"id" `                         // 唯一标识
	MenuID    string              `yaml:"-" binding:"required" json:"menu_id"`  // 菜单ID
	Code      string              `yaml:"code" binding:"required" json:"code"`  // 动作编号
	Name      string              `yaml:"name" binding:"required" json:"name"`  // 动作名称
	Resources MenuActionResources `yaml:"resources,omitempty" json:"resources"` // 资源列表
}

// MenuActionQueryParam 查询条件
type MenuActionQueryParam struct {
	PaginationParam
	MenuID string   // 菜单ID
	IDs    []string // 唯一标识列表
}

// MenuActionQueryOptions 查询可选参数项
type MenuActionQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// MenuActionQueryResult 查询结果
type MenuActionQueryResult struct {
	Data       MenuActions
	PageResult *PaginationResult
}

// MenuActions 菜单动作管理列表
type MenuActions []*MenuAction

// ToMap 转换为map
func (a MenuActions) ToMap() map[string]*MenuAction {
	m := make(map[string]*MenuAction)
	for _, item := range a {
		m[item.Code] = item
	}
	return m
}

// FillResources 填充资源数据
func (a MenuActions) FillResources(mResources map[string]MenuActionResources) {
	for i, item := range a {
		a[i].Resources = mResources[item.ID]
	}
}

// ToMenuIDMap 转换为菜单ID映射
func (a MenuActions) ToMenuIDMap() map[string]MenuActions {
	m := make(map[string]MenuActions)
	for _, item := range a {
		m[item.MenuID] = append(m[item.MenuID], item)
	}
	return m
}
