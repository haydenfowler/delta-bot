# Testing Environment Configuration
environment    = "testing"
aws_region     = "ap-southeast-1"
project_name   = "delta-bot"

# VPC Configuration (different CIDR to avoid conflicts)
vpc_cidr             = "10.1.0.0/16"
public_subnet_cidrs  = ["10.1.1.0/24", "10.1.2.0/24"]

# ECS Configuration (smaller for testing)
app_name             = "delta-bot"
ecr_repository_name  = "delta-bot-testing"
task_cpu            = 256
task_memory         = 512
desired_count       = 1
container_port      = 8080

# NewRelic Configuration (set via environment variables)
# newrelic_license_key = "your-testing-newrelic-license-key"
# newrelic_app_name    = "delta-bot-testing"

# Trading Configuration (set via environment variables)
# binance_api_key    = "your-testing-binance-api-key"
# binance_secret_key = "your-testing-binance-secret-key"

# Trading Configuration (ALWAYS dry run for safety)
dry_run = true  # 🛡️ ENFORCED: Testing environment never executes real trades

# Monitoring (reduced for testing)
enable_newrelic_monitoring = false
log_retention_days        = 7