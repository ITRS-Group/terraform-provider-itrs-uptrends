---
page_title: "monitorgroup_membership Resource - itrs-uptrends"
subcategory: ""
description: |-

---

# itrs-uptrends_monitorgroup_membership (Resource)
  Manages monitor group memberships in the Uptrends monitoring platform.  
  A list of relevant fields and their meaning can be found in the [API documentation for monitor groups](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/MonitorGroup) and the [Uptrends support knowledge base](https://www.uptrends.com/support/kb/api/monitorgroup-api).

## Example usage

```terraform
# Add the monitor to the monitor group
resource "itrs-uptrends_monitorgroup_membership" "monitorgroup_membership_example" {
  provider        = itrs-uptrends.uptrendsauthenticated
  monitorgroup_id = itrs-uptrends_monitorgroup.monitorgroup_example.id
  monitor_id      = itrs-uptrends_monitor.certificate_monitor.id
  depends_on      = [itrs-uptrends_monitorgroup.monitorgroup_example, itrs-uptrends_monitor.certificate_monitor]
}

# Create a monitor group
resource "itrs-uptrends_monitorgroup" "monitorgroup_example" {
  ...
}

# Create a monitor
resource "itrs-uptrends_monitor" "certificate_monitor" {
  ...
}
```

## Use cases

Monitor group memberships are primarily used to organize, manage permissions, and group monitors.

## Related resources

- [itrs-uptrends_monitorgroup](monitorgroup.md) - Create and manage monitor groups
- [itrs-uptrends_monitor](monitor.md) - Create and manage individual monitors

## Schema

### Required

- `monitor_id` (String) The unique identifier of the monitor to add to the group.
- `monitorgroup_id` (String) The unique identifier of the monitor group to add the monitor to.

### Read-only

- `id` (String) The unique identifier of the membership (composite key in format `monitor_id:monitorgroup_id`).

## Import

Import is supported using the following syntax:

```shell
# Monitor group membership can be imported by specifying the composite identifier monitor_id:monitorgroup_id.
terraform import itrs-uptrends_monitorgroup_membership.example "046a727c-7a90-4776-9e41-ab050bdda5dc:046a727c-7a90-4776-9e41-ab050bdda5dc"
```

## Notes

- The `monitor_id` and `monitorgroup_id` fields are immutable and require resource replacement when changed.
- Each membership is a separate resource instance.
- The resource automatically handles the composite ID format (`monitor_id:monitorgroup_id`).
- Make sure both the monitor and monitor group exist before creating memberships.
- Use `depends_on` to ensure proper resource creation order.
- Removing a membership does not delete the monitor or monitor group, only the association between them.