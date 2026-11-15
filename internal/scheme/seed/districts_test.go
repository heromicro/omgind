package seed

import (
	"compress/gzip"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/heromicro/omgind/internal/gen/mainent/enttest"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
	"testing"
)

func TestDistrictImportIsIdempotent(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:district-seed?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	parent := "gn:1"
	rows := []District{
		{ID: parent, Name: "Country", CountryCode: "US", SourceID: "1", TreeID: 1, TreeLevel: 1, TreeLeft: 1, TreeRight: 4},
		{ID: "gn:2", ParentID: &parent, Name: "County", CountryCode: "US", SourceID: "2", TreeID: 1, TreeLevel: 2, TreeLeft: 2, TreeRight: 3, IsLeaf: true},
	}
	file, err := os.Create(filepath.Join(t.TempDir(), "seed.json.gz"))
	if err != nil {
		t.Fatal(err)
	}
	compressed := gzip.NewWriter(file)
	if err := json.NewEncoder(compressed).Encode(rows); err != nil {
		t.Fatal(err)
	}
	if err := compressed.Close(); err != nil {
		t.Fatal(err)
	}
	file.Close()
	count, err := ImportDistricts(context.Background(), client, file.Name())
	if err != nil || count != 2 {
		t.Fatalf("count=%d err=%v", count, err)
	}
	count, err = ImportDistricts(context.Background(), client, file.Name())
	if err != nil || count != 0 {
		t.Fatalf("repeat count=%d err=%v", count, err)
	}
	child, err := client.SysDistrict.Get(context.Background(), "gn:2")
	if err != nil {
		t.Fatal(err)
	}
	if child.ParentID == nil || *child.ParentID != parent || child.TreeLevel == nil || *child.TreeLevel != 2 {
		t.Fatal("tree metadata was not imported")
	}
}

func TestDistrictValidationRejectsBrokenTree(t *testing.T) {
	missing := "missing"
	if err := ValidateDistricts([]District{{ID: "a", Name: "Orphan", ParentID: &missing, TreeID: 1, TreeLevel: 2, TreeLeft: 1, TreeRight: 2}}); err == nil {
		t.Fatal("orphan accepted")
	}
}

func TestCollectedSnapshot(t *testing.T) {
	path := "../../../configs/data/districts/world.json.gz"
	manifestBytes, err := os.ReadFile("../../../configs/data/districts/manifest.json")
	if err != nil {
		t.Fatal(err)
	}
	var manifest struct {
		SHA256  string `json:"sha256"`
		Records int    `json:"records"`
	}
	if err := json.Unmarshal(manifestBytes, &manifest); err != nil {
		t.Fatal(err)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if fmt.Sprintf("%x", sha256.Sum256(content)) != manifest.SHA256 {
		t.Fatal("snapshot checksum mismatch")
	}
	file, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	reader, err := gzip.NewReader(file)
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()
	var rows []District
	if err := json.NewDecoder(reader).Decode(&rows); err != nil {
		t.Fatal(err)
	}
	if err := ValidateDistricts(rows); err != nil {
		t.Fatal(err)
	}
	if len(rows) != manifest.Records {
		t.Fatal("manifest record count mismatch")
	}
	roots := 0
	counties := map[string]int{}
	for _, row := range rows {
		if row.ParentID == nil {
			roots++
		}
		if row.AdminLevel == 3 {
			counties[row.CountryCode]++
		}
	}
	if roots < 240 || len(rows) < 30000 || counties["CN"] < 2500 || counties["DE"] < 300 {
		t.Fatalf("incomplete snapshot: roots=%d records=%d ADM3=%v", roots, len(rows), counties)
	}
}

func TestCollectedSnapshotImport(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:world-district-import?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	count, err := ImportDistricts(context.Background(), client, "../../../configs/data/districts/world.json.gz")
	if err != nil {
		t.Fatal(err)
	}
	actual, err := client.SysDistrict.Query().Count(context.Background())
	if err != nil || actual != count || count < 30000 {
		t.Fatalf("imported=%d actual=%d error=%v", count, actual, err)
	}
}
