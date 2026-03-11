variable "region" {
  default = "us-east-1"
}

variable "project_name" {
  default = "student-phase2"
}

variable "ami" {
  description = "Ubuntu AMI ID"
  type        = string
}

variable "instance_type" {
  default = "t2.micro"
}

variable "key_name" {
  description = "AWS key pair name"
  type        = string
}

variable "my_ip" {
  description = "Your public IP in CIDR format"
  type        = string
}

variable "db_name" {
  default = "STUDENTS"
}

variable "db_username" {
  default = "nodeapp"
}

variable "db_password" {
  description = "RDS password"
  type        = string
  sensitive   = true
}

variable "db_instance_class" {
  default = "db.t3.micro"
}

variable "secret_name" {
  default = "student-phase2-db-secret"
}
