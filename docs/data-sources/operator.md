---
page_title: "itrs-uptrends_operator Data Source - itrs-uptrends"
subcategory: ""
description: |-
  Look up an operator by GUID or full name to consume its attributes in downstream resources.
---

# itrs-uptrends_operator (Data Source)

Use this data source to refresh operator metadata (email, duty status, role, etc.) before feeding it into other resources.

## Example Usage

```terraform
data "itrs-uptrends_operator" "by_name" {
  full_name = "Jane Support"
}

data "itrs-uptrends_operator" "by_id" {
  id = "9a0bbe79-4944-42f4-ace8-dfac33af9a64"
}
```

## Schema

### Optional
- `id` (String) Operator GUID. Provide this or `full_name`.
- `full_name` (String) Operator full name. Provide this or `id`. If the name is not unique it is going to give an error.

### Read-Only
- `backup_email` (String) Backup email.
- `default_dashboard` (String) Default dashboard configuration.
- `email` (String) Primary email.
- `is_account_administrator` (Boolean) Whether operator is account admin.
- `is_on_duty` (Boolean) Whether operator is currently on duty.
- `mobile_phone` (String) Mobile phone number.
- `operator_role` (String) Assigned role (defaults to “Unspecified”).
- `sms_provider` (String) SMS provider.
