package client

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

// MonitorGroupClient is used to interact with the monitor group API.
type MonitorGroupClient struct {
	client     *resty.Client
	baseURL    string
	authHeader string
}

// NewMonitorGroupClient creates a new APIClient instance with the provided baseURL and authHeader.
// The baseURL should be the full endpoint URL, e.g. "https://api.uptrends.com/v4/MonitorGroup".
func NewMonitorGroupClient(baseURL, authHeader, version, platform string) *MonitorGroupClient {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Authorization": authHeader,
	})
	client.SetBaseURL(baseURL)
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &MonitorGroupClient{
		client:     client,
		baseURL:    baseURL,
		authHeader: authHeader,
	}
}

func (c *MonitorGroupClient) GetMonitorGroup(monitorGroupGuid string) (models.MonitorGroupResponse, string, error) {
	var result models.MonitorGroupResponse
	url := fmt.Sprintf("%s/%s", c.baseURL, monitorGroupGuid)

	resp, err := c.client.R().
		SetResult(&result).
		Get(url)

	if err != nil {
		return models.MonitorGroupResponse{}, "", err
	}
	if !resp.IsSuccess() {
		return models.MonitorGroupResponse{}, resp.String(), fmt.Errorf("failed to get monitor group: %s", resp.Status())
	}

	return result, resp.String(), nil
}

func (c *MonitorGroupClient) CreateMonitorGroup(payload models.MonitorGroupRequest) (models.MonitorGroupResponse, int, string, error) {
	var result models.MonitorGroupResponse
	marshalRequestData, err := json.Marshal(payload)

	if err != nil {
		return models.MonitorGroupResponse{}, 0, "", fmt.Errorf("failed to marshal request data: %v", err)
	}

	resp, err := c.client.R().
		SetBody(marshalRequestData).
		SetResult(&result).
		Post(c.baseURL)

	if err != nil {
		return models.MonitorGroupResponse{}, 0, resp.String(), fmt.Errorf("failed to execute HTTP request: %v", err)
	}
	return result, resp.StatusCode(), resp.String(), nil
}

func (c *MonitorGroupClient) UpdateMonitorGroup(payload models.MonitorGroupRequest, monitorGroupGuid string) (int, string, error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, monitorGroupGuid)
	marshalRequestData, err := json.Marshal(payload)

	if err != nil {
		return 0, "", fmt.Errorf("failed to marshal request data: %v", err)
	}

	resp, err := c.client.R().
		SetBody(marshalRequestData).
		Put(url)

	if err != nil {
		return 0, "", err
	}

	return resp.StatusCode(), resp.String(), nil
}

func (c *MonitorGroupClient) DeleteMonitorGroup(monitorGroupGuid string) (int, string, error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, monitorGroupGuid)

	resp, err := c.client.R().
		Delete(url)

	if err != nil {
		return 0, "", err
	}

	return resp.StatusCode(), resp.String(), nil
}
