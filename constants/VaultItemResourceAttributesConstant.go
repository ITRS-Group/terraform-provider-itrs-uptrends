package constants

import "github.com/itrs-group/terraform-provider-itrs-uptrends/helpers"

var allVaultItemAttributes = []string{
	"vault_section_id", // Required within the schema
	"vault_item_type",  // Required within the schema
	"name",             // Required within the schema
	"notes",
}

// VaultItemResourceAttributes defines the required and optional attributes for each vault item type.
// This mapping is used to validate resource configurations during Create and Update operations.
// This helps ensure that the correct fields are provided for each type of vault item in Terraform.
var VaultItemResourceAttributes = map[string]helpers.ResourceAttributes{
	"CredentialSet": {
		RequiredAttributes: []string{},
		OptionalAttributes: append([]string{"password_wo", "password_wo_version", "username"}, allVaultItemAttributes...), // The password_wo is optional for the Update
	},
	"Certificate": {
		RequiredAttributes: []string{},
		OptionalAttributes: append([]string{"value_wo", "value_wo_version"}, allVaultItemAttributes...), // The value is optional for the Update
	},
	"CertificateArchive": {
		RequiredAttributes: []string{},
		OptionalAttributes: append([]string{"certificate_archive"}, allVaultItemAttributes...), // The certificate_archive is optional for the Update
	},
	"File": {
		RequiredAttributes: []string{},
		OptionalAttributes: append([]string{"file"}, allVaultItemAttributes...), // The file is optional for the Update
	},
	"OneTimePassword": {
		RequiredAttributes: []string{},
		OptionalAttributes: append([]string{"one_time_password"}, allVaultItemAttributes...), // The one_time_password is optional for the Update
	},
}

var VaultItemResourceAttributesCreate = map[string]helpers.ResourceAttributes{
	"CredentialSet": {
		RequiredAttributes: []string{"password_wo"},
		OptionalAttributes: append([]string{"username", "password_wo_version"}, allVaultItemAttributes...),
	},
	"Certificate": {
		RequiredAttributes: []string{"value_wo"},
		OptionalAttributes: append([]string{"value_wo_version"}, allVaultItemAttributes...),
	},
	"CertificateArchive": {
		RequiredAttributes: []string{"certificate_archive"},
		OptionalAttributes: append([]string{}, allVaultItemAttributes...),
	},
	"File": {
		RequiredAttributes: []string{"file"},
		OptionalAttributes: append([]string{}, allVaultItemAttributes...),
	},
	"OneTimePassword": {
		RequiredAttributes: []string{"one_time_password"},
		OptionalAttributes: append([]string{}, allVaultItemAttributes...),
	},
}
