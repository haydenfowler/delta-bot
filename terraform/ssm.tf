# SSM Parameters for sensitive data
resource "aws_ssm_parameter" "new_relic_license_key" {
  name  = "/${var.project_name}/new-relic-license-key"
  type  = "SecureString"
  value = var.new_relic_license_key

  tags = {
    Name = "${var.project_name}-new-relic-license-key"
  }
}

resource "aws_ssm_parameter" "binance_api_key" {
  name  = "/${var.project_name}/binance-api-key"
  type  = "SecureString"
  value = var.binance_api_key

  tags = {
    Name = "${var.project_name}-binance-api-key"
  }
}

resource "aws_ssm_parameter" "binance_secret_key" {
  name  = "/${var.project_name}/binance-secret-key"
  type  = "SecureString"
  value = var.binance_secret_key

  tags = {
    Name = "${var.project_name}-binance-secret-key"
  }
}