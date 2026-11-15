package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/internal/gen/mainent/enttest"
	"github.com/heromicro/omgind/pkg/config"
	"github.com/heromicro/omgind/pkg/global"
	_ "github.com/mattn/go-sqlite3"
	"net/http/httptest"
	"testing"
)

func TestTenantMembershipAndOnboarding(t *testing.T) {
	previous := global.CFG
	global.CFG = &config.AppConfig{}
	global.CFG.Root.UserName = "root"
	defer func() { global.CFG = previous }()
	client := enttest.Open(t, "sqlite3", "file:tenant-membership?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	ctx := context.Background()
	organ := client.OrgOrgan.Create().SetName("Organization").SaveX(ctx)
	client.OrgStaff.Create().SetOrgID(organ.ID).SetUserID("alice").SaveX(ctx)
	for _, test := range []struct {
		name, user, org, path, method string
		want                          int
	}{
		{"member", "alice", organ.ID, "/api/v2/org-depts", "GET", 200},
		{"nonmember", "bob", organ.ID, "/api/v2/org-depts", "GET", 403},
		{"global operation without selection", "alice", "", "/api/v2/sys-districts", "GET", 200},
		{"menu without selection", "alice", "", "/api/v2/pub/current/menutree", "GET", 200},
		{"no selection", "alice", "", "/api/v2/org-depts", "GET", 403},
		{"root needs selection", "root", "", "/api/v2/org-depts", "GET", 403},
		{"list own organizations", "alice", "", "/api/v2/pub/current/organs", "GET", 200},
		{"create first organization", "alice", "", "/api/v2/org-organs", "POST", 200},
	} {
		t.Run(test.name, func(t *testing.T) {
			r := gin.New()
			r.Use(func(c *gin.Context) { ginx.SetUserID(c, test.user); c.Next() })
			r.Use(TenantMiddleware(client))
			r.Any(test.path, func(c *gin.Context) { c.Status(200) })
			req := httptest.NewRequest(test.method, test.path, nil)
			req.Header.Set("X-Org-ID", test.org)
			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)
			if rec.Code != test.want {
				t.Fatalf("status=%d want=%d", rec.Code, test.want)
			}
		})
	}
}
