package govalidator

import (
	"encoding/json"
	"strconv"
)

// ToString convert the input to a string.
func ToString(obj interface{}) string {
	res, _ := json.Marshal(obj)
	return string(res)
}

// ToFloat convert the input string to a float, or 0.0 if the input is not a float.
func ToFloat(str string) float64 {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0
	}
	return res
}

// ToInt convert the input string to an integer, or 0 if the input is not an integer.
func ToInt(str string) int64 {
	res, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		return 0
	}
	return res
}

// ToBoolean convert the input string to a boolean.
func ToBoolean(str string) bool {
	res, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return res
}
