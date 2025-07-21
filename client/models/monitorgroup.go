package client

type MonitorGroupResponse struct {
	MonitorGroupGuid        string `json:"MonitorGroupGuid"`
	Description             string `json:"Description"`
	IsAll                   bool   `json:"IsAll"`
	IsQuotaUnlimited        *bool  `json:"IsQuotaUnlimited,omitempty"`
	BasicMonitorQuota       *int   `json:"BasicMonitorQuota,omitempty"`
	BrowserMonitorQuota     *int   `json:"BrowserMonitorQuota,omitempty"`
	TransactionMonitorQuota *int   `json:"TransactionMonitorQuota,omitempty"`
	ApiMonitorQuota         *int   `json:"ApiMonitorQuota,omitempty"`
	ClassicQuota            *int   `json:"ClassicQuota,omitempty"`
}

type MonitorGroupRequest struct {
	MonitorGroupGuid        *string `json:"MonitorGroupGuid,omitempty"`
	Description             string  `json:"Description"`
	IsQuotaUnlimited        *bool   `json:"IsQuotaUnlimited,omitempty"`
	BasicMonitorQuota       *int    `json:"BasicMonitorQuota,omitempty"`
	BrowserMonitorQuota     *int    `json:"BrowserMonitorQuota,omitempty"`
	TransactionMonitorQuota *int    `json:"TransactionMonitorQuota,omitempty"`
	ApiMonitorQuota         *int    `json:"ApiMonitorQuota,omitempty"`
}
