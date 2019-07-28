variable "environment" {}

module "bucket" {
  source = "s3-module"
  bucket = "${var.environment}.example.com"
}

module "sqs" {
  source = "sqs-module"
  name   = "${var.environment}"
}

module "rds" {
  source        = "rds-module"
  name          = "${var.environment}"
  create_record = true
}
