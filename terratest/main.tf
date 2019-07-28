provider "aws" {
  region = "eu-west-2"
}

terraform {
  backend "s3" {
    bucket = "infra-as-code-the-missing-parts"
    key    = "terratest/terraform.tfstate"
    region = "eu-west-1"
  }
}

variable "cidr_block" {}

resource "aws_vpc" "selected" {
  cidr_block = "${var.cidr_block}"
}
