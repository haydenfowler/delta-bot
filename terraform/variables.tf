# General Configuration
variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = "production"
}

variable "project_name" {
  description = "Project name"
  type        = string
  default     = "delta-bot"
}

# VPC Configuration
variable "vpc_cidr" {
  description = "CIDR block for VPC"
  type        = string
  default     = "10.0.0.0/16"
}

variable "public_subnet_cidrs" {
  description = "CIDR blocks for public subnets"
  type        = list(string)
  default     = ["10.0.1.0/24", "10.0.2.0/24"]
}

# ECS Configuration
variable "cpu" {
  description = "CPU units for the task (256, 512, 1024, 2048, 4096)"
  type        = number
  default     = 256
}

variable "memory" {
  description = "Memory (MB) for the task"
  type        = number
  default     = 512
}

variable "desired_count" {
  description = "Desired number of tasks"
  type        = number
  default     = 1
}

variable "container_port" {
  description = "Port the container listens on"
  type        = number
  default     = 8080
}

# Application Configuration
variable "new_relic_license_key" {
  description = "NewRelic license key"
  type        = string
  sensitive   = true
}

variable "new_relic_app_name" {
  description = "NewRelic application name"
  type        = string
  default     = "delta-bot"
}

variable "dry_run" {
  description = "Run in dry-run mode"
  type        = bool
  default     = true
}

variable "binance_api_key" {
  description = "Binance API key"
  type        = string
  sensitive   = true
}

variable "binance_secret_key" {
  description = "Binance secret key"
  type        = string
  sensitive   = true
}

variable "min_profit_threshold" {
  description = "Minimum profit threshold for trades"
  type        = number
  default     = 0.5
}

variable "max_trade_amount" {
  description = "Maximum trade amount"
  type        = number
  default     = 1000
}