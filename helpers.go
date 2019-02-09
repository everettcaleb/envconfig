package envconfig

import (
	"fmt"
	"reflect"
	"strings"
)

// getStructTagAsBool gets a struct tag (by name) as a boolean or returns a default value if not specified
func getStructTagAsBool(f *reflect.StructField, tag string, defaultValue bool) (bool, error) {
	str, ok := f.Tag.Lookup(tag)
	if ok {
		v, err := parseFriendlyBool(str)
		if err != nil {
			return false, fmt.Errorf("Tag %s must have boolean string value", tag)
		}
		return v, nil
	}
	return defaultValue, nil
}

// bitsOf gets the number of bits of the reflected variable kind
func bitsOf(k reflect.Kind) int {
	switch k {
	case reflect.Int8, reflect.Uint8:
		return 8
	case reflect.Int16, reflect.Uint16:
		return 16
	case reflect.Float32, reflect.Int32, reflect.Uint32:
		return 32
	case reflect.Float64, reflect.Int, reflect.Int64, reflect.Uint, reflect.Uint64:
		return 64
	default:
		return 0
	}
}

// parseFriendlyBool parses a multitude of values as boolean case-insensitively: "yes"/"no", "true"/"false", "on"/"off",
// "t"/"f", "y"/"n", "1"/"0" to allow for more expressive configuration settings
func parseFriendlyBool(s string) (bool, error) {
	switch strings.ToLower(s) {
	case "yes", "true", "on", "t", "y", "1":
		return true, nil
	case "no", "false", "off", "f", "n", "0":
		return false, nil
	default:
		return false, fmt.Errorf("Expected boolean value, got: %s", s)
	}
}
