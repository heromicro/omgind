package migrate

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/gotidy/ptr"
	"github.com/jszwec/csvutil"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
	"github.com/xuri/excelize/v2"
	"gopkg.in/yaml.v2"

	"github.com/heromicro/omgind/cmd/omgind/common"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/sysdict"
	"github.com/heromicro/omgind/internal/gen/ent/sysdictitem"
	"github.com/heromicro/omgind/internal/gen/ent/sysdistrict"
	"github.com/heromicro/omgind/internal/gen/ent/sysmenu"
	"github.com/heromicro/omgind/internal/gen/ent/sysmenuaction"
	"github.com/heromicro/omgind/internal/gen/ent/sysmenuactionresource"
	"github.com/heromicro/omgind/internal/schema/repo"
)

func init() {

	// PersistentFlags() for all flags only in current cmd

	CmdLoad.PersistentFlags().String("datafile", "", "data file path")
	CmdLoad.MarkPersistentFlagRequired("datafile")

}

var CmdLoad = &cobra.Command{
	Use:   "load",
	Short: "load file to db data",
	RunE: func(cmd *cobra.Command, args []string) error {

		log.Println(" ----- starting of load ====== ")

		// greenOnWhite := chalk.Green.NewStyle().WithBackground(chalk.White)
		redOnWhite := chalk.Red.NewStyle().WithBackground(chalk.White)
		// cyanOnBlue := chalk.Cyan.NewStyle().WithBackground(chalk.Blue)
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
			log.Println(redOnWhite, " ----- === === ", err, chalk.Reset)
			return err
		}

		_, err = os.Stat(datafile)
		if err != nil && os.IsNotExist(err) {
			return fmt.Errorf("%s not exist", datafile)
		}

		format, err := cmd.Flags().GetString("format")
		if err != nil {
			log.Println(redOnWhite, " ----- === === ", err, chalk.Reset)
			return err
		}

		tablename, err := cmd.Flags().GetString("tablename")
		if err != nil {
			log.Println(redOnWhite, " tablename err: ", err, chalk.Reset)
			return err
		}

		ctx := context.Background()

		switch strings.ToLower(format) {
		case "xlsx":
			excf, err := excelize.OpenFile(datafile)
			if err != nil {
				log.Println("failed to open ", datafile)
				return err
			}
			defer excf.Close()

			// excf.GetRows()
		case "yml":
			fallthrough
		case "yaml":
			switch tablename {
			case "dict":
				load_dict_data(ctx, eclient, datafile)

			case "menu":
				load_menu_data(ctx, eclient, datafile)

			default:
				log.Println(whiteOnGreen, " missing on tablename: ", tablename, chalk.Reset)
			}
		case "csv":

			switch tablename {
			case "district":
				load_district_data(ctx, eclient, datafile)

			default:
				log.Println(whiteOnGreen, " missing on tablename: ", tablename, chalk.Reset)
			}

		default:
			log.Println(whiteOnGreen, " missing on format: ", format, chalk.Reset)
			return errors.New(fmt.Sprintf("%s not supported", format))

		}

		log.Println(" ----- ending of load ====== ")
		return nil
	},
}

func load_district_data(ctx context.Context, eclient *ent.Client, datafile string) error {

	greenOnWhite := chalk.Green.NewStyle().WithBackground(chalk.White)
	redOnWhite := chalk.Red.NewStyle().WithBackground(chalk.White)
	cyanOnBlue := chalk.Cyan.NewStyle().WithBackground(chalk.Blue)
	whiteOnGreen := chalk.Cyan.NewStyle().WithBackground(chalk.Green)

	bytes, err := ioutil.ReadFile(datafile)
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

		var pdistrict *ent.SysDistrict = nil

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

		create_district_input := repoSysDistrict.ToEntCreateSysDistrictInput(&district)

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

func load_dict_data(ctx context.Context, eclient *ent.Client, filename string) error {

	greenOnWhite := chalk.Blue.NewStyle().WithBackground(chalk.White)
	redOnWhite := chalk.Red.NewStyle().WithBackground(chalk.White)
	cyanOnBlue := chalk.Cyan.NewStyle().WithBackground(chalk.Blue)
	whiteOnGreen := chalk.Cyan.NewStyle().WithBackground(chalk.Green)
	cyanOnYellow := chalk.White.NewStyle().WithBackground(chalk.Magenta)

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(redOnWhite, " ----- === === ", err)
		return err
	}
	var dicts []*schema.Dict

	err = yaml.Unmarshal(bytes, &dicts)
	if err != nil {
		return err
	}

	repoDict := repo.Dict{
		EntCli: eclient,
	}

	repoDictItem := repo.DictItem{
		EntCli: eclient,
	}

	var dictBulk []*ent.SysDictCreate
	var dictItemBulk []*ent.SysDictItemCreate

	err = repo.WithTx(ctx, eclient, func(tx *ent.Tx) error {

		for _, d := range dicts {
			for _, di := range d.Items {
				log.Println(cyanOnYellow, " di :: ", di.JSONString(), chalk.Reset)
				dictItemInput := repoDictItem.ToEntCreateSysDictItemInput(di)
				log.Println(greenOnWhite, " di :: ", dictItemInput.DictID, chalk.Reset)

				create_dictitem := tx.SysDictItem.Create().SetID(di.ID).SetInput(*dictItemInput)

				log.Println(greenOnWhite, " di :: ", create_dictitem, chalk.Reset)

				dictItemBulk = append(dictItemBulk, create_dictitem)
			}

			log.Println(cyanOnYellow, " di :: ", d.JSONString(), chalk.Reset)

			d.Items = nil
			dictInput := repoDict.ToEntCreateSysDictInput(d)

			create_dict := tx.SysDict.Create().SetID(d.ID).SetInput(*dictInput)
			dictBulk = append(dictBulk, create_dict)
		}

		err = tx.SysDict.CreateBulk(dictBulk...).OnConflict(
			sql.ConflictColumns(sysdict.FieldID),
		).UpdateNewValues().Exec(ctx)
		if err != nil {
			log.Println(greenOnWhite, " creatbulk dict ", err, chalk.Reset)
			return err
		}
		log.Println(greenOnWhite, " finish creatbulk dict ", chalk.Reset)

		err = tx.SysDictItem.CreateBulk(dictItemBulk...).OnConflict(
			sql.ConflictColumns(sysdictitem.FieldID),
		).UpdateNewValues().Exec(ctx)
		if err != nil {
			log.Println(cyanOnBlue, " creatbulk dict item ", err, chalk.Reset)

			return err
		}

		return nil
	})

	if err != nil {
		log.Println(whiteOnGreen, " creatbulk ", err, chalk.Reset)
		return err
	}

	return nil
}

