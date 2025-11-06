package converters

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	jsonmodels "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
	tfsdkmodels "github.com/itrs-group/terraform-provider-itrs-uptrends/provider/models"
)

func UpdateStateConversion(vaultItem *jsonmodels.VaultItemResponse) tfsdkmodels.VaultItemResourceModel {
	var state tfsdkmodels.VaultItemResourceModel
	state.ID = types.StringValue(vaultItem.VaultItemGuid)
	state.Name = types.StringValue(vaultItem.Name)
	state.VaultSectionID = types.StringValue(vaultItem.VaultSectionGuid)
	state.VaultItemType = types.StringValue(vaultItem.VaultItemType)
	state.Notes = types.StringValue(vaultItem.Notes)

	switch vaultItem.VaultItemType {
	case "CredentialSet":
		if vaultItem.UserName != nil {
			state.UserName = types.StringValue(*vaultItem.UserName)
		} else {
			emptyString := ""
			state.UserName = types.StringValue(emptyString)
		}
		// Password is write only and we don't store it
	case "Certificate":
		// Value is write only and we don't store it
	case "CertificateArchive":
		state.CertificateArchive = &tfsdkmodels.CertificateArchiveModel{
			Issuer:    types.StringValue(*vaultItem.CertificateArchive.Issuer),
			NotBefore: types.StringValue(*vaultItem.CertificateArchive.NotBefore),
			NotAfter:  types.StringValue(*vaultItem.CertificateArchive.NotAfter),
			// ArchiveData is write only and we don't store it
			// Password is write only and we don't store it
		}
	case "File":
		state.File = &tfsdkmodels.FileModel{
			Data: types.StringValue(vaultItem.File.Data),
			Name: types.StringValue(vaultItem.File.Name),
		}
	case "OneTimePassword":
		state.OneTimePassword = &tfsdkmodels.OneTimePasswordModel{
			Digits:        types.Int64Value(int64(vaultItem.OneTimePassword.Digits)),
			Period:        types.Int64Value(int64(vaultItem.OneTimePassword.Period)),
			HashAlgorithm: types.StringValue(vaultItem.OneTimePassword.HashAlgorithm),
			// Secret is write only and we don't store it
			// SecretEncodingMethod is write only and we don't store it
		}
	default:
		state.UserName = types.StringNull()
	}

	return state
}
