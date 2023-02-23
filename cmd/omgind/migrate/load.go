package migrate

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/gotidy/ptr"
	"github.com/heromicro/omgind/cmd/omgind/common"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/sysdistrict"
	"github.com/heromicro/omgind/internal/schema/repo"
	"github.com/jszwec/csvutil"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
	"github.com/xuri/excelize/v2"
)

var CmdLoad = &cobra.Command{
	Use:   "load",
	Short: "load file to db data",
	RunE: func(cmd *cobra.Command, args []string) error {

		log.Println(" ----- starting of load ====== ")

		greenOnWhite := chalk.Green.NewStyle().WithBackground(chalk.White)
		redOnWhite := chalk.Red.NewStyle().WithBackground(chalk.White)
		cyanOnBlue := chalk.Cyan.NewStyle().WithBackground(chalk.Blue)
		whiteOnGreen := chalk.Cyan.NewStyle().WithBackground(chalk.Green)

		configFile, err := cmd.Flags().GetString("conf")
		eclient, cleanup, err := common.MakeEntClient(configFile)

		if err != nil {
			log.Println(redOnWhite, " eclient : ", eclient, chalk.Reset)
			return err
		}
		defer cleanup()

		datafile, err := cmd.Flags().GetString("datafile")
		if err != nil {
			return err
		}

		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return err
		}

		switch strings.ToLower(format) {
		case "xlsx":
			excf, err := excelize.OpenFile(datafile)
			if err != nil {
				log.Println("failed to open ", datafile)
				return err
			}
			defer excf.Close()

			// excf.GetRows()

		case "csv":

			bytes, err := ioutil.ReadFile(datafile)
			if err != nil {

				return err
			}

			var districtRaws []DistrictRaw
			if err := csvutil.Unmarshal(bytes, &districtRaws); err != nil {
				log.Println("=== -- error:", err)
				return err
			}

			// Pids        int64   `csv:"pids,omitempty"`
			// NeedCustoms bool    `csv:"need_customs"`     // 需要通关
			// TreePath    string  `csv:"t_path,omitempty"` // 树形父级id拼接
			serviceSysDistrict := repo.SysDistrict{
				EntCli: eclient,
			}
			ctx := context.Background()

			// total := 0

			for _, item := range districtRaws {

				if item.IsDel {
					log.Println(cyanOnBlue, " === --== skilp item ", item, chalk.Reset)
					continue
				}

				log.Println(greenOnWhite, " ======== -------- item ", item, chalk.Reset)

				// if total > 2 {
				// 	break
				// }
				// total += 1

				district := schema.SysDistrict{
					TreeID:     &item.TreeID,
					TreeLeft:   &item.TreeLeft,
					TreeRight:  &item.TreeRight,
					TreeLevel:  &item.TreeLevel,
					Name:       item.Name,
					Initials:   &item.Initials,
					Pinyin:     &item.Pinyin,
					Longitude:  &item.Longitude,
					Latitude:   &item.Latitude,
					ZipCode:    &item.ZipCode,
					StCode:     &item.StCode,
					MergeName:  &item.MergeName,
					MergeSname: &item.MergeSname,
					Extra:      &item.Extra,
					Sort:       item.Sort,
					IsDel:      item.IsDel,
					IsDirect:   item.IsDirect,
					Suffix:     &item.Suffix,
					Sname:      item.Sname,
					IsHot:      item.IsHot,
					IsMain:     &item.Mpoint,
					IsReal:     item.IsReal,
				}

				if item.Status == 1 {
					district.IsActive = ptr.Bool(true)
				}

				// log.Println(cyanOnBlue, "===--  : ", item.Ids, " ", *district.MergeSname, " ", item.MergeSname, chalk.Reset)

				var pdistrict *ent.SysDistrict = nil

				if item.TreeLevel > 1 {
					if item.ParentID != nil && *item.ParentID != "" {
						query_district := serviceSysDistrict.EntCli.SysDistrict.Query()
						query_district = query_district.Where(sysdistrict.IDEQ(*item.ParentID))
						pdistrict, _ = query_district.First(ctx)
						if pdistrict != nil {
							district.ParentID = pdistrict.ID
							if pdistrict.TreePath == nil {
								district.TreePath = &pdistrict.ID
							} else {
								district.TreePath = ptr.String(strings.Join([]string{*pdistrict.TreePath, pdistrict.ID}, "/"))
							}
							// pdistrict.Update().SetIsLeaf(false).Save(ctx)
						}
					}
				}

				create_district := serviceSysDistrict.EntCli.SysDistrict.Create()
				create_district_input := serviceSysDistrict.ToEntCreateSysDistrictInput(&district)
				log.Println(" ----- ===== --- create_district_input.ParentID : [", *create_district_input.ParentID, "] ")

				if item.ParentID == nil || *item.ParentID == "" {
					create_district_input.ParentID = nil
				}

				create_district = create_district.SetInput(*create_district_input)

				create_district = create_district.SetIsLeaf(true)

				// left value and right value `
				// https://blog.csdn.net/yilovexing/article/details/107066591`
				if pdistrict == nil {
					log.Println(" --- ---- ===== create_district ")
					// 无父级节点, 左值:1, 右值:2
					create_district = create_district.SetTreeLeft(1).SetTreeRight(2)

					// create_district_input.TreeLeft = ptr.Int64(1)
					// create_district_input.TreeRight = ptr.Int64(*create_district_input.TreeLeft + 1)
					create_district = create_district.SetTreeLevel(1)

					log.Println(" --- ---- ===== create_district ", create_district)
				} else {
					// 有父级节点, 左值:pid.right, 右值: pid.right + 1
					create_district = create_district.SetTreeLeft(*pdistrict.TreeRight).SetTreeRight(*pdistrict.TreeRight + 1)
					create_district = create_district.SetTreeLevel(*pdistrict.TreeLevel + 1)

					pdistrict.Update().SetIsLeaf(false).Save(ctx)

				}

				create_district = create_district.SetID(item.ID)

				log.Println(" --------- ==== create_district, ", create_district)
				sch_district, err := create_district.Save(ctx)
				if sch_district != nil && pdistrict != nil {
					// // 修复被破坏平衡的其他节点的左值。大于 parent_id 右值的所有节点的左值加 2。

					log.Println(whiteOnGreen, " ======  sch_district.ID  ====== ", sch_district.ID)
					update_district_l := serviceSysDistrict.EntCli.SysDistrict.Update()
					update_district_l = update_district_l.Where(sysdistrict.IDNEQ(sch_district.ID))

					update_district_l = update_district_l.Where(sysdistrict.TreeLeftGT(*pdistrict.TreeRight))

					var count int = 0
					count, err = update_district_l.AddTreeLeft(2).Save(ctx)
					log.Println(whiteOnGreen, " ====== count ====== ", count)
					log.Println(whiteOnGreen, " ====== err ====== ", err)

					// // 修复被破坏平衡的其他节点的右值。大于等于 parent_id 右值的所有节点的右值加 2
					update_district_r := serviceSysDistrict.EntCli.SysDistrict.Update()
					update_district_r = update_district_r.Where(sysdistrict.IDNEQ(sch_district.ID))

					update_district_r = update_district_r.Where(sysdistrict.TreeRightGTE(*pdistrict.TreeRight))

					count, err = update_district_r.AddTreeRight(2).Save(ctx)
					log.Println(whiteOnGreen, " ====== count ====== ", count)
					log.Println(whiteOnGreen, " uuuu ====== err ====== ", err)
				}

				if err != nil {
					log.Println(redOnWhite, " ====== save failed ====== ", district)
					log.Println(cyanOnBlue, " ====== save failed ====== ", err)

					continue
				}

				log.Println(cyanOnBlue, " ====== -------- ====== ", sch_district.ID, " ", *sch_district.MergeName, " ", *sch_district.MergeSname)
			}

			// eclient.SysDistrict.CreateBulk(districts...)

			// log.Println(" ======== -------- ====== ", districts)

		default:
			return errors.New(fmt.Sprintf("%s not supported", format))
		}

		log.Println(" ----- ending of load ====== ")
		return nil
	},
}

func init() {

	// PersistentFlags() for all flags only in current cmd
	Cmd.PersistentFlags().String("format", "", "format")
	Cmd.MarkPersistentFlagRequired("format")

	Cmd.PersistentFlags().String("datafile", "", "data file path")
	Cmd.MarkPersistentFlagRequired("datafile")

	Cmd.PersistentFlags().String("tablename", "", "table name")
	Cmd.MarkPersistentFlagRequired("tablename")

}

// go run cmd/omgind/main.go migrate load --conf configs/config.toml --format=csv --tablename=districts --datafile=./scripts/sql/district_full.cvs
