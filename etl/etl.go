
package etl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// Crypto struct to map API response
type Crypto struct {
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price_usd,string"`
}

// Global variable to store data
var cryptoData []Crypto
var mu sync.Mutex

// Extract function - fetches data from an API
func Extract() ([]Crypto, error) {
	url := "https://api.coinlore.net/api/tickers/" // Free crypto API
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	// Transform: Parse JSON
	var data struct {
		Data []Crypto `json:"data"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return data.Data, nil
}

// Load function - stores data in memory
func Load() {
	mu.Lock()
	defer mu.Unlock()

	data, err := Extract()
	if err != nil {
		fmt.Println("ETL Error:", err)
		return
	}
	cryptoData = data
	fmt.Println("Crypto data updated successfully!")
}

// GetCryptoData - returns stored crypto data
func GetCryptoData() []Crypto {
	mu.Lock()
	defer mu.Unlock()
	return cryptoData
}
