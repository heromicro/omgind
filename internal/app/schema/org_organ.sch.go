package schema

import "time"

// OrgOrgan 组织管理对象

type OrgOrganShow struct {
	ID     string  `json:"id"`                       // 唯一标识
	Name   string  `json:"name" binding:"required"`  // 名称
	Sname  string  `json:"sname" binding:"required"` // 短名称
	IdenNo *string `json:"iden_no"`                  // 执照号

	Haddr *SysAddress `json:"haddr"` // 总部地址
}

type OrgOrgan struct {
	OrgOrganShow
	Code string `json:"code"` // 助记码

	OwnerId string `json:"owner_id"` // 所有者user.id

	Memo     *string `json:"memo"`                         //
	IsActive *bool   `json:"is_active" binding:"required"` // 状态

	Sort int `json:"sort,omitempty"`

	Creator   string     `json:"creator"`    // 创建者
	CreatedAt *time.Time `json:"created_at"` // 创建时间
	UpdatedAt *time.Time `json:"updated_at"` // 更新时间

}

// OrgOrganQueryParam 查询条件
type OrgOrganQueryParam struct {
	PaginationParam
	QueryValue string  `form:"queryValue"` // 模糊查询
	Code       *string `form:"code"`       //
	Name       string  `form:"name"`       //
	IsActive   *bool   `form:"is_active"`  //
	OwnerId    string  `form:"owner_id"`   //

	CreatedAt_Order string `form:"created_at__order"` // asc, desc
	IsActive_Order  string `form:"is_active__order"`  // asc desc
	Sort_Order      string `form:"sort__order"`       // asc desc

	CountryID_Order  string `form:"country_id__order"`  // asc, desc
	ProvinceID_Order string `form:"province_id__order"` // asc, desc
	CityID_Order     string `form:"city_id__order"`     // asc, desc
	CountyID_Order   string `form:"county_id__order"`   // asc, desc
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
