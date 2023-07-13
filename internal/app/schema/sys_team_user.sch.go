package schema

import (
	"github.com/heromicro/omgind/pkg/helper/json"
)

// SysTeamUser 地址管理对象
type SysTeamUser struct {
	TimeMixin `yaml:"-"`

	ID     string `json:"id"`                         // 唯一标识
	TeamID string `json:"team_id" binding:"required"` // systeam.id
	UserID string `json:"user_id" binding:"required"` // sysuser.id

}

func (a *SysTeamUser) String() string {
	return json.MarshalToString(a)
}

// SysTeamUserQueryParam 查询条件
type SysTeamUserQueryParam struct {
	PaginationParam
	QueryValue string `form:"q"` // 模糊查询
}

// SysTeamUserQueryOptions 查询可选参数项
type SysTeamUserQueryOptions struct {
	OrderFields []*OrderField // 排序字段

	FieldsAll      bool     // all fields
	FieldsIncludes []string // includes fields
	FieldsExcludes []string // exlcudes fields
}

// SysTeamUserQueryResult 查询结果
type SysTeamUserQueryResult struct {
	Data       SysTeamUsers
	PageResult *PaginationResult
}

// SysTeamUsers 地址管理列表
type SysTeamUsers []*SysTeamUser
