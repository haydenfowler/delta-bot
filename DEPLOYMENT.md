# 🚀 AWS Deployment Guide

Complete guide for deploying Delta Bot to AWS ECS Fargate using Terraform.

## 📋 Prerequisites

Before starting, you'll need:
- An AWS account
- A computer with terminal access (Mac/Linux/Windows WSL)
- Basic command line knowledge

## 🔧 Step 1: Install Required Tools

### Install AWS CLI
```bash
# macOS (using Homebrew)
brew install awscli

# Linux
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install

# Windows (PowerShell as Administrator)
msiexec.exe /i https://awscli.amazonaws.com/AWSCLIV2.msi
```

### Install Terraform
```bash
# macOS (using Homebrew)
brew install terraform

# Linux (using snap)
sudo snap install terraform

# Or download from: https://www.terraform.io/downloads.html
```

### Verify Installations
```bash
aws --version
terraform --version
```

## 🔑 Step 2: Configure AWS Credentials

### Option A: AWS CLI Configure (Recommended)
```bash
aws configure
```

You'll be prompted for:
- **AWS Access Key ID**: From your AWS IAM user
- **AWS Secret Access Key**: From your AWS IAM user  
- **Default region name**: `ap-southeast-1` (or your preferred region)
- **Default output format**: `json`

### Option B: Environment Variables
```bash
export AWS_ACCESS_KEY_ID="your-access-key"
export AWS_SECRET_ACCESS_KEY="your-secret-key"
export AWS_DEFAULT_REGION="ap-southeast-1"
```

### Getting AWS Credentials

If you don't have AWS credentials:

1. **Sign in to AWS Console**: https://console.aws.amazon.com/
2. **Navigate to IAM** → Users → Create User
3. **Attach policies**: `AdministratorAccess` (for simplicity) 
4. **Create Access Key** → Save the Access Key ID and Secret Access Key

⚠️ **Security Note**: In production, use least-privilege IAM policies instead of `AdministratorAccess`.

## ⚙️ Step 3: Configure Terraform Variables

1. **Copy the example variables file:**
   ```bash
   cp terraform/terraform.tfvars.example terraform/terraform.tfvars
   ```

2. **Edit `terraform/terraform.tfvars`** with your values:
   ```hcl
   # AWS Configuration
   aws_region  = "ap-southeast-1"
   environment = "production"

   # Application Configuration  
   new_relic_license_key = "your-actual-newrelic-license-key"
   new_relic_app_name    = "delta-bot"

   # Trading Configuration (keep dry_run = true for testing)
   dry_run               = true
   binance_api_key       = "your-binance-api-key"
   binance_secret_key    = "your-binance-secret-key"
   min_profit_threshold  = 0.5
   max_trade_amount      = 1000

   # ECS Configuration
   cpu           = 256
   memory        = 512
   desired_count = 1
   ```

## 🏗️ Step 4: Initialize and Plan Infrastructure

1. **Initialize Terraform:**
   ```bash
   make tf-init
   ```

2. **Format and validate:**
   ```bash
   make tf-fmt
   make tf-validate
   ```

3. **Plan the deployment:**
   ```bash
   make tf-plan
   ```
   
   Review the output to see what AWS resources will be created.

## 🐳 Step 5: Build and Push Docker Image

1. **Create ECR repository first:**
   ```bash
   # Apply just the ECR repository
   cd terraform
   terraform apply -target=aws_ecr_repository.delta_bot
   ```

2. **Get ECR repository URL:**
   ```bash
   make tf-output
   ```
   
   Look for `ecr_repository_url` in the output.

3. **Build and tag your image:**
   ```bash
   # Build the Docker image
   docker build -t delta-bot .
   
   # Tag for ECR (replace with your ECR URL from step 2)
   docker tag delta-bot:latest 123456789012.dkr.ecr.ap-southeast-1.amazonaws.com/delta-bot:latest
   ```

4. **Login to ECR and push:**
   ```bash
   # Get login command (replace region if different)
   aws ecr get-login-password --region ap-southeast-1 | docker login --username AWS --password-stdin 123456789012.dkr.ecr.ap-southeast-1.amazonaws.com
   
   # Push the image
   docker push 123456789012.dkr.ecr.ap-southeast-1.amazonaws.com/delta-bot:latest
   ```

## 🚀 Step 6: Deploy Infrastructure

1. **Deploy everything:**
   ```bash
   make tf-apply
   ```
   
   Type `yes` when prompted to confirm the deployment.

2. **Wait for deployment** (5-10 minutes for all resources)

3. **Get deployment outputs:**
   ```bash
   make tf-output
   ```

## ✅ Step 7: Test Your Deployment

1. **Get the application URL:**
   ```bash
   make tf-output | grep application_url
   ```

2. **Test the health endpoint:**
   ```bash
   # Replace with your actual URL from step 1
   curl http://your-load-balancer-url.amazonaws.com/health
   ```

   You should see:
   ```json
   {"status":"healthy","timestamp":"2025-08-03T..."}
   ```

3. **Check CloudWatch logs:**
   - Go to AWS Console → CloudWatch → Log Groups
   - Find `/ecs/delta-bot`
   - View recent logs to see your application starting

## 📊 Step 8: Monitor Your Application

### AWS Console Monitoring
- **ECS Console**: Check service status and task health
- **CloudWatch**: View logs and metrics
- **ALB Console**: Monitor load balancer health

### NewRelic Monitoring
- Your app should appear in NewRelic dashboard
- View APM metrics, errors, and custom events

## 🛠️ Common Commands

```bash
# View all outputs
make tf-output

# Check service status
aws ecs describe-services --cluster delta-bot-cluster --services delta-bot-service

# View logs
aws logs tail /ecs/delta-bot --follow

# Update the application (after code changes)
# 1. Build and push new image
# 2. Force new deployment
aws ecs update-service --cluster delta-bot-cluster --service delta-bot-service --force-new-deployment
```

## 🧹 Cleanup (When Done Testing)

**⚠️ This will delete ALL AWS resources and cost you nothing further:**

```bash
make tf-destroy
```

Type `yes` to confirm deletion.

## 💰 Cost Estimate

**Monthly costs for minimal setup:**
- ALB: ~$16/month
- ECS Fargate (256 CPU, 512 MB): ~$13/month  
- CloudWatch Logs: ~$1/month
- **Total: ~$30/month**

## 🐛 Troubleshooting

### Common Issues

1. **"No credentials found"**
   - Run `aws configure` or set environment variables
   - Verify with `aws sts get-caller-identity`

2. **Docker push fails**
   - Ensure you're logged into ECR: `aws ecr get-login-password`
   - Check your ECR repository URL is correct

3. **ECS tasks not starting**
   - Check CloudWatch logs for error messages
   - Verify your secrets are set correctly in terraform.tfvars

4. **Health check failing**
   - Ensure your app starts on port 8080
   - Check security group allows ALB → ECS communication

### Get Help
- Check AWS CloudWatch logs first
- Verify your `terraform.tfvars` configuration
- Ensure all required environment variables are set

## 🎯 What You Built

Your infrastructure includes:
- ✅ **VPC with public subnets** across 2 AZs
- ✅ **Application Load Balancer** with health checks  
- ✅ **ECS Fargate cluster** running your containerized app
- ✅ **IAM roles** with least-privilege permissions
- ✅ **CloudWatch logging** and monitoring alarms
- ✅ **SSM Parameter Store** for secure secret management
- ✅ **ECR repository** for your Docker images

**🎉 Congratulations! Your Delta Bot is now running on AWS!**