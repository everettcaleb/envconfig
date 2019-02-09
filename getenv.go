package envconfig

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var env = os.Getenv

func setStringFromEnv(v reflect.Value, name string) bool {
	s := env(name)
	if len(s) == 0 {
		return false
	}
	v.SetString(s)
	return true
}

func setStringSliceFromEnv(v reflect.Value, name string) bool {
	s := env(name)
	if len(s) == 0 {
		return false
	}
	v.Set(reflect.ValueOf(strings.Split(s, ":")))
	return true
}

func setIntFromEnv(v reflect.Value, name string) (bool, error) {
	s := env(name)
	if len(s) == 0 {
		return false, nil
	}

	val, err := strconv.ParseInt(s, 10, bitsOf(v.Kind()))
	if err != nil {
		return false, err
	}
	v.SetInt(val)
	return true, nil
}

func setUintFromEnv(v reflect.Value, name string) (bool, error) {
	s := env(name)
	if len(s) == 0 {
		return false, nil
	}

	val, err := strconv.ParseUint(s, 10, bitsOf(v.Kind()))
	if err != nil {
		return false, err
	}
	v.SetUint(val)
	return true, nil
}

func setFloatFromEnv(v reflect.Value, name string) (bool, error) {
	s := env(name)
	if len(s) == 0 {
		return false, nil
	}

	val, err := strconv.ParseFloat(s, bitsOf(v.Kind()))
	if err != nil {
		return false, err
	}
	v.SetFloat(val)
	return true, nil
}

func setBoolFromEnv(v reflect.Value, name string) (bool, error) {
	s := env(name)
	if len(s) == 0 {
		return false, nil
	}

	b, err := parseFriendlyBool(s)
	if err != nil {
		return false, err
	}
	v.SetBool(b)
	return true, nil
}

// setValueFromEnv gets an environment variable by name and assigns it to a reflected value
// First return value is true if the field was set, false otherwise
// Second return value is an error for if something went wrong (invalid format for parsing, etc)
func setValueFromEnv(v reflect.Value, name string) (bool, error) {
	k := v.Kind()

	switch k {
	case reflect.Bool:
		return setBoolFromEnv(v, name)

	case reflect.Float32, reflect.Float64:
		return setFloatFromEnv(v, name)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return setIntFromEnv(v, name)

	case reflect.Ptr:
		return false, fmt.Errorf("invalid kind, pointer-to-pointer is not supported and single pointer is resolved by unmarshalStructValue")

	case reflect.Slice:
		if v.Type().Elem().Kind() == reflect.String {
			return setStringSliceFromEnv(v, name), nil
		}
		return false, fmt.Errorf("invalid kind, for slices only string slices are currently supported")

	case reflect.String:
		return setStringFromEnv(v, name), nil

	case reflect.Struct:
		return false, fmt.Errorf("invalid kind, struct must be processed by unmarshalStructValue")

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return setUintFromEnv(v, name)

	default:
		return false, fmt.Errorf("invalid kind")
	}
}
