output "cidr_block" {
  value     = "${aws_vpc.selected.cidr_block}"
  sensitive = true
}
