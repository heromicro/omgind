package migrate

// tree_id,tree_left,tree_right,tree_level,ids,name,initials,pinyin,longitude,latitude,area_code,zip_code,stcode,mname,msname,extra,sort,status,is_del,pids,updated_time,created_time,suffix,sname,is_hot,mpoint,need_customs,is_r,is_d,t_path,crtd_at,uptd_at

type DistrictRaw struct {
	ID          string  `csv:"id"`
	ParentID    *string `csv:"pid"`
	TreeID      int32   `csv:"tree_id"`
	TreeLeft    int64   `csv:"tree_left"`
	TreeRight   int64   `csv:"tree_right"`
	TreeLevel   int32   `csv:"tree_level"`
	IsLeaf      bool    `csv:"is_leaf"`
	Name        string  `csv:"name"`
	Sname       string  `csv:"sname,omitempty"`
	Abbr        string  `csv:"abbr,omitempty"`
	Initials    string  `csv:"initials,omitempty"`
	Pinyin      string  `csv:"pinyin,omitempty"`
	Longitude   float64 `csv:"longitude,omitempty"`
	Latitude    float64 `csv:"latitude,omitempty"`
	AreaCode    string  `csv:"area_code,omitempty"`
	ZipCode     string  `csv:"zip_code,omitempty"`
	StCode      string  `csv:"stcode,omitempty"`
	MergeName   string  `csv:"mname,omitempty"`
	MergeSname  string  `csv:"msname,omitempty"`
	Extra       string  `csv:"extra,omitempty"`
	Sort        int32   `csv:"sort,omitempty"`
	IsActive    bool    `csv:"is_active,omitempty"`
	IsDel       bool    `csv:"is_del,omitempty"`
	IsMain      bool    `csv:"is_m,omitempty"`
	Suffix      string  `csv:"suffix,omitempty"`
	IsHot       bool    `csv:"is_hot"`
	IsDirect    bool    `csv:"is_d"`
	Status      int32   `csv:"status"`           // status,1:有效
	Mpoint      bool    `csv:"mpoint,omitempty"` // 主要城市
	NeedCustoms bool    `csv:"need_customs"`     // 需要通关
	IsReal      bool    `csv:"is_r"`             // 是否实际区域
	TreePath    string  `csv:"t_path,omitempty"` // 树形父级id拼接
	// Ids         int     `csv:"ids,omitempty"`    // for old
	// Pids        int     `csv:"pids,omitempty"`   // for old

}
