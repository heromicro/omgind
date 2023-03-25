package schema

// TreeMixin 部门管理对象
type TreeMixin struct {
	IsLeaf    *bool   `json:"is_leaf,omitempty"`    // 是否是子叶
	TreeID    *int64  `json:"tree_id,omitempty"`    // 树id
	TreeLevel *int32  `json:"tree_level,omitempty"` // 层级
	TreeLeft  *int64  `json:"tree_left,omitempty"`  // 层级
	TreeRight *int64  `json:"tree_right,omitempty"` // 层级
	TreePath  *string `json:"tree_path,omitempty"`  // 层级

}

// OrgMixin 部门管理对象
type OrgMixin struct {
	OrgID *string       `json:"org_id" binding:"required"` // 企业id
	Org   *OrgOrganShow `json:"org"`                       //
}
