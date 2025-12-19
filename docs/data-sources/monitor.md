---
page_title: "itrs-uptrends_monitor Data Source - itrs-uptrends"
subcategory: ""
description: |-
  Fetch monitor configuration and metadata by GUID or name for reuse across resources.
---

# itrs-uptrends_monitor (Data Source)

Look up a monitor by `id` or `name`. The data source populates all monitor fields (type, intervals, checkpoints, etc.) so you can reference them elsewhere without re-fetching manually.

## Example Usage

```terraform
data "itrs-uptrends_monitor" "by_name" {
  name = "FullPageCheck monitor"
}

data "itrs-uptrends_monitor" "by_id" {
  id = "78b0ca43-46ec-44c2-93bb-7ad3ffa5ce19"
}
```

## Schema

### Optional
- `id` (String) Monitor GUID. Provide this or `name`.
- `name` (String) Monitor name. Provide this or `id`. If the name is not unique it is going to give an error.

### Read-Only (highlights)
- **Core flags & intervals**
  - `monitor_type`, `monitor_mode`, `generate_alert`, `is_active`
  - `check_interval`, `check_interval_seconds`
  - `use_primary_checkpoints_only`, `use_w3c_total_time`
- **Checkpoint selection**
  - `selected_checkpoints` object (`checkpoints`, `regions`, `exclude_locations`)
- **Content & requests**
  - `notes`, `custom_metrics`, `custom_fields`
  - `request_headers`, `block_urls`, `user_agent`
  - `username`, `password_wo`, `name_for_phone_alerts`
  - `self_service_transaction_script`, `multi_step_api_transaction_script`
- **Authentication & throttling**
  - `authentication_type`, `throttling_options` (type/value/speed/latency)
- **DNS / network**
  - `dns_bypasses`, `dns_server`, `dns_query`, `dns_expected_result`, `dns_test_value`
  - `ip_version`, `port`, `network_address`, `database_name`
- **Browser / HTTP**
  - `browser_type`, `browser_window_dimensions` (is_mobile, width, height, pixel_ratio, mobile_device)
  - `http_method`, `http_version`, `tls_version`, `request_body`, `url`
- **Certificates & security**
  - `certificate_name`, `certificate_organization`, `certificate_organizational_unit`
  - `certificate_serial_number`, `certificate_fingerprint`
  - `certificate_issuer_*` fields, `certificate_expiration_warning_days`
  - `check_certificate_errors`, `ignore_external_elements`
- **Concurrency & error handling**
  - `use_concurrent_monitoring`
  - `concurrent_unconfirmed_error_threshold`, `concurrent_confirmed_error_threshold`
  - `error_conditions` (type/value/percentage/level/match/effect)
- **Other**
  - `block_google_analytics`, `block_uptrends_rum`
  - `predefined_variables`
  - `created_date`, `postman_collection_json`
  - `initial_monitor_group_id_wo`

Refer to the monitor resource documentation for a better understanding of the fields.
