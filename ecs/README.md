# Infrastructure as Code for 1on1 Playground

Welcome to the 1on1 Playground project infrastructure setup! This guide outlines the steps to deploy your environment on AWS using ECS with Fargate in the `ap-southeast-2` region.

## Prerequisites

Before starting, ensure you have the following:

- **AWS CLI:** Configured with the correct credentials.
- **Docker:** For building and pushing the Docker image.
- **Docker Image:** Available in Docker Hub or AWS ECR.

## Infrastructure Components

Our AWS setup includes:

- **ECS (Elastic Container Service):** Using Fargate for serverless task execution.
- **ECR (Elastic Container Registry):** To store Docker images.
- **VPC (Virtual Private Cloud):** Utilizing the default VPC.
- **Subnets:** Using at least two for high availability.
- **Security Group:** Defining ECS task network rules.
- **Application Load Balancer (ALB):** To distribute incoming traffic.

## Step-by-Step Setup

### 1. Build and Push the Docker Image

Navigate to the `ecs` directory and run the `build-push-ecr.sh` script:

```bash
cd ecs
./build-push-ecr.sh
```

This script creates the ECR repository and pushes your Docker image.

### 2. Deploy AWS Infrastructure with Terraform

Navigate to the `terraform` subdirectory:

```bash
cd terraform
```

Execute the following Terraform commands:

```bash
terraform init
terraform plan
terraform apply
```

### 3. Testing and Verification

After Terraform successfully applies:

- **Access the Application**: Visit the ALB DNS URL (`1on1-playground-alb-635238929.ap-southeast-2.elb.amazonaws.com`) in your browser.
- **ECS Console Check**: Verify the ECS service and tasks are running correctly.
- **ALB Monitoring**: In AWS Console, check the target group for proper ECS task routing.

### DNS Configuration

Optionally, configure a DNS record to point to the ALB's DNS name for easier access.

## Considerations

- **Security Settings**: Review and maintain security group settings for ALB and ECS.
- **Health Checks**: Regularly monitor the target group's health checks.
- **SSL/TLS**: Utilize AWS Certificate Manager for secure HTTPS communication.

## Testing the AI-Enhanced One-on-One Meeting Assistant Endpoints

### Testing `generate-questions` Endpoint

This endpoint generates a set of questions for a one-on-one meeting. It does not require any input parameters and can be tested using a simple GET request.

**Command:**

```bash
curl http://localhost:1323/generate-questions
```

**Expected Response:** A list of questions formatted in Markdown, designed to facilitate a one-on-one meeting.

### Testing `generate-report` Endpoint

This endpoint generates a summary report based on the responses from a one-on-one meeting. It requires the meeting responses to be sent as URL-encoded text in the POST request.

**Command:**

```bash
curl -X POST http://localhost:1323/generate-report \
     -H "Content-Type: application/x-www-form-urlencoded" \
     -d "responses=The%20new%20marketing%20project%20is%20shaping%20up%20well%2C%20facing%20some%20team%20alignment%20challenges.%20Interested%20in%20improving%20leadership%20skills.%20Creating%20content%20strategies%20is%20most%20satisfying.%20Time%20management%20is%20a%20current%20hurdle.%20Enjoy%20the%20team%27s%20collaborative%20spirit.%20Keen%20on%20exploring%20product%20management.%20Hiking%20helps%20maintain%20work-life%20balance.%20Interested%20in%20learning%20about%20R%26D.%20Proud%20of%20the%20positive%20client%20feedback%20on%20our%20last%20campaign.%20Would%20advise%20past%20self%20to%20be%20open%20to%20feedback.%20Interested%20in%20a%20monthly%20team%20recognition%20program.%20Digital%20marketing%20analytics%20was%20a%20fascinating%20learning%20area.%20Planning%20a%20road%20trip%20next%20month.%20Feel%20empowered%20at%20level%208%2C%20seeking%20more%20decision-making%20opportunities.%20%27Never%20stop%20learning%27%20has%20been%20pivotal%20advice."
```

**Expected Response:** A comprehensive report summarizing key points and actionable items from the provided responses, formatted in plain text or Markdown.

**Note:** The responses in the `generate-report` command are an example. You should replace them with actual responses from a one-on-one meeting when testing this endpoint.

## Notes

- The setup utilizes the default VPC and subnets.
- The ALB ensures the application can handle traffic directed to the ECS service.
- Monitor AWS costs and usage, particularly for Fargate tasks and the ALB.

## Conclusion

Your 1on1 Playground application is now live on AWS ECS using Fargate. This infrastructure-as-code approach ensures a manageable and scalable environment for your application.
