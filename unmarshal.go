package envconfig

import (
	"fmt"
	"reflect"
)

// Unmarshal dumps environment variable values into a struct (pass a *struct)
// It will automatically look up environment variables that are named in a struct tag.
// For example, if you tag a field in your struct with `env:"MY_VAR"` it will
// be filled with the value of that environment variable. If you would like an error
// to be returned if the field is not set, use the tag `required:"true"`. For boolean,
// fields, any of the following values are valid (any case): "true", "false", "yes", "no",
// "t", "f", "y", "n", "1", "0". Empty/unset environment variable values will not overwrite
// the struct fields that are already set.
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

	// Iterate over fields and look for
	for fi := 0; fi < et.NumField(); fi++ {
		f := et.Field(fi)

		// Look up the environment variable name
		envName, ok := f.Tag.Lookup("env")
		if !ok {
			// If it's not tagged with one, it's safe to skip to the next field
			continue
		}

		// See if it's marked as required
		required, err := getFieldTagBool(&f, "required", false)
		if err != nil {
			return err
		}

		// Let's get the environment variable and try to set it
		isSet, err := getEnvAndSet(&v, &f, fi, envName)
		if err != nil {
			return err
		}

		// Check if it's required and not set
		if required && !isSet {
			return fmt.Errorf("the environment variable %s is required but was not specified", envName)
		}
	}

	return nil
}
