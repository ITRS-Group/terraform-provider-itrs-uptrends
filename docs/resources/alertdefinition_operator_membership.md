---
page_title: "alertdefinition_operator_membership Resource - itrs-uptrends"
subcategory: ""
description: |-
  Manages operator memberships for alert definition escalation levels in the Uptrends monitoring platform. 
  A list of relevant fields and their meaning can be found in the [API documentation for alert definitions](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/AlertDefinition) and the [Uptrends support knowledge base](https://www.uptrends.com/support/kb/api/alert-definition-api).
---

# alertdefinition_operator_membership (Resource)

## Example usage

```terraform
# Add the operator to escalation level 3 of the alert definition
resource "alertdefinition_operator_membership" "membership124" {
  provider            = itrs-uptrends.uptrendsauthenticated
  depends_on          = [operator.alertdefinition_operator_membership, alertdefinition.alertdefinition_operator_membership]
  alertdefinition_id  = alertdefinition.alertdefinition_operator_membership.id
  operator_id         = operator.alertdefinition_operator_membership.id
  escalationlevel     = 3
}

# Create an alert definition
resource "alertdefinition" "alertdefinition_operator_membership" {
  ...
}

# Create an operator
resource "operator" "alertdefinition_operator_membership" {
  ...
}
```

## Use cases

Alert definition operator memberships are commonly used for applying the same alerting rules to multiple operators.

## Related resources

- [alertdefinition](alertdefinition.md) - Create and manage alert definitions
- [operator](operator.md) - Create and manage operators
- [alertdefinition_monitor_membership](alertdefinition_monitor_membership.md) - Add monitors to alert definitions
- [alertdefinition_operatorgroup_membership](alertdefinition_operatorgroup_membership.md) - Add operator groups to alert definition escalation levels

## Schema

### Required

- `alertdefinition_id` (String) The unique identifier of the alert definition.
- `operator_id` (String) The unique identifier of the operator to assign to the escalation level.
- `escalationlevel` (Integer) The escalation level (1, 2, or 3) to assign the operator to. The escalation levels are part of the alert definition and they cannot be created separately. Check the `alertdefinition` resource for more details.

### Read-only

- `id` (String) The unique identifier of the membership (composite key in format `alertdefinition_id:operator_id:escalationlevel`).

## Import

Import is supported using the following syntax:

```shell
# Alert definition operator membership can be imported by specifying the composite identifier alertdefinition_id:operator_id:escalationlevel.
terraform import alertdefinition_operator_membership.example "046a727c-7a90-4776-9e41-ab050bdda5dc:046a727c-7a90-4776-9e41-ab050bdda5dc:3"
```

## Notes

- The `alertdefinition_id`, `operator_id`, and `escalationlevel` fields are immutable and require resource replacement when changed.
- Each membership is a separate resource instance.
- The resource automatically handles the composite ID format (`alertdefinition_id:operator_id:escalationlevel`).
- Make sure both the alert definition and operator exist before creating memberships.
- Use `depends_on` to ensure proper resource creation order.
- Removing a membership does not delete the alert definition or operator, only the association between them.
