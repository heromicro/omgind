package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/heromicro/omgind/internal/app"
	"github.com/heromicro/omgind/pkg/global"
)

const (
	configFile = "../../../configs/config.toml"
	modelFile  = "../../../configs/model.conf"
	apiPrefix  = "/api/"
)

var engine *gin.Engine

func init() {
	// 初始化配置文件
	global.MustLoad(configFile)

	global.CFG.System.RunMode = "test"
	global.CFG.Log.Level = 2
	global.CFG.JWTAuth.Enable = false
	global.CFG.Casbin.Enable = false
	global.CFG.Casbin.Model = modelFile
	// global.CFG.Gorm.Debug = false
	// global.CFG.Gorm.DBType = "sqlite3"

	injector, _, err := app.BuildInjector()
	if err != nil {
		panic(err)
	}
	engine = injector.Engine
}

// ResID 响应唯一标识
type ResID struct {
	ID string `json:"id,omitempty"`
}

func toReader(v interface{}) io.Reader {
	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(v)
	return buf
}

func parseReader(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func parseOK(r io.Reader) error {
	var status struct {
		Status string `json:"status"`
	}
	err := parseReader(r, &status)
	if err != nil {
		return err
	}
	if status.Status != "OK" {
		return errors.New("not OK")
	}
	return nil
}

func newPageParam(extra ...map[string]string) map[string]string {
	data := map[string]string{
		"current":  "1",
		"pageSize": "1",
	}

	if len(extra) > 0 {
		for k, v := range extra[0] {
			data[k] = v
		}
	}

	return data
}

type PaginationResult struct {
	Total    int64 `json:"total"`
	Current  int   `json:"current"`
	PageSize int   `json:"pageSize"`
}

type PageResult struct {
	List       interface{}       `json:"list"`
	Pagination *PaginationResult `json:"pagination"`
}

func parsePageReader(r io.Reader, v interface{}) error {
	result := &PageResult{List: v}
	return parseReader(r, result)
}

func newPostRequest(formatRouter string, v interface{}, args ...interface{}) *http.Request {
	req, _ := http.NewRequest("POST", fmt.Sprintf(formatRouter, args...), toReader(v))
	return req
}

func newPutRequest(formatRouter string, v interface{}, args ...interface{}) *http.Request {
	req, _ := http.NewRequest("PUT", fmt.Sprintf(formatRouter, args...), toReader(v))
	return req
}

func newDeleteRequest(formatRouter string, args ...interface{}) *http.Request {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf(formatRouter, args...), nil)
	return req
}

func newGetRequest(formatRouter string, params map[string]string, args ...interface{}) *http.Request {
	values := make(url.Values)
	for k, v := range params {
		values.Set(k, v)
	}

	urlStr := fmt.Sprintf(formatRouter, args...)
	if len(values) > 0 {
		urlStr += "?" + values.Encode()
	}

	req, _ := http.NewRequest("GET", urlStr, nil)
	return req
}