func load_menu_data(ctx context.Context, eclient *ent.Client, filename string) error {

	greenOnWhite := chalk.Green.NewStyle().WithBackground(chalk.White)
	redOnWhite := chalk.Red.NewStyle().WithBackground(chalk.White)
	cyanOnBlue := chalk.Cyan.NewStyle().WithBackground(chalk.Blue)
	whiteOnGreen := chalk.Cyan.NewStyle().WithBackground(chalk.Green)

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(redOnWhite, " ----- === === ", err)
		return err
	}
	var menus []*schema.Menu

	err = yaml.Unmarshal(bytes, &menus)
	if err != nil {
		return err
	}

	repoMenu := &repo.Menu{
		EntCli: eclient,
	}

	repoAction := &repo.MenuAction{
		EntCli: eclient,
	}

	repoResource := &repo.MenuActionResource{
		EntCli: eclient,
	}

	var menuBulk []*ent.SysMenuCreate
	var actionBulk []*ent.SysMenuActionCreate
	var resourceBulk []*ent.SysMenuActionResourceCreate

	err = repo.WithTx(ctx, eclient, func(tx *ent.Tx) error {

		for _, m := range menus {

			for _, a := range m.Actions {

				for _, r := range a.Resources {
					resourceInput := repoResource.ToEntCreateSysMenuActionResourceInput(r)
					create_resoure := tx.SysMenuActionResource.Create().SetID(r.ID).SetInput(*resourceInput)
					resourceBulk = append(resourceBulk, create_resoure)
				}
				a.Resources = nil

				actionInput := repoAction.ToEntCreateSysMenuActionInput(a)
				create_action := tx.SysMenuAction.Create().SetID(a.ID).SetInput(*actionInput)
				actionBulk = append(actionBulk, create_action)
			}
			m.Actions = nil
			if m.ParentID != nil && *m.ParentID == "" {
				m.ParentID = nil
			}

			menuInput := repoMenu.ToEntCreateSysMenuInput(m)
			create_menu := tx.SysMenu.Create().SetID(m.ID).SetInput(*menuInput)
			menuBulk = append(menuBulk, create_menu)
		}

		err = tx.SysMenu.CreateBulk(menuBulk...).OnConflict(
			sql.ConflictColumns(sysmenu.FieldID),
		).UpdateNewValues().Exec(ctx)
		if err != nil {
			log.Println(cyanOnBlue, " creatbulk menu ", err, chalk.Reset)
			return err
		}
		log.Println(greenOnWhite, " finish creatbulk menu ")

		err = tx.SysMenuAction.CreateBulk(actionBulk...).OnConflict(
			sql.ConflictColumns(sysmenuaction.FieldID),
		).UpdateNewValues().Exec(ctx)
		if err != nil {
			log.Println(greenOnWhite, " creatbulk menu action ", err, chalk.Reset)
			return err
		}
		log.Println(greenOnWhite, " finish creatbulk action ")

		err = tx.SysMenuActionResource.CreateBulk(resourceBulk...).OnConflict(
			sql.ConflictColumns(sysmenuactionresource.FieldID),
		).UpdateNewValues().Exec(ctx)
		if err != nil {
			log.Println(cyanOnBlue, " creatbulk menu action resource ", err, chalk.Reset)
			return err
		}
		log.Println(greenOnWhite, " finish creatbulk action resource ")

		return nil
	})

	if err != nil {
		log.Println(whiteOnGreen, " creatbulk ", err, chalk.Reset)
		return err
	}

	return nil
}

// go run cmd/omgind/main.go migrate load --conf configs/config.toml --format=csv --tablename=districts --datafile=./scripts/sql/district_full.csv
// go run cmd/omgind/main.go migrate load --conf configs/config.toml --format=yaml --tablename=menu --datafile=./configs/menu.2023-03-28_64619.yaml
