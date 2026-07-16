package client

// MonitorGroupAuthorization represents a monitor group authorization in the API.
type MonitorGroupAuthorization struct {
	AuthorizationId   string `json:"AuthorizationId,omitempty"`
	AuthorizationType string `json:"AuthorizationType"`
	OperatorGuid      string `json:"OperatorGuid,omitempty"`
	OperatorGroupGuid string `json:"OperatorGroupGuid,omitempty"`
}
