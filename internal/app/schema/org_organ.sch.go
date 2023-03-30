package schema

// OrgOrgan 组织管理对象

type OrgOrganShow struct {
	ID     string  `json:"id"`                       // 唯一标识
	Name   string  `json:"name" binding:"required"`  // 名称
	Sname  string  `json:"sname" binding:"required"` // 短名称
	IdenNo *string `json:"iden_no"`                  // 执照号

	Haddr *SysAddress `json:"haddr"` // 总部地址
}

type OrgOrgan struct {
	TimeMixin `yaml:"-"`
	OrgOrganShow
	Code string `json:"code"` // 助记码

	OwnerId string `json:"owner_id"` // 所有者user.id

	Memo     *string `json:"memo"`                         //
	IsActive *bool   `json:"is_active" binding:"required"` // 状态

	Sort int `json:"sort,omitempty"`

	Creator string `json:"creator"` // 创建者

}

// OrgOrganQueryParam 查询条件
type OrgOrganQueryParam struct {
	PaginationParam

	QueryValue string  `form:"queryValue" json:"queryValue"` // 模糊查询
	Code       *string `form:"code" json:"code"`             //
	Name       string  `form:"name" json:"name"`             //
	IsActive   *bool   `form:"is_active" json:"is_active"`   //
	OwnerId    string  `form:"owner_id" json:"owner_id"`     //

	CreatedAt_Order string `form:"created_at__order" json:"created_at__order"` // asc, desc
	IsActive_Order  string `form:"is_active__order" json:"is_active__order"`   // asc desc
	Sort_Order      string `form:"sort__order" json:"sort__order"`             // asc desc

	// example: asc
	// example: desc
	CountryID_Order  string `form:"country_id__order" json:"country_id__order"`   // asc, desc
	ProvinceID_Order string `form:"province_id__order" json:"province_id__order"` // asc, desc
	CityID_Order     string `form:"city_id__order" json:"city_id__order"`         // asc, desc
	CountyID_Order   string `form:"county_id__order" json:"county_id__order"`     // asc, desc
}

// OrgOrganQueryOptions 查询可选参数项
type OrgOrganQueryOptions struct {
	OrderFields []*OrderField // 排序字段

	FieldsAll      bool     `default:"true"` // all fields
	FieldsIncludes []string // includes fields
	FieldsExcludes []string // exlcudes fields
}

// OrgOrganQueryResult 查询结果
type OrgOrganQueryResult struct {
	Data       OrgOrgans
	PageResult *PaginationResult
}

// OrgOrgans 组织管理列表
type OrgOrgans []*OrgOrgan
