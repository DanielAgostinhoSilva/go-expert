package environment

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	ENV_TAG = "env"
)

func Unmarshal(stc any) error {
	return unmarshalStruct(reflect.ValueOf(stc).Elem())
}

func unmarshalStruct(value reflect.Value) error {
	fieldType := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := fieldType.Field(i)
		fieldValue := value.Field(i)

		if !fieldValue.CanSet() {
			continue
		}
		if fieldValue.Kind() == reflect.Struct {
			if err := unmarshalStruct(fieldValue); err != nil {
				return err
			}
		}

		envTag := field.Tag.Get(ENV_TAG)
		if envTag != "" && envTag != "-" {
			if strings.Contains(envTag, " ") {
				return fmt.Errorf("invalid parameter '%s': The parameters should be separated by ','", envTag)
			}

			envValue := os.Getenv(envTag)
			if envValue == "" {
				return fmt.Errorf("environment variable '%s' not set", envTag)
			}

			if err := setFieldValue(fieldValue, envValue); err != nil {
				return fmt.Errorf("error setting field '%s': %v", field.Name, err)
			}
		}
	}

	return nil
}

func setFieldValue(fieldValue reflect.Value, envValue string) error {
	switch fieldValue.Kind() {
	case reflect.String:
		fieldValue.SetString(envValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, err := strconv.ParseInt(envValue, 10, 64)
		if err != nil {
			return fmt.Errorf("error converting '%s' to int: %v", envValue, err)
		}
		fieldValue.SetInt(intValue)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintValue, err := strconv.ParseUint(envValue, 10, 64)
		if err != nil {
			return fmt.Errorf("error converting '%s' to uint: %v", envValue, err)
		}
		fieldValue.SetUint(uintValue)
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(envValue)
		if err != nil {
			return fmt.Errorf("error converting '%s' to bool: %v", envValue, err)
		}
		fieldValue.SetBool(boolValue)
	case reflect.Float32, reflect.Float64:
		floatValue, err := strconv.ParseFloat(envValue, 64)
		if err != nil {
			return fmt.Errorf("error converting '%s' to float: %v", envValue, err)
		}
		fieldValue.SetFloat(floatValue)
	default:
		return fmt.Errorf("unsupported field type '%s'", fieldValue.Type().String())
	}

	return nil
}
