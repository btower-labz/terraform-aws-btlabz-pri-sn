resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"
  tags = {
    Name = "${var.name_prefix}-vpc"
  }
}

module "subnet" {
  source = "../../"
  cidr   = "10.0.1.0/24"
  az     = data.aws_availability_zone.az.name
  vpc_id = aws_vpc.main.id
  tags = {
    Name = "${var.name_prefix}-subnet"
  }
}
