package dump

import (
	"context"
	"log"
	"math"
	"os"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/sysmenu"
	"github.com/heromicro/omgind/internal/gen/mainent/sysmenuaction"
	"github.com/heromicro/omgind/internal/gen/mainent/sysmenuactionresource"
	"github.com/heromicro/omgind/internal/scheme/repo"
	"github.com/ttacon/chalk"
	"gopkg.in/yaml.v2"
)

func Dump_menu(ctx context.Context, eclient *mainent.Client, datafile string) error {

	redOnWhite := chalk.Red.NewStyle().WithBackground(chalk.White)
	cyanOnBlue := chalk.Cyan.NewStyle().WithBackground(chalk.Blue)
	// greenOnWhite := chalk.Green.NewStyle().WithBackground(chalk.White)
	// whiteOnGreen := chalk.Cyan.NewStyle().WithBackground(chalk.Green)

	fp, err := os.OpenFile(datafile, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Println(redOnWhite, " can not open file: ", err, chalk.Reset)
		return err
	}
	defer fp.Close()

	query := eclient.SysMenu.Query()
	total, err := query.Count(ctx)
	if err != nil {
		return err
	}

	log.Println(cyanOnBlue, " total: ", total, chalk.Reset)

	pageSize := 10
	page := int(math.Ceil(float64(total) / float64(pageSize)))
	log.Println(cyanOnBlue, " page: ", page, chalk.Reset)

	for i := 0; i < page; i++ {
		offset := i * pageSize
		if offset > total {
			log.Println(" ------ ======= ", offset, " ======= ", total)
			break
		}

		query := eclient.SysMenu.Query()

		query = query.Order(mainent.Asc(sysmenu.FieldLevel), mainent.Asc(sysmenu.FieldSort))

		r_menus, err := query.Limit(pageSize).Offset(offset).All(ctx)
		if err != nil {
			return err
		}
		var sch_menus []*schema.Menu

		for _, menu := range r_menus {

			r_actions, err := eclient.SysMenuAction.Query().Where(sysmenuaction.MenuIDEQ(menu.ID)).Order(mainent.Asc(sysmenuaction.FieldID)).All(ctx)
			if err != nil {
				return err
			}

			var sch_actions []*schema.MenuAction

			for _, action := range r_actions {
				r_resourceses, err := eclient.SysMenuActionResource.Query().Where(sysmenuactionresource.ActionID(action.ID)).Order(mainent.Asc(sysmenuactionresource.FieldID)).All(ctx)
				if err != nil {
					return err
				}
				sch_resourceses := repo.ToSchemaSysMenuActionResources(r_resourceses)
				sch_action := repo.ToSchemaSysMenuAction(action)
				sch_action.Resources = sch_resourceses
				sch_actions = append(sch_actions, sch_action)
			}

			sch_menu := repo.ToSchemaSysMenu(menu)
			sch_menu.Actions = sch_actions

			sch_menus = append(sch_menus, sch_menu)

		}

		// log.Println(" -------- ====== ", sch_dicts[0])
		data, err := yaml.Marshal(sch_menus)
		if err != nil {
			return err
		}

		wc, err := fp.Write(data)
		if err != nil {
			log.Println(redOnWhite, " faild to write ", err, chalk.Reset)
			return err
		}

		log.Println(cyanOnBlue, "write : ", wc, chalk.Reset)
	}

	return nil
}
