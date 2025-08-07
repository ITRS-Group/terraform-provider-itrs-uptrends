---
page_title: "operatorgroup Resource - itrs-uptrends"
subcategory: ""
description: |-
  Manages operator groups in the Uptrends monitoring platform.  
  A list of relevant fields and their meaning can be found in the [API documentation for operator groups](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/OperatorGroup) and the [Uptrends support knowledge base](https://www.uptrends.com/support/kb/api/operator-group-api).
---

# operatorgroup (Resource)

## Example usage

```terraform
# Manage an operator group.
resource "operatorgroup" "operatorgroup123" {
  description = "Operator group description"
  provider    = itrs-uptrends.uptrendsauthenticated
}
```

## Use cases

Operator groups are typically used to organize operators who share similar roles or responsibilities within the monitoring platform.

## Related resources

- [operatorgroup_membership](operatorgroup_membership.md) - Add operators to operator groups
- [operatorgroup_permission](operatorgroup_permission.md) - Manage permissions for operator groups
- [operator](operator.md) - Manage individual operators

## Schema

### Required

- `description` (String) The description of the operator group.

### Read-only

- `id` (String) The unique identifier of the operator group.

## Import

Import is supported using the following syntax:

```shell
# Operator group can be imported by specifying the unique identifier.
terraform import operatorgroup.operatorgroup123 "046a727c-7a90-4776-9e41-ab050bdda5dc"
```

## Notes

- The `id` field is automatically generated and managed by the Uptrends platform.