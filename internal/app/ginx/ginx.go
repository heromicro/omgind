package ginx

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gotidy/ptr"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/json"
	"github.com/heromicro/omgind/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// 定义上下文中的键
const (
	prefix           = "omgind"
	UserIDKey        = prefix + "/user-id"
	ReqBodyKey       = prefix + "/req-body"
	ResBodyKey       = prefix + "/res-body"
	LoggerReqBodyKey = prefix + "/logger-req-body"
)

// GetToken 获取用户令牌
func GetToken(c *gin.Context) string {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

// GetUserID 获取用户ID
func GetUserID(c *gin.Context) string {
	return c.GetString(UserIDKey)
}

// SetUserID 设定用户ID
func SetUserID(c *gin.Context, userID string) {
	c.Set(UserIDKey, userID)
}

// GetBody Get request body
func GetBody(c *gin.Context) []byte {
	if v, ok := c.Get(ReqBodyKey); ok {
		if b, ok := v.([]byte); ok {
			return b
		}
	}
	return nil
}

// ParseJSON 解析请求JSON
func ParseJSON(c *gin.Context, obj any) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		fmt.Printf(" 00000 ---- %+v \n", err)
		return errors.Wrap400Response(err, fmt.Sprintf("解析请求参数发生错误 - %s", err.Error()))
	}
	return nil
}

// ParseQuery 解析Query参数
func ParseQuery(c *gin.Context, obj any) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("解析请求参数发生错误 - %s", err.Error()))
	}
	return nil
}

// ParseForm 解析Form请求
func ParseForm(c *gin.Context, obj any) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("解析请求参数发生错误 - %s", err.Error()))
	}
	return nil
}

// ResOK 响应OK
func ResOK(c *gin.Context, message string) {
	ResSuccess(c, nil, message)
}

// ResList 响应列表数据
func ResList(c *gin.Context, v any) {
	ResSuccess(c, schema.ListResult{List: v})
}

// ResPage 响应分页数据
func ResPage(c *gin.Context, v any, pr *schema.PaginationResult) {
	list := schema.ListResult{
		List:       v,
		Pagination: pr,
	}
	ResSuccess(c, list)
}

// ResSuccess 响应成功
func ResSuccess(c *gin.Context, v any, message ...string) {
	msg := ""
	if len(message) > 0 {
		msg = message[0]
	}
	if v != nil {
		ResJSON(c, http.StatusOK, schema.StatusResult{Code: schema.CodeOK, Message: msg, Burden: v})
	} else {
		ResJSON(c, http.StatusOK, schema.StatusResult{Code: schema.CodeOK, Message: msg})
	}
}

// ResCodeMessageBurden 响应成功
func ResCodeMessageBurden(c *gin.Context, code int, message string, burden any) {
	ResJSON(c, http.StatusOK, schema.CodeMessageBurden{Code: code, Message: message, Burden: burden})
}

// ResSuccess 响应成功
func ResCodeMessage(c *gin.Context, code int, message string) {
	ResJSON(c, http.StatusOK, schema.CodeMessageBurden{Code: code, Message: message, Burden: nil})
}

// ResJSON 响应JSON数据
func ResJSON(c *gin.Context, status int, v any) {
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	c.Set(ResBodyKey, buf)
	c.Data(status, "application/json; charset=utf-8", buf)
	c.Abort()
}

// ResError 响应错误
func ResError(c *gin.Context, err error, status ...int) {
	ctx := c.Request.Context()
	var res *errors.ResponseError

	if err != nil {
		if e, ok := err.(*errors.ResponseError); ok {
			res = e
		} else {
			res = errors.UnWrapResponse(errors.ErrInternalServer)
			res.ERR = err
		}
	} else {
		res = errors.UnWrapResponse(errors.ErrInternalServer)
	}

	if len(status) > 0 {
		res.StatusCode = &status[0]
	}

	if err := res.ERR; err != nil {
		if res.Message == "" {
			res.Message = err.Error()
		}

		if status := res.StatusCode; *status >= 400 && *status < 500 {
			logger.WithContext(ctx).Warnf(err.Error())
		} else if *status >= 500 {
			logger.WithContext(logger.NewStackContext(ctx, err)).Errorf(err.Error())
		}
	}

	eitem := schema.ErrorItem{
		Code:    res.Code,
		Message: res.Message,
	}
	ResJSON(c, *res.StatusCode, schema.ErrorResult{Error: eitem})
}

func ResErrorCode(c *gin.Context, code int, err error, status ...int) {
	ctx := c.Request.Context()
	var res *errors.ResponseError

	if err != nil {
		if e, ok := err.(*errors.ResponseError); ok {
			res = e
		} else {
			res = errors.UnWrapResponse(errors.ErrInternalServer)
			res.ERR = err
		}
	} else {
		res = errors.UnWrapResponse(errors.ErrInternalServer)
	}

	if len(status) > 0 {
		res.StatusCode = &status[0]
	} else {
		res.StatusCode = ptr.Int(http.StatusOK)
	}

	if err != nil {
		if res.Message == "" {
			res.Message = err.Error()
		}
		if status := res.StatusCode; status != nil {
			if *status >= 400 && *status < 500 {
				logger.WithContext(ctx).Warnf(err.Error())
			} else {
				logger.WithContext(logger.NewStackContext(ctx, err)).Errorf(err.Error())
			}
		}
	}

	logger.WithContext(ctx).Warnf(err.Error())
	ResJSON(c, *res.StatusCode, schema.StatusResult{Code: schema.CodeEnum(code), Message: res.Message})
}
