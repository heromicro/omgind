package schema

// ----------------------------------------MenuActionResource--------------------------------------

// MenuActionResource 菜单动作关联资源对象
type MenuActionResource struct {
	ID       string `yaml:"-" json:"id" `                            // 唯一标识
	ActionID string `yaml:"-" json:"action_id"`                      // 菜单动作ID
	Method   string `yaml:"method" binding:"required" json:"method"` // 资源请求方式(支持正则)
	Path     string `yaml:"path" binding:"required" json:"path"`     // 资源请求路径（支持/:id匹配）
}

// MenuActionResourceQueryParam 查询条件
type MenuActionResourceQueryParam struct {
	PaginationParam
	MenuID  string   // 菜单ID
	MenuIDs []string // 菜单ID列表
}

// MenuActionResourceQueryOptions 查询可选参数项
type MenuActionResourceQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// MenuActionResourceQueryResult 查询结果
type MenuActionResourceQueryResult struct {
	Data       MenuActionResources
	PageResult *PaginationResult
}

// MenuActionResources 菜单动作关联资源管理列表
type MenuActionResources []*MenuActionResource

// ToMap 转换为map
func (a MenuActionResources) ToMap() map[string]*MenuActionResource {
	m := make(map[string]*MenuActionResource)
	for _, item := range a {
		m[item.Method+item.Path] = item
	}
	return m
}

// ToActionIDMap 转换为动作ID映射
func (a MenuActionResources) ToActionIDMap() map[string]MenuActionResources {
	m := make(map[string]MenuActionResources)
	for _, item := range a {
		m[item.ActionID] = append(m[item.ActionID], item)
	}
	return m
}
