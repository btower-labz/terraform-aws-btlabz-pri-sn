# Terraform inputs and outputs.

## Requirements

| Name | Version |
|------|---------|
| terraform | >= 0.12 |

## Providers

| Name | Version |
|------|---------|
| aws | n/a |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| az | Subnet's availability zone. | `string` | n/a | yes |
| cidr | Subnet's CIDR range. | `string` | n/a | yes |
| name | Subnet name. Will be used as the 'Name' tag value. | `string` | `"private-subnet"` | no |
| rt\_name | Private route table name. Will be used as the 'Name' tag value. Default is subnet name. | `string` | `""` | no |
| tags | Additional tags. | `map(string)` | `{}` | no |
| vpc\_id | VPC identifier to create the subnet in. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| rt\_id | Private subnet route idenrifier. |
| subnet\_id | Private subnet identifier. |

