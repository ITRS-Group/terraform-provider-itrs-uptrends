---
page_title: "alertdefinition_monitorgroup_membership Resource - itrs-uptrends"
subcategory: ""
description: |-

---

# itrs-uptrends_alertdefinition_monitorgroup_membership (Resource)
  Manages monitor group memberships for alert definitions in the Uptrends monitoring platform.
  A list of relevant fields and their meaning can be found in the [API documentation for alert definitions](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/AlertDefinition) and the [Uptrends support knowledge base](https://www.uptrends.com/support/kb/api/alert-definition-api).

## Example usage

```terraform
# Add the monitor group to the alert definition
resource "itrs-uptrends_alertdefinition_monitorgroup_membership" "example" {
  alertdefinition_id  = itrs-uptrends_alertdefinition.example.id
  monitorgroup_id     = itrs-uptrends_monitorgroup.example.id
}

# Create an alert definition
resource "itrs-uptrends_alertdefinition" "example" {
  name      = "Alert Definition Resource Test"
  is_active = true
}

# Create a monitor group
resource "itrs-uptrends_monitorgroup" "example" {
  description = "Monitor Group Resource Test"
}
```

## Use cases

Alert definition monitor group memberships are commonly used for applying alerting rules to all monitors within a specific group at once.

## Related resources

- [itrs-uptrends_alertdefinition](alertdefinition.md) - Create and manage alert definitions
- [itrs-uptrends_monitorgroup](monitorgroup.md) - Create and manage monitor groups
- [itrs-uptrends_alertdefinition_monitor_membership](alertdefinition_monitor_membership.md) - Add individual monitors to alert definitions

## Schema

### Required

- `alertdefinition_id` (String) The unique identifier of the alert definition.
- `monitorgroup_id` (String) The unique identifier of the monitor group to assign to the alert definition.

### Read-only

- `id` (String) The unique identifier of the membership (composite key in format `alertdefinition_id:monitorgroup_id`).

## Import

Import is supported using the following syntax:

```shell
# Alert definition monitor group membership can be imported by specifying the composite identifier alertdefinition_id:monitorgroup_id.
terraform import itrs-uptrends_alertdefinition_monitorgroup_membership.example "046a727c-7a90-4776-9e41-ab050bdda5dc:046a727c-7a90-4776-9e41-ab050bdda5dc"
```

## Notes

- The `alertdefinition_id` and `monitorgroup_id` fields are immutable and require resource replacement when changed.
- Each membership is a separate resource instance.
- The resource automatically handles the composite ID format (`alertdefinition_id:monitorgroup_id`).
- Make sure both the alert definition and monitor group exist before creating memberships.
