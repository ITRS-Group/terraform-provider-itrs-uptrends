package client

// IntegrationAuthorization represents an authorization (permission) granted on
// a custom integration to an operator or operator group.
type IntegrationAuthorization struct {
	AuthorizationId   string `json:"AuthorizationId,omitempty"`
	AuthorizationType string `json:"AuthorizationType"`
	OperatorGuid      string `json:"OperatorGuid,omitempty"`
	OperatorGroupGuid string `json:"OperatorGroupGuid,omitempty"`
}
