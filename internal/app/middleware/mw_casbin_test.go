package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/heromicro/omgind/internal/app/contextx"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/pkg/config"
	"github.com/heromicro/omgind/pkg/global"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCasbinTenantAuthorization(t *testing.T) {
	previous := global.CFG
	global.CFG = &config.AppConfig{}
	defer func() { global.CFG = previous }()
	global.CFG.Casbin.Enable = true
	global.CFG.Root.UserName = "root"
	enforcer, err := casbin.NewSyncedEnforcer("../../../configs/rbac_model.conf")
	if err != nil {
		t.Fatal(err)
	}
	_, _ = enforcer.AddGroupingPolicy("alice", "editor", "org-a")
	_, _ = enforcer.AddPolicy("editor", "org-a", "/api/v2/things/:id", "GET")
	for _, test := range []struct {
		name, user, tenant, method string
		want                       int
	}{
		{"granted", "alice", "org-a", "GET", 200},
		{"other tenant", "alice", "org-b", "GET", 403},
		{"other method", "alice", "org-a", "DELETE", 403},
		{"no role", "bob", "org-a", "GET", 403},
		{"unauthenticated", "", "org-a", "GET", 401},
	} {
		t.Run(test.name, func(t *testing.T) {
			router := gin.New()
			router.Use(func(c *gin.Context) {
				ginx.SetUserID(c, test.user)
				c.Request = c.Request.WithContext(contextx.WithTenant(c.Request.Context(), test.tenant))
				c.Next()
			})
			router.Use(CasbinMiddleware(enforcer))
			router.Any("/api/v2/things/:id", func(c *gin.Context) { c.Status(200) })
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, httptest.NewRequest(test.method, "/api/v2/things/123", nil))
			if recorder.Code != test.want {
				t.Fatalf("status=%d want=%d", recorder.Code, test.want)
			}
		})
	}
	_, _ = enforcer.RemoveGroupingPolicy("alice", "editor", "org-a")
	allowed, err := enforcer.Enforce("alice", "org-a", "/api/v2/things/123", http.MethodGet)
	if err != nil || allowed {
		t.Fatal("revoked role still allowed")
	}
}
