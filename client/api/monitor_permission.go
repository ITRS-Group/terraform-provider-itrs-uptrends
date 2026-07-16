package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

var _ interfaces.IMonitorPermission = (*MonitorPermission)(nil)

type MonitorPermission struct {
	client  *resty.Client
	baseURL string
}

func NewMonitorPermission(baseURL, authHeader, version, platform string) *MonitorPermission {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &MonitorPermission{
		client:  client,
		baseURL: baseURL,
	}
}

func (c *MonitorPermission) GetMonitorAuthorizations(monitorGuid string) ([]models.MonitorAuthorization, error) {
	url := fmt.Sprintf("%s/%s/Authorizations", c.baseURL, monitorGuid)
	var authorizations []models.MonitorAuthorization
	resp, err := c.client.R().
		SetResult(&authorizations).
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to retrieve monitor authorizations: %s %s", resp.Status(), resp.Body())
	}
	return authorizations, nil
}

func (c *MonitorPermission) CreateMonitorAuthorization(monitorGuid string, auth models.MonitorAuthorization) (*models.MonitorAuthorization, error) {
	url := fmt.Sprintf("%s/%s/Authorizations", c.baseURL, monitorGuid)
	var created []models.MonitorAuthorization
	resp, err := c.client.R().
		SetBody(auth).
		SetResult(&created).
		Post(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to create monitor authorization: %s %s", resp.Status(), resp.Body())
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

	return nil, fmt.Errorf("created authorization not found in response for monitor %q", monitorGuid)
}

func (c *MonitorPermission) DeleteMonitorAuthorization(monitorGuid, authorizationGuid string) error {
	url := fmt.Sprintf("%s/%s/Authorizations/%s", c.baseURL, monitorGuid, authorizationGuid)
	resp, err := c.client.R().Delete(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("failed to delete monitor authorization: %s %s", resp.Status(), resp.Body())
	}
	return nil
}
