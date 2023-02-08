package app

import (
	"fmt"

	"github.com/heromicro/omgind/pkg/global"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func InitInfluxDB() (influxdb2.Client, func(), error) {
	cfg := global.CFG.InfluxDB

	fmt.Println(" ------- influx db dsn ", cfg.DSN())

	client := influxdb2.NewClient(cfg.DSN(), cfg.Token)

	return client, func() {
		client.Close()
	}, nil
}
