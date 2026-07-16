package client

import "encoding/base64"

func GenerateBasicAuthHeader(username string, password string) string {
	// Combine username and password with a colon
	credentials := username + ":" + password

	// Encode the credentials to Base64
	encoded := base64.StdEncoding.EncodeToString([]byte(credentials))

	// Return the Authorization header
	return "Basic " + encoded
}
