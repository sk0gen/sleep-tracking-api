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


# To-Do List for Sleep Tracker API Project

## Backend Development

- <s>Automatic migrations on `make run`</s>
- [X] GRPC - Implemented simple GRPC server with LoginUser and GetUserSleepLogs methods
- [ ] Sleep analysis - Patterns/Sleep time per week calculation etc...
- [ ] Tracing/Observability/Metrics
- [X] API integration tests - Half done. No full coverage in Sleep-logs api
- [X] Graceful shutdown
- [X] Zap logger
- [ ] Export sleep data to file?
- [ ] Extract sleep data - Cursor query for GRPC
- [X] Swagger
