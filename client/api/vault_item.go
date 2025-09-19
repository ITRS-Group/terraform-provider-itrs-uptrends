package client

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type VaultItem struct {
	Client  *resty.Client
	BaseUrl string
}

// NewVaultItem creates and returns a new VaultItem client.
// Parameters:
//   - baseUrl: The base URL of the API (e.g., "https://api.uptrends.com/v4").
//   - authHeader: The value for the Authorization header.
func NewVaultItem(baseURL, authenticationHeader, version, platform string) *VaultItem {
	client := resty.New()

	// Set User Properties
	client.SetHeaders(map[string]string{
		"accept":        "application/json",
		"Content-Type":  "application/json",
		"authorization": authenticationHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &VaultItem{
		Client:  client,
		BaseUrl: baseURL,
	}
}

// GetVaultItem retrieves a specific vault section by its ID.
func (a *VaultItem) GetVaultItem(vaultItemID string) (*models.VaultItemResponse, error, string) {
	var vs models.VaultItemResponse
	url := fmt.Sprintf("%s/%s", a.BaseUrl, vaultItemID)
	resp, err := a.Client.R().
		SetResult(&vs).
		Get(url)
	if err != nil {
		return nil, err, ""
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("failed to get vault item: %s", resp.Status()), resp.String()
	}
	return &vs, nil, resp.String()
}

func (a *VaultItem) CreateVaultItem(requestData models.VaultItemRequest) (models.VaultItemResponse, int, error, string) {
	var vaultItemResponse models.VaultItemResponse
	marshalRequestData, err := json.Marshal(requestData)

	if err != nil {
		return models.VaultItemResponse{}, 0, fmt.Errorf("failed to marshal request data: %v", err), ""
	}

	resp, err := a.Client.R().
		SetBody(marshalRequestData).
		SetResult(&vaultItemResponse).
		Post(a.BaseUrl)

	if err != nil {
		return models.VaultItemResponse{}, 0, fmt.Errorf("failed to execute HTTP request: %v", err), resp.String()
	}

	// Return the response details
	return vaultItemResponse, resp.StatusCode(), nil, resp.String()
}

func (a *VaultItem) UpdateVaultItem(vaultItemID string, requestBody models.VaultItemRequest) (int, string, error) {
	url := fmt.Sprintf("%s/%s", a.BaseUrl, vaultItemID)
	marshalRequestData, err := json.Marshal(requestBody)

	if err != nil {
		return 0, "", fmt.Errorf("failed to marshal request data: %v", err)
	}

	resp, err := a.Client.R().
		SetBody(marshalRequestData).
		Patch(url)

	if err != nil {
		return 0, "", err
	}

	return resp.StatusCode(), resp.String(), nil
}

func (a *VaultItem) DeleteVaultItem(vaultItemID string) (int, string, error) {
	url := fmt.Sprintf("%s/%s", a.BaseUrl, vaultItemID)
	resp, err := a.Client.R().
		Delete(url)

	if err != nil {
		return 0, "", err
	}

	return resp.StatusCode(), resp.String(), nil
}
