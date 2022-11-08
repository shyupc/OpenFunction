package util

import (
	"os"
	"reflect"
)

func InterfaceIsNil(val interface{}) bool {

	if val == nil {
		return true
	}

	return reflect.ValueOf(val).IsNil()
}

func GetEnvOrDefault(key, defaultValue string) string {
	value, found := os.LookupEnv(key)
	if found {
		return value
	}
	return defaultValue
}
