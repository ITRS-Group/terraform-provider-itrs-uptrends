package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

var _ interfaces.IOperatorGroup = (*OperatorGroup)(nil)

// OperatorGroup encapsulates the methods to interact with the API.
type OperatorGroup struct {
	client  *resty.Client
	baseURL string
}

// NewOperatorGroup creates a new API client instance.
func NewOperatorGroup(baseURL, authHeader, version, platform string) *OperatorGroup {
	client := resty.New()
	// Set common headers; Basic Auth header is provided as a string.
	client.SetHeaders(map[string]string{
		"accept":        "application/json",
		"authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &OperatorGroup{
		client:  client,
		baseURL: baseURL,
	}
}

// GetOperatorGroups lists all operator groups.
func (api *OperatorGroup) GetOperatorGroups() ([]models.OperatorGroupResponse, int, string, error) {
	var groups []models.OperatorGroupResponse

	resp, err := api.client.R().
		SetResult(&groups).
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
		return nil, statusCode, responseBody, fmt.Errorf("failed to list operator groups: %s", resp.Status())
	}

	return groups, statusCode, responseBody, nil
}

// CreateOperatorGroup creates a new operator group.
func (api *OperatorGroup) CreateOperatorGroup(description string) (*models.OperatorGroupResponse, error, string) {
	var opGroup models.OperatorGroupResponse
	resp, err := api.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{"Description": description}).
		SetResult(&opGroup).
		Post(api.baseURL)
	if err != nil {
		return nil, err, ""
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("failed to create operator group: %s", resp.Status()), ""
	}
	return &opGroup, nil, resp.String()
}

// GetOperatorGroup retrieves a specific operator group by its ID.
func (api *OperatorGroup) GetOperatorGroup(operatorGroupId string) (*models.OperatorGroupResponse, error, string) {
	var opGroup models.OperatorGroupResponse
	url := fmt.Sprintf("%s/%s", api.baseURL, operatorGroupId)
	resp, err := api.client.R().
		SetResult(&opGroup).
		Get(url)
	if err != nil {
		return nil, err, ""
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("failed to get operator group: %s", resp.Status()), resp.String()
	}
	return &opGroup, nil, resp.String()
}

// UpdateOperatorGroup updates an OperatorGroup with a given description and operatorID
func (a *OperatorGroup) UpdateOperatorGroup(description string, operatorGroupID string) (string, error) {
	payload := map[string]interface{}{
		"OperatorGroupGuid":     operatorGroupID,
		"IsAdministratorsGroup": false,
		"IsEveryone":            false,
		"Description":           description,
	}

	url := fmt.Sprintf("%s/%s", a.baseURL, operatorGroupID)

	// Execute PUT request
	resp, err := a.client.R().
		SetHeader("Accept", "application/json").
		SetBody(payload).
		Put(url)

	if err != nil {
		return "", err
	}

	// Always return an empty string as requested
	return resp.String(), err
}

func (api *OperatorGroup) DeleteOperatorGroup(operatorGroupId string) (error, string) {
	url := fmt.Sprintf("%s/%s", api.baseURL, operatorGroupId)
	resp, err := api.client.R().
		Delete(url)
	if err != nil {
		return err, ""
	}
	return err, resp.String()
}
