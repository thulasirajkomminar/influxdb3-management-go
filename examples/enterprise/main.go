package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/komminarlabs/influxdb3-management-go/enterprise"
)

func main() {
	// Get environment variables
	baseURL := os.Getenv("INFLUXDB_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8086"
	}

	token := os.Getenv("INFLUXDB_TOKEN")
	database := os.Getenv("INFLUXDB_DATABASE")
	if database == "" {
		database = "example_db"
	}

	if token == "" {
		log.Fatal("Please set INFLUXDB_TOKEN environment variable")
	}

	// Create client
	client, err := enterprise.NewClient(baseURL)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	fmt.Println("🏢 InfluxDB Enterprise API Example")
	fmt.Println("==================================")

	// Example 1: Check health
	fmt.Println("\n❤️  Checking server health...")
	healthResp, err := client.GetHealthV1(ctx, func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to check health: %v", err)
	}
	defer healthResp.Body.Close()

	if healthResp.StatusCode == 200 {
		fmt.Printf("✅ Server is healthy (status: %d)\n", healthResp.StatusCode)
	} else {
		fmt.Printf("⚠️  Health check completed with status: %d\n", healthResp.StatusCode)
	}

	// Example 2: Write some sample data
	fmt.Println("\n📝 Writing sample data...")
	lineProtocolData := fmt.Sprintf("temperature,location=datacenter1 value=24.8 %d", time.Now().UnixNano())

	writeParams := &enterprise.PostV2WriteParams{
		Db:        database,
		Precision: enterprise.PrecisionWriteCompatibilityNs,
	}

	writeResp, err := client.PostV2WriteWithTextBody(ctx, writeParams, lineProtocolData, func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "text/plain")
		return nil
	})
	if err != nil {
		log.Printf("Failed to write data: %v", err)
	} else {
		defer writeResp.Body.Close()
		if writeResp.StatusCode == 204 {
			fmt.Printf("✅ Data written successfully (status: %d)\n", writeResp.StatusCode)
		} else {
			fmt.Printf("⚠️  Write completed with status: %d\n", writeResp.StatusCode)
		}
	}

	// Example 3: Configure database
	fmt.Println("\n⚙️  Configuring database...")
	dbConfig := enterprise.PostConfigureDatabaseJSONRequestBody{
		Db: database,
	}

	configResp, err := client.PostConfigureDatabase(ctx, dbConfig, func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	})
	if err != nil {
		log.Printf("Failed to configure database: %v", err)
	} else {
		defer configResp.Body.Close()
		if configResp.StatusCode == 200 || configResp.StatusCode == 201 {
			fmt.Printf("✅ Database configured successfully (status: %d)\n", configResp.StatusCode)
		} else {
			fmt.Printf("⚠️  Database configuration completed with status: %d\n", configResp.StatusCode)
		}
	}

	// Example 4: Enterprise-specific - Configure processing engine trigger
	fmt.Println("\n🚀 Configuring processing engine trigger...")
	triggerConfig := enterprise.PostConfigureProcessingEngineTriggerJSONRequestBody{
		Db:                   database,
		TriggerName:          "example-trigger",
		PluginFilename:       "example-plugin.py",
		TriggerSpecification: "every:1m", // Every minute
	}

	triggerResp, err := client.PostConfigureProcessingEngineTrigger(ctx, triggerConfig, func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	})
	if err != nil {
		log.Printf("Failed to configure processing engine trigger: %v", err)
	} else {
		defer triggerResp.Body.Close()
		if triggerResp.StatusCode == 200 || triggerResp.StatusCode == 201 {
			fmt.Printf("✅ Processing engine trigger configured successfully (status: %d)\n", triggerResp.StatusCode)
		} else {
			fmt.Printf("⚠️  Processing engine trigger configuration completed with status: %d\n", triggerResp.StatusCode)
		}
	}

	fmt.Println("\n🎉 Enterprise API example completed!")
	fmt.Println("\nFor more examples, see: https://github.com/komminarlabs/influxdb3-management-go/tree/main/examples")
}
