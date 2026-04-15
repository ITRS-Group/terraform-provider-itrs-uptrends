package client

import (
	"fmt"
	"net/http"
)

// userAgentRoundTripper injects the custom User-Agent header.
type userAgentRoundTripper struct {
	rt       http.RoundTripper
	version  string
	platform string
}

// RoundTrip sets the User-Agent header on each request.
func (u *userAgentRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	ua := fmt.Sprintf("itrs-uptrends-terraform-provider/%s (%s)", u.version, u.platform)
	req.Header.Set("User-Agent", ua)
	return u.rt.RoundTrip(req)
}

// NewHTTPClient creates an HTTP client that automatically includes the custom header.
func NewHTTPClient(version, platform string) *http.Client {
	return &http.Client{
		Transport: &userAgentRoundTripper{
			rt:       http.DefaultTransport,
			version:  version,
			platform: platform,
		},
	}
}
