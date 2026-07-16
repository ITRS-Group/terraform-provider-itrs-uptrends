---
page_title: "itrs-uptrends_vault_section Data Source - itrs-uptrends"
subcategory: ""
description: |-
  Fetch a vault section by GUID or name to reference it in vault item creation or other resources.
---

# itrs-uptrends_vault_section (Data Source)

Use this data source to retrieve a vault section before creating items or linking other resources.

## Example Usage

```terraform
data "itrs-uptrends_vault_section" "by_name" {
  name = "Shared Vault"
}

data "itrs-uptrends_vault_section" "by_id"{
    id = "9a0bbe79-4944-42f4-ace8-dfac33af9a64"
} 
```

## Schema

### Optional
- `id` (String) Vault section GUID. Provide this or `name`.
- `name` (String) Vault section name. Provide this or `id`. If the name is not unique it is going to give an error.

