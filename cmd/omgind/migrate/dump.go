package migrate

import "github.com/spf13/cobra"

var CmdDump = &cobra.Command{
	Use:   "dump",
	Short: "dump db data to file",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
