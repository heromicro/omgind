package migrate

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang-module/carbon/v2"
	strformat "github.com/jossef/format"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"

	"github.com/heromicro/omgind/cmd/omgind/common"
	"github.com/heromicro/omgind/cmd/omgind/migrate/dump"
)

// postgres
//
// \copy (select * from district_full order by tree_id asc, ids asc) to 'omgind/scripts/sql/district_full.csv' delimiter ',' CSV HEADER;
// \copy district_full TO 'Path/file_name.csv' delimiter ',' CSV HEADER;

// go run cmd/omgind/main.go migrate dump --conf=./configs/config.dev.toml --format=yml --tablename=dict --datadir=./configs/data
// go run cmd/omgind/main.go migrate dump --conf=./configs/config.dev.toml --format=yml --tablename=menu --datadir=./configs/data

func init() {

	CmdDump.PersistentFlags().String("datadir", "", "data file path")
	CmdDump.MarkPersistentFlagRequired("datadir")

}

var CmdDump = &cobra.Command{
	Use:   "dump",
	Short: "dump db data to file",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println(" ----- starting of dump ====== ")

		redOnWhite := chalk.Red.NewStyle().WithBackground(chalk.White)
		greenOnWhite := chalk.Green.NewStyle().WithBackground(chalk.White)
		// cyanOnBlue := chalk.Cyan.NewStyle().WithBackground(chalk.Blue)
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

			sts := carbon.Now().StartOfDay().Timestamp()
			nts := carbon.Now().Timestamp()
			dt := carbon.Now().ToDateString()
			datafile := filepath.Join(datadir, strformat.String(`{tn}.{dt}_{s}.yaml`, strformat.Items{"tn": tablename, "dt": dt, "s": nts - sts}))

			switch tablename {
			case "dict":

				dump.Dump_dict(ctx, eclient, datafile)

			case "menu":
				dump.Dump_menu(ctx, eclient, datafile)

			case "role":
				dump.Dump_role(ctx, eclient, datafile)

			case "role_menu":
				dump.Dump_role_menu(ctx, eclient, datafile)

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
