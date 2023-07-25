package load

import (
	"context"
	"io/ioutil"
	"log"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/sysmenu"
	"github.com/heromicro/omgind/internal/gen/mainent/sysmenuaction"
	"github.com/heromicro/omgind/internal/gen/mainent/sysmenuactionresource"
	"github.com/heromicro/omgind/internal/scheme/repo"
	"github.com/ttacon/chalk"
	"gopkg.in/yaml.v2"
)

func Load_menu_data(ctx context.Context, eclient *mainent.Client, filename string) error {

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

	var menuBulk []*mainent.SysMenuCreate
	var actionBulk []*mainent.SysMenuActionCreate
	var resourceBulk []*mainent.SysMenuActionResourceCreate

	err = repo.WithTx(ctx, eclient, func(tx *mainent.Tx) error {

		for _, m := range menus {

			for _, a := range m.Actions {

				for _, r := range a.Resources {
					resourceInput := repo.ToEntCreateSysMenuActionResourceInput(r)
					create_resoure := tx.SysMenuActionResource.Create().SetID(r.ID).SetInput(*resourceInput)
					resourceBulk = append(resourceBulk, create_resoure)
				}
				a.Resources = nil

				actionInput := repo.ToEntCreateSysMenuActionInput(a)
				create_action := tx.SysMenuAction.Create().SetID(a.ID).SetInput(*actionInput)
				actionBulk = append(actionBulk, create_action)
			}
			m.Actions = nil

			if m.ParentID != nil && *m.ParentID == "" {
				m.ParentID = nil
			}
			menuInput := repo.ToEntCreateSysMenuInput(m)
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

		err = tx.SysMenuAction.CreateBulk(actionBulk...).OnConflict(
			sql.ConflictColumns(sysmenuaction.FieldID),
		).UpdateNewValues().Exec(ctx)
		if err != nil {
			log.Println(greenOnWhite, " creatbulk menu action ", err, chalk.Reset)
			return err
		}

		err = tx.SysMenuActionResource.CreateBulk(resourceBulk...).OnConflict(
			sql.ConflictColumns(sysmenuactionresource.FieldID),
		).UpdateNewValues().Exec(ctx)
		if err != nil {
			log.Println(cyanOnBlue, " creatbulk menu action resource ", err, chalk.Reset)
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
