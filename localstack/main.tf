provider "aws" {
  access_key                  = "mock_access_key"
  region                      = "eu-west-1"
  secret_key                  = "mock_secret_key"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  endpoints {
    sqs = "http://localhost:4576"
  }
}

resource "aws_sqs_queue" "terraform_queue" {
  name = "terraform-example-queue"
}
