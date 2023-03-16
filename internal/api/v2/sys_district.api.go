package api_v2

import (
	"encoding/json"
	"io/ioutil"
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
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.SysDistrictSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResPage2(c, result.Data, result.PageResult)
}

// Get 查询指定数据
func (a *SysDistrict) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.SysDistrictSrv.Get(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess2(c, item)
}

// View 查询指定数据详情
func (a *SysDistrict) View(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.SysDistrictSrv.View(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess2(c, item)
}

// Get 所有的行政区
func (a *SysDistrict) GetAllSubDistricts(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.SysDistrictQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}
	var pid string = c.Param("id")
	params.Current = 1
	params.PageSize = 1000

	log.Println(" --- ---- === ==== pid ", pid)
	log.Println(" ----- -- ==== === URL ", c.Request.URL)

	result, err := a.SysDistrictSrv.GetAllSubDistricts(ctx, pid, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage2(c, result.Data, result.PageResult)
}

// QueryTree 查询菜单树
func (a *SysDistrict) QueryTree(c *gin.Context) {

	ctx := c.Request.Context()
	var params schema.SysDistrictQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}
	var pid string = c.Param("id")

	if params.ParentID == nil {
		// params.ParentID = ptr.String("")
	}

	params.PageSize = 100

	result, err := a.SysDistrictSrv.GetTree(ctx, pid, params)

	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResList3(c, result)
}

// Create 创建数据
func (a *SysDistrict) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.SysDistrict
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	item.Creator = ginx.GetUserID(c)
	result, err := a.SysDistrictSrv.Create(ctx, item)
	if err != nil {
		// log.Println(" -------- ==== district create error ", err)
		ginx.ResError(c, err)
		return
	}

	ginx.ResSuccess2(c, result)
}

// Update 更新数据
func (a *SysDistrict) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.SysDistrict
	body, _ := ioutil.ReadAll(c.Request.Body)
	log.Println(" ---- --- body ", string(body))
	err := json.Unmarshal(body, &item)
	// err := ginx.ParseJSON(c, &item)
	log.Println(" ---- --- err ", err)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	result, err := a.SysDistrictSrv.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess2(c, result)
}

// Delete 删除数据
func (a *SysDistrict) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SysDistrictSrv.Delete(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK2(c, "成功删除数据")
}

// Enable 启用数据
func (a *SysDistrict) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SysDistrictSrv.UpdateActive(ctx, c.Param("id"), true)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK2(c, "启用成功")
}

// Disable 禁用数据
func (a *SysDistrict) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SysDistrictSrv.UpdateActive(ctx, c.Param("id"), false)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK2(c, "禁用成功")
}
