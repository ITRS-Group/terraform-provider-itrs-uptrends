package client

import models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"

type IVaultSectionPermission interface {
	GetVaultSectionAuthorizations(vaultSectionGuid string) ([]models.VaultSectionAuthorization, error)
	CreateVaultSectionAuthorization(vaultSectionGuid string, auth models.VaultSectionAuthorization) (*models.VaultSectionAuthorization, error)
	DeleteVaultSectionAuthorization(vaultSectionGuid, authorizationGuid string) error
}
