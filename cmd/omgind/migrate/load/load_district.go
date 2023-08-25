package load

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/gotidy/ptr"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/sysdistrict"
	"github.com/heromicro/omgind/internal/scheme/repo"
	"github.com/jszwec/csvutil"
	"github.com/ttacon/chalk"
)

func Load_district_data(ctx context.Context, eclient *mainent.Client, datafile string) error {

	greenOnWhite := chalk.Green.NewStyle().WithBackground(chalk.White)
	redOnWhite := chalk.Red.NewStyle().WithBackground(chalk.White)
	cyanOnBlue := chalk.Cyan.NewStyle().WithBackground(chalk.Blue)
	whiteOnGreen := chalk.Cyan.NewStyle().WithBackground(chalk.Green)

	bytes, err := os.ReadFile(datafile)
	if err != nil {
		log.Println(redOnWhite, " ----- === === ", err)
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
	repoSysDistrict := repo.SysDistrict{
		EntCli: eclient,
	}

	// total := 0

	// log.Println(" ----------- --- 222222 ")
	// log.Println(" ----------- --- 222222 districtRaws ", districtRaws)

	for _, item := range districtRaws {

		if item.IsDel {
			log.Println(cyanOnBlue, " === --== skilp item ", item, chalk.Reset)
			continue
		}

		log.Println(cyanOnBlue, " ==== ==== ======= =========  ", chalk.Reset)
		log.Println(cyanOnBlue, " ==== ==== -------- item.name ", item.Name, chalk.Reset)
		log.Println(cyanOnBlue, " ==== ==== ======= =========  ", chalk.Reset)

		// if total > 2 {
		// 	break
		// }
		// total += 1

		district := schema.SysDistrict{
			// TreeID: &item.TreeID,
			// TreeLeft: &item.TreeLeft,
			// TreeRight:  &item.TreeRight,
			// TreeLevel:  &item.TreeLevel,
			Name:       item.Name,
			NameEn:     item.NameEn,
			Initials:   &item.Initials,
			Pinyin:     &item.Pinyin,
			Longitude:  &item.Longitude,
			Latitude:   &item.Latitude,
			ZipCode:    &item.ZipCode,
			AreaCode:   &item.AreaCode,
			StCode:     &item.StCode,
			MergeName:  &item.MergeName,
			MergeSname: &item.MergeSname,
			Extra:      &item.Extra,
			Sort:       item.Sort,
			IsDel:      item.IsDel,
			IsDirect:   item.IsDirect,
			Suffix:     &item.Suffix,
			Sname:      item.Sname,
			SnameEn:    item.SnameEn,
			IsHot:      item.IsHot,
			IsMain:     &item.Mpoint,
			IsReal:     item.IsReal,
			ParentID:   item.ParentID,
		}
		district.TreeID = &item.TreeID

		if item.Status == 1 {
			district.IsActive = ptr.Bool(true)
		}

		// log.Println(cyanOnBlue, "===--  : ", item.Ids, " ", *district.MergeSname, " ", item.MergeSname, chalk.Reset)

		var pdistrict *mainent.SysDistrict = nil

		if item.TreeLevel > 1 {
			log.Println(" ---------- ======== ", item.ParentID)

			if item.ParentID != nil && *item.ParentID != "" {

				query_district := repoSysDistrict.EntCli.SysDistrict.Query()
				query_district = query_district.Where(sysdistrict.IDEQ(*item.ParentID))
				pdistrict, _ = query_district.First(ctx)
				if pdistrict != nil {
					district.ParentID = &pdistrict.ID
					if pdistrict.TreePath == nil {
						district.TreePath = &pdistrict.ID
					} else {
						district.TreePath = ptr.String(strings.Join([]string{*pdistrict.TreePath, pdistrict.ID}, "/"))
					}
					// pdistrict.Update().SetIsLeaf(false).Save(ctx)
				}
			}
		}

		create_district := repoSysDistrict.EntCli.SysDistrict.Create()

		create_district_input := repo.ToEntCreateSysDistrictInput(&district)

		// log.Println(" ----- ===== --- create_district_input : [", create_district_input, "] ")
		// log.Println(" ----- ===== --- create_district_input.ParentID : [", create_district_input.ParentID, "] ")
		log.Println(" ----- ===== --- item.ParentID : [", item.ParentID, "] ")
		log.Println(" ----- ===== --- item.ParentID : [", item.ParentID, "] ")

		if item.ParentID == nil || *item.ParentID == "" {
			create_district_input.ParentID = nil
		} else {
			create_district_input.ParentID = item.ParentID
		}
		// log.Println(" ----- ===== --- create_district_input.ParentID : [", create_district_input.ParentID, "] ")

		if item.AreaCode == "" {
			create_district_input.AreaCode = nil
		}
		if item.ZipCode == "" {
			create_district_input.ZipCode = nil
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
			log.Println(" ----- ===== --- 1111 pdistrict.ID : [", pdistrict.ID, "] ")

			create_district = create_district.SetTreeLeft(*pdistrict.TreeRight).SetTreeRight(*pdistrict.TreeRight + 1)
			create_district = create_district.SetTreeLevel(*pdistrict.TreeLevel + 1)

			log.Println(" ----- ===== --- 2222 pdistrict.ID : [", pdistrict.TreeRight, "] ")

		}

		create_district = create_district.SetID(item.ID)

		log.Println(" --------- ==== create_district, ", create_district)
		sch_district, err := create_district.Save(ctx)
		if err != nil {
			log.Println(redOnWhite, " ====== save failed ====== ", district.TreeLevel)
			log.Println(redOnWhite, " ====== save failed ====== ", district.ParentID)
			log.Println(cyanOnBlue, " ====== save failed ====== ", err)
			break
		}

		if pdistrict != nil {

			log.Println(greenOnWhite, " ====------ - ---  pdistrict.TreeRight ", pdistrict.TreeRight, chalk.Reset)
			pdistrict, err = pdistrict.Update().SetIsLeaf(false).Save(ctx)
			// log.Println(greenOnWhite, " ====------ - --- update pdistrict ", err, chalk.Reset)

			if err != nil {
				log.Println(redOnWhite, " ====------ - --- ", err, chalk.Reset)
			}
		}

		if sch_district != nil && pdistrict != nil {
			// // 修复被破坏平衡的其他节点的左值。大于 parent_id 右值的所有节点的左值加 2。

			log.Println(whiteOnGreen, " ======  sch_district.ID  ====== ", sch_district.ID)
			update_district_l := repoSysDistrict.EntCli.SysDistrict.Update()
			update_district_l = update_district_l.Where(sysdistrict.IDNEQ(sch_district.ID))
			update_district_l = update_district_l.Where(sysdistrict.TreeIDEQ(*pdistrict.TreeID))
			update_district_l = update_district_l.Where(sysdistrict.TreeLeftGT(*pdistrict.TreeRight))

			var count int = 0
			count, err = update_district_l.AddTreeLeft(2).Save(ctx)
			log.Println(whiteOnGreen, " ====== count ====== ", count)
			log.Println(whiteOnGreen, " ====== err ====== ", err)

			// // 修复被破坏平衡的其他节点的右值。大于等于 parent_id 右值的所有节点的右值加 2
			update_district_r := repoSysDistrict.EntCli.SysDistrict.Update()
			update_district_r = update_district_r.Where(sysdistrict.IDNEQ(sch_district.ID))
			update_district_r = update_district_r.Where(sysdistrict.TreeIDEQ(*pdistrict.TreeID))
			update_district_r = update_district_r.Where(sysdistrict.TreeRightGTE(*pdistrict.TreeRight))

			count, err = update_district_r.AddTreeRight(2).Save(ctx)
			log.Println(whiteOnGreen, " ====== count ====== ", count)
			log.Println(whiteOnGreen, " uuuu ====== err ====== ", err)
		}

		log.Println(cyanOnBlue, " ====== -------- ====== ", sch_district.ID, " ", *sch_district.MergeName, " ", *sch_district.MergeSname)
	}

	// eclient.SysDistrict.CreateBulk(districts...)

	// log.Println(" == ====== -------- ====== ", districts)

	return nil
}
