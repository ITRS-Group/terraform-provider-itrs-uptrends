package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

// ICheckpoint defines the operations for retrieving checkpoints and regions.
type ICheckpoint interface {
	GetCheckpoints() (models.CheckpointResponse, int, string, error)
	GetCheckpointRegions() ([]models.CheckpointRegionResponse, int, string, error)
}
