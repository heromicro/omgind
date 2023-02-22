package main

import (
	"context"
	"log"

	"github.com/heromicro/omgind/cmd/omgind/common"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli/v2"
)

func DumpCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "dump",
		Usage: "运行web服务",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "配置文件(.json,.yaml,.toml)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {

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
}

// go run -ldflags "-X main.VERSION=$(RELEASE_TAG)" ./cmd/${APP}/main.go dump -c ./configs/config.toml -m ./configs/model.conf --menu ./configs/menu.yaml
