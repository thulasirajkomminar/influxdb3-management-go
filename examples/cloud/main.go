package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/komminarlabs/influxdb3-management-go/cloud"
)

func main() {
	// Get environment variables
	baseURL := os.Getenv("INFLUXDB3_BASE_URL")
	if baseURL == "" {
		baseURL = "https://console.influxdata.com/api/v0"
	}

	accountID := os.Getenv("INFLUXDB3_ACCOUNT_ID")
	clusterID := os.Getenv("INFLUXDB3_CLUSTER_ID")
	token := os.Getenv("INFLUXDB3_TOKEN")

	if accountID == "" || clusterID == "" || token == "" {
		log.Fatal("Please set INFLUXDB3_ACCOUNT_ID, INFLUXDB3_CLUSTER_ID, and INFLUXDB3_TOKEN environment variables")
	}

	// Create client
	client, err := cloud.NewClientWithResponses(baseURL, cloud.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	}), cloud.WithHTTPClient(&http.Client{}))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Convert string IDs to UuidV4 type
	accountUUID, err := uuid.Parse(accountID)
	if err != nil {
		log.Fatalf("Invalid account ID format: %v", err)
	}

	clusterUUID, err := uuid.Parse(clusterID)
	if err != nil {
		log.Fatalf("Invalid cluster ID format: %v", err)
	}

	fmt.Println("🌤️  InfluxDB Cloud Dedicated Management API Example")
	fmt.Println("==================================================")

	// Example 1: List databases
	fmt.Println("\n📋 Listing databases...")
	resp, err := client.GetClusterDatabasesWithResponse(ctx, accountUUID, clusterUUID, func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to list databases: %v", err)
	}

	if resp.StatusCode() == 200 {
		fmt.Printf("✅ Successfully retrieved databases (status: %d)\n", resp.StatusCode())
		for _, db := range *resp.JSON200 {
			fmt.Printf(" - %s\n", db.Name)
		}
	} else {
		fmt.Printf("⚠️  Request completed with status: %d\n", resp.StatusCode())
	}

	fmt.Println("\n🎉 Cloud Dedicated API example completed!")
	fmt.Println("\nFor more examples, see: https://github.com/komminarlabs/influxdb3-management-go/tree/main/examples")
}
