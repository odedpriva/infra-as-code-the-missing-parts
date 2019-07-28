module "rds" {
  source        = "rds-module"
  name          = "${var.environment}"
  create_record = false
}
