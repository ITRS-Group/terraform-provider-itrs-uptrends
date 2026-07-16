---
page_title: "escalation_level_integration Resource - itrs-uptrends"
subcategory: ""
description: |-
  Manages integrations attached to alert definition escalation levels.
---

# itrs-uptrends_escalation_level_integration (Resource)

Attaches an integration (e.g. Slack, PagerDuty, email) to a specific escalation level of an alert definition. Each escalation level can have multiple integrations.

A list of relevant fields and their meaning can be found in the [API documentation for alert definitions](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/AlertDefinition) and the [Uptrends support knowledge base](https://www.uptrends.com/support/kb/api/alert-definition-api).

## Example Usage

### Email integration

```terraform
resource "itrs-uptrends_escalation_level_integration" "email" {
  provider             = itrs-uptrends.uptrendsauthenticated
  alertdefinition_id   = itrs-uptrends_alertdefinition.example.id
  escalation_level_id  = 1
  integration_guid     = "d48e6070-98cb-4f45-9d6b-bb4426f2cba3"
  is_active_wo            = true
  send_ok_alerts_wo       = true
  send_reminder_alerts_wo = true
}
```

### Phone integration (send_ok_alerts_wo must be false)

```terraform
resource "itrs-uptrends_escalation_level_integration" "phone" {
  provider             = itrs-uptrends.uptrendsauthenticated
  alertdefinition_id   = itrs-uptrends_alertdefinition.example.id
  escalation_level_id  = 2
  integration_guid     = "7bcb0fae-85c8-4bbe-b117-fd346f21c7be"
  is_active_wo            = true
  send_ok_alerts_wo       = false
  send_reminder_alerts_wo = true
}
```

### Slack integration with variable values

```terraform
resource "itrs-uptrends_escalation_level_integration" "slack" {
  provider             = itrs-uptrends.uptrendsauthenticated
  alertdefinition_id   = itrs-uptrends_alertdefinition.example.id
  escalation_level_id  = 1
  integration_guid     = "your-slack-integration-guid"
  is_active_wo            = true
  send_ok_alerts_wo       = true
  send_reminder_alerts_wo = false

  variable_values = {
    channel = "#alerts"
  }
}
```

## Use Cases

- Attach a Slack or Teams integration to escalation level 1 so the team is notified immediately on error.
- Attach a PagerDuty integration to escalation level 2 for on-call escalation.
- Manage multiple integrations per escalation level as separate Terraform resources.

## Related Resources

- [itrs-uptrends_alertdefinition](alertdefinition.md) - Create and manage alert definitions.

## Schema

### Required

- `alertdefinition_id` (String) The GUID of the alert definition. Changing this forces a new resource.
- `escalation_level_id` (Number) The escalation level ID (1-4). Changing this forces a new resource.
- `integration_guid` (String) The GUID of the integration to attach. Changing this forces a new resource.

### Optional

- `variable_values` (Map of String) Key-value variable values for the integration (e.g. Slack channel name).
- `extra_email_addresses` (List of String) Additional email addresses to notify. Only applicable to Email integrations — do not provide for other types.
- `status_hub_service_list` (Block List) Status hub service mappings. Only applicable to Statushub integrations — do not provide for other types. Each block contains:
  - `monitor_guid` (String, Required) The GUID of the monitor.
  - `integration_service_guid` (String, Required) The GUID of the integration service.

### Write-Only
- `is_active_wo` (Boolean) Whether the integration is active.
- `is_active_wo_version` (Int64) Version of the `is_active_wo` field. Increment this value to re-send `is_active_wo` without changing other attributes.
- `send_ok_alerts_wo` (Boolean) Whether to send OK recovery alerts. Must be `false` for Phone and GenericWebhook integrations.
- `send_ok_alerts_wo_version` (Int64) Version of the `send_ok_alerts_wo` field. Increment this value to re-send `send_ok_alerts_wo` without changing other attributes.
- `send_reminder_alerts_wo` (Boolean) Whether to send reminder alerts. Must be `false` for GenericWebhook integrations.
- `send_reminder_alerts_wo_version` (Int64) Version of the `send_reminder_alerts_wo` field. Increment this value to re-send `send_reminder_alerts_wo` without changing other attributes.

### Read-Only

- `id` (String) Composite identifier in format `alertdefinition_id:escalation_level_id:integration_guid`.
- `integration_services` (List of String) Integration service GUIDs returned by the API.

## Import

Import is supported using the following syntax:

```shell
# Import using the composite identifier alertdefinition_id:escalation_level_id:integration_guid
terraform import itrs-uptrends_escalation_level_integration.example "046a727c-7a90-4776-9e41-ab050bdda5dc:1:a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

## Notes

- The `alertdefinition_id`, `escalation_level_id`, and `integration_guid` fields are immutable and require resource replacement when changed.
- The three boolean fields (`is_active_wo`, `send_ok_alerts_wo`, `send_reminder_alerts_wo`) are required by the API on create (POST). Always provide them explicitly.
- On update (PATCH), optional fields can be changed in place. If an optional field is removed from configuration, its previous value is preserved because all optional fields are marked as Computed.
- The integration GUID must reference an existing integration configured in the Uptrends account.
- The `extra_email_addresses` and `status_hub_service_list` fields are only sent to the API when explicitly provided. Do not provide them for integration types where they are not applicable — the API will reject the request.

