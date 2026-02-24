package client

type AccountResponse struct {
	AccountID               string        `json:"AccountId"`
	ExpirationDate          string        `json:"ExpirationDate"`
	MonitorQuota            MonitorQuota  `json:"MonitorQuota"`
	OperatorQuota           OperatorQuota `json:"OperatorQuota"`
	RemainingMessageCredits int           `json:"RemainingMessageCredits"`
}

type MonitorQuota struct {
	BasicMonitors             int `json:"BasicMonitors"`
	BasicMonitorsInUse        int `json:"BasicMonitorsInUse"`
	BrowserMonitors           int `json:"BrowserMonitors"`
	BrowserMonitorsInUse      int `json:"BrowserMonitorsInUse"`
	ApiMonitoringCredits      int `json:"ApiMonitoringCredits"`
	ApiMonitoringCreditsInUse int `json:"ApiMonitoringCreditsInUse"`
	TransactionCredits        int `json:"TransactionCredits"`
	TransactionCreditsInUse   int `json:"TransactionCreditsInUse"`
}

type OperatorQuota struct {
	OperatorsInUse int `json:"OperatorsInUse"`
}
