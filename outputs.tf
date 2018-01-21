output "subnet-id" {
  description = "Private subnet identifier."
  value       = "${aws_subnet.main.id}"
}

output "rt-id" {
  description = "Private subnet route idenrifier."
  value       = "${aws_route_table.main.id}"
}
