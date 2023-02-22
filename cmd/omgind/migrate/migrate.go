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

}

// go run -ldflags "-X main.VERSION=$(RELEASE_TAG)" ./cmd/${APP}/main.go dump -c ./configs/config.toml -m ./configs/model.conf --menu ./configs/menu.yaml
