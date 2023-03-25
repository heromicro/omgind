package schema

import "time"

// OrgDept 部门管理对象
type OrgDept struct {
	ID       string   `json:"id"`       // 唯一标识
	ParentID *string  `json:"pid"`      // pid
	Parent   *OrgDept `json:"parent"`   // parent
	Children OrgDepts `json:"children"` // children

	Name      string  `json:"name" binding:"required"` // 名称
	Code      string  `json:"code"`                    // 助记码
	MergeName *string `json:"merge_name,omitempty"`    // 带前缀全名称

	Memo     *string `json:"memo"`                         //
	IsActive *bool   `json:"is_active" binding:"required"` // 状态
	Sort     int     `json:"sort,omitempty"`

	OrgMixin
	TreeMixin

	Creator   string     `json:"creator"`    // 创建者
	CreatedAt *time.Time `json:"created_at"` // 创建时间
	UpdatedAt *time.Time `json:"updated_at"` // 更新时间

}

// OrgDeptQueryParam 查询条件
type OrgDeptQueryParam struct {
	PaginationParam

	QueryValue string `form:"queryValue"` // 模糊查询

	Name     string `form:"name" json:"name"`           //
	Code     string `form:"code" json:"code"`           //
	IsActive *bool  `form:"is_active" json:"is_active"` //
	OrgID    string `form:"org_id" json:"org_id"`       //

	ParentID *string `form:"pid" json:"pid"`                      // pid
	IsLeaf   *bool   `form:"is_leaf" json:"is_leaf"`              // 是否是子叶
	IsReal   *bool   `form:"is_real" json:"is_real"`              // 是否真实
	IsShow   *bool   `form:"is_show,default=true" json:"is_show"` // 是否显示

	TreeID    *int64 `form:"tree_id" json:"tree_id"`       // 树id
	TreeLevel *int32 `form:"tree_level" json:"tree_level"` // tree_level

	TreeLeft    *int64 `form:"tree_left" json::"tree_left"`        // tree_left 结束
	TreeLeft_St *int64 `form:"tree_left__st" json:"tree_left__st"` // tree_left 结束
	TreeLeft_Ed *int64 `form:"tree_left__ed" json:"tree_left__ed"` // tree_left 结束

	TreeRight    *int64 `form:"tree_right" json:"tree_right"`         // tree_right 结束
	TreeRight_St *int64 `form:"tree_right__st" json:"tree_right__st"` // tree_right 结束
	TreeRight_Ed *int64 `form:"tree_right__ed" json:"tree_right__ed"` // tree_right 结束

	// example: "asc"
	// example: "desc"
	TreeID_Order    string `form:"tree_id__order" json:"tree_id__order"`       // asc desc
	TreeLevel_Order string `form:"tree_level__order" json:"tree_level__order"` // 层级 asc desc
	TreeLeft_Order  string `form:"tree_left__order" json:"tree_left__order"`   // 左值 asc desc

	// example: "asc"
	// example: "desc"
	Name_Order string `form:"name__order" json:"name__order"` // asc desc
	// example: "asc"
	// example: "desc"
	CreatedAt_Order string `form:"created_at__order" json:"created_at__order"` // asc, desc
	// example: "asc"
	// example: "desc"
	IsActive_Order string `form:"is_active__order" json:"is_active__order"` //  asc desc
	// example: "asc"
	// example: "desc"
	Sort_Order string `form:"sort__order" json:"sort__order"` // asc desc

	ParentParentID *string `form:"p_pid" json:"p_pid"` // parent.pid

}

// OrgDeptQueryOptions 查询可选参数项
type OrgDeptQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// OrgDeptQueryResult 查询结果
type OrgDeptQueryResult struct {
	Data       OrgDepts
	PageResult *PaginationResult
}

// OrgDeptQueryTreeResult 查询结果
type OrgDeptQueryTreeResult struct {
	Top  OrgDepts `json:"top"`
	Subs OrgDepts `json:"subs"`
}

// OrgDepts 部门管理列表
type OrgDepts []*OrgDept
