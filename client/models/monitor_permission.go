package client

// MonitorAuthorization represents a monitor authorization in the API.
type MonitorAuthorization struct {
	AuthorizationId   string `json:"AuthorizationId,omitempty"`
	AuthorizationType string `json:"AuthorizationType"`
	OperatorGuid      string `json:"OperatorGuid,omitempty"`
	OperatorGroupGuid string `json:"OperatorGroupGuid,omitempty"`
}
