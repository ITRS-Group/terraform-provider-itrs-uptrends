resource "vault_section" "section" {
  provider          = itrsuptrends.uptrendsauthenticated
  name             = "Section Name"
}

# Import example:
# Import States available in the Uptrends APP for downloading as a tf file:
import {
  to = vault_section.section_imported
  id  = "${vault_section.section.id}" # Replace with the actual ID (e.g. "046a727c-7a90-4776-9e41-ab050bdda5dc")
  provider          = itrsuptrends.uptrendsauthenticated
}