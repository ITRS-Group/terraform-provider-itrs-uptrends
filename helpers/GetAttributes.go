package helpers

type ResourceAttributes struct {
	RequiredAttributes []string
	OptionalAttributes []string
}

func GetRequiredAttributes(resourceType string, resourceAttributes map[string]ResourceAttributes) []string {
	attr, ok := resourceAttributes[resourceType]
	if !ok {
		return []string{}
	}
	return attr.RequiredAttributes
}

func GetAllowedAttributes(resourceType string, resourceAttributes map[string]ResourceAttributes) []string {
	attr := resourceAttributes[resourceType]
	allowed := make([]string, 0, len(attr.RequiredAttributes)+len(attr.OptionalAttributes))
	allowed = append(allowed, attr.RequiredAttributes...)
	allowed = append(allowed, attr.OptionalAttributes...)
	return allowed
}
