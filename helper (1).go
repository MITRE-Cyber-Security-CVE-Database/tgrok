package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// GetResponse sends a request to the Grok 3 API and prints the response
func GetResponse(input string, params Params, isQuiet bool) {
	if !isQuiet {
		fmt.Print("Thinking")
		go func() {
			for {
				fmt.Print(".")
				time.Sleep(500 * time.Millisecond)
			}
		}()
	}

	// Prepare request payload
	payload := map[string]interface{}{
		"prompt": input,
		"model":  "grok-3", // Placeholder model name
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		PrintError(fmt.Sprintf("Error marshaling payload: %v", err))
		return
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", params.Url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		PrintError(fmt.Sprintf("Error creating request: %v", err))
		return
	}
	req.Header.Set("Content-Type", "application/json")
	if params.ApiKey != "" {
		req.Header.Set("Authorization", "Bearer "+params.ApiKey)
	}

	// Send request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if !isQuiet {
		fmt.Println()
	}
	if err != nil {
		PrintError(fmt.Sprintf("Error sending request: %v", err))
		return
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		PrintError(fmt.Sprintf("Error reading response: %v", err))
		return
	}

	// Parse response (assuming JSON with a "response" field)
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		PrintError(fmt.Sprintf("Error parsing response: %v", err))
		return
	}
	response, ok := result["response"].(string)
	if !ok {
		PrintError("Invalid response format")
		return
	}

	// Print response
	fmt.Println("\n" + response)
}