package helpers

import (
	"fmt"
	"sort"
	"strings"
)

// ValidateRequiredAttributes validates that all attributes in requiredAttributes are provided in the config.
//
// resourceName: e.g. "vault_item"
// resourceType: e.g. "CredentialSet"
// config: the Terraform configuration struct.
// requiredAttributes: attributes that must be provided.
func ValidateRequiredAttributes(resourceName, resourceType string, config interface{}, requiredAttributes []string) error {
	providedAttrs := buildProvidedAttributes(config)

	// Check for missing required attributes.
	providedSet := make(map[string]bool, len(providedAttrs))
	for _, a := range providedAttrs {
		providedSet[a] = true
	}
	var missing []string
	for _, req := range requiredAttributes {
		if !providedSet[req] {
			missing = append(missing, req)
		}
	}

	// If any errors, build a combined message.
	if len(missing) > 0 {
		sort.Strings(missing)
		errMsg := fmt.Sprintf(
			"\n--------------------------------------------------\n"+
				"Resource: %s\nType:     %s\n"+
				"--------------------------------------------------\n"+
				"❌ The following attributes are required but missing:\n"+
				"    - %s\n"+
				"--------------------------------------------------",
			resourceName, resourceType,
			strings.Join(missing, "\n    - "),
		)
		return fmt.Errorf("%s", errMsg)
	}
	return nil
}
