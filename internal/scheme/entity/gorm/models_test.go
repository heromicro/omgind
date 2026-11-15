package entity

import (
	"encoding/json"
	"github.com/heromicro/omgind/internal/gen/mainent/migrate"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strings"
	"testing"
)

func TestModelsMatchEntColumns(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	models := []interface{}{&OrgDept{}, &OrgOrgan{}, &OrgPosition{}, &OrgStaff{}, &SysAddress{}, &SysAnnex{}, &SysDict{}, &SysDictItem{}, &SysDistrict{}, &SysJwtBlock{}, &SysLogging{}, &SysMenu{}, &SysMenuAction{}, &SysMenuActionResource{}, &SysRole{}, &SysRoleMenu{}, &SysTeam{}, &SysTeamUser{}, &SysUser{}, &SysUserRole{}}
	for _, model := range models {
		scope := db.NewScope(model)
		table := scope.TableName()
		t.Run(table, func(t *testing.T) {
			if err := db.AutoMigrate(model).Error; err != nil {
				t.Fatal(err)
			}
			found := false
			for _, expected := range migrate.Tables {
				if expected.Name != table {
					continue
				}
				found = true
				for _, column := range expected.Columns {
					if !db.Dialect().HasColumn(table, column.Name) {
						t.Errorf("missing column %s", column.Name)
					}
				}
			}
			if !found {
				t.Fatal("missing Ent table")
			}
		})
	}
}

func TestUserPasswordIsNotSerialized(t *testing.T) {
	data, err := json.Marshal(SysUser{Password: "secret-value"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(data), "secret-value") {
		t.Fatal("password exposed in JSON")
	}
}
