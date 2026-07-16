---
page_title: "monitor_permission Resource - itrs-uptrends"
subcategory: ""
description: |-
  Manages authorization permissions for monitors in the Uptrends monitoring platform.
---

# itrs-uptrends_monitor_permission (Resource)

Manages authorization permissions for monitors in the Uptrends monitoring platform.
A list of relevant fields and their meaning can be found in the [API documentation for monitors](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/Monitor) and the [Uptrends support knowledge base](https://www.uptrends.com/support/kb/api/monitor-api).

## Example usage

### Grant an operator view access to monitor data

```terraform
resource "itrs-uptrends_monitor_permission" "operator_view_data" {
  provider    = itrs-uptrends.uptrendsauthenticated
  monitor_id  = itrs-uptrends_monitor.example.id
  permission  = "ViewMonitorData"
  operator_id = itrs-uptrends_operator.example.id
}
```

### Grant an operator group edit access to a monitor

```terraform
resource "itrs-uptrends_monitor_permission" "group_edit" {
  provider         = itrs-uptrends.uptrendsauthenticated
  monitor_id       = itrs-uptrends_monitor.example.id
  permission       = "EditMonitor"
  operatorgroup_id = itrs-uptrends_operatorgroup.example.id
}
```

## Use cases

Monitor permissions control which operators or operator groups can view monitor data, view monitor settings, or edit monitor settings for a specific monitor.

## Related resources

- [itrs-uptrends_monitor](monitor.md) - Create and manage monitors
- [itrs-uptrends_operator](operator.md) - Manage operators
- [itrs-uptrends_operatorgroup](operatorgroup.md) - Manage operator groups

## Schema

### Required

- `monitor_id` (String) The GUID of the monitor.
- `permission` (String) The monitor authorization type. Valid values: `ViewMonitorData`, `ViewMonitor`, `EditMonitor`.

### Optional

- `operator_id` (String) The GUID of the operator. Provide this or `operatorgroup_id`, not both.
- `operatorgroup_id` (String) The GUID of the operator group. Provide this or `operator_id`, not both.

### Read-Only

- `id` (String) The unique identifier of the authorization (composite key in format `monitor_id:authorization_id`).

## Import

Import is supported using the following syntax:

```shell
# Monitor permission can be imported by specifying the composite identifier monitor_id:authorization_id.
terraform import itrs-uptrends_monitor_permission.example "monitor-guid:authorization-guid"
```

## Notes

- All attributes are immutable, so changing any value requires resource replacement.
- Exactly one of `operator_id` or `operatorgroup_id` must be provided.
- Removing a permission does not delete the monitor, operator, or operator group, only the association.
