package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

// OperatorPermission represents a client for interacting with the Uptrends API.
type OperatorPermission struct {
	client  *resty.Client
	baseUrl string
}

// NewOperatorPermission creates and returns a new NewOperatorPermission instance.
// Parameters:
//   - baseUrl: The base URL of the API (e.g., "https://api.uptrends.com/v4/Operator").
//   - authHeader: The value for the Authorization header.
func NewOperatorPermission(baseUrl, authHeader, version, platform string) *OperatorPermission {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &OperatorPermission{
		client:  client,
		baseUrl: baseUrl,
	}
}

func (uc *OperatorPermission) AssignOperatorPermission(operatorGuid, permission string) error {
	url := fmt.Sprintf("%s/%s/Authorization/%s", uc.baseUrl, operatorGuid, permission)

	resp, err := uc.client.R().Post(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("failed to assign operator: %s %s", resp.Status(), resp.Body())
	}
	return nil
}

func (uc *OperatorPermission) GetOperatorPermission(operatorGuid string) (models.OperatorPermissionResponse, error) {
	url := fmt.Sprintf("%s/%s/Authorization", uc.baseUrl, operatorGuid)
	var permissions models.OperatorPermissionResponse
	resp, err := uc.client.R().
		SetResult(&permissions).
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to retrieve permissions: %s %s", resp.Status(), resp.Body())
	}

	return permissions, nil
}

func (uc *OperatorPermission) DeleteOperatorPermission(operatorGuid, permission string) error {
	url := fmt.Sprintf("%s/%s/Authorization/%s", uc.baseUrl, operatorGuid, permission)
	resp, err := uc.client.R().Delete(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("failed to delete permission: %s %s", resp.Status(), resp.Body())
	}
	return nil
}
