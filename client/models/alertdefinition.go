package client

// AlertDefinitionItem represents the alert definition object.
type AlertDefinitionItem struct {
	AlertDefinitionGuid string `json:"AlertDefinitionGuid"`
	Hash                string `json:"Hash"`
	AlertName           string `json:"AlertName"`
	IsActive            bool   `json:"IsActive"`
}

// EscalationLevel represents an escalation level object.
type EscalationLevel struct {
	EscalationMode      string `json:"EscalationMode"`
	Id                  int    `json:"Id"`
	ThresholdErrorCount int    `json:"ThresholdErrorCount"`
	ThresholdMinutes    int    `json:"ThresholdMinutes"`
	IsActive            bool   `json:"IsActive"`
	Message             string `json:"Message"`
	NumberOfReminders   int    `json:"NumberOfReminders"`
	ReminderDelay       int    `json:"ReminderDelay"`
	IncludeTraceRoute   bool   `json:"IncludeTraceRoute"`
	Hash                string `json:"Hash"`
}

// CreateAlertDefinitionRequest is the payload for creating an alert definition.
type CreateAlertDefinitionRequest struct {
	AlertName string `json:"AlertName"`
	IsActive  bool   `json:"IsActive"`
}

// UpdateAlertDefinitionRequest is the payload for updating an alert definition.
type UpdateAlertDefinitionRequest struct {
	AlertName string `json:"AlertName,omitempty"`
	IsActive  bool   `json:"IsActive,omitempty"`
}

// PatchEscalationLevelRequest is the payload for updating an escalation level.
type PatchEscalationLevelRequest struct {
	ThresholdErrorCount int    `json:"ThresholdErrorCount,omitempty"`
	IsActive            bool   `json:"IsActive,omitempty"`
	Message             string `json:"Message,omitempty"`
	NumberOfReminders   int    `json:"NumberOfReminders,omitempty"`
	ReminderDelay       int    `json:"ReminderDelay,omitempty"`
	IncludeTraceRoute   bool   `json:"IncludeTraceRoute,omitempty"`
}
