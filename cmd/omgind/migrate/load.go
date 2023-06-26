package migrate

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
	"github.com/xuri/excelize/v2"

	"github.com/heromicro/omgind/cmd/omgind/common"
	"github.com/heromicro/omgind/cmd/omgind/migrate/load"
)

func init() {

	// PersistentFlags() for all flags only in current cmd

	CmdLoad.PersistentFlags().String("datafile", "", "data file path")
	CmdLoad.MarkPersistentFlagRequired("datafile")

}

// go run cmd/omgind/main.go migrate load --conf configs/config.toml --format=csv --tablename=districts --datafile=./scripts/sql/district_full.csv
// go run cmd/omgind/main.go migrate load --conf configs/config.toml --format=yaml --tablename=menu --datafile=./configs/menu.2023-03-28_64619.yaml

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
				load.Load_dict_data(ctx, eclient, datafile)

			case "menu":
				load.Load_menu_data(ctx, eclient, datafile)

			case "role":
				load.Load_role_data(ctx, eclient, datafile)

			case "role_menu":
				load.Load_role_menu_data(ctx, eclient, datafile)

			default:
				log.Println(whiteOnGreen, " missing on tablename: ", tablename, chalk.Reset)
			}
		case "csv":

			switch tablename {
			case "district":
				load.Load_district_data(ctx, eclient, datafile)

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
