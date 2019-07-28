terraform {
  required_version = "~> 0.12"
}

provider "aws" {
  region = "eu-west-2"
}

variable "cidr_block" {
  default = "172.31.0.0/16"
}

resource "aws_vpc" "selected" {
  cidr_block = var.cidr_block
}

