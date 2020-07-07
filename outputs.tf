output "subnet_id" {
  description = "Private subnet identifier."
  value       = aws_subnet.main.id
}

output "rt_id" {
  description = "Private subnet route idenrifier."
  value       = aws_route_table.main.id
}
