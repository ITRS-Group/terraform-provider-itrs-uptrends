# Changelog

All notable changes to this provider are documented in this file.

## [1.6.0]

### Added
- Added support for additional monitor group quota models in `itrs-uptrends_monitorgroup`:
  - `unified_credits_quota`
  - `classic_quota`
- Added quota usage attributes to the `itrs-uptrends_monitorgroup` data source, including:
  - `used_unified_credits_quota`
  - `used_classic_quota`
  - `used_transaction_monitor_quota`
  - `used_browser_monitor_quota`
  - `used_api_monitor_quota`
  - `used_basic_monitor_quota`

### Fixed
- Fixed an issue where validation for certain fields incorrectly fails when using Terraform variables or other computed expressions, instead of literal values. These fields now correctly accept unknown or computed values:
  - `escalation_levels` in `itrs-uptrends_alertdefinition`
  - `check_interval` and `check_interval_seconds` in `itrs-uptrends_monitor`

## [1.5.0]

### Added
- Added support for RUM management:
  - New resource: `itrs-uptrends_rum_website`
  - New data source: `itrs-uptrends_rum_website`

### Fixed
- Resolved monitor and vault item data source schema mismatch that caused `Value Conversion Error`.

## [1.4.0]

### Added
- Write-only version counters (`_wo_version`) for secret-style attributes.  
  Terraform cannot detect diffs to write-only sensitive values (for example, `password_wo`). Previously, updating only these fields did not trigger a resource update.
  You must now increase the matching `_wo_version` field when the related `_wo` value changes.

- New data sources for major Uptrends objects.These allow referencing existing infrastructure by name or description instead of hard-coded GUIDs:  
  - `itrs-uptrends_monitor`
  - `itrs-uptrends_operator`
  - `itrs-uptrends_alertdefinition`
  - `itrs-uptrends_monitorgroup`
  - `itrs-uptrends_operatorgroup`
  - `itrs-uptrends_vault_item`
  - `itrs-uptrends_vault_section`
  - `itrs-uptrends_checkpoint`
  - `itrs-uptrends_checkpoint_region`

### Changed
- `_wo_version` counters are now available for:
  - **Operator**: `password_wo_version`
  - **Monitor**: `password_wo_version`  
    (`initial_monitor_group_id_wo` does not have a version counter)
  - **Vault item**:
    - `value_wo_version`
    - `password_wo_version`
    - `certificate_archive.password_wo_version`
    - `certificate_archive.archive_data_wo_version`
    - `one_time_password.secret_wo_version`
    - `one_time_password.secret_encoding_method_wo_version`

### Notes
-  `_wo_version` attributes are state-only fields and are not sent to the API.
- Always increase the relevant `_wo_version` value when updating its matching `_wo` attribute.
- Data sources accept `id` or `name` (and for groups also `description`) and return read-only attributes for the matching object.
- Checkpoint data sources support lists of names, or no filter, to return multiple IDs or metadata entries.

## [1.3.0]

Intermediate release.

## [1.2.0]

### Added
- `check_interval_seconds` for the monitor resource. Use this instead of `check_interval` (minutes) when more frequent checks are needed.
- `initial_monitor_group_id_wo` for the monitor resource, to assign the monitor to a specific monitor group at creation time.

## [1.1.1]

### Fixed
- Monitor validation now correctly accepts Terraform variables.  
  Previously, variables could be treated as missing values during validation.

## [1.1.0]
### Changed
- Terraform resource names were updated to align with HashiCorp naming recommendations.

### Added
- Postman monitor type support in the monitor resource.
- `use_w3c_total_time` and `http_version` attributes added to the monitor schema.

## [1.0.2]

### Added
- Provider documentation published in the Terraform Registry (using the old resource names at that time).
## [1.0.1]

### Added
- Monitor validation in Terraform.

## [1.0.0]

### Added
- Initial release of the Uptrends Terraform provider.
- Initial supported resources:
  - `alertdefinition_monitor_membership_resource`
  - `alertdefinition_operator_membership_resource`
  - `alertdefinition_operatorgroup_membership_resource`
  - `alertdefinition_resource`
  - `monitor_resource`
  - `monitorgroup_membership_resource`
  - `monitorgroup_resource`
  - `operator_permission_resource`
  - `operator_resource`
  - `operatorgroup_permission_resource`
  - `operatorgroup_resource`
  - `vault_item_resource`
  - `vault_section_resource`
