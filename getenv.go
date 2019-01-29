package envconfig

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var env = os.Getenv

func parseFriendlyBool(s string) (bool, error) {
	switch strings.ToLower(s) {
	case "yes", "true", "t", "y", "1":
		return true, nil
	case "no", "false", "f", "n", "0":
		return false, nil
	default:
		return false, fmt.Errorf("Expected boolean value, got: %s", s)
	}
}

func getEnvString(name string) (string, bool) {
	s := env(name)
	return s, len(s) > 0
}

func getEnvInt(name string, base int, bits int) (int64, bool, error) {
	s := env(name)
	if len(s) == 0 {
		return 0, false, nil
	}

	val, err := strconv.ParseInt(s, base, bits)
	return val, err == nil, err
}

func getEnvUint(name string, base int, bits int) (uint64, bool, error) {
	s := env(name)
	if len(s) == 0 {
		return 0, false, nil
	}

	val, err := strconv.ParseUint(s, base, bits)
	return val, err == nil, err
}

func getEnvFloat(name string, bits int) (float64, bool, error) {
	s := env(name)
	if len(s) == 0 {
		return 0, false, nil
	}

	val, err := strconv.ParseFloat(s, bits)
	return val, err == nil, err
}

func getEnvBool(name string) (bool, bool, error) {
	s := env(name)
	if len(s) == 0 {
		return false, false, nil
	}

	b, err := parseFriendlyBool(s)
	return b, err == nil, err
}

func getFieldTagBool(f *reflect.StructField, tag string, defaultValue bool) (bool, error) {
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
	case reflect.Float32, reflect.Int, reflect.Int32, reflect.Uint, reflect.Uint32:
		return 32
	case reflect.Float64, reflect.Int64, reflect.Uint64:
		return 64
	default:
		return 0
	}
}

// getEnvAndSet gets an environment variable by name and assigns it to the field belonging to value "obj"
// First return value is true if the field was set, false otherwise
// Second return value is an error for if something went wrong (invalid format for parsing, etc)
func getEnvAndSet(obj *reflect.Value, f *reflect.StructField, fieldIndex int, envName string) (bool, error) {
	k := f.Type.Kind()
	switch k {
	case reflect.Bool:
		v, isSet, err := getEnvBool(envName)
		if isSet {
			obj.Field(fieldIndex).SetBool(v)
		}
		return isSet, err

	case reflect.Float32, reflect.Float64:
		v, isSet, err := getEnvFloat(envName, bitsOf(k))
		if isSet {
			obj.Field(fieldIndex).SetFloat(v)
		}
		return isSet, err

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, isSet, err := getEnvInt(envName, 10, bitsOf(k))
		if isSet {
			obj.Field(fieldIndex).SetInt(v)
		}
		return isSet, err

	case reflect.String:
		v, isSet := getEnvString(envName)
		if isSet {
			obj.Field(fieldIndex).SetString(v)
		}
		return isSet, nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, isSet, err := getEnvUint(envName, 10, bitsOf(k))
		if isSet {
			obj.Field(fieldIndex).SetUint(v)
		}
		return isSet, err

	default:
		return false, fmt.Errorf("invalid kind")
	}
}
