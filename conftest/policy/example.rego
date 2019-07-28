package main

cidrsRFC1918 = [
  "10.0.0.0/8",
  "172.16.0.0/12",
  "192.168.0.0/16"
]

deny[msg] {
  not checkRFC1918(input.variables.cidr_block.value)
  msg = "cidr_block variables does not meet RFC 1918 "
}


checkRFC1918(ip){
  net.cidr_contains(cidrsRFC1918[_], ip)
}