package api_v2

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/app/service"
	"github.com/heromicro/omgind/internal/gen/mainent/orgorgan"
)

// OrgOrganSet 注入OrgOrgan
var OrgOrganSet = wire.NewSet(wire.Struct(new(OrgOrgan), "*"))

// OrgOrgan 组织管理
type OrgOrgan struct {
	OrgOrganSrv *service.OrgOrgan
}

// Query 查询数据
func (a *OrgOrgan) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.OrgOrganQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResErrorCode(c, -1000, err)
		return
	}

	params.Pagination = true
	result, err := a.OrgOrganSrv.Query(ctx, params)

	if err != nil {
		ginx.ResErrorCode(c, -1010, err)
		return
	}

	ginx.ResPage(c, result.Data, result.PageResult)
}

// QuerySelectPage 查询选择数据
func (a *OrgOrgan) QuerySelectPage(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.OrgOrganQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		// log.Printf(" ------- ------ %+v \n", err)
		ginx.ResErrorCode(c, -1000, err)
		return
	}

	// log.Printf(" ------- ------ %+v \n", params)

	result, err := a.OrgOrganSrv.QuerySelectPage(ctx, params, schema.OrgOrganQueryOptions{
		OrderFields:    schema.NewOrderFields(schema.NewOrderField("id", schema.OrderByDESC)),
		FieldsAll:      false,
		FieldsIncludes: []string{orgorgan.FieldID, orgorgan.FieldName},
	})
	if err != nil {
		ginx.ResErrorCode(c, -1020, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

// Get 查询指定数据
func (a *OrgOrgan) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.OrgOrganSrv.Get(ctx, c.Param("id"))
	if err != nil {
		ginx.ResErrorCode(c, -1030, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// View 查询指定数据
func (a *OrgOrgan) View(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.OrgOrganSrv.View(ctx, c.Param("id"))
	if err != nil {
		ginx.ResErrorCode(c, -1040, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Create 创建数据
func (a *OrgOrgan) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.OrgOrgan
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResErrorCode(c, -1000, err)
		return
	}

	result, err := a.OrgOrganSrv.Create(ctx, item)
	if err != nil {
		ginx.ResErrorCode(c, -1050, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Update 更新数据
func (a *OrgOrgan) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.OrgOrgan
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResErrorCode(c, -1000, err)
		return
	}

	result, err := a.OrgOrganSrv.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginx.ResErrorCode(c, -1060, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Delete 删除数据
func (a *OrgOrgan) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.OrgOrganSrv.Delete(ctx, c.Param("id"))
	if err != nil {
		ginx.ResErrorCode(c, -1070, err)
		return
	}
	ginx.ResOK(c, "成功删除数据")
}

// Enable 启用数据
func (a *OrgOrgan) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.OrgOrganSrv.UpdateActive(ctx, c.Param("id"), true)
	if err != nil {
		ginx.ResErrorCode(c, -1080, err)
		return
	}
	ginx.ResOK(c, "启用成功")
}

// Disable 禁用数据
func (a *OrgOrgan) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.OrgOrganSrv.UpdateActive(ctx, c.Param("id"), false)
	if err != nil {
		ginx.ResErrorCode(c, -1090, err)
		return
	}
	ginx.ResOK(c, "启用成功")
}
