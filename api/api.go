package api

import (
	"github.com/gin-gonic/gin"
	"github.com/joselrnz/go-api/etl"
	"net/http"
)

// initializes the API
func StartServer() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "The Crypto ETL API! Use /cryptos to get data."})
	})

	// API endpoint to get crypto data
	r.GET("/cryptos", func(c *gin.Context) {
		data := etl.GetCryptoData()
		c.JSON(http.StatusOK, data)
	})

	// Start API server
	r.Run(":8080") // 
}
