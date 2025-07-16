package client

import models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"

// IVaultSection defines the interface for managing vault sections.
type IVaultSection interface {
	GetVaultSection(VaultSectionGuid string) (*models.VaultSection, error, string)
	CreateVaultSection(name string) (*models.VaultSection, error, string)
	UpdateVaultSection(VaultSectionGuid string, name string) (string, error)
	DeleteVaultSection(VaultSectionGuid string) (error, string)
}
