package scheme

import (
	"github.com/heromicro/omgind/pkg/config"
	"github.com/heromicro/omgind/pkg/config/option"
	"path/filepath"
	"testing"
)

func TestGormSQLiteClient(t *testing.T) {
	cfg := &config.AppConfig{Gorm: option.GormConfig{DBType: "sqlite3", MaxIdleConns: 1}, Sqlite3: option.Sqlite3Config{Path: filepath.Join(t.TempDir(), "nested", "test.db")}}
	db, cleanup, err := NewGormClient(cfg)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()
	if err := db.Exec("CREATE TABLE factory_probe (id INTEGER PRIMARY KEY)").Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Exec("INSERT INTO factory_probe (id) VALUES (1)").Error; err != nil {
		t.Fatal(err)
	}
}

func TestClientConfigValidation(t *testing.T) {
	if _, cleanup, err := NewEntClient(nil); err == nil {
		t.Fatal("expected Ent config error")
	} else {
		cleanup()
	}
	if _, cleanup, err := NewGormClient(&config.AppConfig{}); err == nil {
		t.Fatal("expected GORM driver error")
	} else {
		cleanup()
	}
	if _, cleanup, err := NewMongoClient(&config.AppConfig{}); err == nil {
		t.Fatal("expected Mongo config error")
	} else {
		cleanup()
	}
}
