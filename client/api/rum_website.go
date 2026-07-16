package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

var _ interfaces.IRumWebsite = (*RumWebsite)(nil)

// RumWebsite encapsulates the methods to interact with the API.
type RumWebsite struct {
	client  *resty.Client
	baseURL string
}

// NewRumWebsite creates a new API client instance.
func NewRumWebsite(baseURL, authHeader, version, platform string) *RumWebsite {
	client := resty.New()
	// Set common headers; Basic Auth header is provided as a string.
	client.SetHeaders(map[string]string{
		"accept":        "application/json",
		"authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &RumWebsite{
		client:  client,
		baseURL: baseURL,
	}
}

// GetRumWebsites lists all rum websites.
func (api *RumWebsite) GetRumWebsites() ([]models.RumWebsite, int, string, error) {
	var rumWebsites []models.RumWebsite

	resp, err := api.client.R().
		SetResult(&rumWebsites).
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
		return nil, statusCode, responseBody, fmt.Errorf("failed to list rum websites: %s", resp.Status())
	}

	return rumWebsites, statusCode, responseBody, nil
}

// CreateRumWebsite creates a new rum website.
func (api *RumWebsite) CreateRumWebsite(request *models.RumWebsite) (*models.RumWebsite, string, error) {
	var rumWebsite models.RumWebsite
	resp, err := api.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&rumWebsite).
		Post(api.baseURL)
	if err != nil {
		return nil, "", err
	}
	if !resp.IsSuccess() {
		return nil, "", fmt.Errorf("failed to create rum website: %s", resp.Status())
	}
	return &rumWebsite, resp.String(), nil
}

// GetRumWebsite retrieves a specific rum website by its ID.
func (api *RumWebsite) GetRumWebsite(rumWebsiteId string) (*models.RumWebsite, string, error) {
	var rumWebsite models.RumWebsite
	url := fmt.Sprintf("%s/%s", api.baseURL, rumWebsiteId)
	resp, err := api.client.R().
		SetResult(&rumWebsite).
		Get(url)
	if err != nil {
		return nil, "", err
	}
	if !resp.IsSuccess() {
		return nil, "", fmt.Errorf("failed to get rum website: %s", resp.Status())
	}
	return &rumWebsite, resp.String(), nil
}

// UpdateRumWebsite updates an RumWebsite with a given description and rumWebsiteID
func (a *RumWebsite) UpdateRumWebsite(request *models.RumWebsite) (string, error) {
	payload := map[string]interface{}{
		"RumWebsiteGuid":     request.RumWebsiteGuid,
		"Description":        request.Description,
		"Url":                request.Url,
		"IsSpa":              request.IsSpa,
		"IncludeUrlFragment": request.IncludeUrlFragment,
	}

	url := fmt.Sprintf("%s/%s", a.baseURL, request.RumWebsiteGuid)

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

func (api *RumWebsite) DeleteRumWebsite(rumWebsiteId string) (string, error) {
	url := fmt.Sprintf("%s/%s", api.baseURL, rumWebsiteId)
	resp, err := api.client.R().
		Delete(url)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}
