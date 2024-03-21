package wirex

import (
	"log"

	"github.com/heromicro/omgind/pkg/global"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func InitInfluxDB() (influxdb2.Client, func(), error) {
	cfg := global.CFG.InfluxDB

	log.Println(" ------- influx db dsn ", cfg.DSN())

	client := influxdb2.NewClient(cfg.DSN(), cfg.Token)

	return client, func() {
		client.Close()
	}, nil
}
