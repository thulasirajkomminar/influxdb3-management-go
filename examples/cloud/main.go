package main

import (
	"context"
	"net/http"

	"github.com/caarlos0/env/v11"
	influxdb3cloud "github.com/thulasirajkomminar/influxdb3-management-go/cloud"
)

type InfluxdbConfig struct {
	AccountId influxdb3cloud.UuidV4 `env:"INFLUXDB3_ACCOUNT_ID"`
	ClusterId influxdb3cloud.UuidV4 `env:"INFLUXDB3_CLUSTER_ID"`
	Token     string                `env:"INFLUXDB3_TOKEN"`
	Url       string                `env:"INFLUXDB3_URL"`
}

func main() {
	cfg := InfluxdbConfig{}
	opts := env.Options{RequiredIfNoDef: true}

	err := env.ParseWithOptions(&cfg, opts)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	client, err := influxdb3cloud.NewClientWithResponses(cfg.Url, influxdb3cloud.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Authorization", "Bearer "+cfg.Token)
		return nil

	}))
	if err != nil {
		panic(err)
	}

	resp, err := client.GetClusterDatabasesWithResponse(ctx, cfg.AccountId, cfg.ClusterId)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode() == 200 {
		for _, database := range *resp.JSON200 {
			println("Database name: " + database.Name)
		}
	}
}
