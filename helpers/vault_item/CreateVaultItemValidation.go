package helpers

import (
	"fmt"

	tfsdkmodels "github.com/itrs-group/terraform-provider-itrs-uptrends/provider/models"
)

func CreateVaultItemValidation(config tfsdkmodels.VaultItemResourceModel) error {

	switch config.VaultItemType.ValueString() {
	case "CredentialSet":
		if config.Password.IsNull() {
			return fmt.Errorf("missing required fields for create: Vault item type 'CredentialSet' requires 'password' to be set")
		}
	case "Certificate":
		if config.Value.IsNull() {
			return fmt.Errorf("missing required fields for create: Vault item type 'Certificate' requires 'value' to be set")
		}
	case "CertificateArchive":
		if config.CertificateArchive.Password.IsNull() {
			return fmt.Errorf("missing required fields for create: Vault item type 'CertificateArchive' requires 'certificate_archive.password' to be set")
		}
	case "File":
		if config.File.Data.IsNull() {
			return fmt.Errorf("missing required fields for create: Vault item type 'File' requires 'file.data' to be set")
		}
		if config.File.Name.IsNull() {
			return fmt.Errorf("missing required fields for create: Vault item type 'File' requires 'file.name' to be set")
		}
	case "OneTimePassword":
		if config.OneTimePassword.Secret.IsNull() {
			return fmt.Errorf("missing required fields for create: Vault item type 'OneTimePassword' requires 'one_time_password.secret' to be set")
		}
		if config.OneTimePassword.SecretEncodingMethod.IsNull() {
			return fmt.Errorf("missing required fields for create: Vault item type 'OneTimePassword' requires 'one_time_password.secret_encoding_method' to be set")
		}
	default:
	}

	return nil
}
