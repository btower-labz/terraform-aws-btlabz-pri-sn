output "vpc_id" {
  value = aws_vpc.main.id
}

output "az" {
  value = data.aws_availability_zone.az.name
}

// Validate required module outputs with infratest

output "subnet_id" {
  value = module.subnet.subnet_id
}

output "rt_id" {
  value = module.subnet.rt_id
}
