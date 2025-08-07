---
page_title: "alertdefinition_monitor_membership Resource - itrs-uptrends"
subcategory: ""
description: |-
  Manages monitor memberships for alert definitions in the Uptrends monitoring platform.
  A list of relevant fields and their meaning can be found in the [API documentation for alert definitions](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/AlertDefinition) and the [Uptrends support knowledge base](https://www.uptrends.com/support/kb/api/alert-definition-api).
---

# alertdefinition_monitor_membership (Resource)

## Example usage

```terraform
# Add the monitor to the alert definition
resource "alertdefinition_monitor_membership" "alertdefinition_monitor_membership_example" {
  provider            = itrs-uptrends.uptrendsauthenticated
  alertdefinition_id  = alertdefinition.alertdefinition_example.id
  monitor_id          = monitor.certificate_monitor.id
  depends_on          = [alertdefinition.alertdefinition_example, monitor.certificate_monitor]
}

# Create an alert definition
resource "alertdefinition" "alertdefinition_example" {
  ...
}

# Create a monitor
resource "monitor" "certificate_monitor" {
  ...
}
```

## Use cases

Alert definition monitor memberships are commonly used for applying the same alerting rules to multiple monitors.


## Related resources

- [alertdefinition](alertdefinition.md) - Create and manage alert definitions
- [monitor](monitor.md) - Create and manage monitors
- [alertdefinition_operator_membership](alertdefinition_operator_membership.md) - Add operators to alert definition escalation levels
- [alertdefinition_operatorgroup_membership](alertdefinition_operatorgroup_membership.md) - Add operator groups to alert definition escalation levels

## Schema

### Required

- `alertdefinition_id` (String) The unique identifier of the alert definition.
- `monitor_id` (String) The unique identifier of the monitor to assign to the alert definition.

### Read-only

- `id` (String) The unique identifier of the membership (composite key in format `alertdefinition_id:monitor_id`).

## Import

Import is supported using the following syntax:

```shell
# Alert definition monitor membership can be imported by specifying the composite identifier alertdefinition_id:monitor_id.
terraform import alertdefinition_monitor_membership.example "046a727c-7a90-4776-9e41-ab050bdda5dc:046a727c-7a90-4776-9e41-ab050bdda5dc"
```

## Notes

- The `alertdefinition_id` and `monitor_id` fields are immutable and require resource replacement when changed.
- Each membership is a separate resource instance.
- The resource automatically handles the composite ID format (`alertdefinition_id:monitor_id`).
- Make sure both the alert definition and monitor exist before creating memberships.
- Use `depends_on` to ensure proper resource creation order.
- Removing a membership does not delete the alert definition or monitor, only the association between them.