---
page_title: "vault_section_permission Resource - itrs-uptrends"
subcategory: ""
description: |-
  Manages authorization permissions for vault sections in the Uptrends monitoring platform.
---

# itrs-uptrends_vault_section_permission (Resource)

Manages authorization permissions for vault sections in the Uptrends monitoring platform.
A list of relevant fields and their meaning can be found in the [API documentation for vault sections](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/VaultSection) and the [Uptrends support knowledge base](https://www.uptrends.com/support/kb/api).

## Example usage

### Grant an operator view access to a vault section

```terraform
resource "itrs-uptrends_vault_section_permission" "operator_view" {
  provider           = itrs-uptrends.uptrendsauthenticated
  vault_section_id   = itrs-uptrends_vault_section.example.id
  permission = "ViewVaultSection"
  operator_id        = itrs-uptrends_operator.example.id
}
```

### Grant an operator group change access to a vault section

```terraform
resource "itrs-uptrends_vault_section_permission" "group_change" {
  provider           = itrs-uptrends.uptrendsauthenticated
  vault_section_id   = itrs-uptrends_vault_section.example.id
  permission = "ChangeVaultSection"
  operatorgroup_id   = itrs-uptrends_operatorgroup.example.id
}
```

## Use cases

Vault section permissions control which operators or operator groups can view or modify secrets stored in a vault section.

## Related resources

- [itrs-uptrends_vault_section](vault_section.md) - Create and manage vault sections
- [itrs-uptrends_operator](operator.md) - Manage operators
- [itrs-uptrends_operatorgroup](operatorgroup.md) - Manage operator groups

## Schema

### Required

- `permission` (String) The authorization type. Valid values: `ViewVaultSection`, `ChangeVaultSection`.
- `vault_section_id` (String) The GUID of the vault section.

### Optional

- `operator_id` (String) The GUID of the operator. Provide this or `operatorgroup_id`, not both.
- `operatorgroup_id` (String) The GUID of the operator group. Provide this or `operator_id`, not both.

### Read-Only

- `id` (String) The unique identifier of the authorization (composite key in format `vault_section_id:authorization_id`).

## Import

Import is supported using the following syntax:

```shell
# Vault section permission can be imported by specifying the composite identifier vault_section_id:authorization_id.
terraform import itrs-uptrends_vault_section_permission.example "vault-section-guid:authorization-guid"
```

## Notes

- All attributes are immutable — changing any value requires resource replacement.
- Exactly one of `operator_id` or `operatorgroup_id` must be provided.
- Removing a permission does not delete the vault section, operator, or operator group — only the association.
