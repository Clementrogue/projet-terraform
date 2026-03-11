variable "public_subnets" {
  type = list(string)
}

variable "private_subnets" {
  type = list(string)
}

variable "app_ami_id" {
  type = string
}

variable "iam_instance_profile_name" {
  type = string
}