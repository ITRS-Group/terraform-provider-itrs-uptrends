package client

// OperatorGroupResponse represents an operator group.
type OperatorGroupResponse struct {
	OperatorGroupGuid     string `json:"OperatorGroupGuid"`
	Description           string `json:"Description"`
	IsEveryone            bool   `json:"IsEveryone,omitempty"`
	IsAdministratorsGroup bool   `json:"IsAdministratorsGroup,omitempty"`
}
