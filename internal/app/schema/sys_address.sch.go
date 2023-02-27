package schema

import "time"

// SysAddress 地址管理对象
type SysAddress struct {
	ID string `json:"id"` // 唯一标识

	Country string `json:"country"` // 国
	Provice string `json:"provice"` // 省
	City    string `json:"city"`    // 区/市
	County  string `json:"county"`  // 区/县

	CountryID string `json:"country_id"` // 国ID
	ProviceID string `json:"provice_id"` // 省/市ID
	CityID    string `json:"city_id"`    // 区/市ID
	CountyID  string `json:"county_id"`  // 区/县ID

	ZipCode string `json:"zip_code"` // 邮政编码
	Daddr   string `json:"daddr"`    // 详细地址
	Name    string `json:"name"`     // 联系人
	Mobile  string `json:"mobile"`   // 电话

	Sort      int        `json:"sort,omitempty"`
	Creator   string     `json:"creator"`    // 创建者
	CreatedAt *time.Time `json:"created_at"` // 创建时间
	UpdatedAt *time.Time `json:"updated_at"` // 更新时间

}

// SysAddressQueryParam 查询条件
type SysAddressQueryParam struct {
	PaginationParam
	QueryValue string `form:"queryValue"` // 模糊查询

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
