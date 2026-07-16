package client

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

var _ interfaces.IMonitorGroupPermission = (*MonitorGroupPermission)(nil)

type MonitorGroupPermission struct {
	client  *resty.Client
	baseURL string
}

func NewMonitorGroupPermission(baseURL, authHeader, version, platform string) *MonitorGroupPermission {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &MonitorGroupPermission{
		client:  client,
		baseURL: baseURL,
	}
}

func (c *MonitorGroupPermission) GetMonitorGroupAuthorizations(monitorGroupGuid string) ([]models.MonitorGroupAuthorization, error) {
	url := fmt.Sprintf("%s/%s/Authorizations", c.baseURL, monitorGroupGuid)
	var authorizations []models.MonitorGroupAuthorization
	resp, err := c.client.R().
		SetResult(&authorizations).
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to retrieve monitor group authorizations: %s %s", resp.Status(), resp.Body())
	}
	return authorizations, nil
}

func (c *MonitorGroupPermission) CreateMonitorGroupAuthorization(monitorGroupGuid string, auth models.MonitorGroupAuthorization) (*models.MonitorGroupAuthorization, error) {
	url := fmt.Sprintf("%s/%s/Authorizations", c.baseURL, monitorGroupGuid)
	resp, err := c.client.R().
		SetBody(auth).
		Post(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to create monitor group authorization: %s %s", resp.Status(), resp.Body())
	}

	created, err := decodeMonitorGroupAuthorizationResponse(resp.Body(), auth)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (c *MonitorGroupPermission) DeleteMonitorGroupAuthorization(monitorGroupGuid, authorizationGuid string) error {
	url := fmt.Sprintf("%s/%s/Authorizations/%s", c.baseURL, monitorGroupGuid, authorizationGuid)
	resp, err := c.client.R().Delete(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("failed to delete monitor group authorization: %s %s", resp.Status(), resp.Body())
	}
	return nil
}

func decodeMonitorGroupAuthorizationResponse(body []byte, requested models.MonitorGroupAuthorization) (*models.MonitorGroupAuthorization, error) {
	var created []models.MonitorGroupAuthorization
	if err := json.Unmarshal(body, &created); err == nil {
		for _, auth := range created {
			if auth.AuthorizationType == requested.AuthorizationType &&
				auth.OperatorGuid == requested.OperatorGuid &&
				auth.OperatorGroupGuid == requested.OperatorGroupGuid {
				return &auth, nil
			}
		}
		if len(created) > 0 {
			return &created[0], nil
		}
		return nil, fmt.Errorf("monitor group authorization response did not include created authorizations")
	}

	var single models.MonitorGroupAuthorization
	if err := json.Unmarshal(body, &single); err != nil {
		return nil, fmt.Errorf("failed to decode monitor group authorization response: %w", err)
	}
	return &single, nil
}
