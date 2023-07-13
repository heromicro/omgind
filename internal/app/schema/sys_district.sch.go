package schema

// SysDistrict 行政区域对象
type SysDistrict struct {
	TimeMixin `yaml:"-"`

	ID       string       `json:"id"`       // 唯一标识
	ParentID *string      `json:"pid"`      // pid
	Parent   *SysDistrict `json:"parent"`   // parent
	Children SysDistricts `json:"children"` // children

	Name       string   `json:"name" binding:"required"`                // 名称
	NameEn     *string  `json:"name_en,omitempty"`                      // 名称[英语]
	Sname      *string  `json:"sname,omitempty"`                        // 短名称[英语]
	SnameEn    *string  `json:"sname_en,omitempty"`                     // 短名称
	Abbr       *string  `json:"abbr,omitempty"`                         // 简称
	Suffix     *string  `json:"suffix,omitempty"`                       // 区域后缀,省/市/区/旗/盟/自治区/
	StCode     *string  `json:"st_code,omitempty"`                      // 统计局区域编码
	Initials   *string  `json:"initials,omitempty"`                     // 简拼
	Pinyin     *string  `json:"pinyin,omitempty"`                       // 简拼
	Longitude  *float64 `json:"longitude,omitempty"`                    // 经度
	Latitude   *float64 `json:"latitude,omitempty"`                     // 经度
	AreaCode   *string  `json:"area_code,omitempty"`                    // 电话区号码
	ZipCode    *string  `json:"zip_code,omitempty"`                     // 邮政编码
	MergeName  *string  `json:"merge_name,omitempty"`                   // 带前缀全名称
	MergeSname *string  `json:"merge_sname,omitempty"`                  // 带前缀简名称
	Extra      *string  `json:"extra,omitempty"`                        // 带前缀简名称
	IsActive   *bool    `json:"is_active,omitempty" binding:"required"` // 状态
	Sort       int32    `json:"sort"`
	IsDel      bool     `json:"is_del"`

	IsMain   *bool `json:"is_main,omitempty"` // 主要城市
	IsHot    bool  `json:"is_hot"`            // 热门城市
	IsReal   bool  `json:"is_real"`           // 是否虚拟区域
	IsDirect bool  `json:"is_direct"`         // 是否是直辖

	TreeMixin

	// Ids       *int       `json:"ids"`        //
	// Pids      *int       `json:"pids"`       //
}

// SysDistrictQueryParam 查询条件
type SysDistrictQueryParam struct {
	PaginationParam

	ParentID   *string `form:"pid" json:"pid"`               // pid
	QueryValue string  `form:"q" json:"q"` // 模糊查询
	Name       string  `form:"name" json:"name"`             // 名称
	NameEN     string  `form:"name_en" json:"name_en"`       // 名称
	Initials   string  `form:"initials" json:"initials"`     // 简拼
	Suffix     string  `form:"suffix" json:"suffix"`         // 区域后缀,省/市/区/旗/盟/自治区/
	IsMain     *bool   `form:"is_main" json:"is_main"`       // 主要城市
	IsHot      *bool   `form:"is_hot" json:"is_hot"`         // 热门城市
	IsDirect   *bool   `form:"is_direct" json:"is_direct"`   // 是否是直辖
	IsLeaf     *bool   `form:"is_leaf" json:"is_leaf"`       // 是否是子叶
	IsActive   *bool   `form:"is_active" json:"is_active"`   // 状态
	IsReal     *bool   `form:"is_real" json:"is_real"`       // 状态

	TreeID    *int64 `form:"tree_id" json:"tree_id"`       // 树id
	TreeLevel *int32 `form:"tree_level" json:"tree_level"` // tree_level

	TreeLeft    *int64 `form:"tree_left" json::"tree_left"`        // tree_left 结束
	TreeLeft_St *int64 `form:"tree_left__st" json:"tree_left__st"` // tree_left 结束
	TreeLeft_Ed *int64 `form:"tree_left__ed" json:"tree_left__ed"` // tree_left 结束

	TreeRight    *int64 `form:"tree_right" json:"tree_right"`         // tree_right 结束
	TreeRight_St *int64 `form:"tree_right__st" json:"tree_right__st"` // tree_right 结束
	TreeRight_Ed *int64 `form:"tree_right__ed" json:"tree_right__ed"` // tree_right 结束

	// example: asc
	// example: desc
	CreatedAt_Order string `form:"created_at__order" json:"created_at__order"` // asc, desc
	Name_Order      string `form:"name__order" json:"name__order"`             // asc desc
	TreeID_Order    string `form:"tree_id__order" json:"tree_id__order"`       // asc desc
	TreeLevel_Order string `form:"tree_level__order" json:"tree_level__order"` // 层级 asc desc
	TreeLeft_Order  string `form:"tree_left__order" json:"tree_left__order"`   // 左值 asc desc

	ParentInitials *string `form:"p_initials" json:"p_initials"`   // parent.initials
	ParentAreaCode *string `form:"p_area_code" json:"p_area_code"` // parent.area_code
	ParentParentID *string `form:"p_pid" json:"p_pid"`             // parent.pid

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

// SysDistrictQueryResult 查询结果
type SysDistrictQueryTreeResult struct {
	Top  SysDistricts `json:"top"`
	Subs SysDistricts `json:"subs"`
}

// SysDistricts 行政区域列表
type SysDistricts []*SysDistrict
