# Delete DynamoDB table
aws dynamodb delete-table \
  --table-name delta-bot-terraform-locks \
  --region ap-southeast-1

# Delete all objects and versions from S3 bucket
aws s3api delete-objects \
  --bucket delta-bot-terraform-state \
  --delete "$(aws s3api list-object-versions \
    --bucket delta-bot-terraform-state \
    --output json \
    --query '{Objects: Versions[].{Key:Key,VersionId:VersionId}}')"

# Delete all delete markers from S3 bucket
aws s3api delete-objects \
  --bucket delta-bot-terraform-state \
  --delete "$(aws s3api list-object-versions \
    --bucket delta-bot-terraform-state \
    --output json \
    --query '{Objects: DeleteMarkers[].{Key:Key,VersionId:VersionId}}')"

# Delete S3 bucket
aws s3 rb s3://delta-bot-terraform-state --region ap-southeast-1