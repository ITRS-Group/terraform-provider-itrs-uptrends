package client

type AlertDefinitionRequest struct {
	AlertName string `json:"AlertName"`
	IsActive  bool   `json:"IsActive"`
}

type AlertDefinitionResponse struct {
	AlertDefinitionGuid string `json:"AlertDefinitionGuid"`
	AlertName           string `json:"AlertName"`
	IsActive            bool   `json:"IsActive"`
}

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
	AlertDefinitionGuid string `json:"AlertDefinitionGuid"`
}
