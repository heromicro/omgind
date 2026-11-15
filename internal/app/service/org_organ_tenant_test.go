package service

import (
	"context"
	"github.com/heromicro/omgind/internal/app/contextx"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent/enttest"
	repo "github.com/heromicro/omgind/internal/scheme/repo/entgo"
	"github.com/heromicro/omgind/pkg/config"
	"github.com/heromicro/omgind/pkg/global"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestCreateOrganizationLinksAuthenticatedOwner(t *testing.T) {
	previous := global.CFG
	global.CFG = &config.AppConfig{}
	defer func() { global.CFG = previous }()
	client := enttest.Open(t, "sqlite3", "file:organization-onboarding?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	srv := &OrgOrgan{EntCli: client, OrgDeptRepo: &repo.OrgDept{EntCli: client}, OrgOrganRepo: &repo.OrgOrgan{EntCli: client}}
	ctx := contextx.NewUserID(context.Background(), "creator")
	active := true
	result, err := srv.Create(ctx, schema.OrgOrgan{OrgOrganShow: schema.OrgOrganShow{Name: "Example organization", Sname: "Example", Haddr: &schema.SysAddress{}}, OwnerId: "spoofed", IsActive: &active})
	if err != nil {
		t.Fatal(err)
	}
	organ, err := client.OrgOrgan.Get(ctx, result.ID)
	if err != nil {
		t.Fatal(err)
	}
	if organ.OwnerID == nil || *organ.OwnerID != "creator" {
		t.Fatal("owner was not bound to authenticated account")
	}
	members, err := client.OrgStaff.Query().All(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(members) != 1 || members[0].UserID == nil || *members[0].UserID != "creator" || members[0].OrgID == nil || *members[0].OrgID != organ.ID {
		t.Fatal("creator membership missing")
	}
}
