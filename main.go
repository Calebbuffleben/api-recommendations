package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func getRecomendations(userID string, productID string) (string, error) {
	url := "https://recommendationengine.googleapis.com/v1beta1/projects/[PROJECT_ID]/locations/global/catalogs/default_catalog/branches/default_branch/recommendations:predict"

	payload := map[string]interface{}{
		"userEvent": map[string]interface{}{
			"visitorId": userID,
			"eventType": "page-view",
			"eventDetail": map[string]interface{}{
				"productID": productID,
			},
		},
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	rep, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer rep.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(rep.Body).Decode(&result)

	// Process the response as needed
	return fmt.Sprintf("Recommendation: %v", result), nil
}

func main() {
	recommendation, err := getRecomendations("123", "456")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(recommendation)
}
