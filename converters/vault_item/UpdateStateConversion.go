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
		}
	case "Certificate":
		if vaultItem.Value != nil {
			state.Value = types.StringValue(*vaultItem.Value)
		}
	case "CertificateArchive":
		state.CertificateArchive = &tfsdkmodels.CertificateArchiveModel{
			Issuer:    types.StringValue(vaultItem.CertificateArchive.Issuer),
			NotBefore: types.StringValue(vaultItem.CertificateArchive.NotBefore),
			NotAfter:  types.StringValue(vaultItem.CertificateArchive.NotAfter),
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
		}
	default:
	}

	return state
}
