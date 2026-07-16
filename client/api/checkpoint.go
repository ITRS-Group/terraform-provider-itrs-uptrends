package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

// Checkpoint is the client for the Checkpoint and CheckpointRegion endpoints.
type Checkpoint struct {
	client              *resty.Client
	checkpointURL       string
	checkpointRegionURL string
}

var _ interfaces.ICheckpoint = (*Checkpoint)(nil)

// NewCheckpoint constructs a new checkpoint client.
func NewCheckpoint(checkpointURL, checkpointRegionURL, authHeader, version, platform string) *Checkpoint {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)

	return &Checkpoint{
		client:              client,
		checkpointURL:       checkpointURL,
		checkpointRegionURL: checkpointRegionURL,
	}
}

// GetCheckpoints returns all checkpoints.
func (c *Checkpoint) GetCheckpoints() (models.CheckpointResponse, int, string, error) {
	var result models.CheckpointResponse

	resp, err := c.client.R().
		SetResult(&result).
		Get(c.checkpointURL)

	statusCode := -1
	responseBody := ""
	if resp != nil {
		statusCode = resp.StatusCode()
		responseBody = resp.String()
	}

	if err != nil {
		return result, statusCode, responseBody, err
	}
	if !resp.IsSuccess() {
		return result, statusCode, responseBody, fmt.Errorf("failed to list checkpoints: %s", resp.Status())
	}

	return result, statusCode, responseBody, nil
}

// GetCheckpointRegions returns all checkpoint regions.
func (c *Checkpoint) GetCheckpointRegions() ([]models.CheckpointRegionResponse, int, string, error) {
	var result []models.CheckpointRegionResponse

	resp, err := c.client.R().
		SetResult(&result).
		Get(c.checkpointRegionURL)

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
		return nil, statusCode, responseBody, fmt.Errorf("failed to list checkpoint regions: %s", resp.Status())
	}

	return result, statusCode, responseBody, nil
}
