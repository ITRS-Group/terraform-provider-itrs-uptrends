package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type EscalationLevelIntegration struct {
	baseURL    string
	authHeader string
	client     *resty.Client
}

var _ interfaces.IEscalationLevelIntegration = (*EscalationLevelIntegration)(nil)

func NewEscalationLevelIntegration(baseURL, authHeader, version, platform string) *EscalationLevelIntegration {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &EscalationLevelIntegration{
		baseURL:    baseURL,
		authHeader: authHeader,
		client:     client,
	}
}

func (c *EscalationLevelIntegration) integrationURL(alertDefinitionGuid string, escalationLevelId int) string {
	return fmt.Sprintf("%s/%s/EscalationLevel/%d/Integration", c.baseURL, alertDefinitionGuid, escalationLevelId)
}

func (c *EscalationLevelIntegration) GetIntegration(alertDefinitionGuid string, escalationLevelId int, integrationGuid string) (*models.EscalationLevelIntegrationResponse, error) {
	var result models.EscalationLevelIntegrationResponse
	url := fmt.Sprintf("%s/%s", c.integrationURL(alertDefinitionGuid, escalationLevelId), integrationGuid)

	resp, err := c.client.R().
		SetResult(&result).
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error fetching integration: %s %s", resp.Status(), resp.Body())
	}
	return &result, nil
}

func (c *EscalationLevelIntegration) AddIntegration(alertDefinitionGuid string, escalationLevelId int, payload models.EscalationLevelIntegrationRequest) (*models.EscalationLevelIntegrationResponse, error) {
	url := c.integrationURL(alertDefinitionGuid, escalationLevelId)

	resp, err := c.client.R().
		SetBody(payload).
		SetResult(&models.EscalationLevelIntegrationResponse{}).
		Post(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error adding integration: %s %s", resp.Status(), resp.Body())
	}
	return resp.Result().(*models.EscalationLevelIntegrationResponse), nil
}

func (c *EscalationLevelIntegration) UpdateIntegration(alertDefinitionGuid string, escalationLevelId int, integrationGuid string, payload models.EscalationLevelIntegrationRequest) error {
	url := fmt.Sprintf("%s/%s", c.integrationURL(alertDefinitionGuid, escalationLevelId), integrationGuid)

	resp, err := c.client.R().
		SetBody(payload).
		Patch(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("error updating integration: %s %s", resp.Status(), resp.Body())
	}
	return nil
}

func (c *EscalationLevelIntegration) RemoveIntegration(alertDefinitionGuid string, escalationLevelId int, integrationGuid string) error {
	url := fmt.Sprintf("%s/%s", c.integrationURL(alertDefinitionGuid, escalationLevelId), integrationGuid)

	resp, err := c.client.R().
		Delete(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("error removing integration: %s %s", resp.Status(), resp.Body())
	}
	return nil
}
