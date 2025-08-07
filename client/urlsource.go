package client

// UrlSource holds the configurable base URL for the API.
type UrlSource struct {
	baseURL string
}

// NewUrlSource instantiates a new APIClient with the given base URL.
func NewUrlSource(baseURL string) *UrlSource {
	return &UrlSource{baseURL: baseURL}
}

// AccountURL returns the full URL for the Account endpoint.
func (c *UrlSource) AccountURL() string {
	return c.baseURL + "/Account"
}

// MonitorURL returns the full URL for the Monitor endpoint.
func (c *UrlSource) MonitorURL() string {
	return c.baseURL + "/Monitor"
}

// MonitorGroupURL returns the full URL for the MonitorGroup endpoint.
func (c *UrlSource) MonitorGroupURL() string {
	return c.baseURL + "/MonitorGroup"
}

// AlertDefinitionURL returns the full URL for the AlertDefinition endpoint.
func (c *UrlSource) AlertDefinitionURL() string {
	return c.baseURL + "/AlertDefinition"
}

// OperatorURL returns the full URL for the Operator endpoint.
func (c *UrlSource) OperatorURL() string {
	return c.baseURL + "/Operator"
}

// OperatorGroupURL returns the full URL for the OperatorGroup endpoint.
func (c *UrlSource) OperatorGroupURL() string {
	return c.baseURL + "/OperatorGroup"
}

// VaultItemURL returns the full URL for the OperatorGroup endpoint.
func (c *UrlSource) VaultItemURL() string {
	return c.baseURL + "/VaultItem"
}

// VaultItemURL returns the full URL for the OperatorGroup endpoint.
func (c *UrlSource) VaultSectionURL() string {
	return c.baseURL + "/VaultSection"
}
