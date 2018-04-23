# Terraform inputs and outputs.


## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|:----:|:-----:|:-----:|
| az | Subnet's availability zone. | string | - | yes |
| cidr | Subnet's CIDR range. | string | - | yes |
| name | Subnet name. Will be used as the 'Name' tag value. | string | `private-subnet` | no |
| rt_name | Private route table name. Will be used as the 'Name' tag value. Default is subnet name. | string | `` | no |
| tags | Additional tags. | map | `<map>` | no |
| vpc_id | VPC identifier to create the subnet in. | string | - | yes |

## Outputs

| Name | Description |
|------|-------------|
| rt_id | Private subnet route idenrifier. |
| subnet_id | Private subnet identifier. |

