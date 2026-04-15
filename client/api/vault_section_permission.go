package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

var _ interfaces.IVaultSectionPermission = (*VaultSectionPermission)(nil)

type VaultSectionPermission struct {
	client  *resty.Client
	baseURL string
}

func NewVaultSectionPermission(baseURL, authHeader, version, platform string) *VaultSectionPermission {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"accept":        "application/json",
		"authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &VaultSectionPermission{
		client:  client,
		baseURL: baseURL,
	}
}

func (c *VaultSectionPermission) GetVaultSectionAuthorizations(vaultSectionGuid string) ([]models.VaultSectionAuthorization, error) {
	url := fmt.Sprintf("%s/%s/Authorization", c.baseURL, vaultSectionGuid)
	var authorizations []models.VaultSectionAuthorization
	resp, err := c.client.R().
		SetResult(&authorizations).
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to retrieve vault section authorizations: %s %s", resp.Status(), resp.Body())
	}
	return authorizations, nil
}

func (c *VaultSectionPermission) CreateVaultSectionAuthorization(vaultSectionGuid string, auth models.VaultSectionAuthorization) (*models.VaultSectionAuthorization, error) {
	url := fmt.Sprintf("%s/%s/Authorization", c.baseURL, vaultSectionGuid)
	var created models.VaultSectionAuthorization
	resp, err := c.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(auth).
		SetResult(&created).
		Post(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to create vault section authorization: %s %s", resp.Status(), resp.Body())
	}
	return &created, nil
}

func (c *VaultSectionPermission) DeleteVaultSectionAuthorization(vaultSectionGuid, authorizationGuid string) error {
	url := fmt.Sprintf("%s/%s/Authorization/%s", c.baseURL, vaultSectionGuid, authorizationGuid)
	resp, err := c.client.R().Delete(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("failed to delete vault section authorization: %s %s", resp.Status(), resp.Body())
	}
	return nil
}
