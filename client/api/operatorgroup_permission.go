package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

// OperatorGroupPermission represents a client for interacting with the Uptrends API.
type OperatorGroupPermission struct {
	client  *resty.Client
	baseUrl string
}

// NewOperatorGroupPermission creates and returns a new NewOperatorGroupPermission instance.
// Parameters:
//   - baseUrl: The base URL of the API (e.g., "https://api.uptrends.com/v4/OperatorGroup").
//   - authHeader: The value for the Authorization header.
func NewOperatorGroupPermission(baseUrl, authHeader, version, platform string) *OperatorGroupPermission {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &OperatorGroupPermission{
		client:  client,
		baseUrl: baseUrl,
	}
}

func (uc *OperatorGroupPermission) AssignOperatorGroupPermission(operatorGroupGuid, permission string) error {
	url := fmt.Sprintf("%s/%s/Authorization/%s", uc.baseUrl, operatorGroupGuid, permission)

	resp, err := uc.client.R().Post(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("failed to assign operator: %s %s", resp.Status(), resp.Body())
	}
	return nil
}

func (uc *OperatorGroupPermission) GetOperatorGroupPermission(operatorGroupGuid string) (models.OperatorGroupPermissionResponse, error) {
	url := fmt.Sprintf("%s/%s/Authorization", uc.baseUrl, operatorGroupGuid)
	var permissions models.OperatorGroupPermissionResponse
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

func (uc *OperatorGroupPermission) DeleteOperatorGroupPermission(operatorGroupGuid, permission string) error {
	url := fmt.Sprintf("%s/%s/Authorization/%s", uc.baseUrl, operatorGroupGuid, permission)
	resp, err := uc.client.R().Delete(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("failed to delete permission: %s %s", resp.Status(), resp.Body())
	}
	return nil
}
