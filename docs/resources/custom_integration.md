---
page_title: "custom_integration Resource - itrs-uptrends"
subcategory: ""
description: |-
  Manages custom integrations in the Uptrends monitoring platform.
---

# itrs-uptrends_custom_integration (Resource)

Manages a custom integration in the Uptrends monitoring platform. Custom integrations let you define a set of HTTP steps that Uptrends executes (for example to deliver an alert to a third party system).

A list of relevant fields and their meaning can be found in the [API documentation for integrations](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/Integration) and the [Uptrends support knowledge base](https://www.uptrends.com/support/kb/api).

## Example Usage

```terraform
resource "itrs-uptrends_custom_integration" "example" {
  provider  = itrs-uptrends.uptrendsauthenticated
  name      = "My Custom Integration"
  is_active = true
  notes     = "Created and managed with Terraform"

  integration_variables = [
    {
      name                             = "webhook_url"
      value                            = "https://example.com/webhook"
      is_value_set_in_escalation_level = false
    },
  ]

  http_step_definitions = [
    {
      usages = ["Alert"]

      steps = [
        {
          name      = "Send alert"
          url       = "https://example.com/webhook"
          method    = "POST"
          body_type = "Raw"
          body      = "{\"message\":\"Alert triggered\"}"
          step_type = "HttpRequest"

          request_headers = [
            {
              key   = "Content-Type"
              value = "application/json"
            },
          ]

          authentication = {
            authentication_type = "Basic"
            user_name           = "api-user"
            password_wo         = "super-secret"
            password_wo_version = 1
          }
        },
      ]
    },
  ]
}
```

## Related Resources

- [itrs-uptrends_custom_integration_permission](custom_integration_permission.md) - Manage which operators or operator groups can use the integration.

## Schema

### Required

- `name` (String) The name of the custom integration.
- `http_step_definitions` (Attributes List) The HTTP step definitions of the custom integration. At least one definition is required, and each definition requires at least one step. See [below](#nested-schema-for-http_step_definitions).

### Optional

- `is_active` (Boolean) Whether the custom integration is active. Defaults to `true`. The API requires a non-null value, so a value is always sent.
- `notes` (String) Notes for the custom integration. Defaults to `""`. The API requires a non-null value, so an empty string is sent when unset.
- `integration_variables` (Attributes List) Variables that can be referenced by the integration steps. Always sent to the API (as an empty array when unset). See [below](#nested-schema-for-integration_variables).

### Read-Only

- `id` (String) The unique identifier (GUID) of the custom integration.

<a id="nested-schema-for-integration_variables"></a>
### Nested Schema for `integration_variables`

- `name` (String, **Required**) The name of the variable.
- `value` (String) The value of the variable.
- `is_value_set_in_escalation_level` (Boolean) Whether the value is set in the escalation level. Defaults to `false`.

<a id="nested-schema-for-http_step_definitions"></a>
### Nested Schema for `http_step_definitions`

- `usages` (List of String) The usages of the HTTP step definition (e.g. `Alert`, `Reminder`, `Ok`).
- `http_step_definition_usage_guid` (String) The usage GUID of the HTTP step definition (computed by the API; leave empty when creating).
- `steps` (Attributes List, **Required**) The HTTP steps to execute. At least one step is required. Each step supports the following:
  - `name`, `url`, `method`, `body`, `body_type`, `vault_file_id`, `step_type`, `client_certificate_vault_item`, `pre_request_script`, `post_response_script`, `calculated_content_type`, `msa_step_http_version` (String)
  - `use_fixed_client_certificate`, `ignore_certificate_errors`, `retry_until_successful` (Boolean)
  - `delay`, `max_attempts`, `retry_wait_milliseconds` (Number)
  - `allowed_tls_versions` (List of String) Allowed TLS versions (e.g. `Tls11`).
  - `multi_part_form` (Attributes List) Multipart form entries with `type`, `key`, `value` (String).
  - `request_headers` (Attributes List) Request headers with `key`, `value` (String).
  - `variables` (Attributes List) Variables captured from the response with `source`, `property`, `name` (String) and `arguments` (List of String).
  - `assertions` (Attributes List) Assertions with `source`, `property`, `comparison`, `target_value` (String).
  - `authentication` (Attributes) Authentication settings with `id`, `authentication_type`, `user_name` (String), `password_wo` (String, sensitive, **write-only** — sent to the API but never stored in state) and `password_wo_version` (Number — increment to re-send a changed `password_wo`).
- `user_defined_functions` (Attributes List) Reusable functions. Each supports `name`, `type`, `regex`, `jwt_signing_key`, `jwt_algorithm`, `hash_key`, `hash_algorithm` (String) and `mappings` (Attributes List of `key`/`value`).

## Import

Import is supported using the following syntax:

```shell
# Custom integration can be imported by specifying the integration GUID.
terraform import itrs-uptrends_custom_integration.example "integration-guid"
```

## Notes

- The `password_wo` field inside `authentication` is write-only: it is sent to the API but never stored in Terraform state. Because Terraform cannot detect changes to write-only values, increment `password_wo_version` whenever you change `password_wo` to trigger an update.
- `http_step_definitions` and `steps` are required (and each must contain at least one item); `authentication` is optional. These blocks are not computed, because the framework does not allow write-only attributes (`password_wo`) nested under a computed attribute. Most leaf fields are optional and computed; when omitted, the value returned by the API is stored in state.
- Updates are sent as a full replace: the provider always sends the entire object (all top-level fields and full nested lists, including `integration_variables` as `[]` when empty) on every apply, because the API performs no server-side defaulting on update.
