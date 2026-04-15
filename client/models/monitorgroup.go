package client

type MonitorGroupResponse struct {
	MonitorGroupGuid            string `json:"MonitorGroupGuid"`
	Description                 string `json:"Description"`
	IsAll                       bool   `json:"IsAll"`
	IsQuotaUnlimited            *bool  `json:"IsQuotaUnlimited,omitempty"`
	BasicMonitorQuota           *int   `json:"BasicMonitorQuota,omitempty"`
	UsedBasicMonitorQuota       *int   `json:"UsedBasicMonitorQuota,omitempty"`
	BrowserMonitorQuota         *int   `json:"BrowserMonitorQuota,omitempty"`
	UsedBrowserMonitorQuota     *int   `json:"UsedBrowserMonitorQuota,omitempty"`
	TransactionMonitorQuota     *int   `json:"TransactionMonitorQuota,omitempty"`
	UsedTransactionMonitorQuota *int   `json:"UsedTransactionMonitorQuota,omitempty"`
	ApiMonitorQuota             *int   `json:"ApiMonitorQuota,omitempty"`
	UsedApiMonitorQuota         *int   `json:"UsedApiMonitorQuota,omitempty"`
	UnifiedCreditsQuota         *int   `json:"UnifiedCreditsQuota,omitempty"`
	UsedUnifiedCreditsQuota     *int   `json:"UsedUnifiedCreditsQuota,omitempty"`
	ClassicQuota                *int   `json:"ClassicQuota,omitempty"`
	UsedClassicQuota            *int   `json:"UsedClassicQuota,omitempty"`
}

type MonitorGroupRequest struct {
	MonitorGroupGuid        *string `json:"MonitorGroupGuid,omitempty"`
	Description             string  `json:"Description"`
	IsQuotaUnlimited        *bool   `json:"IsQuotaUnlimited,omitempty"`
	BasicMonitorQuota       int     `json:"BasicMonitorQuota,omitempty"`
	BrowserMonitorQuota     int     `json:"BrowserMonitorQuota,omitempty"`
	TransactionMonitorQuota int     `json:"TransactionMonitorQuota,omitempty"`
	ApiMonitorQuota         int     `json:"ApiMonitorQuota,omitempty"`
	ClassicQuota            int     `json:"ClassicQuota,omitempty"`
	UnifiedCreditsQuota     int     `json:"UnifiedCreditsQuota,omitempty"`
}
