package api_v2

import (
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/app/service"
)

// SysDistrictSet 注入SysDistrict
var SysDistrictSet = wire.NewSet(wire.Struct(new(SysDistrict), "*"))

// SysDistrict 行政区域
type SysDistrict struct {
	SysDistrictSrv *service.SysDistrict
}

// Query 查询数据
func (a *SysDistrict) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.SysDistrictQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResErrorCode(c, -1000, err)
		return
	}

	params.Pagination = true
	result, err := a.SysDistrictSrv.Query(ctx, params)
	if err != nil {
		ginx.ResErrorCode(c, -1010, err)
		return
	}

	ginx.ResPage(c, result.Data, result.PageResult)
}

// Get 查询指定数据
func (a *SysDistrict) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.SysDistrictSrv.Get(ctx, c.Param("id"))
	if err != nil {
		ginx.ResErrorCode(c, -1020, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// View 查询指定数据详情
func (a *SysDistrict) View(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.SysDistrictSrv.View(ctx, c.Param("id"))
	if err != nil {
		ginx.ResErrorCode(c, -1030, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Get 所有的行政区
func (a *SysDistrict) GetAllSubs(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.SysDistrictQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResErrorCode(c, -1000, err)
		return
	}
	var pid string = c.Param("id")
	params.Current = 1
	params.PageSize = 1000

	log.Println(" --- ---- === ==== pid ", pid)
	log.Println(" ----- -- ==== === URL ", c.Request.URL)

	result, err := a.SysDistrictSrv.GetAllSubs(ctx, pid, params)
	if err != nil {
		ginx.ResErrorCode(c, -1015, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

// QueryTree 查询菜单树
func (a *SysDistrict) QueryTree(c *gin.Context) {

	ctx := c.Request.Context()
	var params schema.SysDistrictQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResErrorCode(c, -1000, err)
		return
	}
	var pid string = c.Param("id")

	if params.ParentID == nil {
		// params.ParentID = ptr.String("")
	}

	params.PageSize = 200

	result, err := a.SysDistrictSrv.GetTree(ctx, pid, params)

	if err != nil {
		ginx.ResErrorCode(c, -1018, err)
		return
	}

	ginx.ResSuccess(c, result)
}

// Create 创建数据
func (a *SysDistrict) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.SysDistrict
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResErrorCode(c, -1000, err)
		return
	}

	result, err := a.SysDistrictSrv.Create(ctx, item)
	if err != nil {
		// log.Println(" -------- ==== district create error ", err)
		ginx.ResErrorCode(c, -1040, err)
		return
	}

	ginx.ResSuccess(c, result)
}

// Update 更新数据
func (a *SysDistrict) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.SysDistrict
	body, _ := io.ReadAll(c.Request.Body)
	log.Println(" ---- --- body ", string(body))
	// err := json.Unmarshal(body, &item)
	err := ginx.ParseJSON(c, &item)
	log.Println(" ---- --- err ", err)
	if err != nil {
		ginx.ResErrorCode(c, -1000, err)
		return
	}
	result, err := a.SysDistrictSrv.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginx.ResErrorCode(c, -1050, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Delete 删除数据
func (a *SysDistrict) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SysDistrictSrv.Delete(ctx, c.Param("id"))
	if err != nil {
		ginx.ResErrorCode(c, -1060, err)
		return
	}
	ginx.ResOK(c, "成功删除数据")
}

// Enable 启用数据
func (a *SysDistrict) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SysDistrictSrv.UpdateActive(ctx, c.Param("id"), true)
	if err != nil {
		ginx.ResErrorCode(c, -1070, err)
		return
	}
	ginx.ResOK(c, "启用成功")
}

// Disable 禁用数据
func (a *SysDistrict) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SysDistrictSrv.UpdateActive(ctx, c.Param("id"), false)
	if err != nil {
		ginx.ResErrorCode(c, -1080, err)
		return
	}
	ginx.ResOK(c, "禁用成功")
}
