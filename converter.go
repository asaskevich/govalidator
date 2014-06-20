// A package of string validators and sanitizers.
package govalidator

import (
	"encoding/json"
	"strconv"
)

// Convert the input to a string.
func ToString(obj interface{}) string {
	res, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(res)
}

// Convert the input string to a float, or 0.0 if the input is not a float.
func ToFloat(str string) float64 {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0
	}
	return res
}

// Convert the input string to an integer, or 0 if the input is not an integer.
func ToInt(str string) int64 {
	res, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		return 0
	}
	return res
}

// Convert the input string to a boolean.
func ToBoolean(str string) bool {
	res, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return res
}
