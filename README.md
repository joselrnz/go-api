# 🚀 Crypto ETL API - Go

A simple Go ETL API that **fetches live cryptocurrency prices, processes the data, and serves it via a REST API**.

## 📌 Features

- 🔄 **Extracts** real-time cryptocurrency prices from [CoinLore API](https://www.coinlore.com/cryptocurrency-data-api).
- 🔧 **Transforms** data into a structured JSON format.
- 💾 **Loads** and serves the data through an API using [Gin](https://github.com/gin-gonic/gin).
- ⚡ Uses **goroutines** for periodic data fetching.
- 🔥 Fast, lightweight, and ideal for cloud deployments.

---

## 🛠 Setup

### **1️⃣ Clone the Repository**

```bash
git clone https://github.com/joselrnz/go-api.git
cd go-api
```

### **2️⃣ Initialize Go Modules**

```bash
go mod tidy
```

### **3️⃣ Run the Application**

```bash
go run main.go
```

---

## 🖥 API Endpoints

| Method | Endpoint   | Description                      |
| ------ | ---------- | -------------------------------- |
| GET    | `/`        | Welcome message                  |
| GET    | `/cryptos` | Fetch live cryptocurrency prices |

### **Example API Response**

```json
[
  {
    "name": "Bitcoin",
    "symbol": "BTC",
    "price": "60000"
  },
  {
    "name": "Ethereum",
    "symbol": "ETH",
    "price": "3500"
  }
]
```

---

## 🛠 Project Structure

```
go-api/
│── etl/
│   ├── etl.go         # Extract, Transform, Load logic
│── api/
│   ├── api.go         # API routes and server setup
│── main.go            # Main application entry point
│── go.mod             # Go modules and dependencies
│── README.md          # Documentation
```

---

## 🔄 How the ETL Process Works

1. **Extract** - Fetches crypto data from an external API.
2. **Transform** - Converts JSON response into structured Go objects.
3. **Load** - Stores the data in memory and serves it via a REST API.

The data **refreshes automatically every 30 seconds** using **Goroutines**.

---

## 🛠️ Building the Binary

To compile the Go application into an executable:

```bash
go build -o crypto-etl-api
./crypto-etl-api
```

---

## 🚀 Deploying to a Server (Optional)

For cloud or containerized deployments:

### **Using Docker**

```bash
docker build -t crypto-etl-api .
docker run -p 8080:8080 crypto-etl-api
```

### **Using a Cloud VM**

1. Upload the compiled `crypto-etl-api` binary to your VM.
2. Run it as a background service:
   ```bash
   nohup ./crypto-etl-api > output.log 2>&1 &
   ```

---

## 👨‍💻 Contributing

Want to improve this project? Follow these steps:

1. **Fork the repository** on GitHub.
2. **Create a new branch** (`git checkout -b feature-branch`).
3. **Commit your changes** (`git commit -m "Added new feature"`).
4. **Push to your fork** (`git push origin feature-branch`).
5. **Submit a pull request**.

---


