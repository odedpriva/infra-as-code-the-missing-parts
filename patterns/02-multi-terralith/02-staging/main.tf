variable "environment" {}

resource "aws_s3_bucket" "bucket" {
  bucket = "${var.environment}.example.com"
}

resource "aws_sqs_queue" "sqs" {
  name = "${var.environment}"
}

resource "aws_db_instance" "rds" {
  name = "${var.environment}"
}

resource "aws_route53_record" "rds" {
  name    = "${var.environment}"
  records = "${aws_db_instance.rds.fqdn}"
}
