package migrate

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate db data",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	Cmd.AddCommand(CmdDump)
	Cmd.AddCommand(CmdLoad)

	// PersistentFlags() for all flags only in current cmd
	Cmd.PersistentFlags().StringP("conf", "c", "", "config file path")
	Cmd.MarkPersistentFlagRequired("conf")

	Cmd.PersistentFlags().String("format", "", "format")
	Cmd.MarkPersistentFlagRequired("format")

}

// go run -ldflags "-X main.VERSION=$(RELEASE_TAG)" ./cmd/${APP}/main.go dump -c ./configs/config.toml -m ./configs/model.conf --menu ./configs/menu.yaml
