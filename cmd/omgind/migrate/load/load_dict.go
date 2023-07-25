package load

import (
	"context"
	"io/ioutil"
	"log"

	"entgo.io/ent/dialect/sql"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/sysdict"
	"github.com/heromicro/omgind/internal/gen/mainent/sysdictitem"
	"github.com/heromicro/omgind/internal/scheme/repo"
	"github.com/ttacon/chalk"
	"gopkg.in/yaml.v2"
)

func Load_dict_data(ctx context.Context, eclient *mainent.Client, filename string) error {

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

	var dictBulk []*mainent.SysDictCreate
	var dictItemBulk []*mainent.SysDictItemCreate

	err = repo.WithTx(ctx, eclient, func(tx *mainent.Tx) error {

		for _, d := range dicts {
			for _, di := range d.Items {
				log.Println(cyanOnYellow, " di :: ", di.String(), chalk.Reset)
				dictItemInput := repo.ToEntCreateSysDictItemInput(di)
				log.Println(greenOnWhite, " di :: ", dictItemInput.DictID, chalk.Reset)

				create_dictitem := tx.SysDictItem.Create().SetID(di.ID).SetInput(*dictItemInput)

				log.Println(greenOnWhite, " di :: ", create_dictitem, chalk.Reset)

				dictItemBulk = append(dictItemBulk, create_dictitem)
			}

			log.Println(cyanOnYellow, " di :: ", d.String(), chalk.Reset)

			d.Items = nil
			dictInput := repo.ToEntCreateSysDictInput(d)

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
