package client

// AssignResponse represents the response returned after assigning a monitor.
type AssignResponse struct {
	AlertDefinition string `json:"AlertDefinition"`
	Monitor         string `json:"Monitor"`
}

// Assignment represents a monitor assignment, which may be an individual monitor or a monitor group.
// Using pointers allows fields to be nil if they are not present.
type Assignment struct {
	MonitorGuid      *string `json:"MonitorGuid,omitempty"`
	MonitorGroupGuid *string `json:"MonitorGroupGuid,omitempty"`
}
