package client

type EscalationLevelIntegrationResponse struct {
	IntegrationGuid      string                  `json:"IntegrationGuid"`
	Name                 string                  `json:"Name"`
	Type                 string                  `json:"Type"`
	VariableValues       map[string]string       `json:"VariableValues,omitempty"`
	ExtraEmailAddresses  string                  `json:"ExtraEmailAddresses,omitempty"`
	StatusHubServiceList []StatusHubServiceEntry `json:"StatusHubServiceList,omitempty"`
	IntegrationServices  []string                `json:"IntegrationServices,omitempty"`
	Hash                 string                  `json:"Hash,omitempty"`
}

type StatusHubServiceEntry struct {
	MonitorGuid            string `json:"MonitorGuid"`
	IntegrationServiceGuid string `json:"IntegrationServiceGuid"`
}

type EscalationLevelIntegrationRequest struct {
	IntegrationGuid               string                  `json:"IntegrationGuid"`
	ExtraEmailAddresses           []string                `json:"ExtraEmailAddresses,omitempty"`
	ExtraEmailAddressesSpecified  *bool                   `json:"ExtraEmailAddressesSpecified,omitempty"`
	StatusHubServiceList          []StatusHubServiceEntry `json:"StatusHubServiceList,omitempty"`
	StatusHubServiceListSpecified *bool                   `json:"StatusHubServiceListSpecified,omitempty"`
	IsActive                      *bool                   `json:"IsActive,omitempty"`
	SendOkAlerts                  *bool                   `json:"SendOkAlerts,omitempty"`
	SendReminderAlerts            *bool                   `json:"SendReminderAlerts,omitempty"`
	VariableValues                map[string]string       `json:"VariableValues,omitempty"`
}
