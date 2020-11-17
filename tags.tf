locals {
  subnet_tags = merge(
    var.tags,
    {
      "Name" = var.name
    },
  )
}

resource "aws_ec2_tag" "subnet" {
  for_each    = local.subnet_tags
  resource_id = aws_subnet.main.id
  key         = each.key
  value       = each.value
}
