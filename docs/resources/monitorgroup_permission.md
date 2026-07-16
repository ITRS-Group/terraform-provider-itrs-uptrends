---
page_title: "monitorgroup_permission Resource - itrs-uptrends"
subcategory: ""
description: |-
  Manages authorization permissions for monitor groups in the Uptrends monitoring platform.
---

# itrs-uptrends_monitorgroup_permission (Resource)

Manages authorization permissions for monitor groups in the Uptrends monitoring platform.
A list of relevant fields and their meaning can be found in the [API documentation for monitor groups](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/MonitorGroup) and the [Uptrends support knowledge base](https://www.uptrends.com/support/kb/api/monitor-api).

## Example usage

### Grant an operator view access to monitor data in a group

```terraform
resource "itrs-uptrends_monitorgroup_permission" "operator_view_data" {
  provider        = itrs-uptrends.uptrendsauthenticated
  monitorgroup_id = itrs-uptrends_monitorgroup.example.id
  permission      = "ViewMonitorDataInGroup"
  operator_id     = itrs-uptrends_operator.example.id
}
```

### Grant an operator group edit access to monitors in a group

```terraform
resource "itrs-uptrends_monitorgroup_permission" "group_edit" {
  provider         = itrs-uptrends.uptrendsauthenticated
  monitorgroup_id  = itrs-uptrends_monitorgroup.example.id
  permission       = "EditMonitorsInGroup"
  operatorgroup_id = itrs-uptrends_operatorgroup.example.id
}
```

## Use cases

Monitor group permissions control which operators or operator groups can view monitor data, view monitor settings, edit monitors, or create and delete monitors in a specific monitor group.

## Related resources

- [itrs-uptrends_monitorgroup](monitorgroup.md) - Create and manage monitor groups
- [itrs-uptrends_operator](operator.md) - Manage operators
- [itrs-uptrends_operatorgroup](operatorgroup.md) - Manage operator groups

## Schema

### Required

- `monitorgroup_id` (String) The GUID of the monitor group.
- `permission` (String) The monitor group authorization type. Valid values: `ViewMonitorDataInGroup`, `ViewMonitorsInGroup`, `EditMonitorsInGroup`, `CreateAndDeleteMonitorsInGroup`.

### Optional

- `operator_id` (String) The GUID of the operator. Provide this or `operatorgroup_id`, not both.
- `operatorgroup_id` (String) The GUID of the operator group. Provide this or `operator_id`, not both.

### Read-Only

- `id` (String) The unique identifier of the authorization (composite key in format `monitorgroup_id:authorization_id`).

## Import

Import is supported using the following syntax:

```shell
# Monitor group permission can be imported by specifying the composite identifier monitorgroup_id:authorization_id.
terraform import itrs-uptrends_monitorgroup_permission.example "monitor-group-guid:authorization-guid"
```

## Notes

- All attributes are immutable, so changing any value requires resource replacement.
- Exactly one of `operator_id` or `operatorgroup_id` must be provided.
- Removing a permission does not delete the monitor group, operator, or operator group, only the association.
