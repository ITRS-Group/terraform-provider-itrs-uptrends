package client

import models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"

// IVaultItem defines the interface for managing vault items.
type IVaultItem interface {
	GetVaultItem(vaultItemID string) (*models.VaultItemResponse, error, string)
	GetVaultItems() ([]models.VaultItemResponse, int, string, error)
	CreateVaultItem(requestData models.VaultItemRequest) (models.VaultItemResponse, int, error, string)
	UpdateVaultItem(vaultItemID string, requestBody models.VaultItemRequest) (int, string, error)
	DeleteVaultItem(vaultItemID string) (int, string, error)
}
