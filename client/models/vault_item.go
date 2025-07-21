package client

type File struct {
	Data string `json:"Data"`
	Name string `json:"Name"`
}

type OneTimePassword struct {
	Secret               *string `json:"Secret,omitempty"` // can be null
	Digits               int     `json:"Digits"`
	Period               int     `json:"Period"`
	HashAlgorithm        string  `json:"HashAlgorithm"`
	SecretEncodingMethod *string `json:"SecretEncodingMethod,omitempty"` // can be null
}

type CertificateArchive struct {
	Issuer      string  `json:"Issuer"`
	NotBefore   string  `json:"NotBefore"`
	NotAfter    string  `json:"NotAfter"`
	Password    *string `json:"Password,omitempty"` // can be null
	ArchiveData string  `json:"ArchiveData"`
}

type VaultItemResponse struct {
	VaultItemGuid      string              `json:"VaultItemGuid"`
	Hash               string              `json:"Hash"`
	Name               string              `json:"Name"`
	VaultSectionGuid   string              `json:"VaultSectionGuid"`
	VaultItemType      string              `json:"VaultItemType"`
	Notes              string              `json:"Notes"`
	VaultItemUsedBy    string              `json:"VaultItemUsedBy"`
	CertificateArchive *CertificateArchive `json:"CertificateArchive,omitempty"`
	File               *File               `json:"FileInfo,omitempty"`
	OneTimePassword    *OneTimePassword    `json:"OneTimePasswordInfo,omitempty"`
	Value              *string             `json:"Value,omitempty"`    // for vault item type: "Certificate"(API), "Certificate public key"(APP)
	UserName           *string             `json:"UserName,omitempty"` // for vault item type: "CredentialSet"(API), "Credential set"(APP)
}

type VaultItemRequest struct {
	Name               string              `json:"Name"`
	VaultSectionGuid   string              `json:"VaultSectionGuid"`
	VaultItemType      string              `json:"VaultItemType"`
	Notes              *string             `json:"Notes,omitempty"`
	CertificateArchive *CertificateArchive `json:"CertificateArchive,omitempty"`
	File               *File               `json:"FileInfo,omitempty"`
	OneTimePassword    *OneTimePassword    `json:"OneTimePasswordInfo,omitempty"`
	Value              *string             `json:"Value,omitempty"`    // for vault item type: "Certificate"(API), "Certificate public key"(APP)
	UserName           *string             `json:"UserName,omitempty"` // for vault item type: "CredentialSet"(API), "Credential set"(APP)
	Password           *string             `json:"Password,omitempty"` // for vault item type: "CredentialSet"(API), "Credential set"(APP)
}
