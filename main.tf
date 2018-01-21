resource "aws_subnet" "main" {
  vpc_id                          = "${var.vpc-id}"
  cidr_block                      = "${var.cidr}"
  availability_zone               = "${var.az}"
  map_public_ip_on_launch         = false
  assign_ipv6_address_on_creation = false

  tags = "${merge(
    var.tags,
    map(
      "Name", "${var.name}"
    )
  )}"
}

resource "aws_route_table" "main" {
  vpc_id = "${var.vpc-id}"

  tags = "${merge(
    var.tags,
    map(
      "Name", "${var.name}"
    )
  )}"
}

resource "aws_route_table_association" "main" {
  subnet_id      = "${aws_subnet.main.id}"
  route_table_id = "${aws_route_table.main.id}"
}
