package helpers

import (
	"fmt"
	"sort"
	"strings"
)

// ValidateAllowedAttributes validates that all provided attributes (via the tfsdk tag) are within allowedAttributes	.
//
// resourceName: e.g. "vault_item"
// resourceType: e.g. "CredentialSet"
// config: the Terraform configuration struct.
// allowedAttributes: attributes (both required and optional) that are allowed.
func ValidateAllowedAttributes(resourceName, resourceType string, config interface{}, allowedAttributes []string) error {
	providedAttrs := buildProvidedAttributes(config)

	providedSet := make(map[string]bool, len(providedAttrs))
	for _, a := range providedAttrs {
		providedSet[a] = true
	}

	// Check for unexpected attributes (provided that aren't allowed).
	allowedMap := make(map[string]bool, len(allowedAttributes))
	for _, a := range allowedAttributes {
		allowedMap[a] = true
	}

	var unexpected []string
	for _, attr := range providedAttrs {
		if !allowedMap[attr] {
			unexpected = append(unexpected, attr)
		}
	}

	// If any errors, build a combined message.
	if len(unexpected) > 0 {
		sort.Strings(unexpected)
		sort.Strings(allowedAttributes)
		errMsg := fmt.Sprintf(
			"\n--------------------------------------------------\n"+
				"Resource: %s\nType:     %s\n"+
				"--------------------------------------------------\n"+
				"❌ The following attributes are not allowed:\n"+
				"    - %s\n"+
				"\n✅ Accepted attributes for this resource type are:\n"+
				"    - %s\n"+
				"--------------------------------------------------",
			resourceName, resourceType,
			strings.Join(unexpected, "\n    - "),
			strings.Join(allowedAttributes, "\n    - "),
		)
		return fmt.Errorf("%s", errMsg)
	}
	return nil
}
