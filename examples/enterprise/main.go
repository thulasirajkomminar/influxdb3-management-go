package main

import (
	"context"
	"net/http"

	"github.com/caarlos0/env/v11"
	influxdb3enterprise "github.com/thulasirajkomminar/influxdb3-management-go/enterprise"
)

type InfluxdbConfig struct {
	Token string `env:"INFLUXDB3_TOKEN"`
	Url   string `env:"INFLUXDB3_URL"`
}

func main() {
	cfg := InfluxdbConfig{}
	opts := env.Options{RequiredIfNoDef: true}

	err := env.ParseWithOptions(&cfg, opts)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	client, err := influxdb3enterprise.NewClientWithResponses(cfg.Url, influxdb3enterprise.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Authorization", "Bearer "+cfg.Token)
		return nil

	}))
	if err != nil {
		panic(err)
	}

	resp, err := client.GetConfigureDatabaseWithResponse(ctx, &influxdb3enterprise.GetConfigureDatabaseParams{
		Format: "json",
	})
	if err != nil {
		panic(err)
	}

	if resp.StatusCode() == 200 {
		for _, database := range *resp.JSON200.Databases {
			println("Database name: " + database)
		}
	}
}
