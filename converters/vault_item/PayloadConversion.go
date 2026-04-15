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
	}

	if !config.Notes.IsNull() {
		notes := config.Notes.ValueString()
		payload.Notes = &notes
	} else {
		emptyString := ""
		payload.Notes = &emptyString
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

	// Certificate specific fields : value is write only and we don't store it
	if !config.Value.IsNull() {
		val := config.Value.ValueString()
		payload.Value = &val
	}

	//CertificateArchive specific fields
	if config.CertificateArchive != nil {
		ca := config.CertificateArchive
		certArchive := &jsonmodels.CertificateArchive{
			Issuer:    ca.Issuer.ValueStringPointer(),
			NotBefore: ca.NotBefore.ValueStringPointer(),
			NotAfter:  ca.NotAfter.ValueStringPointer(),
		}
		if !ca.Password.IsNull() {
			p := ca.Password.ValueString()
			certArchive.Password = &p
		}
		if !ca.ArchiveData.IsNull() {
			ad := ca.ArchiveData.ValueString()
			certArchive.ArchiveData = &ad
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
