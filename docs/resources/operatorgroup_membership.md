---
page_title: "operatorgroup_membership Resource - itrs-uptrends"
subcategory: ""
description: |-
  Manages operator group memberships in the Uptrends monitoring platform.  
  A list of relevant fields and their meaning can be found in the [API documentation for operator groups](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/OperatorGroup) and the [Uptrends support knowledge base](https://www.uptrends.com/support/kb/api/operator-group-api).
---

# operatorgroup_membership (Resource)

## Example usage

```terraform
# Add the operator to the operator group
resource "operatorgroup_membership" "example_membership" {
  provider         = itrs-uptrends.uptrendsauthenticated
  operator_id      = operator.example_operator.id
  operatorgroup_id = operatorgroup.example_group.id
  depends_on       = [operator.example_operator, operatorgroup.example_group]
}

# Create an operator group
resource "operatorgroup" "example_group" {
  ...
}

# Create an operator
resource "operator" "example_operator" {
  ...
}
```

## Use cases

Operator group memberships enable efficient organization, permission management, and notification control by associating operators with specific groups in the Uptrends monitoring platform.

## Related resources

- [operatorgroup](operatorgroup.md) - Create and manage operator groups
- [operatorgroup_permission](operatorgroup_permission.md) - Manage permissions for operator groups
- [operator](operator.md) - Create and manage individual operators
- [operator_permission](operator_permission.md) - Manage permissions for individual operators.

## Schema

### Required

- `operator_id` (String) The unique identifier of the operator to add to the group.
- `operatorgroup_id` (String) The unique identifier of the operator group to add the operator to.

### Read-only

- `id` (String) The unique identifier of the membership (composite key in format `operator_id:operatorgroup_id`).

## Import

Import is supported using the following syntax:

```shell
# Operator group membership can be imported by specifying the composite identifier operator_id:operatorgroup_id.
terraform import operatorgroup_membership.example "046a727c-7a90-4776-9e41-ab050bdda5dc:046a727c-7a90-4776-9e41-ab050bdda5dc"
```

## Notes

- The `operator_id` and `operatorgroup_id` fields are immutable and require resource replacement when changed.
- Each membership is a separate resource instance.
- The resource automatically handles the composite ID format (`operator_id:operatorgroup_id`).
- Make sure both the operator and operator group exist before creating memberships.
- Use `depends_on` to ensure proper resource creation order.
- Removing a membership does not delete the operator or operator group, only the association between them.
