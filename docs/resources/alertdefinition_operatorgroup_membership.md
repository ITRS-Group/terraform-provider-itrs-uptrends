---
page_title: "alertdefinition_operatorgroup_membership Resource - itrs-uptrends"
subcategory: ""
description: |-
  Manages operator group memberships for alert definition escalation levels in the Uptrends monitoring platform. 
  A list of relevant fields and their meaning can be found in the [API documentation for alert definitions](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/AlertDefinition) and the [Uptrends support knowledge base](https://www.uptrends.com/support/kb/api/alert-definition-api).
---

# alertdefinition_operatorgroup_membership (Resource)

## Example usage

```terraform
# Add the operator group to escalation level 3 of the alert definition
resource "alertdefinition_operatorgroup_membership" "example_membership123" {
  provider            = itrs-uptrends.uptrendsauthenticated
  depends_on          = [operatorgroup.alertdefinition_operatorgroup_membership, alertdefinition.alertdefinition_operatorgroup_membership]
  alertdefinition_id  = alertdefinition.alertdefinition_operatorgroup_membership.id
  operatorgroup_id    = operatorgroup.alertdefinition_operatorgroup_membership.id
  escalationlevel     = 3
}

# Create an alert definition
resource "alertdefinition" "alertdefinition123" {
  ...
}

# Create an operator group
resource "operatorgroup" "operatorgroup123" {
  ...
}
```

## Use cases

Alert definition operator group memberships are commonly used for applying the same alerting rules to multiple operator groups.

## Related resources

- [alertdefinition](alertdefinition.md) - Create and manage alert definitions
- [operatorgroup](operatorgroup.md) - Create and manage operator groups
- [alertdefinition_monitor_membership](alertdefinition_monitor_membership.md) - Add monitors to alert definitions
- [alertdefinition_operator_membership](alertdefinition_operator_membership.md) - Add individual operators to alert definition escalation levels

## Schema

### Required

- `alertdefinition_id` (String) The unique identifier of the alert definition.
- `operatorgroup_id` (String) The unique identifier of the operator group to assign to the escalation level.
- `escalationlevel` (Integer) The escalation level (1, 2, or 3) to assign the operator group to. The escalation levels are part of the alert definition and they cannot be created separately. Check the `alertdefinition` resource for more details.

### Read-only

- `id` (String) The unique identifier of the membership (composite key in format `alertdefinition_id:operatorgroup_id:escalationlevel`).

## Import

Import is supported using the following syntax:

```shell
# Alert definition operator group membership can be imported by specifying the composite identifier alertdefinition_id:operatorgroup_id:escalationlevel.
terraform import alertdefinition_operatorgroup_membership.example "046a727c-7a90-4776-9e41-ab050bdda5dc:046a727c-7a90-4776-9e41-ab050bdda5dc:3"
```

## Notes

- The `alertdefinition_id`, `operatorgroup_id`, and `escalationlevel` fields are immutable and require resource replacement when changed.
- Each membership is a separate resource instance.
- The resource automatically handles the composite ID format (`alertdefinition_id:operatorgroup_id:escalationlevel`).
- Make sure both the alert definition and operator group exist before creating memberships.
- Use `depends_on` to ensure proper resource creation order.
- Removing a membership does not delete the alert definition or operator group, only the association between them.