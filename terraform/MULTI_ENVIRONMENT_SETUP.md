# Multi-Environment Terraform Setup

This guide explains how to deploy and manage multiple environments (production and testing) using Terraform.

## 🏗️ Architecture Overview

The setup supports:
- **Production Environment**: Full-scale deployment with monitoring
- **Testing Environment**: Smaller, cost-optimized deployment
- **Isolated State**: Separate Terraform state files per environment
- **Environment-Specific Configs**: Different settings per environment

## 📁 Directory Structure

```
terraform/
├── environments/
│   ├── production.tfvars      # Production configuration
│   ├── testing.tfvars         # Testing configuration
│   ├── backend-production.conf # Production state backend
│   └── backend-testing.conf   # Testing state backend
├── *.tf files                 # Terraform resources
└── terraform.tfvars.example  # Template (deprecated)
```

## 🚀 Prerequisites

1. **AWS CLI configured** with appropriate permissions
2. **Terraform installed** (>= 1.0)
3. **S3 bucket for state storage** (create manually first)
4. **DynamoDB table for locking** (create manually first)

### Create State Storage Resources

```bash
# Create S3 bucket for Terraform state
aws s3 mb s3://delta-bot-terraform-state --region ap-southeast-1

# Enable versioning
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
```

## 🛠️ Environment Setup

### Step 1: Set Environment Variables

Create environment-specific `.env` files:

**Production (.env.production):**
```bash
export TF_VAR_newrelic_license_key="your-production-newrelic-key"
export TF_VAR_newrelic_app_name="delta-bot-production"
export TF_VAR_binance_api_key="your-production-binance-key"
export TF_VAR_binance_secret_key="your-production-binance-secret"
```

**Testing (.env.testing):**
```bash
export TF_VAR_newrelic_license_key="your-testing-newrelic-key"
export TF_VAR_newrelic_app_name="delta-bot-testing"
export TF_VAR_binance_api_key="your-testing-binance-key"
export TF_VAR_binance_secret_key="your-testing-binance-secret"
```

### Step 2: Initialize Environments

**Production:**
```bash
# Source production environment
source .env.production

# Initialize Terraform with production backend
make tf-init-prod

# Plan and apply
make tf-plan-prod
make tf-apply-prod
```

**Testing:**
```bash
# Source testing environment
source .env.testing

# Initialize Terraform with testing backend
make tf-init-test

# Plan and apply
make tf-plan-test
make tf-apply-test
```

## 📊 Environment Differences

| Resource | Production | Testing |
|----------|------------|---------|
| **VPC CIDR** | 10.0.0.0/16 | 10.1.0.0/16 |
| **ECS CPU** | 512 | 256 |
| **ECS Memory** | 1024 MB | 512 MB |
| **Desired Count** | 2 tasks | 1 task |
| **Monitoring** | Full NewRelic | Disabled |
| **Log Retention** | 30 days | 7 days |
| **ECR Repository** | delta-bot | delta-bot-testing |
| **Trading Mode** | Real trades (`dry_run=false`) | **Shadow only** (`dry_run=true`) |

## 🔄 Workflow Examples

### Deploy to Testing First
```bash
# 1. Test your changes
source .env.testing
make tf-plan-test
make tf-apply-test

# 2. Verify deployment works
curl https://your-testing-alb-url/health

# 3. Deploy to production
source .env.production
make tf-plan-prod
make tf-apply-prod
```

### Environment-Specific Operations

**Switch between environments:**
```bash
# Work with testing
source .env.testing
make tf-init-test

# Work with production  
source .env.production
make tf-init-prod
```

**View environment outputs:**
```bash
# After initialization
terraform output -json
```

## 🛡️ Best Practices

### 1. **Always Deploy to Testing First**
```bash
# ✅ Good workflow
make tf-apply-test  # Test first
make tf-apply-prod  # Then production

# ❌ Avoid direct production changes
make tf-apply-prod  # Without testing
```

### 2. **Use Environment Variables for Secrets**
```bash
# ✅ Good - environment variables
export TF_VAR_binance_api_key="secret"

# ❌ Bad - hardcoded in tfvars
echo 'binance_api_key = "secret"' >> production.tfvars
```

### 3. **Separate ECR Repositories**
- Production: `delta-bot`
- Testing: `delta-bot-testing`

### 4. **Resource Naming Convention**
Resources are automatically tagged and named with environment:
- Production: `delta-bot-production-*`
- Testing: `delta-bot-testing-*`

## 🔍 Troubleshooting

### State File Issues
```bash
# If state gets corrupted
terraform refresh -var-file="environments/production.tfvars"

# Force unlock (use carefully)
terraform force-unlock <LOCK_ID>
```

### Backend Configuration Issues
```bash
# Reconfigure backend
terraform init -reconfigure -backend-config="environments/backend-production.conf"
```

### Variable Issues
```bash
# Check what variables are set
terraform console
> var.environment
```

## 📈 Cost Optimization

**Testing Environment** is configured for minimal cost:
- Smaller ECS tasks (256 CPU, 512 MB memory)
- Single task instance
- Reduced log retention (7 days)
- NewRelic monitoring disabled

**Production Environment** is optimized for reliability:
- Larger ECS tasks (512 CPU, 1024 MB memory)  
- Multiple task instances (2)
- Extended log retention (30 days)
- Full monitoring enabled

## 🛡️ Trading Safety Features

### **Environment-Specific Trading Modes**

- **Testing Environment**: 
  - `dry_run = true` **ENFORCED** by Terraform
  - Never executes real trades - only shadow/simulation mode
  - Safe for testing strategies and code changes

- **Production Environment**:
  - `dry_run = false` (configurable)
  - Can execute real trades when enabled
  - Defaults to `true` for safety

### **Safety Workflow**
```bash
# 1. Always test in shadow mode first
make tf-apply-test    # dry_run=true enforced

# 2. Verify strategy works without risk  
curl https://testing-alb.../health

# 3. Only then deploy to production
make tf-apply-prod    # dry_run=false (if configured)
```

## 🔐 Security Considerations

1. **Never commit `.env.*` files** to git
2. **Use SSM Parameter Store** for secrets in production
3. **Enable S3 bucket encryption** for state files
4. **Use IAM roles** with minimal permissions
5. **Regular credential rotation**
6. **Test strategies in dry-run mode first**

## 📋 Migration from Single Environment

If you have existing single-environment setup:

1. **Backup existing state:**
   ```bash
   cp terraform.tfstate terraform.tfstate.backup
   ```

2. **Create environment configs:**
   ```bash
   cp terraform.tfvars.example environments/production.tfvars
   # Edit with production values
   ```

3. **Initialize with backend:**
   ```bash
   make tf-init-prod
   ```

4. **Import existing resources:**
   ```bash
   terraform import -var-file="environments/production.tfvars" aws_vpc.main vpc-xxxxxxxx
   # Repeat for all resources
   ```