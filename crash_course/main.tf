provider "aws" {
  region = "eu-west-1"
}

variable "queue_name" {
  default = "developemt"
}

resource "aws_sqs_queue" "development" {
  name = "${var.queue_name}"
}

provider "datadog" {
  api_key = "my_api_key"
  app_key = "my_app_key"
}

resource "datadog_monitor" "sqs" {
  name  = "monitor for sqs"
  type  = "metric alert"
  query = "avg(last_1h):avg:aws.sqs.number_of_messages_sent > 1000"
}
