variable "region" {
  default = "us-east-1"
}

variable "project_name" {
  default = "student-phase6"
}

variable "key_name" {
  type = string
}

variable "my_ip" {
  type = string
}

variable "instance_type" {
  default = "t3.medium"
}

variable "docker_image" {
  type = string
}

variable "db_host" {
  type = string
}

variable "db_user" {
  default = "nodeapp"
}

variable "db_password" {
  type      = string
  sensitive = true
}

variable "db_name" {
  default = "STUDENTS"
}
