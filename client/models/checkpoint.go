package client

// CheckpointResponse represents the response from GET /Checkpoint.
type CheckpointResponse struct {
	Data []Checkpoint `json:"Data"`
}

// Checkpoint is an individual checkpoint entry.
type Checkpoint struct {
	Id         int                  `json:"Id"`
	Type       string               `json:"Type"`
	Attributes CheckpointAttributes `json:"Attributes"`
	Links      *CheckpointLinks     `json:"Links,omitempty"`
}

// CheckpointAttributes contains details for a checkpoint.
type CheckpointAttributes struct {
	CheckpointName      string   `json:"CheckpointName"`
	Code                string   `json:"Code"`
	Ipv4Addresses       []string `json:"Ipv4Addresses"`
	IpV6Addresses       []string `json:"IpV6Addresses"`
	IsPrimaryCheckpoint bool     `json:"IsPrimaryCheckpoint"`
	SupportsIpv6        bool     `json:"SupportsIpv6"`
	HasHighAvailability bool     `json:"HasHighAvailability"`
}

// CheckpointLinks represents hyperlink references for the checkpoint.
type CheckpointLinks struct {
	Self string `json:"Self"`
}

// CheckpointRegionResponse represents a region returned by GET /CheckpointRegion.
type CheckpointRegionResponse struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}
