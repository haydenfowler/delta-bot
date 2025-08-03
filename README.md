# Delta Bot - Triangular Arbitrage Trading Bot

A modular triangular arbitrage trading bot built in Go with clean architecture principles, supporting local development with Docker and deployment to AWS ECS Fargate.

## Features

- 🏗️ Clean, modular architecture with dependency injection
- ⚡ High-performance Echo HTTP framework with middleware
- 🐳 Docker support for local development
- ☁️ AWS ECS Fargate deployment with Terraform
- 📊 NewRelic integration with automatic request tracing
- 🔄 Triangular arbitrage detection and execution
- 🧪 Dry-run mode for safe testing
- 📡 Real-time market data via Binance WebSocket

## Project Structure

```
delta-bot/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── config/              # Configuration management
│   ├── logger/              # Logging and NewRelic integration
│   ├── server/              # Echo HTTP server implementation
│   ├── middleware/          # HTTP middleware (NewRelic, logging, etc.)
│   ├── handlers/            # HTTP request handlers
│   ├── routes/              # Route definitions and setup
│   ├── exchange/            # Exchange integrations (Binance)
│   ├── arb/                 # Arbitrage detection logic
│   └── executor/            # Trade execution engine
├── deployments/
│   └── terraform/           # Infrastructure as Code
├── docker/
│   ├── Dockerfile
│   └── docker-compose.yml
└── Makefile                 # Development commands
```

## Quick Start

### Prerequisites

- Go 1.21+
- Docker and Docker Compose
- Make

### Setup

1. **Clone and setup:**
   ```bash
   git clone <repository-url>
   cd delta-bot
   make deps
   ```

2. **Create environment file:**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Configure environment variables:**
   ```env
   # Server Configuration
   PORT=8080
   LOG_LEVEL=INFO

   # NewRelic Configuration
   NEW_RELIC_LICENSE_KEY=your_license_key_here
   NEW_RELIC_APP_NAME=delta-bot

   # Trading Configuration
   DRY_RUN=true
   BINANCE_API_KEY=your_binance_api_key
   BINANCE_SECRET_KEY=your_binance_secret_key

   # Arbitrage Configuration
   MIN_PROFIT_THRESHOLD=0.5
   MAX_TRADE_AMOUNT=1000
   ```

4. **Run the application:**
   ```bash
   # Run locally
   make run
   
   # Or run with Docker
   make docker-run
   ```

5. **Test the health endpoint:**
   ```bash
   make health
   # or
   curl http://localhost:8080/health
   ```

## Development Commands

```bash
make help          # Show available commands
make deps          # Download dependencies
make build         # Build the application
make run           # Build and run
make dev           # Run with live reload (requires air)
make test          # Run tests
make test-coverage # Run tests with coverage
make lint          # Run linter
make fmt           # Format code
make clean         # Clean build artifacts

# Docker commands
make docker-build          # Build Docker image
make docker-run            # Run with docker-compose
make docker-run-detached   # Run in background
make docker-down           # Stop containers
make docker-logs           # View container logs
make docker-restart        # Restart containers
make docker-clean          # Clean up images and containers
```

## API Endpoints

### Health Check
```http
GET /health
```

Response:
```json
{
  "status": "healthy",
  "timestamp": "2024-01-01T12:00:00Z"
}
```

## Configuration

The application uses environment variables for configuration with sensible defaults:

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | HTTP server port |
| `LOG_LEVEL` | `INFO` | Logging level |
| `NEW_RELIC_LICENSE_KEY` | `` | NewRelic license key |
| `NEW_RELIC_APP_NAME` | `delta-bot` | NewRelic application name |
| `DRY_RUN` | `true` | Enable dry-run mode |
| `MIN_PROFIT_THRESHOLD` | `0.5` | Minimum profit threshold (%) |
| `MAX_TRADE_AMOUNT` | `1000` | Maximum trade amount |

## Architecture

The application follows clean architecture principles with clear separation of concerns:

### HTTP Layer
- **Echo Framework**: High-performance HTTP router with middleware support
- **Middleware**: NewRelic tracing, request logging, CORS, recovery
- **Handlers**: Business logic separated from HTTP concerns
- **Routes**: Centralized route configuration

### Core Components
- **Config**: Environment-based configuration management
- **Logger**: Structured logging with NewRelic integration
- **Server**: Echo server wrapper with graceful shutdown

### Middleware Stack
1. **Recovery**: Panic recovery and error handling
2. **CORS**: Cross-origin request support
3. **NewRelic**: Automatic request tracing and performance monitoring
4. **Request Logging**: Structured HTTP request/response logging

## Development Roadmap

- [x] **Step 1**: Basic HTTP server with health endpoint and NewRelic integration
- [x] **Step 2**: Docker containerization for local development
- [ ] **Step 3**: Terraform infrastructure for AWS ECS Fargate
- [ ] **Step 4**: Enhanced NewRelic logging and custom events
- [ ] **Step 5**: Core bot structure (exchange, arbitrage, executor packages)
- [ ] **Step 6**: Simulation mode with mock data
- [ ] **Step 7**: Real-time market data integration
- [ ] **Step 8**: Live trade execution engine
- [ ] **Step 9**: Environment-driven configuration flags
- [ ] **Step 10**: Production monitoring and alerting

## Docker Usage

The application includes a complete Docker setup for local development:

### Quick Start with Docker

**🚀 Development (Recommended):**
```bash
# Build and run with basic NewRelic APM
make docker-run

# Run in background
make docker-run-detached

# View logs
make docker-logs

# Stop containers
make docker-down
```

**📊 Production/Staging (Full Monitoring):**
```bash
# Requires NEW_RELIC_LICENSE_KEY in .env
make docker-run-monitoring      # Includes NewRelic sidecars + log forwarding
make docker-logs-monitoring     # View all logs
make docker-down-monitoring     # Stop all monitoring
```

**💡 How it works:** Single `docker-compose.yml` with [profiles](https://docs.docker.com/compose/profiles/) - sidecars only start when using the `monitoring` profile.

### Docker Features
- **Multi-stage build**: Optimized image size with Go build stage
- **Health checks**: Built-in health monitoring
- **Environment variables**: Full configuration via .env file
- **Development ready**: Hot reload support and logging
- **Production ready**: Minimal Alpine-based final image
- **NewRelic sidecars**: Infrastructure monitoring and log forwarding

### NewRelic Monitoring

The application includes comprehensive NewRelic monitoring via sidecar containers:

#### Monitoring Components
1. **NewRelic Infrastructure Agent**: System-level monitoring (CPU, memory, disk, network)
2. **Fluent Bit Log Forwarder**: Structured log forwarding to NewRelic Logs
3. **Application Performance Monitoring**: Built-in Go agent integration

#### Running with Monitoring
```bash
# Development with basic monitoring
make docker-run

# Production-like monitoring with all sidecars
make docker-run-monitoring

# Background with full monitoring
make docker-run-monitoring-detached

# View all monitoring logs
make docker-logs-monitoring
```

#### Monitoring Setup
```bash
# Set your NewRelic license key
export NEW_RELIC_LICENSE_KEY=your_license_key_here

# Run with monitoring
make docker-run-monitoring
```

#### What Gets Monitored
- ✅ **Application metrics**: Request latency, throughput, errors
- ✅ **Infrastructure metrics**: CPU, memory, disk, network usage  
- ✅ **Container metrics**: Docker container performance
- ✅ **Structured logs**: Application logs with context and metadata
- ✅ **Custom events**: Trading events, arbitrage opportunities
- ✅ **Health checks**: Application and infrastructure health

## License

MIT License - see LICENSE file for details.