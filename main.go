package main

import (
	"context"
	"fmt"
	"log"
	"time"

	recommendationengine "cloud.google.com/go/recommendationengine/apiv1"
	recommendationpb "google.golang.org/genproto/googleapis/cloud/recommendationengine/v1beta1"
)

// getRecommendations fetches recommendations for a specific user and context.
func getRecommendations(projectID, eventStoreID, placement string, userID string) {
	ctx := context.Background()

	// Create a recommendation engine client
	client, err := recommendationengine.NewPredictionApiClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Specify the full resource name of the placement
	placementName := fmt.Sprintf(
		"projects/%s/locations/global/catalogs/default_catalog/eventStores/%s/placements/%s", 
		projectID, eventStoreID, placement,
	)

	// Create the predict request
	req := &recommendationpb.PredictRequest{
		Placement: placementName,
		UserEvent: &recommendationpb.UserEvent{
			EventType: "detail-page-view", // Example: user views a product
			UserInfo: &recommendationpb.UserInfo{
				userId: userID,
			},
			EventTime: &recommendationpb.Timestamp{
				Seconds: time.Now().Unix(),
			},
		},
		PageSize: 5, // Number of recommendations to return
	}

	// Call the predictions API
	resp, err := client.Predict(ctx, req)
	if err != nil {
		log.Fatalf("Failed to predict: %v", err)
	}

	// Print the recommendations
	fmt.Println("Recommendations:")
	for _, result := range resp.Results {
		fmt.Printf("- Item ID: %s\n", result.Id)
	}
}

func main() {
	// Example variables (replace with your actual values)
	projectID := "your-project-id"
	eventStoreID := "your-event-store-id"
	placement := "your-placement"
	userID := "your-user-id"

	// Fetch and display recommendations
	getRecommendations(projectID, eventStoreID, placement, userID)
}
