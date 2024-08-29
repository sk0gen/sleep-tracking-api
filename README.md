# 💤 SleepTrack API
SleepTrack API is a RESTful API built with Go, designed to help users track and analyze their sleep patterns. This API is part of a health tech initiative aimed at improving personal wellness through data-driven insights.


## 🚀 Features
User Registration & Authentication: Secure user signup and login with JWT-based authentication.
Log Sleep Data: Easily log sleep start and end times, along with sleep quality ratings.
View Sleep Logs: Retrieve and manage past sleep logs.
Sleep Trends Analysis: Analyze sleep patterns over time to help users understand and improve their sleep quality.


## 🛠️ Tech Stack
Go: The powerful and efficient programming language for building scalable backends.
Gin: A lightweight and fast HTTP framework for building web applications.
Sqlc: Type-safe code generator from SQL
PostgreSQL: A robust and reliable relational database for storing user and sleep data.
JWT: For secure user authentication and authorization.


## 📦 Project Structure
```
├── cmd/                # Entry point of the application
├── internal/           # Core application logic
│   ├── api/            # HTTP handlers and routing
│   ├── config/         # Loads configuration and contains strong type representation of env variables
│   ├── database/       # Database access and queries
│   ├── gapi/           # GRPC api definition
│   ├── logging/        # Initialize Zap logger
│   ├── pagination/     # Pagination request models
│   ├── pb              # Implementation of GRPC service
│   ├── proto/          # Definition of GRPC service
│   ├── token/          # JWT authentication creation and validation
├── util/               # Utility classes/functions used across modules
└── README.md           # Project documentation
```

# 📝 Getting Started
Prerequisites
* Go (version 1.19+)
* Docker

Development tools (optional):
* Protobuf compiler (brew install protobuf)
* Cobra CLI generator (go install github.com/spf13/cobra/cobra)
* Protoc-gen-go (go install github.com/golang/protobuf/protoc-gen-go)
* Protoc-gen-go-grpc (go install google.golang.org/grpc/cmd/protoc-gen-go-grpc)
* Swag-go (go install github.com/swaggo/swag/cmd/swag)

### Installation
1. Clone the repository:
```
git clone https://github.com/sk0gen/sleep-tracking-api.git
cd sleep-tracking-api
```
2. Install dependencies:
```
go mod tidy
```
3. Create a `.env` file with environment variables:
```
mv .env.example .env
```
4. Set up the PostgreSQL database:
```
make docker-up
```
5. Update database schema
```
make migration_up
```
6. Start the server:
```
make run
```

### Usage
The SleepTrack API can be accessed and tested through Swagger UI. Once the server is running, you can explore and interact with the API endpoints at (If the example configuration is used)

```
http://localhost:8080/swagger/index.html
```
This Swagger interface provides a user-friendly way to view all available endpoints, their parameters, and even test the API directly from your browser.


# To-Do List for Sleep Tracker API Project

## Backend Development

- [X] GRPC - Implemented simple GRPC server with LoginUser and GetUserSleepLogs methods
- [X] Zap logger
- [X] Graceful shutdown
- [X] API integration tests - Half done. No full coverage in Sleep-logs api
- [X] Swagger
- [ ] Tracing/Observability/Metrics
- [ ] Export sleep data to file?
- [ ] Extract sleep data - Cursor query for GRPC
- [ ] Sleep analysis - Patterns/Sleep time per week calculation etc...