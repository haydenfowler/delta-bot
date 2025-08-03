# Delta Bot Architecture

## Overview

Delta Bot is built using clean architecture principles with a focus on maintainability, testability, and separation of concerns. The application is structured to support both local development and production deployment on AWS ECS Fargate.

## Architecture Layers

### 1. HTTP Layer (`internal/server`, `internal/handlers`, `internal/routes`, `internal/middleware`)

The HTTP layer handles all web-related concerns using the Echo framework:

- **Echo Server**: High-performance HTTP router with built-in middleware support
- **Handlers**: Pure business logic functions that return HTTP responses
- **Routes**: Centralized route configuration and middleware setup
- **Middleware**: Cross-cutting concerns like logging, monitoring, and security

#### Middleware Stack (in order):
1. **Recovery**: Catches panics and returns 500 errors gracefully
2. **CORS**: Handles cross-origin requests for browser compatibility
3. **NewRelic**: Automatic request tracing and APM integration
4. **Request Logging**: Structured logging of all HTTP requests/responses

### 2. Configuration Layer (`internal/config`)

Environment-based configuration management:
- Loads from `.env` files with fallback defaults
- Type-safe configuration structs
- Supports all deployment environments (local, staging, production)

### 3. Logging Layer (`internal/logger`)

Unified logging with observability:
- Structured logging to stdout
- NewRelic integration for custom events and error tracking
- Application performance monitoring (APM)
- Heartbeat events for health monitoring

### 4. Business Logic Layer (Future: `internal/exchange`, `internal/arb`, `internal/executor`)

Core trading functionality (to be implemented):
- **Exchange**: Integration with Binance API (REST + WebSocket)
- **Arbitrage**: Triangle detection and profit calculation
- **Executor**: Order placement and execution logic

## Design Principles

### 1. Dependency Injection
- Dependencies are injected through constructors
- Interfaces are used for testability
- No global state or singletons

### 2. Separation of Concerns
- HTTP concerns separate from business logic
- Configuration isolated from application logic
- Logging and monitoring as cross-cutting concerns

### 3. Error Handling
- Structured error handling with context
- Graceful degradation when external services fail
- Proper HTTP status codes and error responses

### 4. Observability
- Request tracing through NewRelic
- Structured logging for debugging
- Health checks for monitoring

## Request Flow

```
Client Request
    ↓
Echo Router
    ↓
Middleware Stack
    ↓
Route Handler
    ↓
Business Logic
    ↓
Response
```

### Detailed Flow:
1. **Client** sends HTTP request
2. **Echo Router** matches route and applies middleware
3. **Recovery Middleware** wraps request in panic handler
4. **CORS Middleware** adds appropriate headers
5. **NewRelic Middleware** starts transaction tracing
6. **Logging Middleware** logs request details
7. **Route Handler** processes business logic
8. **Response** returned through middleware stack in reverse

## Configuration Management

### Environment Variables
- `PORT`: HTTP server port (default: 8080)
- `LOG_LEVEL`: Logging verbosity (default: INFO)
- `NEW_RELIC_LICENSE_KEY`: APM license key (optional)
- `NEW_RELIC_APP_NAME`: Application name in NewRelic

### Configuration Loading Priority:
1. Environment variables
2. `.env` file
3. Default values

## Error Handling Strategy

### HTTP Errors
- 400 Bad Request: Client input validation errors
- 404 Not Found: Resource not found
- 500 Internal Server Error: Unexpected server errors
- 503 Service Unavailable: External service failures

### Logging Levels
- **INFO**: Normal operation events
- **ERROR**: Error conditions that need attention
- **DEBUG**: Detailed diagnostic information (development only)

### NewRelic Integration
- Automatic error tracking and alerting
- Custom events for business metrics
- Performance monitoring and bottleneck detection

## Testing Strategy

### Unit Tests
- Handler functions with mocked dependencies
- Business logic with isolated unit tests
- Configuration parsing and validation

### Integration Tests
- End-to-end HTTP request/response testing
- Database and external API integration
- Middleware functionality verification

### Load Testing
- Performance benchmarks for high-frequency trading
- Concurrent request handling
- Memory and CPU profiling

## Security Considerations

### Current Measures
- CORS middleware for browser security
- Panic recovery to prevent information leakage
- Structured logging to avoid sensitive data exposure

### Future Enhancements
- API key authentication for trading endpoints
- Rate limiting for DDoS protection
- Input validation and sanitization
- HTTPS enforcement in production

## Monitoring and Observability

### Health Checks
- `/health` endpoint returns application status
- NewRelic heartbeat events for monitoring
- Graceful shutdown handling

### Metrics and Alerts
- Request latency and throughput
- Error rates and patterns
- Trading performance metrics
- Infrastructure resource usage

## Deployment Architecture

### Local Development
- Docker Compose for service orchestration
- Hot reload for rapid development
- Local environment variable configuration

### Production (AWS ECS Fargate)
- Containerized application deployment
- Auto-scaling based on CPU/memory usage
- Load balancing across multiple instances
- CloudWatch logs integration
- NewRelic APM for performance monitoring