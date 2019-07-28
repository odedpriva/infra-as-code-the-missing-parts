
# development
resource "aws_s3_bucket" "development" {
  bucket = "development.example.com"
}

resource "aws_sqs_queue" "development" {
  name = "development"
}

# staging
resource "aws_s3_bucket" "staging" {
  bucket = "staging.example.com"
}

resource "aws_sqs_queue" "staging" {
  name = "staging"
}

# production
resource "aws_s3_bucket" "production" {
  bucket = "production.example.com"
}

resource "aws_sqs_queue" "productoin" {
  name = "productoin"
}
