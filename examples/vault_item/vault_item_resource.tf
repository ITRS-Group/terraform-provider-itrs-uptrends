# Type: CredentialSet
resource "vault_item" "credential_set" {
  provider          = itrsuptrends.uptrendsauthenticated
  name             = "CredentialSet Vault Name"
  vault_section_id = "046a727c-7a90-4776-9e41-ab050bdda5dc" # Replace with your actual section ID
  vault_item_type  = "CredentialSet"
  notes            = "your notes here" # Optional, can be ""
  username         = "username"
  password_wo         = "password" # WriteOnly, not saved in the state
}

# Type: Certificate
resource "vault_item" "certificate" {
  provider          = itrsuptrends.uptrendsauthenticated
  name             = "Certificate Vault Name"
  vault_section_id = "046a727c-7a90-4776-9e41-ab050bdda5dc" # Replace with your actual section ID
  vault_item_type  = "Certificate"
  notes            = "your notes here" # Optional, can be ""
  value = var.vault_value # Check and update value inside variables.tf
}

# Type: CertificateArchive
resource "vault_item" "certificate_archive" {
  provider          = itrsuptrends.uptrendsauthenticated
  name              = "CertificateArchive Vault Name"
  vault_section_id  = "046a727c-7a90-4776-9e41-ab050bdda5dc" # Replace with your actual section ID
  vault_item_type   = "CertificateArchive"
  notes             = "your notes here" # Optional, can be ""
  certificate_archive = {
    password_wo     = "password" # WriteOnly, not saved in the state
    archive_data = var.vault_certificate_archive # Check and update value inside variables.tf
  }
}

# Type: File
resource "vault_item" "file" {
  provider          = itrsuptrends.uptrendsauthenticated
  name              = "File Vault Name"
  vault_section_id  = "046a727c-7a90-4776-9e41-ab050bdda5dc" # Replace with your actual section ID
  vault_item_type   = "File"
  notes             = "your notes here" # Optional, can be ""
  file = {
    data = var.vault_file_data # Check and update value inside variables.tf
    name = "file_name.txt" # The name can be anyhting, not just the file name
  }
}

# Type: OneTimePassword
resource "vault_item" "one_time_password" {
  provider              = itrsuptrends.uptrendsauthenticated
  name                  = "OneTimePassword Vault Name"
  vault_section_id      = "046a727c-7a90-4776-9e41-ab050bdda5dc" # Replace with your actual section ID
  vault_item_type       = "OneTimePassword"
  notes                 = "your notes here" # Optional, can be ""
  one_time_password = {
    secret_wo                = "JBSWY3DPEHPK3PXP" # WriteOnly, not saved in the state
    digits                = 6
    period                = 30
    hash_algorithm        = "SHA256"
    secret_encoding_method = "Base32"
  }
}

# Import example:
# Import States available in the Uptrends APP for downloading as a tf file:
import {
  to = vault_item.credential_set_imported
  id = "${vault_item.credential_set.id}" # Replace with the actual ID (e.g. "046a727c-7a90-4776-9e41-ab050bdda5dc")
  provider          = itrsuptrends.uptrendsauthenticated
}