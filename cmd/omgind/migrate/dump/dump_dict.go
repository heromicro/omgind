package dump

import (
	"context"
	"log"
	"math"
	"os"

	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/sysdict"
	"github.com/heromicro/omgind/internal/gen/ent/sysdictitem"
	"github.com/heromicro/omgind/internal/scheme/repo"
	"github.com/ttacon/chalk"
	"gopkg.in/yaml.v2"
)

func Dump_dict(ctx context.Context, eclient *ent.Client, datafile string) error {

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

	query := eclient.SysDict.Query()
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

		query := eclient.SysDict.Query().WithItems(func(sdiq *ent.SysDictItemQuery) {
			sdiq.Select(sysdictitem.FieldID, sysdictitem.FieldLabel, sysdictitem.FieldValue, sysdictitem.FieldDictID, sysdictitem.FieldSort, sysdictitem.FieldMemo)
		})
		query = query.Order(ent.Asc(sysdict.FieldSort))

		r_dicts, err := query.Limit(pageSize).Offset(offset).All(ctx)
		if err != nil {
			return err
		}

		// log.Println(" -------- ====== ", r_dicts[0].String())

		sch_dicts := repo.ToSchemaSysDicts(r_dicts)
		// log.Println(" -------- ====== ", sch_dicts[0])
		data, err := yaml.Marshal(sch_dicts)
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
