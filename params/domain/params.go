package domain

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type Params struct {
}

func (p *Params) GetQueryParams(req *http.Request, params interface{}) {
	// Get all Query Parameters as a map
	queryParams := req.URL.Query()

	// Get the mutable value of the structure
	valueOf := reflect.ValueOf(params).Elem()

	// Iterate over the fields of the structure
	for i := 0; i < valueOf.NumField(); i++ {
		// Get the name and type of each field
		fieldValue := valueOf.Field(i)
		fieldName := valueOf.Type().Field(i).Tag.Get("json")

		values, exists := queryParams[fieldName]
		if exists && len(values) > 0 {
			// Set the new value of the field based on the data type
			setFieldValue(fieldValue, values)
		}
	}
}

func setFieldValue(fieldValue reflect.Value, values []string) {
	switch fieldValue.Kind() {
	case reflect.String:
		// If there are multiple values, concatenate them with a separator
		fieldValue.SetString(strings.Join(values, ","))
	case reflect.Ptr:
		// If it's a pointer, create a new value of the underlying type and set the value
		elem := reflect.New(fieldValue.Type().Elem()).Elem()
		setFieldValue(elem, values)
		fieldValue.Set(elem.Addr())
	case reflect.Int:
		// Try to convert the first value to int
		if len(values) > 0 {
			intValue, err := strconv.Atoi(values[0])
			if err == nil {
				fieldValue.SetInt(int64(intValue))
			}
		}
	case reflect.Bool:
		// Try to convert the first value to bool
		if len(values) > 0 {
			boolValue, err := strconv.ParseBool(values[0])
			if err == nil {
				fieldValue.SetBool(boolValue)
			}
		}
	case reflect.Slice:
		// Set the field as a slice of strings
		fieldValue.Set(reflect.ValueOf(values))
	}
}
