package main

import (
	"fmt"
	"time"
	"github.com/joselrnz/go-api/api"
	"github.com/joselrnz/go-api/etl" 
)

func main() {
	fmt.Println("Starting Crypto ETL API...")

	// Run ETL every 30 seconds using a goroutine
	go func() {
		for {
			etl.Load()
			time.Sleep(30 * time.Second)
		}
	}()

	// Start API server
	api.StartServer()
}
