# Backend configuration for multiple environments
# This will be configured per environment using backend config files

terraform {
  backend "s3" {
    # These values will be provided via backend config files
    # bucket         = "your-terraform-state-bucket"
    # key            = "environments/{environment}/terraform.tfstate"
    # region         = "ap-southeast-1"
    # encrypt        = true
    # dynamodb_table = "terraform-state-locks"
  }
}

# S3 bucket for storing Terraform state (create this manually first)
# resource "aws_s3_bucket" "terraform_state" {
#   bucket = "${var.project_name}-terraform-state-${random_string.bucket_suffix.result}"
# }

# resource "random_string" "bucket_suffix" {
#   length  = 8
#   special = false
#   upper   = false
# }

# resource "aws_s3_bucket_versioning" "terraform_state" {
#   bucket = aws_s3_bucket.terraform_state.id
#   versioning_configuration {
#     status = "Enabled"
#   }
# }

# resource "aws_s3_bucket_server_side_encryption_configuration" "terraform_state" {
#   bucket = aws_s3_bucket.terraform_state.id
#   rule {
#     apply_server_side_encryption_by_default {
#       sse_algorithm = "AES256"
#     }
#   }
# }

# resource "aws_dynamodb_table" "terraform_locks" {
#   name           = "${var.project_name}-terraform-locks"
#   billing_mode   = "PAY_PER_REQUEST"
#   hash_key       = "LockID"

#   attribute {
#     name = "LockID"
#     type = "S"
#   }
# }