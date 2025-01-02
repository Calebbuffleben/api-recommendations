package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Fetches recommendations for a given user and product using Google's Recommendations AI API.
func getRecommendations(userID string, productID string) (string, error) {
	// URL for the Recommendations AI API
	url := "https://recommendationengine.googleapis.com/v1beta1/projects/[PROJECT_ID]/locations/global/catalogs/default_catalog/branches/default_branch/recommendations:predict"

	// Payload to send to the API
	// This defines the user event for which recommendations are being fetched.
	payload := map[string]interface{}{
		"userEvent": map[string]interface{}{
			"visitorId": userID,      // Unique identifier for the user.
			"eventType": "page-view", // Event type, such as viewing a product or adding to cart.
			"eventDetail": map[string]interface{}{
				"productID": productID, // The ID of the product related to this event.
			},
		},
	}
	// Convert the payload to JSON format
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	// Set the content type header to indicate JSON payload
	req.Header.Set("Content-Type", "application/json")

	// Initialize an HTTP client to send the request
	client := &http.Client{}

	// Send the request and receive a response
	rep, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer rep.Body.Close() // Ensure the response body is closed to free resources

	// Decode the JSON response into a map for processing
	var result map[string]interface{}
	json.NewDecoder(rep.Body).Decode(&result)

	// Convert the result into a string for now; adapt this based on your application's needs
	return fmt.Sprintf("Recommendation: %v", result), nil
}

func main() {
	// Example user and product IDs to test the API
	userID := "user1234"
	productID := "product123"

	// Call the getRecommendations function to fetch recommendations
	recommendation, err := getRecommendations(userID, productID)
	if err != nil {
		// Print an error message if the API call fails
		fmt.Println("Error:", err)
		return
	}

	// Print the recommendation response
	fmt.Println(recommendation)
}
