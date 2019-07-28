module "sqs" {
  source = "sqs-module"
  name   = "${var.environment}"
}
