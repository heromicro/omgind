package migrate

import (
	"log"

	"github.com/heromicro/omgind/cmd/omgind/common"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

// postgres
//
// \copy (select * from district_full order by tree_id asc, ids asc) to 'omgind/scripts/sql/district_full.cvs' delimiter ',' CSV HEADER;
// \copy district_full TO 'Path/file_name.csv' delimiter ',' CSV HEADER;

var CmdDump = &cobra.Command{
	Use:   "dump",
	Short: "dump db data to file",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println(" ----- starting of dump ====== ")

		redOnWhite := chalk.Red.NewStyle().WithBackground(chalk.White)

		configFile, err := cmd.Flags().GetString("conf")

		eclient, cleanup, err := common.MakeEntClient(configFile)
		if err != nil {
			log.Println(redOnWhite, " eclient : ", eclient, chalk.Reset)
			return err
		}
		defer cleanup()

		log.Println(" ------- ending of dump ----- ")
		return nil
	},
}

func init() {

	// PersistentFlags() for all flags only in current cmd
	// Cmd.PersistentFlags().StringP("conf", "c", "", "config file path")
	// Cmd.MarkPersistentFlagRequired("conf")

}
