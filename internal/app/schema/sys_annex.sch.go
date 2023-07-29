package schema

import (
	"time"

	"github.com/heromicro/omgind/pkg/helper/json"
)

// SysAnnex 对象
type SysAnnex struct {
	ID        string     `json:"id"`                           // 唯一标识
	Name      string     `json:"name" binding:"required"`      // 文件名称
	FilePath  string     `json:"file_path" binding:"required"` // 文件路径
	CreatedAt *time.Time `json:"created_at"`                   // 创建时间
	UpdatedAt *time.Time `json:"updated_at"`                   // 更新时间

}

func (a *SysAnnex) String() string {
	return json.MarshalToString(a)
}

// SysAnnexQueryParam 查询条件
type SysAnnexQueryParam struct {
	PaginationParam
	Q string `form:"q"` // 模糊查询

	BasicOrderParam
	TimeOrderParam
}

// SysAnnexQueryOptions 查询可选参数项
type SysAnnexQueryOptions struct {
	OrderFields []*OrderField // 排序字段

	FieldsAll      bool     // all fields
	FieldsIncludes []string // includes fields
	FieldsExcludes []string // exlcudes fields
}

// SysAnnexQueryResult 查询结果
type SysAnnexQueryResult struct {
	Data       SysAnnexes
	PageResult *PaginationResult
}

// SysAnnexes 列表
type SysAnnexes []*SysAnnex
