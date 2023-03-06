package schema

import "time"

// SysDistrict 行政区域对象
type SysDistrict struct {
	ID         string       `json:"id"`                           // 唯一标识
	ParentID   *string      `json:"pid"`                          // pid
	Parent     *SysDistrict `json:"parent"`                       // parent
	Name       string       `json:"name" binding:"required"`      // 名称
	NameEn     *string      `json:"name_en" `                     // 名称[英语]
	Sname      *string      `json:"sname"`                        // 短名称[英语]
	SnameEn    *string      `json:"sname_en"`                     // 短名称
	Abbr       *string      `json:"abbr"`                         // 简称
	Suffix     *string      `json:"suffix"`                       // 区域后缀,省/市/区/旗/盟/自治区/
	StCode     *string      `json:"st_code"`                      // 统计局区域编码
	Initials   *string      `json:"initials"`                     // 简拼
	Pinyin     *string      `json:"pinyin"`                       // 简拼
	Longitude  *float64     `json:"longitude"`                    // 经度
	Latitude   *float64     `json:"latitude"`                     // 经度
	AreaCode   *string      `json:"area_code"`                    // 电话区号码
	ZipCode    *string      `json:"zip_code"`                     // 邮政编码
	MergeName  *string      `json:"merge_name"`                   // 带前缀全名称
	MergeSname *string      `json:"merge_sname"`                  // 带前缀简名称
	Extra      *string      `json:"extra"`                        // 带前缀简名称
	IsActive   *bool        `json:"is_active" binding:"required"` // 状态
	Sort       int32        `json:"sort"`
	IsDel      bool         `json:"is_del"`

	IsMain   *bool `json:"is_main"`   // 主要城市
	IsHot    bool  `json:"is_hot"`    // 热门城市
	IsReal   bool  `json:"is_real"`   // 是否虚拟区域
	IsDirect bool  `json:"is_direct"` // 是否是直辖

	IsLeaf    *bool   `json:"isLeaf"`     // 是否是子叶
	TreeID    *int32  `json:"tree_id"`    // 树id
	TreeLevel *int32  `json:"tree_level"` // 层级
	TreeLeft  *int64  `json:"tree_left"`  // 层级
	TreeRight *int64  `json:"tree_right"` // 层级
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

	ParentID   *string `form:"pid"`        // pid
	QueryValue string  `form:"queryValue"` // 模糊查询
	Name       string  `form:"name"`       // 名称
	NameEN     string  `form:"name_en"`    // 名称
	Initials   string  `form:"initials"`   // 简拼
	Suffix     string  `form:"suffix"`     // 区域后缀,省/市/区/旗/盟/自治区/
	IsMain     *bool   `form:"is_main"`    // 主要城市
	IsHot      *bool   `form:"is_hot"`     // 热门城市
	IsDirect   *bool   `form:"is_direct"`  // 是否是直辖
	IsLeaf     *bool   `form:"isLeaf"`     // 是否是子叶
	IsActive   *bool   `form:"is_active"`  // 状态
	IsReal     *bool   `form:"is_real"`    // 状态

	TreeID    *int32 `form:"tree_id"`    // 树id
	TreeLevel *int32 `form:"tree_level"` // tree_level

	TreeLeft    *int64 `form:"tree_left"`     // tree_left 结束
	TreeLeft_St *int64 `form:"tree_left__st"` // tree_left 结束
	TreeLeft_Ed *int64 `form:"tree_left__ed"` // tree_left 结束

	TreeRight    *int64 `form:"tree_right"`     // tree_right 结束
	TreeRight_St *int64 `form:"tree_right__st"` // tree_right 结束
	TreeRight_Ed *int64 `form:"tree_right__ed"` // tree_right 结束

	CreatedAt_Order string `form:"created_at__order"` // asc, desc
	Name_Order      string `form:"name__order"`       // asc desc
	TreeID_Order    string `form:"tree_id__order"`    // asc desc
	TreeLevel_Order string `form:"tree_level__order"` // 层级 asc desc
	TreeLeft_Order  string `form:"tree_left__order"`  // 左值 asc desc

	ParentInitials *string `form:"p_initials"`  // parent.initials
	ParentAreaCode *string `form:"p_area_code"` // parent.area_code
	ParentParentID *string `form:"p_pid"`       // parent.pid

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
