package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

var _ interfaces.ICustomIntegrationPermission = (*CustomIntegrationPermission)(nil)

// CustomIntegrationPermission encapsulates the methods to interact with the
// /Integration/{integrationGuid}/Authorizations endpoints of the API.
type CustomIntegrationPermission struct {
	client  *resty.Client
	baseURL string
}

// NewCustomIntegrationPermission creates a new API client instance. baseURL is
// expected to be the /Integration endpoint.
func NewCustomIntegrationPermission(baseURL, authHeader, version, platform string) *CustomIntegrationPermission {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &CustomIntegrationPermission{
		client:  client,
		baseURL: baseURL,
	}
}

// GetIntegrationAuthorizations gets the authorizations of an integration.
func (c *CustomIntegrationPermission) GetIntegrationAuthorizations(integrationGuid string) ([]models.IntegrationAuthorization, error) {
	url := fmt.Sprintf("%s/%s/Authorizations", c.baseURL, integrationGuid)
	var authorizations []models.IntegrationAuthorization
	resp, err := c.client.R().
		SetResult(&authorizations).
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to retrieve integration authorizations: %s %s", resp.Status(), resp.Body())
	}
	return authorizations, nil
}

// CreateIntegrationAuthorization creates an authorization for an integration.
// The API may add additional dependent authorizations; the returned value is
// the authorization matching the requested type and operator/operator group.
func (c *CustomIntegrationPermission) CreateIntegrationAuthorization(integrationGuid string, auth models.IntegrationAuthorization) (*models.IntegrationAuthorization, error) {
	url := fmt.Sprintf("%s/%s/Authorizations", c.baseURL, integrationGuid)
	var created []models.IntegrationAuthorization
	resp, err := c.client.R().
		SetBody(auth).
		SetResult(&created).
		Post(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to create integration authorization: %s %s", resp.Status(), resp.Body())
	}

	for i := range created {
		if created[i].AuthorizationType != auth.AuthorizationType {
			continue
		}
		if auth.OperatorGuid != "" && created[i].OperatorGuid == auth.OperatorGuid {
			return &created[i], nil
		}
		if auth.OperatorGroupGuid != "" && created[i].OperatorGroupGuid == auth.OperatorGroupGuid {
			return &created[i], nil
		}
	}

	return nil, fmt.Errorf("created authorization not found in response for integration %q", integrationGuid)
}

// DeleteIntegrationAuthorization deletes an integration authorization.
func (c *CustomIntegrationPermission) DeleteIntegrationAuthorization(integrationGuid, authorizationGuid string) error {
	url := fmt.Sprintf("%s/%s/Authorizations/%s", c.baseURL, integrationGuid, authorizationGuid)
	resp, err := c.client.R().Delete(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("failed to delete integration authorization: %s %s", resp.Status(), resp.Body())
	}
	return nil
}
