package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VERSION = "7.0.0"

var Cmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VERSION)
	},
}
