package client

// AlertDefinitionOperatorMembershipResponse represents the response structure returned from the POST call.
type AlertDefinitionOperatorMembershipResponse struct {
	AlertDefinition string `json:"AlertDefinition"`
	Escalationlevel int    `json:"Escalationlevel"`
	Operator        string `json:"Operator"`
}

// GetMembershipResponse represents an element in the slice returned from the GET call.
type GetMembershipResponse struct {
	OperatorGuid      string `json:"OperatorGuid"`
	OperatorGroupGuid string `json:"OperatorGroupGuid"`
}
