package schema

import "time"

// SysDistrict 行政区域对象
type SysDistrict struct {
	ID         string   `json:"id"`                           // 唯一标识
	ParentID   string   `json:"pid"`                          // pid
	Name       string   `json:"name" binding:"required"`      // 名称
	Sname      string   `json:"sname" binding:"required"`     // 短名称
	Abbr       *string  `json:"abbr"`                         // 简称
	Suffix     *string  `json:"suffix"`                       // 区域后缀,省/市/区/旗/盟/自治区/
	StCode     *string  `json:"st_code"`                      // 统计局区域编码
	Initials   *string  `json:"initials"`                     // 简拼
	Pinyin     *string  `json:"pinyin"`                       // 简拼
	Longitude  *float64 `json:"longitude"`                    // 经度
	Latitude   *float64 `json:"latitude"`                     // 经度
	AreaCode   string   `json:"area_code"`                    // 电话区号码
	ZipCode    *string  `json:"zip_code"`                     // 邮政编码
	MergeName  *string  `json:"merge_name"`                   // 带前缀全名称
	MergeSname *string  `json:"merge_sname"`                  // 带前缀简名称
	Extra      *string  `json:"extra"`                        // 带前缀简名称
	IsActive   *bool    `json:"is_active" binding:"required"` // 状态
	Sort       int32    `json:"sort"`
	IsDel      bool     `json:"is_del"`

	IsMain   *bool `json:"is_main"`   // 主要城市
	IsHot    bool  `json:"is_hot"`    // 热门城市
	IsReal   bool  `json:"is_real"`   // 是否虚拟区域
	IsDirect bool  `json:"is_direct"` // 是否是直辖

	IsLeaf    *bool   `json:"is_leaf"`    // 是否是子叶
	TreeID    *int32  `json:"tree_id"`    // 树id
	TreeLevel *int32  `json:"tree_level"` // 层级
	TreeLeft  *int32  `json:"tree_left"`  // 层级
	TreeRight *int32  `json:"tree_right"` // 层级
	TreePath  *string `json:"tree_path"`  // 层级

	Creator   string     `json:"creator"`    // 创建者
	CreatedAt *time.Time `json:"created_at"` // 创建时间
	UpdatedAt *time.Time `json:"updated_at"` // 更新时间
	// Ids       *int       `json:"ids"`        //
	// Pids      *int       `json:"pids"`       //
}

// SysDistrictQueryParam 查询条件
type SysDistrictQueryParam struct {
	PaginationParam
	QueryValue string `form:"queryValue"` // 模糊查询
	Name       string `form:"name"`       // 名称
	Initials   string `form:"initials"`   // 简拼
	Suffix     string `form:"suffix"`     // 区域后缀,省/市/区/旗/盟/自治区/
	IsMain     *bool  `form:"is_main"`    // 主要城市
	IsHot      *bool  `form:"is_hot"`     // 热门城市
	IsDirect   *bool  `form:"is_direct"`  // 是否是直辖
	IsLeaf     *bool  `form:"is_leaf"`    // 是否是子叶
	TreeID     *int32 `form:"tree_id"`    // 树id

	CreatedAt_Order string `form:"created_at__order"` // asc, desc
	Name_Order      string `form:"name__order"`       // asc desc
}

// SysDistrictQueryOptions 查询可选参数项
type SysDistrictQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// SysDistrictQueryResult 查询结果
type SysDistrictQueryResult struct {
	Data       SysDistricts
	PageResult *PaginationResult
}

// SysDistricts 行政区域列表
type SysDistricts []*SysDistrict
