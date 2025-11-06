package helpers

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

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
		// Handle types.String
		if field.Type() == reflect.TypeOf(types.String{}) {
			ts := field.Interface().(types.String)
			if !ts.IsNull() {
				providedAttrs = append(providedAttrs, tag)
			}
			continue
		}
		// Handle types.Int64
		if field.Type() == reflect.TypeOf(types.Int64{}) {
			ti := field.Interface().(types.Int64)
			if !ti.IsNull() {
				providedAttrs = append(providedAttrs, tag)
			}
			continue
		}
		// Handle types.Bool
		if field.Type() == reflect.TypeOf(types.Bool{}) {
			tbool := field.Interface().(types.Bool)
			if !tbool.IsNull() {
				providedAttrs = append(providedAttrs, tag)
			}
			continue
		}
		// Handle types.Float64
		if field.Type() == reflect.TypeOf(types.Float64{}) {
			tf := field.Interface().(types.Float64)
			if !tf.IsNull() {
				providedAttrs = append(providedAttrs, tag)
			}
			continue
		}
		// Handle types.Number
		if field.Type() == reflect.TypeOf(types.Number{}) {
			tn := field.Interface().(types.Number)
			if !tn.IsNull() {
				providedAttrs = append(providedAttrs, tag)
			}
			continue
		}
		// Handle types.List
		if field.Type() == reflect.TypeOf(types.List{}) {
			tl := field.Interface().(types.List)
			if !tl.IsNull() {
				providedAttrs = append(providedAttrs, tag)
			}
			continue
		}
		// Handle types.Set
		if field.Type() == reflect.TypeOf(types.Set{}) {
			ts := field.Interface().(types.Set)
			if !ts.IsNull() {
				providedAttrs = append(providedAttrs, tag)
			}
			continue
		}
		// Handle types.Map
		if field.Type() == reflect.TypeOf(types.Map{}) {
			tm := field.Interface().(types.Map)
			if !tm.IsNull() {
				providedAttrs = append(providedAttrs, tag)
			}
			continue
		}
		// Handle types.Object
		if field.Type() == reflect.TypeOf(types.Object{}) {
			to := field.Interface().(types.Object)
			if !to.IsNull() {
				providedAttrs = append(providedAttrs, tag)
			}
			continue
		}
		// For pointer types (e.g. nested blocks), check if non-nil.
		if field.Kind() == reflect.Ptr && !field.IsNil() {
			providedAttrs = append(providedAttrs, tag)
			continue
		}
		// ... handle additional tfsdk types if needed.
	}
	return providedAttrs
}
