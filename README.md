# go-api# 🚀 Crypto ETL API - Go 

A simple Go ETL API that fetches cryptocurrency prices, processes the data, and serves it via a REST API.

## 📌 Features
- Extracts real-time crypto prices from [CoinLore API](https://www.coinlore.com/cryptocurrency-data-api)
- Transforms data into JSON format
- Loads and serves the data through an API using [Gin](https://github.com/gin-gonic/gin)
- Uses **goroutines** for periodic data fetching

## 🛠 Setup
```bash
git clone https://github.com/yourusername/crypto-etl-api.git
cd crypto-etl-api
go mod tidy
