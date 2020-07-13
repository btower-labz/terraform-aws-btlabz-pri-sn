package main
import input

resource_whitelist = {
  "aws_subnet",
  "aws_vpc",
  "aws_route_table",
  "aws_route_table_association"
}

module_whitelist = {
  "../"
}

resource_deletions[r] {
  input.resource_changes[_].change.actions[_] == "delete"
  r := input.resource_changes[_].address
}

resource_changes[r] {
  input.resource_changes[_].change.actions[_] != "no-op"
  r := input.resource_changes[_].type
  not resource_whitelist[r]
}

module_changes[m] {
  input.resource_changes[_].change.actions[_] != "no-op"
  r := split(input.resource_changes[_].address, ".")
  r[0] == "module"
  m := input.configuration.root_module.module_calls[r[1]].source
  not module_whitelist[m]

}

deny[msg] {
  count(resource_changes) > 0
  msg := sprintf("Resource type is not preapproved: 5s", [resource_changes[_]])
}

deny[msg] {
  count(module_changes) > 0
  msg := sprintf("Module is not preapproved: %s", [module_changes[_]])
}

deny[msg] {
  count(resource_deletions) > 0
  msg := sprintf("Deletions are not allowed: %s", [resource_deletions[_]])
}

