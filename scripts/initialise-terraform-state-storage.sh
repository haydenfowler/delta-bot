# Create S3 bucket for Terraform state
aws s3 mb s3://delta-bot-terraform-state --region ap-southeast-1

# Enable versioning on the bucket
aws s3api put-bucket-versioning \
  --bucket delta-bot-terraform-state \
  --versioning-configuration Status=Enabled

# Create DynamoDB table for state locking
aws dynamodb create-table \
  --table-name delta-bot-terraform-locks \
  --attribute-definitions AttributeName=LockID,AttributeType=S \
  --key-schema AttributeName=LockID,KeyType=HASH \
  --billing-mode PAY_PER_REQUEST \
  --region ap-southeast-1