package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

var _ interfaces.ICustomIntegration = (*CustomIntegration)(nil)

// CustomIntegration encapsulates the methods to interact with the
// /Integration/Custom endpoints of the API.
type CustomIntegration struct {
	client  *resty.Client
	baseURL string
}

// NewCustomIntegration creates a new API client instance. baseURL is expected
// to be the /Integration endpoint.
func NewCustomIntegration(baseURL, authHeader, version, platform string) *CustomIntegration {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &CustomIntegration{
		client:  client,
		baseURL: baseURL,
	}
}

func (c *CustomIntegration) customURL() string {
	return fmt.Sprintf("%s/Custom", c.baseURL)
}

// GetCustomIntegrations lists all custom integrations.
func (c *CustomIntegration) GetCustomIntegrations() ([]models.CustomIntegration, error) {
	var integrations []models.CustomIntegration
	resp, err := c.client.R().
		SetResult(&integrations).
		Get(c.customURL())
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to list custom integrations: %s %s", resp.Status(), resp.Body())
	}
	return integrations, nil
}

// CreateCustomIntegration creates a new custom integration.
func (c *CustomIntegration) CreateCustomIntegration(integration *models.CustomIntegration) (*models.CustomIntegration, error) {
	var created models.CustomIntegration
	resp, err := c.client.R().
		SetBody(integration).
		SetResult(&created).
		Post(c.customURL())
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to create custom integration: %s %s", resp.Status(), resp.Body())
	}
	return &created, nil
}

// GetCustomIntegration returns the details of the specified custom integration.
func (c *CustomIntegration) GetCustomIntegration(integrationGuid string) (*models.CustomIntegration, error) {
	var integration models.CustomIntegration
	url := fmt.Sprintf("%s/%s", c.customURL(), integrationGuid)
	resp, err := c.client.R().
		SetResult(&integration).
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to get custom integration: %s %s", resp.Status(), resp.Body())
	}
	return &integration, nil
}

// UpdateCustomIntegration updates the specified custom integration.
func (c *CustomIntegration) UpdateCustomIntegration(integrationGuid string, integration *models.CustomIntegration) (*models.CustomIntegration, error) {
	url := fmt.Sprintf("%s/%s", c.customURL(), integrationGuid)
	resp, err := c.client.R().
		SetBody(integration).
		Put(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to update custom integration: %s %s", resp.Status(), resp.Body())
	}
	return c.GetCustomIntegration(integrationGuid)
}

// DeleteCustomIntegration deletes the specified custom integration.
func (c *CustomIntegration) DeleteCustomIntegration(integrationGuid string) error {
	url := fmt.Sprintf("%s/%s", c.customURL(), integrationGuid)
	resp, err := c.client.R().Delete(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("failed to delete custom integration: %s %s", resp.Status(), resp.Body())
	}
	return nil
}
