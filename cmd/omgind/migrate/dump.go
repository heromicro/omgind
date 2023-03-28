package migrate

import (
	"context"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"

	strformat "github.com/jossef/format"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
	"gopkg.in/yaml.v2"

	"github.com/heromicro/omgind/cmd/omgind/common"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/sysdict"
	"github.com/heromicro/omgind/internal/gen/ent/sysdictitem"
	"github.com/heromicro/omgind/internal/schema/repo"
)

// postgres
//
// \copy (select * from district_full order by tree_id asc, ids asc) to 'omgind/scripts/sql/district_full.csv' delimiter ',' CSV HEADER;
// \copy district_full TO 'Path/file_name.csv' delimiter ',' CSV HEADER;

// go run cmd/omgind/main.go migrate dump --conf=./configs/config.dev.toml --format=yml --tablename=dict --datadir=./configs/

var CmdDump = &cobra.Command{
	Use:   "dump",
	Short: "dump db data to file",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println(" ----- starting of dump ====== ")

		redOnWhite := chalk.Red.NewStyle().WithBackground(chalk.White)
		greenOnWhite := chalk.Green.NewStyle().WithBackground(chalk.White)
		cyanOnBlue := chalk.Cyan.NewStyle().WithBackground(chalk.Blue)
		whiteOnGreen := chalk.Cyan.NewStyle().WithBackground(chalk.Green)

		configFile, err := cmd.Flags().GetString("conf")

		eclient, cleanup, err := common.MakeEntClient(configFile)
		if err != nil {
			log.Println(redOnWhite, " eclient : ", eclient, chalk.Reset)
			return err
		}
		defer cleanup()

		datadir, err := cmd.Flags().GetString("datadir")
		if err != nil {
			// log.Println(redOnWhite, " ----- === === ", err, chalk.Reset)
			return err
		}

		_, err = os.Stat(datadir)
		if err != nil && os.IsNotExist(err) {
			log.Println(redOnWhite, datadir, " exist, choose another.", datadir, err, chalk.Reset)
			return err
		}

		format, err := cmd.Flags().GetString("format")
		if err != nil {
			log.Println(redOnWhite, "format err: ", err, chalk.Reset)
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
		case "yaml":
			fallthrough
		case "yml":
			switch tablename {
			case "dict":

				dt := time.Now().Format("2006-01-02_11_11_11")
				datafile := filepath.Join(datadir, strformat.String(`{tn}.{dt}.yaml`, strformat.Items{"tn": tablename, "dt": dt}))

				fp, err := os.OpenFile(datafile, os.O_RDWR|os.O_CREATE, 0777)
				if err != nil {
					log.Println(redOnWhite, " can not open file: ", err, chalk.Reset)
					return err
				}
				defer fp.Close()

				query := eclient.SysDict.Query()
				total, err := query.Count(ctx)

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
			default:
				log.Println(whiteOnGreen, " missing on tablename: ", tablename, chalk.Reset)

			}
		case "csv":

		default:
			log.Println(whiteOnGreen, " missing on format: ", format, chalk.Reset)

		}

		log.Println(greenOnWhite, " ------- ending of dump ----- ", chalk.Reset)

		return nil
	},
}

func init() {

	CmdDump.PersistentFlags().String("datadir", "", "data file path")
	CmdDump.MarkPersistentFlagRequired("datadir")

	CmdDump.PersistentFlags().String("tablename", "", "table name")
	CmdDump.MarkPersistentFlagRequired("tablename")

}
