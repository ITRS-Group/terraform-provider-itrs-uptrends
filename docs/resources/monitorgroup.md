---
page_title: "monitorgroup Resource - itrs-uptrends"
subcategory: ""
description: |-

---

# itrs-uptrends_monitorgroup (Resource)
  Manages monitor groups in the Uptrends monitoring platform.  
  A list of relevant fields and their meaning can be found in the [API documentation for monitor groups](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/MonitorGroup) and the [Uptrends support knowledge base](https://www.uptrends.com/support/kb/api/monitorgroup-api).

## Example usage

```terraform
# Monitor group with limited quota (available only for Enterprise accounts)
resource "itrs-uptrends_monitorgroup" "limited_quota_group" {
  description = "Monitor group with quota"
  is_quota_unlimited = false
  unified_credits_quota = 70
  provider = itrs-uptrends.uptrendsauthenticated
}

# Monitor group with unlimited quota
resource "itrs-uptrends_monitorgroup" "unlimited_quota_group" {
  description = "Monitor group with unlimited quota"
  is_quota_unlimited = true
  provider = itrs-uptrends.uptrendsauthenticated
}
```

## Use cases

Monitor groups are primarily used to organize monitors.

## Related resources

- [itrs-uptrends_monitor](monitor.md) - Create and manage individual monitors
- [itrs-uptrends_monitorgroup_membership](monitorgroup_membership.md) - Add monitors to monitor groups

## Schema

### Required

- `description` (String) The description of the monitor group.

### Optional

- `is_quota_unlimited` (Boolean) Whether the monitor group has unlimited quotas. Defaults to true.
- `basic_monitor_quota` (Integer) The quota for basic monitors (HTTP, HTTPS, DNS, etc.).
- `browser_monitor_quota` (Integer) The quota for browser-based monitors (Full Page Check, Transaction).
- `transaction_monitor_quota` (Integer) The quota for transaction monitors.
- `api_monitor_quota` (Integer) The quota for API monitors.
- `unified_credits_quota` (Integer) The unified credits quota (single-bucket quota model).
- `classic_quota` (Integer) The classic quota (single-bucket quota model).

### Read-only

- `id` (String) The unique identifier of the monitor group.

## Import

Import is supported using the following syntax:

```shell
# Monitor group can be imported by specifying the unique identifier.
terraform import itrs-uptrends_monitorgroup.example "046a727c-7a90-4776-9e41-ab050bdda5dc"
```

## Quota Systems

Uptrends has three different quota models:

- Classic quota: one bucket represented by `classic_quota`.
- Bucketed quota: four buckets represented by:
  - `basic_monitor_quota`
  - `browser_monitor_quota`
  - `transaction_monitor_quota`
  - `api_monitor_quota`
- Unified credits quota: one bucket represented by `unified_credits_quota`.

Each account uses exactly one quota system. Depending on the account configuration, only the fields for that active system are relevant.

## Notes

- The `id` field is automatically generated and managed by the Uptrends platform.
- When `is_quota_unlimited` is true, all quota fields are ignored.
- Quota fields are only applicable when `is_quota_unlimited` is false.
- Use quota fields from the account's active quota system only (`classic_quota`, the four bucketed quotas, or `unified_credits_quota`).