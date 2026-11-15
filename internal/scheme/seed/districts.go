package seed

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/heromicro/omgind/internal/gen/mainent"
)

// District preserves source identifiers separately from national statistical codes.
type District struct {
	ID          string   `json:"id"`
	ParentID    *string  `json:"parent_id"`
	Name        string   `json:"name"`
	CountryCode string   `json:"country_code"`
	SourceCode  string   `json:"source_code"`
	SourceID    string   `json:"source_id"`
	AdminLevel  int      `json:"admin_level"`
	TreeID      int64    `json:"tree_id"`
	TreeLevel   int32    `json:"tree_level"`
	TreePath    string   `json:"tree_path"`
	TreeLeft    int64    `json:"tree_left"`
	TreeRight   int64    `json:"tree_right"`
	IsLeaf      bool     `json:"is_leaf"`
	Latitude    *float64 `json:"latitude"`
	Longitude   *float64 `json:"longitude"`
}

// ImportDistricts initializes an empty table atomically; existing installations are preserved.
func ImportDistricts(ctx context.Context, client *mainent.Client, path string) (int, error) {
	count, err := client.SysDistrict.Query().Count(ctx)
	if err != nil || count > 0 {
		return 0, err
	}
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	reader, err := gzip.NewReader(file)
	if err != nil {
		return 0, err
	}
	defer reader.Close()
	var rows []District
	if err := json.NewDecoder(reader).Decode(&rows); err != nil {
		return 0, err
	}
	if err := ValidateDistricts(rows); err != nil {
		return 0, err
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	// Recheck inside the transaction; startup should run once per database migration job.
	count, err = tx.SysDistrict.Query().Count(ctx)
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, nil
	}
	const batchSize = 200
	for start := 0; start < len(rows); start += batchSize {
		end := start + batchSize
		if end > len(rows) {
			end = len(rows)
		}
		builders := make([]*mainent.SysDistrictCreate, 0, end-start)
		for index, row := range rows[start:end] {
			builder := tx.SysDistrict.Create().SetID(row.ID).SetName(row.Name).SetAbbr(row.CountryCode).
				SetExtra("geonames:" + row.SourceID).SetIsActive(true).SetIsDel(false).SetIsReal(true).
				SetIsMain(true).SetIsDirect(false).SetIsHot(false).SetSort(int32(start + index)).
				SetTreeID(row.TreeID).SetTreeLevel(row.TreeLevel).SetTreePath(row.TreePath).
				SetTreeLeft(row.TreeLeft).SetTreeRight(row.TreeRight).SetIsLeaf(row.IsLeaf).
				SetNillableParentID(row.ParentID).SetNillableLatitude(row.Latitude).SetNillableLongitude(row.Longitude)
			builders = append(builders, builder)
		}
		if _, err := tx.SysDistrict.CreateBulk(builders...).Save(ctx); err != nil {
			return 0, fmt.Errorf("district batch %d: %w", start, err)
		}
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return len(rows), nil
}

// ValidateDistricts rejects duplicate identifiers, broken parent links and invalid tree intervals.
func ValidateDistricts(rows []District) error {
	if len(rows) == 0 {
		return fmt.Errorf("district seed is empty")
	}
	seen := make(map[string]District, len(rows))
	for _, row := range rows {
		if row.ID == "" || len(row.ID) > 36 || row.Name == "" || len([]rune(row.Name)) > 128 {
			return fmt.Errorf("invalid district %q", row.ID)
		}
		if _, ok := seen[row.ID]; ok {
			return fmt.Errorf("duplicate district %q", row.ID)
		}
		if row.TreeLeft <= 0 || row.TreeRight <= row.TreeLeft || row.TreeLevel < 1 {
			return fmt.Errorf("invalid tree interval for %s", row.ID)
		}
		if row.ParentID != nil {
			parent, ok := seen[*row.ParentID]
			if !ok || parent.TreeID != row.TreeID || parent.TreeLevel+1 != row.TreeLevel || parent.TreeLeft >= row.TreeLeft || parent.TreeRight <= row.TreeRight {
				return fmt.Errorf("invalid parent for %s", row.ID)
			}
		} else if row.TreeLevel != 1 {
			return fmt.Errorf("invalid root %s", row.ID)
		}
		seen[row.ID] = row
	}
	return nil
}
