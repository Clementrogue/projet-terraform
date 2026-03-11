variable "region" {
  default = "us-east-1"
}

variable "ami" {
  description = "Ubuntu AMI"
}

variable "instance_type" {
  default = "t2.micro"
}

variable "key_name" {
  description = "AWS key pair"
}
