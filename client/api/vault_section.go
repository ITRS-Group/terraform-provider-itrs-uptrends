// // filepath: c:\Users\ioneus\source\repos\Terraform\client\api\vaultsection.go
package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

var _ interfaces.IVaultSection = (*VaultSection)(nil)

// VaultSection encapsulates the methods to interact with the VaultSection API.
type VaultSection struct {
	client  *resty.Client
	baseURL string
}

// NewVaultSection creates and returns a new VaultSection API client.
func NewVaultSection(baseURL, authHeader, version, platform string) *VaultSection {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"accept":        "application/json",
		"authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &VaultSection{
		client:  client,
		baseURL: baseURL,
	}
}

// GetVaultSections lists all vault sections.
func (api *VaultSection) GetVaultSections() ([]models.VaultSection, int, string, error) {
	var sections []models.VaultSection

	resp, err := api.client.R().
		SetResult(&sections).
		Get(api.baseURL)

	statusCode := -1
	responseBody := ""
	if resp != nil {
		statusCode = resp.StatusCode()
		responseBody = resp.String()
	}

	if err != nil {
		return nil, statusCode, responseBody, err
	}
	if !resp.IsSuccess() {
		return nil, statusCode, responseBody, fmt.Errorf("failed to list vault sections: %s", resp.Status())
	}

	return sections, statusCode, responseBody, nil
}

// GetVaultSection retrieves a specific vault section by its ID.
func (api *VaultSection) GetVaultSection(VaultSectionGuid string) (*models.VaultSection, error, string) {
	var vs models.VaultSection
	url := fmt.Sprintf("%s/%s", api.baseURL, VaultSectionGuid)
	resp, err := api.client.R().
		SetResult(&vs).
		Get(url)
	if err != nil {
		return nil, err, ""
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("failed to get vault section: %s", resp.Status()), resp.String()
	}
	return &vs, nil, resp.String()
}

// CreateVaultSection creates a new vault section.
func (api *VaultSection) CreateVaultSection(name string) (*models.VaultSection, error, string) {
	var createdVS models.VaultSection
	resp, err := api.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{"Name": name}).
		SetResult(&createdVS).
		Post(api.baseURL)
	if err != nil {
		return nil, err, ""
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("failed to create vault section: %s", resp.Status()), ""
	}
	return &createdVS, nil, resp.String()
}

// UpdateVaultSection updates an existing vault section.
func (api *VaultSection) UpdateVaultSection(vaultSectionID string, name string) (string, error) {
	payload := map[string]interface{}{
		"VaultSectionGuid": vaultSectionID,
		"Name":             name,
	}
	url := fmt.Sprintf("%s/%s", api.baseURL, vaultSectionID)

	resp, err := api.client.R().
		SetHeader("Accept", "application/json").
		SetBody(payload).
		Put(url)

	if err != nil {
		return "", err
	}

	return resp.String(), err
}

// DeleteVaultSection deletes a vault section.
func (api *VaultSection) DeleteVaultSection(vaultSectionID string) (error, string) {
	url := fmt.Sprintf("%s/%s", api.baseURL, vaultSectionID)
	resp, err := api.client.R().
		Delete(url)
	if err != nil {
		return err, ""
	}
	return err, resp.String()
}
