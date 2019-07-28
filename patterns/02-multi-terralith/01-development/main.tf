variable "environment" {}

resource "aws_s3_bucket" "bucket" {
  bucket = "${var.environment}.example.com"
}

resource "aws_sqs_queue" "sqs" {
  name = "${var.environment}"
}

# db - using docker

