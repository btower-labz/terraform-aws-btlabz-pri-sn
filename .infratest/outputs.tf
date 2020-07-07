output "vpc_id" {
  value = aws_vpc.main.id
}

output "az" {
  value = data.aws_availability_zone.az.name
}
