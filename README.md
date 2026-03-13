# Phase 6 — Deployment of Application on AWS ECS using Terraform

## Project Overview

This phase focuses on deploying a containerized application on **Amazon ECS (Elastic Container Service)** using **Terraform** as Infrastructure as Code.

The objective is to provision a complete AWS infrastructure capable of running a Docker container on an ECS cluster using EC2 instances.

The infrastructure includes:

* A **VPC**
* Public **subnets**
* **Internet Gateway**
* **Route tables**
* **Security groups**
* **ECS Cluster**
* **EC2 instances running ECS agent**
* **Auto Scaling Group**
* **Task Definition**
* **ECS Service**

The deployed container exposes a web application accessible through the public IP of the ECS instance.

---

# Architecture

The infrastructure deployed by Terraform follows this architecture:

```
Internet
   │
   │
Internet Gateway
   │
   ▼
VPC (10.2.0.0/16)
   │
   ├── Public Subnet A (10.2.1.0/24)
   │        │
   │        └── ECS EC2 Instance
   │                │
   │                └── Docker Container (Application)
   │
   └── Public Subnet B (10.2.2.0/24)

```

Components used:

| Component           | Description                            |
| ------------------- | -------------------------------------- |
| VPC                 | Custom network for the infrastructure  |
| Subnets             | Public subnets where EC2 instances run |
| Internet Gateway    | Allows internet access                 |
| Security Group      | Controls inbound/outbound traffic      |
| ECS Cluster         | Container orchestration service        |
| Launch Template     | Configuration of EC2 instances         |
| Auto Scaling Group  | Ensures ECS instance availability      |
| ECS Task Definition | Defines container configuration        |
| ECS Service         | Ensures tasks remain running           |

---

# Prerequisites

Before running this project, ensure the following tools are installed:

* Terraform >= 1.3
* AWS CLI
* Docker
* AWS account
* SSH key pair

Configure AWS credentials:

```bash
aws configure
```

---

# Project Structure

```
phase6-ecs/
│
├── main.tf
├── variables.tf
├── terraform.tfvars
├── outputs.tf
└── README.md
```

---

# Terraform Deployment

Initialize Terraform:

```bash
terraform init
```

Validate configuration:

```bash
terraform validate
```

Check execution plan:

```bash
terraform plan
```

Deploy infrastructure:

```bash
terraform apply
```

Terraform will provision:

* VPC
* Networking
* ECS Cluster
* EC2 instance
* Docker container service

---

# Accessing the Application

After deployment, Terraform outputs the public subnet and cluster information.

Retrieve the public IP of the ECS instance:

```bash
aws ec2 describe-instances \
--filters "Name=tag:Name,Values=student-phase6-ecs-instance" \
--query "Reservations[*].Instances[*].PublicIpAddress" \
--output text
```

Then open the application in a browser:

```
http://PUBLIC_IP
```

or

```
http://PUBLIC_IP:81
```

depending on the configured port mapping.

---

# Verifying ECS Deployment

Check ECS cluster:

```bash
aws ecs list-clusters
```

Check running tasks:

```bash
aws ecs list-tasks --cluster student-phase6-cluster
```

Check ECS service:

```bash
aws ecs describe-services \
--cluster student-phase6-cluster \
--services student-phase6-service
```

---

# Accessing the ECS Instance

SSH connection:

```bash
ssh -i vockey.pem ec2-user@PUBLIC_IP
```

Verify running containers:

```bash
docker ps
```

Check container logs:

```bash
docker logs CONTAINER_ID
```

---

# Troubleshooting

If the application is not accessible:

### Check ECS tasks

```bash
aws ecs list-tasks --cluster student-phase6-cluster
```

### Check ECS service events

```bash
aws ecs describe-services \
--cluster student-phase6-cluster \
--services student-phase6-service \
--query "services[0].events"
```

### Verify container status on EC2

```bash
docker ps -a
```

### Check if port is listening

```bash
sudo ss -tulpn | grep :80
```

---

# Destroy Infrastructure

To remove all resources:

```bash
terraform destroy
```

---

# Technologies Used

* **AWS ECS**
* **EC2**
* **Terraform**
* **Docker**
* **AWS CLI**

---

# Learning Outcomes

Through this phase, the following skills were developed:

* Infrastructure as Code with Terraform
* Container deployment on ECS
* AWS networking configuration
* ECS task and service management
* Debugging containerized applications
