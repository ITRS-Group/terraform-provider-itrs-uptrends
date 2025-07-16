package converters

import (
	jsonmodels "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
	tfsdkmodels "github.com/itrs-group/terraform-provider-itrs-uptrends/provider/models"
)

func PayloadConversion(config tfsdkmodels.VaultItemResourceModel) jsonmodels.VaultItemRequest {
	payload := &jsonmodels.VaultItemRequest{
		Name:             config.Name.ValueString(),
		VaultSectionGuid: config.VaultSectionID.ValueString(),
		VaultItemType:    config.VaultItemType.ValueString(),
		Notes:            config.Notes.ValueString(),
	}

	// CredentialSet specific fields
	if !config.UserName.IsNull() {
		user := config.UserName.ValueString()
		payload.UserName = &user
	}
	if !config.Password.IsNull() {
		pass := config.Password.ValueString()
		payload.Password = &pass
	}

	// Certificate specific fields
	if !config.Value.IsNull() {
		val := config.Value.ValueString()
		payload.Value = &val
	}

	//CertificateArchive specific fields
	if config.CertificateArchive != nil {
		ca := config.CertificateArchive
		certArchive := &jsonmodels.CertificateArchive{
			Issuer:      ca.Issuer.ValueString(),
			NotBefore:   ca.NotBefore.ValueString(),
			NotAfter:    ca.NotAfter.ValueString(),
			ArchiveData: ca.ArchiveData.ValueString(),
		}
		if !ca.Password.IsNull() {
			p := ca.Password.ValueString()
			certArchive.Password = &p
		}
		payload.CertificateArchive = certArchive
	}

	if config.File != nil {
		f := config.File
		file := &jsonmodels.File{
			Data: f.Data.ValueString(),
			Name: f.Name.ValueString(),
		}
		payload.File = file
	}

	if config.OneTimePassword != nil {
		otp := config.OneTimePassword
		otpModel := &jsonmodels.OneTimePassword{
			Digits:        int(otp.Digits.ValueInt64()),
			Period:        int(otp.Period.ValueInt64()),
			HashAlgorithm: otp.HashAlgorithm.ValueString(),
		}
		if !otp.SecretEncodingMethod.IsNull() {
			sem := otp.SecretEncodingMethod.ValueString()
			otpModel.SecretEncodingMethod = &sem
		}
		if !otp.Secret.IsNull() {
			s := otp.Secret.ValueString()
			otpModel.Secret = &s
		}

		payload.OneTimePassword = otpModel
	}

	return *payload
}
