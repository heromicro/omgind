package migrate

import (
	"log"

	"github.com/heromicro/omgind/cmd/omgind/common"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

var CmdLoad = &cobra.Command{
	Use:   "load",
	Short: "load file to db data",
	RunE: func(cmd *cobra.Command, args []string) error {

		redOnWhite := chalk.Red.NewStyle().WithBackground(chalk.White)

		log.Println(" ----- ====== ---- ====== ")
		cf := c.String("conf")

		if true {
			return nil
		}

		// _, cleanup, err := common.MakeEntClient(cf)
		eclient, cleanup, err := common.MakeEntClient(cf)
		if err != nil {
			log.Println(redOnWhite, " eclient : ", eclient, chalk.Reset)
			return err
		}

		defer cleanup()

		log.Println(" ------- ======== ----- ")

		return nil
	},
}
