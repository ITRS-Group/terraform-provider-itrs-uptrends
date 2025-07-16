package tfsdkmodels

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type VaultItemResourceModel struct {
	ID                 types.String             `tfsdk:"id"`
	Name               types.String             `tfsdk:"name"`
	VaultSectionID     types.String             `tfsdk:"vault_section_id"`
	VaultItemType      types.String             `tfsdk:"vault_item_type"`
	Notes              types.String             `tfsdk:"notes"`
	Value              types.String             `tfsdk:"value"`
	UserName           types.String             `tfsdk:"username"`
	Password           types.String             `tfsdk:"password_wo"`
	CertificateArchive *CertificateArchiveModel `tfsdk:"certificate_archive"`
	File               *FileModel               `tfsdk:"file"`
	OneTimePassword    *OneTimePasswordModel    `tfsdk:"one_time_password"`
}

type OneTimePasswordModel struct {
	Secret               types.String `tfsdk:"secret_wo"`
	Digits               types.Int64  `tfsdk:"digits"`
	Period               types.Int64  `tfsdk:"period"`
	HashAlgorithm        types.String `tfsdk:"hash_algorithm"`
	SecretEncodingMethod types.String `tfsdk:"secret_encoding_method"`
}

type FileModel struct {
	Data types.String `tfsdk:"data"`
	Name types.String `tfsdk:"name"`
}

type CertificateArchiveModel struct {
	Issuer      types.String `tfsdk:"issuer"`
	NotBefore   types.String `tfsdk:"not_before"`
	NotAfter    types.String `tfsdk:"not_after"`
	Password    types.String `tfsdk:"password_wo"`
	ArchiveData types.String `tfsdk:"archive_data"`
}
