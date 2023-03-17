package schema

import "time"

// OrgOrgan 组织管理对象
type OrgOrgan struct {
	ID        string     `json:"id"`                          // 唯一标识
	Name      string     `json:"name" binding:"required"`     // 名称
	Sname     string     `json:"sname" binding:"required"`    // 短名称
	OwnerID   string     `json:"owner_id" binding:"required"` // 所有者id
	Creator   string     `json:"creator"`                     // 创建者
	CreatedAt *time.Time `json:"created_at"`                  // 创建时间
	UpdatedAt *time.Time `json:"updated_at"`                  // 更新时间

}

// OrgOrganQueryParam 查询条件
type OrgOrganQueryParam struct {
	PaginationParam
	QueryValue string `form:"queryValue"` // 模糊查询
}

// OrgOrganQueryOptions 查询可选参数项
type OrgOrganQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// OrgOrganQueryResult 查询结果
type OrgOrganQueryResult struct {
	Data       OrgOrgans
	PageResult *PaginationResult
}

// OrgOrgans 组织管理列表
type OrgOrgans []*OrgOrgan
