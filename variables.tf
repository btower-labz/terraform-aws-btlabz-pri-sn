variable "vpc_id" {
  description = "VPC identifier to create the subnet in."
  type        = "string"
}

variable "cidr" {
  description = "Subnet's CIDR range."
  type        = "string"
}

variable "az" {
  description = "Subnet's availability zone."
  type        = "string"
}

variable "tags" {
  description = "Additional tags."
  type        = "map"
  default     = {}
}

variable "name" {
  description = "Subnet name. Will be used as the 'Name' tag value."
  type        = "string"
  default     = "private-subnet"
}

variable "rt_name" {
  description = "Private route table name. Will be used as the 'Name' tag value. Default is subnet name."
  type        = "string"
  default     = ""
}
