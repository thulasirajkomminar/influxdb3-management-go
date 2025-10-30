package main

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/hashicorp/go-retryablehttp"
	influxdb3 "github.com/thulasirajkomminar/influxdb3-management-go"
)

type InfluxdbConfig struct {
	AccountId influxdb3.UuidV4 `env:"INFLUXDB_ACCOUNT_ID"`
	BaseURL   string           `env:"INFLUXDB_BASE_URL"`
	ClusterId influxdb3.UuidV4 `env:"INFLUXDB_CLUSTER_ID"`
	Token     string           `env:"INFLUXDB_TOKEN"`
}

func main() {
	cfg := InfluxdbConfig{}
	opts := env.Options{RequiredIfNoDef: true}

	err := env.ParseWithOptions(&cfg, opts)
	if err != nil {
		panic(err)
	}

	retryClient := retryablehttp.NewClient()
	retryClient.Backoff = retryablehttp.LinearJitterBackoff
	retryClient.RetryWaitMin = 1 * time.Second
	retryClient.RetryWaitMax = 5 * time.Second
	retryClient.RetryMax = 3

	client, err := influxdb3.NewClientWithResponses(cfg.BaseURL, influxdb3.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Authorization", "Bearer "+cfg.Token)
		return nil
	}), influxdb3.WithHTTPClient(retryClient.StandardClient()))
	if err != nil {
		panic(err)
	}

	resp, err := client.GetClusterDatabases(context.Background(), cfg.AccountId, cfg.ClusterId)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		bodyString := string(bodyBytes)
		println(bodyString)
	}
}
