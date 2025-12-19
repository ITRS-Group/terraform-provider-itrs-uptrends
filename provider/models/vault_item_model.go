package tfsdkmodels

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Any attribute that is added to the model must be added to the VaultItemResourceModelForValidation struct. In the VaultItemResourceModelForValidation struct, all attributes must be from types (eg.: types.String, types.Bool, types.Int64, types.List, types.Object).
type VaultItemResourceModel struct {
	ID                 types.String             `tfsdk:"id"`
	Name               types.String             `tfsdk:"name"`
	VaultSectionID     types.String             `tfsdk:"vault_section_id"`
	VaultItemType      types.String             `tfsdk:"vault_item_type"`
	Notes              types.String             `tfsdk:"notes"`
	Value              types.String             `tfsdk:"value_wo"`
	ValueVersion       types.Int64              `tfsdk:"value_wo_version"`
	UserName           types.String             `tfsdk:"username"`
	Password           types.String             `tfsdk:"password_wo"`
	PasswordVersion    types.Int64              `tfsdk:"password_wo_version"`
	CertificateArchive *CertificateArchiveModel `tfsdk:"certificate_archive"`
	File               *FileModel               `tfsdk:"file"`
	OneTimePassword    *OneTimePasswordModel    `tfsdk:"one_time_password"`
}

type OneTimePasswordModel struct {
	Secret                      types.String `tfsdk:"secret_wo"`
	SecretVersion               types.Int64  `tfsdk:"secret_wo_version"`
	Digits                      types.Int64  `tfsdk:"digits"`
	Period                      types.Int64  `tfsdk:"period"`
	HashAlgorithm               types.String `tfsdk:"hash_algorithm"`
	SecretEncodingMethod        types.String `tfsdk:"secret_encoding_method_wo"`
	SecretEncodingMethodVersion types.Int64  `tfsdk:"secret_encoding_method_wo_version"`
}

type FileModel struct {
	Data types.String `tfsdk:"data"`
	Name types.String `tfsdk:"name"`
}

type CertificateArchiveModel struct {
	Issuer             types.String `tfsdk:"issuer"`
	NotBefore          types.String `tfsdk:"not_before"`
	NotAfter           types.String `tfsdk:"not_after"`
	Password           types.String `tfsdk:"password_wo"`
	PasswordVersion    types.Int64  `tfsdk:"password_wo_version"`
	ArchiveData        types.String `tfsdk:"archive_data_wo"`
	ArchiveDataVersion types.Int64  `tfsdk:"archive_data_wo_version"`
}

type VaultItemResourceModelForValidation struct {
	ID                 types.String `tfsdk:"id"`
	Name               types.String `tfsdk:"name"`
	VaultSectionID     types.String `tfsdk:"vault_section_id"`
	VaultItemType      types.String `tfsdk:"vault_item_type"`
	Notes              types.String `tfsdk:"notes"`
	Value              types.String `tfsdk:"value_wo"`
	ValueVersion       types.Int64  `tfsdk:"value_wo_version"`
	UserName           types.String `tfsdk:"username"`
	Password           types.String `tfsdk:"password_wo"`
	PasswordVersion    types.Int64  `tfsdk:"password_wo_version"`
	CertificateArchive types.Object `tfsdk:"certificate_archive"`
	File               types.Object `tfsdk:"file"`
	OneTimePassword    types.Object `tfsdk:"one_time_password"`
}
