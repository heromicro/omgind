package web

import (
	"context"

	"github.com/heromicro/omgind/cmd/omgind/version"
	"github.com/heromicro/omgind/internal/app"
	"github.com/heromicro/omgind/pkg/logger"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "web",
	Short: "run web server",
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := logger.NewTagContext(context.Background(), "__main__")

		configFile, err := cmd.Flags().GetString("conf")
		if err != nil {
			return err
		}
		casbinModel, err := cmd.Flags().GetString("model")
		if err != nil {
			return err
		}

		wwwdir, err := cmd.Flags().GetString("www")

		initMenuFile, err := cmd.Flags().GetString("menu")

		return app.Run(ctx,
			app.SetConfigFile(configFile),
			app.SetModelFile(casbinModel),
			app.SetWWWDir(wwwdir),
			app.SetMenuFile(initMenuFile),
			app.SetVersion(version.VERSION))
	},
}

func init() {

	// Flags() all flags in current cmd and parent cmds
	// Cmd.Flags().,

	// PersistentFlags() for all flags only in current cmd
	Cmd.PersistentFlags().StringP("conf", "c", "", "config file path")
	Cmd.MarkPersistentFlagRequired("conf")

	// PersistentFlags() all flags only in current cmd
	Cmd.PersistentFlags().StringP("model", "m", "", "casbin model file path")
	Cmd.MarkPersistentFlagRequired("model")

	// PersistentFlags() all flags only in current cmd
	Cmd.PersistentFlags().String("www", "", "www folder")

	// PersistentFlags() all flags only in current cmd
	Cmd.PersistentFlags().String("menu", "", "menu file path")

}
