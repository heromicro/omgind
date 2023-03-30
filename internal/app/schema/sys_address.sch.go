package schema

// SysAddress 地址管理对象
type SysAddress struct {
	TimeMixin `yaml:"-"`

	ID string `json:"id"` // 唯一标识

	Country  string `json:"country"`  // 国
	Province string `json:"province"` // 省
	City     string `json:"city"`     // 区/市
	County   string `json:"county"`   // 区/县

	CountryID  string `json:"country_id"`  // 国ID
	ProvinceID string `json:"province_id"` // 省/市ID
	CityID     string `json:"city_id"`     // 区/市ID
	CountyID   string `json:"county_id"`   // 区/县ID

	ZipCode   string  `json:"zip_code,omitempty"`   // 邮政编码
	Daddr     string  `json:"daddr"`                // 详细地址
	LastName  string  `json:"last_name,omitempty"`  // 联系人姓
	FirstName string  `json:"first_name,omitempty"` // 联系人名
	AreaCode  *string `json:"area_code,omitempty"`  // 电话区号码
	Mobile    string  `json:"mobile,omitempty"`     // 电话

	IsActive *bool  `json:"is_active,omitempty"` // 状态
	Sort     int    `json:"sort,omitempty"`
	Creator  string `json:"creator,omitempty"` // 创建者

}

// SysAddressQueryParam 查询条件
type SysAddressQueryParam struct {
	PaginationParam
	QueryValue string `form:"queryValue"` // 模糊查询

	CountryID  string  `form:"country_id" json:"country_id"`   //
	ProvinceID string  `form:"province_id" json:"province_id"` //
	CityID     string  `form:"city_id" json:"city_id"`         //
	CountyID   string  `form:"county_id" json:"county_id"`     //
	AreaCode   *string `form:"area_code" json:"area_code"`
	Mobile     *string `form:"mobile" json:"mobile"`

	// example: asc
	// example: desc
	CreatedAt_Order  string `form:"created_at__order" json:"created_at__order"`   // asc, desc
	ID_Order         string `form:"id__order" json:"id__order"`                   // asc, desc
	CountryID_Order  string `form:"country_id__order" json:"country_id__order"`   // asc, desc
	ProvinceID_Order string `form:"province_id__order" json:"province_id__order"` // asc, desc
	CityID_Order     string `form:"city_id__order" json:"city_id__order"`         // asc, desc
	CountyID_Order   string `form:"county_id__order" json:"county_id__order"`     // asc, desc

}

// SysAddressQueryOptions 查询可选参数项
type SysAddressQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// SysAddressQueryResult 查询结果
type SysAddressQueryResult struct {
	Data       SysAddresses
	PageResult *PaginationResult
}

// SysAddresses 地址管理列表
type SysAddresses []*SysAddress
