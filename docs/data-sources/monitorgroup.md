---
page_title: "itrs-uptrends_monitorgroup Data Source - itrs-uptrends"
subcategory: ""
description: |-
  Retrieve a monitor group by GUID or description to reference quotas or membership.
---

# itrs-uptrends_monitorgroup (Data Source)

Use this data source to fetch monitor group metadata so you can configure dependent resources without duplicating values.

## Example Usage

```terraform
data "itrs-uptrends_monitorgroup" "by_description" {
  description = "Production monitors"
}

data "itrs-uptrends_monitorgroup" "by_id" {
  id = "17efa6f6-1f20-4d0f-bb23-f41d9a372b90"
}
```

## Schema

### Optional
- `id` (String) Monitor group GUID. Provide this or `description`.
- `description` (String) Monitor group description. Provide this or `id`. If the description is not unique it is going to give an error.

### Read-Only
- `api_monitor_quota` (Number) API monitor quota.
- `used_api_monitor_quota` (Number) Consumed API monitor quota.
- `browser_monitor_quota` (Number) Browser monitor quota.
- `used_browser_monitor_quota` (Number) Consumed browser monitor quota.
- `transaction_monitor_quota` (Number) Transaction monitor quota.
- `used_transaction_monitor_quota` (Number) Consumed transaction monitor quota.
- `basic_monitor_quota` (Number) Basic monitor quota.
- `used_basic_monitor_quota` (Number) Consumed basic monitor quota.
- `unified_credits_quota` (Number) Unified credits quota.
- `used_unified_credits_quota` (Number) Consumed unified credits quota.
- `classic_quota` (Number) Classic quota.
- `used_classic_quota` (Number) Consumed classic quota.
- `is_all` (Boolean) Whether group covers all monitors.
- `is_quota_unlimited` (Boolean) Whether quotas are unlimited.

## Quota Systems

Uptrends has three different quota models:

- Classic quota: one bucket represented by `classic_quota` and `used_classic_quota`.
- Bucketed quota: four buckets represented by:
  - `basic_monitor_quota` / `used_basic_monitor_quota`
  - `browser_monitor_quota` / `used_browser_monitor_quota`
  - `transaction_monitor_quota` / `used_transaction_monitor_quota`
  - `api_monitor_quota` / `used_api_monitor_quota`
- Unified credits quota: one bucket represented by `unified_credits_quota` and `used_unified_credits_quota`.

Each account uses exactly one quota system. Depending on the account configuration, only the fields for that active system are relevant.
