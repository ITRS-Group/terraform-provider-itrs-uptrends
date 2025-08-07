---
page_title: "vault_item Resource - itrs-uptrends"
subcategory: ""
description: |-
  Manages vault items in the Uptrends monitoring platform.  
  A list of relevant fields and their meaning can be found in the [API documentation for vault items](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/Vault) and the [Uptrends support knowledge base](https://www.uptrends.com/support/kb/api/vault-api).
---

# vault_item (Resource)

## Vault Item types

The following vault item types are supported:

- `CredentialSet` - Stores username and password credentials for authentication purposes.
- `Certificate` - Stores certificate data for SSL/TLS monitoring.
- `CertificateArchive` - Stores certificate archive files with password protection.
- `File` - Stores file data with metadata for file-based monitoring.
- `OneTimePassword` - Stores OTP configuration for TOTP/HOTP authentication.

## Example usage - CredentialSet

```terraform
resource "vault_item" "credential_set" {
  provider          = itrs-uptrends.uptrendsauthenticated
  name             = "Credential set vault name"
  vault_section_id = vault_section.section.id
  vault_item_type  = "CredentialSet"
  notes            = "your notes here"
  username         = "username"
  password_wo      = "password"
}

resource "vault_section" "section" {
  ...
}
```

**Required for CredentialSet:**
- `name` - The name of the vault item
- `vault_section_id` - The ID of the vault section this item belongs to
- `vault_item_type` - Must be set to `"CredentialSet"`
- `username` - The username for the credential set
- `password_wo` - The password (write-only, not stored in state)

## Example usage -  Certificate

```terraform
resource "vault_item" "certificate" {
  provider          = itrs-uptrends.uptrendsauthenticated
  name             = "Certificate vault name"
  vault_section_id = vault_section.section.id
  vault_item_type  = "Certificate"
  notes            = "your notes here"
  value_wo         = var.vault_value
}

resource "vault_section" "section" {
  ...
}
```

**Required for Certificate:**
- `name` - The name of the vault item
- `vault_section_id` - The ID of the vault section this item belongs to
- `vault_item_type` - Must be set to `"Certificate"`
- `value_wo` – The certificate value (write-only, not stored in state). For instructions on how to create a Base64-encoded certificate value suitable for this attribute, refer to the [example guide](https://github.com/ITRS-Group/terraform-provider-itrs-uptrends/blob/main/examples/vault_item/README.md#certificate-value-for-attribute-value).

## Example usage -  CertificateArchive

```terraform
resource "vault_item" "certificate_archive" {
  provider          = itrs-uptrends.uptrendsauthenticated
  name              = "Certificate archive vault name"
  vault_section_id  = vault_section.section.id
  vault_item_type   = "CertificateArchive"
  notes             = "your notes here"
  certificate_archive = {
    password_wo     = "password"
    archive_data_wo = var.vault_archive_data
  }
}

resource "vault_section" "section" {
  ...
}
```

**Required for CertificateArchive:**
- `name` - The name of the vault item
- `vault_section_id` - The ID of the vault section this item belongs to
- `vault_item_type` - Must be set to `"CertificateArchive"`
- `certificate_archive.password_wo` - Password for the certificate archive (write-only)
- `certificate_archive.archive_data_wo` - Base64 encoded certificate archive data (write-only). For instructions on how to create a Base64-encoded certificate archive value suitable for this attribute, refer to the [example guide](https://github.com/ITRS-Group/terraform-provider-itrs-uptrends/blob/main/examples/vault_item/README.md#certificate-archive-value-for-attribute-certificate_archivearchive_data).

**Read-only for CertificateArchive:**
- `certificate_archive.issuer` - The certificate issuer
- `certificate_archive.not_before` - The certificate start date
- `certificate_archive.not_after` - The certificate expiration date

## Example usage -  File

```terraform
resource "vault_item" "file" {
  provider          = itrs-uptrends.uptrendsauthenticated
  name              = "File vault name"
  vault_section_id  = vault_section.section.id
  vault_item_type   = "File"
  notes             = "your notes here"
  file = {
    data = var.vault_file_data
    name = "file_name.txt"
  }
}

resource "vault_section" "section" {
  ...
}
```

**Required for File:**
- `name` - The name of the vault item
- `vault_section_id` - The ID of the vault section this item belongs to
- `vault_item_type` - Must be set to `"File"`
- `file.data` - The file data content. For instructions on how to create a valid value for this attribute, refer to the [example guide](https://github.com/ITRS-Group/terraform-provider-itrs-uptrends/blob/main/examples/vault_item/README.md#file-data-value-for-attribute-filedata).
- `file.name` - The name of the file (can be any name, not just the original filename)

## Example usage -  OneTimePassword

```terraform
resource "vault_item" "one_time_password" {
  provider              = itrs-uptrends.uptrendsauthenticated
  name                  = "One time password vault name"
  vault_section_id      = vault_section.section.id
  vault_item_type       = "OneTimePassword"
  notes                 = "your notes here"
  one_time_password = {
    secret_wo                = "JBSWY3DPEHPK3PXP"
    digits                   = 6
    period                   = 30
    hash_algorithm           = "SHA256"
    secret_encoding_method_wo = "Base32"
  }
}

resource "vault_section" "section" {
  ...
}
```

**Required for OneTimePassword:**
- `name` - The name of the vault item
- `vault_section_id` - The ID of the vault section this item belongs to
- `vault_item_type` - Must be set to `"OneTimePassword"`
- `one_time_password.secret_wo` - The secret key for OTP generation (write-only)
- `one_time_password.digits` - Number of digits in the OTP (must be 6, 7, or 8)
- `one_time_password.period` - Time period in seconds for OTP generation (minimum 1)
- `one_time_password.hash_algorithm` - Hash algorithm for OTP generation (SHA256, SHA512, or SHA1)

**Optional for OneTimePassword:**
- `one_time_password.secret_encoding_method_wo` - Encoding method for the secret (Base32 or Hex, write-only)

## Common attributes

All vault item types share these common attributes:

### Required

- `name` (String) The name of the vault item.
- `vault_item_type` (String) The type of vault item. Must be one of: `CredentialSet`, `Certificate`, `CertificateArchive`, `File`, `OneTimePassword`.
- `vault_section_id` (String) The ID of the vault section this item belongs to.

### Optional

- `notes` (String) Optional notes for the vault item.

### Read-only

- `id` (String) The unique identifier of the vault item.

## Write-only fields

Several fields are marked as write-only (suffix `_wo`) and are not stored in the Terraform state for security reasons:

- `password_wo` - Password fields
- `value_wo` - Certificate values
- `secret_wo` - OTP secrets
- `archive_data_wo` - Certificate archive data
- `secret_encoding_method_wo` - OTP secret encoding method

These fields are required during creation and optional for updates but will not be visible in the Terraform state or plan output.

## Validation rules

### OneTimePassword validation

- `digits` must be one of: 6, 7, or 8
- `period` must be at least 1
- `hash_algorithm` must be one of: SHA256, SHA512, SHA1
- `secret_encoding_method_wo` must be one of: Base32, Hex

### Vault Item type validation

The `vault_item_type` field is immutable and requires resource replacement when changed.

## Import

Import is supported using the following syntax:

```shell
# Vault item can be imported by specifying the unique identifier.
terraform import vault_item.example "046a727c-7a90-4776-9e41-ab050bdda5dc"
```

## Notes

- The `vault_item_type` field cannot be changed after creation and will trigger a resource replacement.
- Write-only fields (marked with `_wo`) are sensitive and not stored in the Terraform state.
- Each vault item type has specific required attributes that must be provided based on the selected type.
- The resource automatically validates that all required attributes for the selected vault item type are provided.
- You can check in the git repo on how to create a valid value for `file.data`, `certificate_archive.archive_data_wo` and `value_wo` (https://github.com/ITRS-Group/terraform-provider-itrs-uptrends/blob/main/examples/vault_item/README.md).