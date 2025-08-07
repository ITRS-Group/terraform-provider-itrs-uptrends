---
page_title: "alertdefinition Resource - itrs-uptrends"
subcategory: ""
description: |-
  Manages alert definitions in the Uptrends monitoring platform.
  A list of relevant fields and their meanings can be found in the [API documentation for alert definitions](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/AlertDefinition), as well as in the Uptrends support knowledge base articles on [alert definition API](https://www.uptrends.com/support/kb/api/alert-definition-api), [alert escalation levels](https://www.uptrends.com/support/kb/alerting/alert-escalation-levels), and [alert reminders in escalations](https://www.uptrends.com/support/kb/alerting/alert-reminders-in-escalations).
---

# alertdefinition (Resource)

## Example usage

```terraform
resource "alertdefinition" "alertdefinition123" {
  name = "Alert Definition Resource"
  is_active = true
  escalation_levels = [
    {
      escalation_mode        = "AlertOnErrorCount"
      include_trace_route    = true
      is_active              = true
      message                = "message"
      number_of_reminders    = 5
      reminder_delay         = 10
      threshold_error_count  = 1
      threshold_minutes      = 5
      id = 1 # refers to "Escalation level 1"
    },
    {
      escalation_mode        = "AlertOnErrorCount"
      include_trace_route    = true
      is_active              = false
      message                = ""
      number_of_reminders    = 5
      reminder_delay         = 10
      threshold_error_count  = 1
      threshold_minutes      = 5
      id = 2 // refers to "Escalation level 2"
    },
    {
      escalation_mode        = "AlertOnErrorCount"
      include_trace_route    = true
      is_active              = false
      message                = ""
      number_of_reminders    = 5
      reminder_delay         = 10
      threshold_error_count  = 1
      threshold_minutes      = 5
      id = 3 // refers to "Escalation level 3"
    }
  ]
  provider = itrs-uptrends.uptrendsauthenticated
}
```

## Use cases

Alert definitions are used to specify when alerts are triggered and to configure multi-level escalations with customizable thresholds.

## Related resources

- [alertdefinition_monitor_membership](alertdefinition_monitor_membership.md) - Add monitors to alert definitions
- [alertdefinition_operator_membership](alertdefinition_operator_membership.md) - Add operators to alert definition escalation levels
- [alertdefinition_operatorgroup_membership](alertdefinition_operatorgroup_membership.md) - Add operator groups to alert definition escalation levels
- [monitor](monitor.md) - Create and manage monitors
- [operator](operator.md) - Create and manage operators
- [operatorgroup](operatorgroup.md) - Create and manage operator groups

## Schema

### Required

- `name` (String) The name of the alert definition.
- `is_active` (Boolean) Whether the alert definition is active.

### Optional

- `escalation_levels` (List) A list of escalation levels associated with the alert definition.  
  **Note:**  
  - The number of escalation levels is determined by your Uptrends account and cannot be changed through Terraform.  
  - You can update the settings (such as thresholds, messages, etc.) for each escalation level, but you cannot add or remove escalation levels from the list.

### Read-only

- `id` (String) The unique identifier of the alert definition.

### Escalation Level attributes

Each escalation level in the `escalation_levels` list must contain:

#### Required

- `id` (Integer) The unique identifier for the escalation level. This value must be an integer between 1 and the total number of escalation levels allowed by your Uptrends account. For example, an `id` of 1 refers to "Escalation level 1", an `id` of 2 refers to "Escalation level 2", and so on.
- `escalation_mode` (String) The escalation mode. Must be one of: `AlertOnErrorCount`, `AlertOnErrorDuration`.
- `threshold_error_count` (Integer) Threshold for error count. Used when escalation mode is `AlertOnErrorCount`.
- `threshold_minutes` (Integer) Threshold for minutes. Used when escalation mode is `AlertOnErrorDuration`.
- `is_active` (Boolean) Whether the escalation level is active.
- `message` (String) Message for the escalation level.
- `number_of_reminders` (Integer) Number of reminders to send.
- `reminder_delay` (Integer) Delay between reminders in minutes.
- `include_trace_route` (Boolean) Whether to include trace route information.

## Import

Import is supported using the following syntax:

```shell
# Alert definition can be imported by specifying the unique identifier.
terraform import alertdefinition.example "046a727c-7a90-4776-9e41-ab050bdda5dc"
```

## Notes

- The `id` field of the alert definition is automatically generated and managed by the Uptrends platform.
- The `escalation_levels` field must contain all your escalation levels. The exact number is determined by your Uptrends account settings and cannot be changed via Terraform.
- Each escalation level must have a unique `id` between 1 and the number of escalation levels.
- Escalation levels can be configured with different modes and thresholds.
- The resource automatically validates escalation level configuration.