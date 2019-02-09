// Package envconfig provides a function called Unmarshal that allows you to dump
// environment variable values into a structure using struct tags
package envconfig

import (
	"fmt"
	"reflect"
)

// unmarshalStructValue iterates over exported fields to set them if they're tagged with `env:"ENV_NAME"`
func unmarshalStructValue(t reflect.Type, v reflect.Value) error {
	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		fv := v.Field(i)

		// Skip an unexported field (we can only unmarshal exported fields)
		if len(f.PkgPath) > 0 {
			continue
		}

		// Handle pointers
		if fv.Kind() == reflect.Ptr {
			// Skip over it if it's nil
			if fv.IsNil() {
				continue
			}

			// We're just going to reset fv here so the code below works as-if pointers don't exist
			fv = fv.Elem()
		}

		// Handle structs (with recursion!)
		if fv.Kind() == reflect.Struct {
			err := unmarshalStructValue(fv.Type(), fv)
			if err != nil {
				return err
			}

			// Since we ignore tags on fields that have their own fields, skip to the next one
			continue
		}

		// Look up the environment variable name
		name, ok := f.Tag.Lookup("env")
		if !ok {
			// If it's not tagged with one, it's safe to skip to the next field
			continue
		}

		// See if it's marked as required
		required, err := getStructTagAsBool(&f, "required", false)
		if err != nil {
			return err
		}

		// Let's get the environment variable and try to set it
		wasSet, err := setValueFromEnv(fv, name)
		if err != nil {
			return err
		}

		// Check if it's required and not set
		if required && !wasSet {
			return fmt.Errorf("the environment variable %s is required but was not specified", name)
		}
	}
	return nil
}

// Unmarshal dumps environment variable values into exported fields of a `struct` (pass a `*struct`)
// It will automatically look up environment variables that are named in a struct tag.
// For example, if you tag a field in your struct with ```env:"MY_VAR"``` it will
// be filled with the value of that environment variable. If you would like an error
// to be returned if the field is not set, use the tag ```required:"true"```. For boolean,
// fields, any of the following values are valid (any case): `"true"/"false"`, `"yes"/"no"`, `"on"/"off"`,
// `"t"/"f"`, `"y"/"n"`, `"1"/"0"`. Empty/unset environment variable values will not overwrite
// the struct fields that are already set. Any pointer fields will be dereferenced once (but are skipped if `nil`).
// Fields of a struct kind are unmarshalled recursively. Slices are supported with ":" separated strings (only string slices are
// supported though)
func Unmarshal(i interface{}) error {
	if i == nil {
		return fmt.Errorf("expected Unmarshal parameter to be a pointer to struct, received nil")
	}

	// Check the parameter type
	it := reflect.TypeOf(i)
	if it.Kind() != reflect.Ptr {
		return fmt.Errorf("expected Unmarshal parameter to be a pointer to struct")
	}

	// Check its containing type
	et := it.Elem()
	if et.Kind() != reflect.Struct {
		return fmt.Errorf("expected Unmarshal parameter to be a pointer to struct")
	}

	// Reflect its contained value
	v := reflect.ValueOf(i).Elem()

	// Iterate over fields and unmarshal the struct
	return unmarshalStructValue(et, v)
}
