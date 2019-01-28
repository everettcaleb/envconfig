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

func getEnvAndSet(obj *reflect.Value, f *reflect.StructField, fieldIndex int, envName string) (bool, error) {
	switch f.Type.Kind() {
	case reflect.Bool:
		v, isSet, err := getEnvBool(envName)
		if isSet {
			obj.Field(fieldIndex).SetBool(v)
		}
		return isSet, err

	case reflect.Float32:
		v, isSet, err := getEnvFloat(envName, 32)
		if isSet {
			obj.Field(fieldIndex).SetFloat(v)
		}
		return isSet, err

	case reflect.Float64:
		v, isSet, err := getEnvFloat(envName, 64)
		if isSet {
			obj.Field(fieldIndex).SetFloat(v)
		}
		return isSet, err

	case reflect.Int:
		v, isSet, err := getEnvInt(envName, 10, 32)
		if isSet {
			obj.Field(fieldIndex).SetInt(v)
		}
		return isSet, err

	case reflect.Int8:
		v, isSet, err := getEnvInt(envName, 10, 8)
		if isSet {
			obj.Field(fieldIndex).SetInt(v)
		}
		return isSet, err

	case reflect.Int16:
		v, isSet, err := getEnvInt(envName, 10, 16)
		if isSet {
			obj.Field(fieldIndex).SetInt(v)
		}
		return isSet, err

	case reflect.Int32:
		v, isSet, err := getEnvInt(envName, 10, 32)
		if isSet {
			obj.Field(fieldIndex).SetInt(v)
		}
		return isSet, err

	case reflect.Int64:
		v, isSet, err := getEnvInt(envName, 10, 64)
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

	case reflect.Uint:
		v, isSet, err := getEnvUint(envName, 10, 32)
		if isSet {
			obj.Field(fieldIndex).SetUint(v)
		}
		return isSet, err

	case reflect.Uint8:
		v, isSet, err := getEnvUint(envName, 10, 8)
		if isSet {
			obj.Field(fieldIndex).SetUint(v)
		}
		return isSet, err

	case reflect.Uint16:
		v, isSet, err := getEnvUint(envName, 10, 16)
		if isSet {
			obj.Field(fieldIndex).SetUint(v)
		}
		return isSet, err

	case reflect.Uint32:
		v, isSet, err := getEnvUint(envName, 10, 32)
		if isSet {
			obj.Field(fieldIndex).SetUint(v)
		}
		return isSet, err

	case reflect.Uint64:
		v, isSet, err := getEnvUint(envName, 10, 64)
		if isSet {
			obj.Field(fieldIndex).SetUint(v)
		}
		return isSet, err

	case reflect.Array, reflect.Ptr, reflect.Slice, reflect.Struct:
		return false, fmt.Errorf("todo, currently unsupported but I definitely want this to support arrays, pointers, slices, and structs in the future")

	default:
		return false, fmt.Errorf("invalid kind")
	}
}
