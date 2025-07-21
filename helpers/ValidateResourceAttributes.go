package helpers

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ValidateResourceAttributes validates that:
//  1. All attributes in requiredAttrs are provided in the config.
//  2. All provided attributes (via the tfsdk tag) are within allowedAttrs.
//
// resourceName: e.g. "vault_item"
// resourceType: e.g. "CredentialSet"
// config: the Terraform configuration struct.
// requiredAttrs: attributes that must be provided.
// allowedAttrs: attributes (both required and optional) that are allowed.
func ValidateResourceAttributes(resourceName, resourceType string, config interface{}, requiredAttrs, allowedAttrs []string) error {
	providedAttrs := buildProvidedAttributes(config)

	// Check for missing required attributes.
	providedSet := make(map[string]bool, len(providedAttrs))
	for _, a := range providedAttrs {
		providedSet[a] = true
	}
	var missing []string
	for _, req := range requiredAttrs {
		if !providedSet[req] {
			missing = append(missing, req)
		}
	}

	// Check for unexpected attributes (provided that aren't allowed).
	allowedMap := make(map[string]bool, len(allowedAttrs))
	for _, a := range allowedAttrs {
		allowedMap[a] = true
	}
	var unexpected []string
	for _, attr := range providedAttrs {
		if !allowedMap[attr] {
			unexpected = append(unexpected, attr)
		}
	}

	// If any errors, build a combined message.
	if len(missing) > 0 || len(unexpected) > 0 {
		sort.Strings(missing)
		sort.Strings(unexpected)
		sort.Strings(allowedAttrs)
		errMsg := fmt.Sprintf("For resource '%s' (type '%s'):", resourceName, resourceType)
		if len(missing) > 0 {
			errMsg += fmt.Sprintf("\nRequired attributes missing:\n\t%v", missing)
		}
		if len(unexpected) > 0 {
			errMsg += fmt.Sprintf("\nAccepted attributes are:\n\t%v\nReceived unexpected attributes:\n\t%v", allowedAttrs, unexpected)
		}
		return fmt.Errorf(errMsg)
	}
	return nil
}

// buildProvidedAttributes uses reflection to iterate over the config struct and collects the attribute names (from the tfsdk tag)
// that are non-null (for types.String or pointer fields).
func buildProvidedAttributes(config interface{}) []string {
	var providedAttrs []string
	v := reflect.ValueOf(config)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		tag := fieldType.Tag.Get("tfsdk")
		if tag == "" {
			continue
		}
		// For types.String:
		if field.Type() == reflect.TypeOf(types.String{}) {
			ts := field.Interface().(types.String)
			if !ts.IsNull() && ts.ValueString() != "" {
				providedAttrs = append(providedAttrs, tag)
			}
		}
		// For pointer types (e.g. nested blocks), check if non-nil.
		if field.Kind() == reflect.Ptr && !field.IsNil() {
			providedAttrs = append(providedAttrs, tag)
		}
		// ... handle additional tfsdk types if needed.
	}
	return providedAttrs
}
