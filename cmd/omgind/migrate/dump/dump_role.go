package dump

import (
	"context"
	"log"
	"math"
	"os"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/entscheme"
	"github.com/heromicro/omgind/internal/gen/entscheme/sysrole"
	"github.com/heromicro/omgind/internal/scheme/repo"
	"github.com/ttacon/chalk"
	"gopkg.in/yaml.v2"
)

func Dump_role(ctx context.Context, eclient *entscheme.Client, datafile string) error {

	redOnWhite := chalk.Red.NewStyle().WithBackground(chalk.White)
	cyanOnBlue := chalk.Cyan.NewStyle().WithBackground(chalk.Blue)

	fp, err := os.OpenFile(datafile, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Println(redOnWhite, " can not open file: ", err, chalk.Reset)
		return err
	}
	defer fp.Close()

	query := eclient.SysRole.Query()
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

		query := eclient.SysRole.Query()

		query = query.Order(entscheme.Asc(sysrole.FieldID), entscheme.Asc(sysrole.FieldSort))

		r_roles, err := query.Limit(pageSize).Offset(offset).All(ctx)
		if err != nil {
			return err
		}

		var sch_roles []*schema.Role = repo.ToSchemaRoles(r_roles)

		// log.Println(" -------- ====== ", sch_dicts[0])
		data, err := yaml.Marshal(sch_roles)
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
