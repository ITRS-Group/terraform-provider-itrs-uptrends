---
page_title: "itrs-uptrends_alertdefinition Data Source - itrs-uptrends"
subcategory: ""
description: |-
  Use this data source to retrieve information about an existing alert definition to use it in another resources.
---

# itrs-uptrends_alertdefinition (Data Source)
Use this data source to look up an existing alert definition by id or name so you can reuse its attributes elsewhere.

## Example Usage

```terraform
data "itrs-uptrends_alertdefinition" "by_name" {
  name = "Alert definition"
}

data "itrs-uptrends_alertdefinition" "by_id" {
  id = "9a0bbe79-4944-42f4-ace8-dfac33af9a64"
}
```

## Schema

### Optional
- `id` (String) Alert definition GUID. Provide this or `name`. If the name is not unique it is going to give an error.
- `name` (String) Alert definition name. Provide this or `id`.

### Read-Only

- `is_active` (Boolean) Whether the alert definition is active.
- `escalation_levels` (List) A list of escalation levels associated with the alert definition.  


### Escalation Level attributes

Each escalation level in the `escalation_levels` list contains:

#### Read-only

- `id` (Integer) The unique identifier for the escalation level.
- `escalation_mode` (String) The escalation mode.
- `threshold_error_count` (Integer) Threshold for error count. Used when escalation mode is `AlertOnErrorCount`.
- `threshold_minutes` (Integer) Threshold for minutes. Used when escalation mode is `AlertOnErrorDuration`.
- `is_active` (Boolean) Whether the escalation level is active.
- `message` (String) Message for the escalation level.
- `number_of_reminders` (Integer) Number of reminders to send.
- `reminder_delay` (Integer) Delay between reminders in minutes.
- `include_trace_route` (Boolean) Whether to include trace route information.