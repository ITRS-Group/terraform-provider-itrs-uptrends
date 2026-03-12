package client

type AlertDefinitionOperatorGroupMembershipResponse struct {
	AlertDefinition string `json:"AlertDefinition"`
	Escalationlevel int    `json:"Escalationlevel"`
	OperatorGroup   string `json:"OperatorGroup"`
}
