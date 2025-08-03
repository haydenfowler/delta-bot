# VPC Outputs
output "vpc_id" {
  description = "VPC ID"
  value       = aws_vpc.main.id
}

output "public_subnet_ids" {
  description = "Public subnet IDs"
  value       = aws_subnet.public[*].id
}

# ECS Outputs
output "ecs_cluster_name" {
  description = "ECS cluster name"
  value       = aws_ecs_cluster.main.name
}

output "ecs_service_name" {
  description = "ECS service name"
  value       = aws_ecs_service.delta_bot.name
}

# Load Balancer Outputs
output "load_balancer_dns" {
  description = "Load balancer DNS name"
  value       = aws_lb.main.dns_name
}

output "application_url" {
  description = "Application URL"
  value       = "http://${aws_lb.main.dns_name}"
}

output "health_check_url" {
  description = "Health check URL"
  value       = "http://${aws_lb.main.dns_name}/health"
}

# CloudWatch Outputs
output "cloudwatch_log_group" {
  description = "CloudWatch log group name"
  value       = aws_cloudwatch_log_group.delta_bot.name
}

# ECR Repository
output "ecr_repository_url" {
  description = "ECR repository URL"
  value       = aws_ecr_repository.delta_bot.repository_url
}