package client

// AssignResponse represents the response returned after assigning a monitor.
type AlertDefinitionMonitorGroupMembershipResponse struct {
	AlertDefinition string `json:"AlertDefinition"`
	MonitorGroup    string `json:"MonitorGroup"`
}

// GetMembershipResponse represents an element in the slice returned from the GET call.
type GetMonitorGroupMembershipResponse struct {
	MonitorGroupGuid *string `json:"MonitorGroupGuid,omitempty"`
}
