---
page_title: "vault_section Resource - itrs-uptrends"
subcategory: ""
description: |-
  Manages vault sections in the Uptrends monitoring platform.  
  A list of relevant fields and their meaning can be found in the [API documentation for vault sections](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/Vault/Vault_GetAllVaultSections) and the [Uptrends support knowledge base](https://www.uptrends.com/support/kb/api/vault-api#vault-sections).
---

# vault_section (Resource)

## Example usage

```terraform
# Manage a vault section.
resource "vault_section" "section" {
  provider = itrs-uptrends.uptrendsauthenticated
  name     = "Section Name"
}
```

## Schema

### Required

- `name` (String) The name of the vault section.

### Read-only

- `id` (String) The unique identifier of the vault section.

## Import

Import is supported using the following syntax:

```shell
# Vault section can be imported by specifying the unique identifier.
terraform import vault_section.section "046a727c-7a90-4776-9e41-ab050bdda5dc"
```
