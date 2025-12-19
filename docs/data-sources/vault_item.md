---
page_title: "itrs-uptrends_vault_item Data Source - itrs-uptrends"
subcategory: ""
description: |-
  Retrieve vault item metadata for credentials, certificates, files, or OTP definitions by GUID or name.
---

# itrs-uptrends_vault_item (Data Source)

Look up a vault item to reuse its configuration in other resources rather than duplicating data.

## Example Usage

```terraform
data "itrs-uptrends_vault_item" "by_name" {
  name = "Database credentials"
}

data "itrs-uptrends_vault_item" "by_id"{
    id = "9a0bbe79-4944-42f4-ace8-dfac33af9a64"
} 
```

## Schema

### Optional
- `id` (String) Vault item GUID. Provide this or `name`.
- `name` (String) Vault item name. Provide this or `id`. If the name is not unique it is going to give an error.

### Read-Only
- `vault_section_id` (String) Section GUID that owns the item.
- `vault_item_type` (String) Type (CredentialSet, Certificate, File, OneTimePassword, etc.).
- `notes` (String) Notes stored on the vault item.
- `username` (String) Username for credential sets.
- `value_wo`, `password_wo` (String) Write-only values rendered as empty strings for safety.
- `certificate_archive`, `file`, `one_time_password` (Objects) Nested metadata (see schema).
