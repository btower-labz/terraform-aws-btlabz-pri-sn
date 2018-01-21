output "subnet-id" {
  value = "${aws_subnet.main.id}"
}

output "rt-id" {
  value = "${aws_route_table.main.id}"
}
