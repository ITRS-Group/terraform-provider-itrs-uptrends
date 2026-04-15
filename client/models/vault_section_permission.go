package client

// VaultSectionAuthorization represents a vault section authorization in the API.
type VaultSectionAuthorization struct {
	AuthorizationId    string `json:"AuthorizationId,omitempty"`
	AuthorizationType  string `json:"AuthorizationType"`
	OperatorGuid       string `json:"OperatorGuid,omitempty"`
	OperatorGroupGuid  string `json:"OperatorGroupGuid,omitempty"`
}
