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
- `api_monitor_quota`, `browser_monitor_quota`, `transaction_monitor_quota`, `basic_monitor_quota` (Number) Quota settings.
- `is_all` (Boolean) Whether group covers all monitors.
- `is_quota_unlimited` (Boolean) Whether quotas are unlimited.
