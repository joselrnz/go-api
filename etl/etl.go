package etl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// Crypto struct to map API response data
// This struct represents a cryptocurrency with its name, symbol, and price.
type Crypto struct {
	Name   string  `json:"name"`                  // Cryptocurrency name
	Symbol string  `json:"symbol"`                // Symbol of the cryptocurrency
	Price  float64 `json:"price_usd,string"`      // Price in USD (string in JSON, converted to float)
}

// Global variables to store extracted cryptocurrency data and ensure thread safety
var cryptoData []Crypto // Stores the extracted crypto data
var mu sync.Mutex       // Mutex for synchronizing access to cryptoData

// Extract function - fetches cryptocurrency data from an external API
func Extract() ([]Crypto, error) {
	url := "https://api.coinlore.net/api/tickers/" // Public cryptocurrency API endpoint
	
	// Send HTTP GET request to fetch data
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err) // Return an error if the request fails
	}
	defer resp.Body.Close() // Ensure response body is closed after function exits

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err) // Return an error if reading fails
	}

	// Transform: Parse JSON response
	var data struct {
		Data []Crypto `json:"data"` // Extract relevant data field from JSON response
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err) // Return an error if parsing JSON fails
	}

	return data.Data, nil // Return parsed cryptocurrency data
}

// Load function - extracts and updates stored crypto data
func Load() {
	mu.Lock()         // Lock mutex to prevent concurrent data modification
	defer mu.Unlock() // Ensure mutex is unlocked after function execution

	// Extract data from the API
	data, err := Extract()
	if err != nil {
		fmt.Println("ETL Error:", err) // Print error if extraction fails
		return
	}
	
	// Store extracted data in global variable
	cryptoData = data
	fmt.Println("Crypto data updated successfully!") // Confirm successful update
}

// GetCryptoData - returns the stored cryptocurrency data
func GetCryptoData() []Crypto {
	mu.Lock()         // Lock mutex for thread-safe access
	defer mu.Unlock() // Unlock mutex after function execution
	return cryptoData // Return stored crypto data
}
