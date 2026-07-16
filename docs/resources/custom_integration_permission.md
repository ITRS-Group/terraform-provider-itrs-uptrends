---
page_title: "custom_integration_permission Resource - itrs-uptrends"
subcategory: ""
description: |-
  Manages authorization permissions for custom integrations in the Uptrends monitoring platform.
---

# itrs-uptrends_custom_integration_permission (Resource)

Manages authorization permissions for custom integrations in the Uptrends monitoring platform. A permission grants an operator or operator group the right to use a custom integration.

A list of relevant fields and their meaning can be found in the [API documentation for integrations](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/Integration) and the [Uptrends support knowledge base](https://www.uptrends.com/support/kb/api).

## Example Usage

### Grant an operator use access

```terraform
resource "itrs-uptrends_custom_integration_permission" "operator_use" {
  provider       = itrs-uptrends.uptrendsauthenticated
  integration_id = itrs-uptrends_custom_integration.example.id
  permission     = "UseIntegration"
  operator_id    = itrs-uptrends_operator.example.id
}
```

### Grant an operator group use access

```terraform
resource "itrs-uptrends_custom_integration_permission" "group_use" {
  provider         = itrs-uptrends.uptrendsauthenticated
  integration_id   = itrs-uptrends_custom_integration.example.id
  permission       = "UseIntegration"
  operatorgroup_id = itrs-uptrends_operatorgroup.example.id
}
```

## Related Resources

- [itrs-uptrends_custom_integration](custom_integration.md) - Create and manage custom integrations.
- [itrs-uptrends_operator](operator.md) - Manage operators.
- [itrs-uptrends_operatorgroup](operatorgroup.md) - Manage operator groups.

## Schema

### Required

- `integration_id` (String) The GUID of the custom integration.
- `permission` (String) The integration authorization type. Valid value: `UseIntegration`.

### Optional

- `operator_id` (String) The GUID of the operator. Provide this or `operatorgroup_id`, not both.
- `operatorgroup_id` (String) The GUID of the operator group. Provide this or `operator_id`, not both.

### Read-Only

- `id` (String) The unique identifier of the authorization (composite key in format `integration_id:authorization_id`).

## Import

Import is supported using the following syntax:

```shell
# Custom integration permission can be imported by specifying the composite identifier integration_id:authorization_id.
terraform import itrs-uptrends_custom_integration_permission.example "integration-guid:authorization-guid"
```

## Notes

- All attributes are immutable, so changing any value requires resource replacement.
- Exactly one of `operator_id` or `operatorgroup_id` must be provided.
- When creating an authorization, the API may add additional dependent authorizations automatically.
- Removing a permission does not delete the integration, operator, or operator group, only the association.
