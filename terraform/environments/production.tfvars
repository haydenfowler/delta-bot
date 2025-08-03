# Production Environment Configuration
environment    = "production"
aws_region     = "ap-southeast-1"
project_name   = "delta-bot"

# VPC Configuration
vpc_cidr             = "10.0.0.0/16"
public_subnet_cidrs  = ["10.0.1.0/24", "10.0.2.0/24"]

# ECS Configuration
app_name             = "delta-bot"
ecr_repository_name  = "delta-bot"
task_cpu            = 512
task_memory         = 1024
desired_count       = 2
container_port      = 8080

# NewRelic Configuration (set via environment variables)
# newrelic_license_key = "your-production-newrelic-license-key"
# newrelic_app_name    = "delta-bot-production"

# Trading Configuration (set via environment variables)
# binance_api_key    = "your-production-binance-api-key"
# binance_secret_key = "your-production-binance-secret-key"

# Trading Configuration
dry_run = false  # ⚠️  REAL TRADING ENABLED: Set to true to disable real trades

# Monitoring
enable_newrelic_monitoring = true
log_retention_days        = 30