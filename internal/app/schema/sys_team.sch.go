package schema

import (
	"github.com/heromicro/omgind/pkg/helper/json"
)

// SysTeam 地址管理对象
type SysTeam struct {
	TimeMixin `yaml:"-"`

	ID   string `json:"id"`   // 唯一标识
	Name string `json:"name"` // 联系人
	Code string `json:"code"` // 电话

}

func (a *SysTeam) String() string {
	return json.MarshalToString(a)
}

// SysTeamQueryParam 查询条件
type SysTeamQueryParam struct {
	PaginationParam
	QueryValue string `form:"queryValue"` // 模糊查询
}

// SysTeamQueryOptions 查询可选参数项
type SysTeamQueryOptions struct {
	OrderFields []*OrderField // 排序字段

	FieldsAll      bool     // all fields
	FieldsIncludes []string // includes fields
	FieldsExcludes []string // exlcudes fields
}

// SysTeamQueryResult 查询结果
type SysTeamQueryResult struct {
	Data       SysTeams
	PageResult *PaginationResult
}

// SysTeams 地址管理列表
type SysTeams []*SysTeam
