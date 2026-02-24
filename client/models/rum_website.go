package client

// RumWebsiteResponse represents a rum website.
type RumWebsite struct {
	RumWebsiteGuid     string `json:"RumWebsiteId"`
	Description        string `json:"Description"`
	Url                string `json:"Url"`
	IsSpa              bool   `json:"IsSpa"`
	IncludeUrlFragment bool   `json:"IncludeUrlFragment"`
	RumScript          string `json:"RumScript"`
}
