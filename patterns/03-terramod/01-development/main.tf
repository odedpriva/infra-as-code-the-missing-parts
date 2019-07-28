variable "environment" {}

module "bucket" {
  bucket = "${var.environment}.example.com"
}

module "sqs" {
  name = "${var.environment}"
}

# db - using docker

