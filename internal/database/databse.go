package database

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/saltcodes/owl-sys-monitor/internal/util"
	"golang.org/x/net/context"
)

func ConnectToInfluxDB() (api.WriteAPIBlocking, error) {
	bucket := util.GetVariableWith("BUCKET")
	org := util.GetVariableWith("ORG")
	token := util.GetVariableWith("INFLUX_TOKEN")
	url := util.GetVariableWith("INFLUX_URL")
	client := influxdb2.NewClient(url, token)
	_, err := client.Health(context.Background())
	return client.WriteAPIBlocking(org, bucket), err
}
