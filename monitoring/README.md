# NewRelic Monitoring Setup

This directory contains configuration files for NewRelic monitoring sidecars.

## Components

### 1. NewRelic Infrastructure Agent (`newrelic-infra`)
- **Purpose**: System-level monitoring (CPU, memory, disk, network)
- **Config**: `newrelic-infra.yml`
- **Image**: `newrelic/infrastructure:latest`

### 2. Fluent Bit Log Forwarder (`fluent-bit`)
- **Purpose**: Structured log forwarding to NewRelic Logs
- **Config**: `fluent-bit.conf` (dev), `fluent-bit.prod.conf` (production)
- **Parsers**: `parsers.conf`
- **Image**: `newrelic/newrelic-fluent-bit-output:latest`

## Configuration Files

### `fluent-bit.conf`
Development configuration for log forwarding:
- Collects application logs from `/var/log/delta-bot/`
- Collects Docker container logs
- Adds service metadata
- Forwards to NewRelic Logs API

### `fluent-bit.prod.conf`
Production configuration with enhanced features:
- Persistent storage for reliability
- Error detection and alerting
- Backup logging to files
- Higher buffer limits and workers

### `parsers.conf`
Log parsing rules for:
- JSON formatted logs
- Structured application logs
- Docker container logs

### `newrelic-infra.yml`
Infrastructure agent configuration:
- Custom attributes for service identification
- Docker integration enabled
- Log forwarding (alternative to Fluent Bit)

## Usage

### Development
```bash
make docker-run
```

### Production-like monitoring
```bash
make docker-run-monitoring
```

### View monitoring logs
```bash
make docker-logs-monitoring
```

## Environment Variables

Required:
- `NEW_RELIC_LICENSE_KEY`: Your NewRelic license key

Optional:
- `ENVIRONMENT`: Environment name (development, staging, production)
- `VERSION`: Application version
- `CLUSTER_NAME`: Cluster identifier
- `AWS_REGION`: AWS region for production deployments

## Monitoring Data

### Infrastructure Metrics
- CPU usage and load
- Memory utilization
- Disk I/O and space
- Network traffic
- Container performance

### Application Logs
- Structured application logs
- Request/response logs
- Error logs with stack traces
- Custom trading events

### Custom Events
- Health check events
- Arbitrage opportunity detection
- Trading execution results
- Performance metrics

## Troubleshooting

### Check Fluent Bit Status
```bash
curl http://localhost:2020/api/v1/health
```

### View Fluent Bit Metrics
```bash
curl http://localhost:2020/api/v1/metrics
```

### Check Infrastructure Agent
```bash
docker logs newrelic-infra
```

### Verify Log Forwarding
Check NewRelic Logs UI for incoming log data with service tag `delta-bot`.

## Production Considerations

1. **Security**: Ensure NewRelic license key is properly secured
2. **Resource limits**: Set appropriate CPU/memory limits for sidecars
3. **Storage**: Configure persistent storage for Fluent Bit in production
4. **Networking**: Ensure outbound HTTPS access to NewRelic APIs
5. **Backup**: Enable backup logging for critical environments