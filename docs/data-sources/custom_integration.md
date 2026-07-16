---
page_title: "itrs-uptrends_custom_integration Data Source - itrs-uptrends"
subcategory: ""
description: |-
  Fetch a custom integration by GUID or name when you need to reference it in other resources.
---

# itrs-uptrends_custom_integration (Data Source)

Look up a custom integration to reuse its GUID and settings without hard-coding values.

## Example Usage

```terraform
data "itrs-uptrends_custom_integration" "by_name" {
  provider = itrs-uptrends.uptrendsauthenticated
  name     = "My Custom Integration"
}

data "itrs-uptrends_custom_integration" "by_id" {
  provider = itrs-uptrends.uptrendsauthenticated
  id       = "f5a3f2a3-1234-5678-9abc-def012345678"
}
```

## Schema

### Optional

- `id` (String) The GUID of the custom integration. Provide this or `name`.
- `name` (String) The name of the custom integration. Provide this or `id`. If the name is not unique an error is returned.

### Read-Only

- `is_active` (Boolean) Whether the custom integration is active.
- `notes` (String) Notes for the custom integration.
- `integration_variables` (Attributes List) Variables that can be referenced by the integration steps (`name`, `value`, `is_value_set_in_escalation_level`).
- `http_step_definitions` (Attributes List) The HTTP step definitions. Each contains `usages`, `http_step_definition_usage_guid`, `steps`, and `user_defined_functions`. Each step exposes its request configuration and an `authentication` block with `id`, `authentication_type` and `user_name`.

## Notes

- The step `authentication` password is write-only on the resource and is **never** exposed by this data source.
