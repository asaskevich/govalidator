// Package govalidator is package of string validators and sanitizers.
// ver. 0.0.1
package govalidator

import (
	"encoding/json"
	"fmt"
	"net"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

// IsEmail check if the string is an email.
func IsEmail(str string) bool {
	// TODO uppercase letters are not supported
	return rxEmail.MatchString(str)
}

// IsURL check if the string is an URL.
func IsURL(str string) bool {
	if str == "" || len(str) >= 2083 {
		return false
	}
	return rxURL.MatchString(str)
}

// IsAlpha check if the string contains only letters (a-zA-Z).
func IsAlpha(str string) bool {
	return rxAlpha.MatchString(str)
}

//IsUTFLetter check if the string contains only unicode letter characters.
//Similar to IsAlpha but for all languages.
func IsUTFLetter(str string) bool {

	for _, c := range str {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true

}

// IsAlphanumeric check if the string contains only letters and numbers.
func IsAlphanumeric(str string) bool {
	return rxAlphanumeric.MatchString(str)
}

// IsUTFLetterNumeric check if the string contains only unicode letters and numbers.
func IsUTFLetterNumeric(str string) bool {

	for _, c := range str {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) { //letters && numbers are ok
			return false
		}
	}
	return true

}

// IsNumeric check if the string contains only numbers.
func IsNumeric(str string) bool {
	return rxNumeric.MatchString(str)
}

// IsUTFNumeric check if the string contains only unicode numbers of any kind.
// Numbers can be 0-9 but also Fractions ¾,Roman Ⅸ and Hangzhou 〩.
func IsUTFNumeric(str string) bool {

	for _, c := range str {
		if unicode.IsNumber(c) == false && c != '-' { //numbers && minus sign are ok
			return false
		}
	}
	return true

}

// IsUTFDigit check if the string contains only unicode radix-10 decimal digits.
func IsUTFDigit(str string) bool {

	for _, c := range str {
		if !unicode.IsDigit(c) && c != '-' { //digits && minus sign are ok
			return false
		}
	}
	return true

}

// IsHexadecimal check if the string is a hexadecimal number.
func IsHexadecimal(str string) bool {
	return rxHexadecimal.MatchString(str)
}

// IsHexcolor check if the string is a hexadecimal color.
func IsHexcolor(str string) bool {
	return rxHexcolor.MatchString(str)
}

// IsRGBcolor check if the string is a valid RGB color in form rgb(RRR, GGG, BBB).
func IsRGBcolor(str string) bool {
	return rxRGBcolor.MatchString(str)
}

// IsLowerCase check if the string is lowercase.
func IsLowerCase(str string) bool {
	return str == strings.ToLower(str)
}

// IsUpperCase check if the string is uppercase.
func IsUpperCase(str string) bool {
	return str == strings.ToUpper(str)
}

// IsInt check if the string is an integer.
func IsInt(str string) bool {
	return rxInt.MatchString(str)
}

// IsFloat check if the string is a float.
func IsFloat(str string) bool {
	return str != "" && rxFloat.MatchString(str)
}

// IsDivisibleBy check if the string is a number that's divisible by another.
// If second argument is not valid integer or zero, it's return false.
// Otherwise, if first argument is not valid integer or zero, it's return true (Invalid string converts to zero).
func IsDivisibleBy(str, num string) bool {
	f, _ := ToFloat(str)
	p := int64(f)
	q, _ := ToInt(num)
	if q == 0 {
		return false
	}
	return (p == 0) || (p%q == 0)
}

// IsNull check if the string is null.
func IsNull(str string) bool {
	return len(str) == 0
}

// IsByteLength check if the string's length (in bytes) falls in a range.
func IsByteLength(str string, min, max int) bool {
	return len(str) >= min && len(str) <= max
}

// IsUUIDv3 check if the string is a UUID version 3.
func IsUUIDv3(str string) bool {
	return rxUUID3.MatchString(str)
}

// IsUUIDv4 check if the string is a UUID version 4.
func IsUUIDv4(str string) bool {
	return rxUUID4.MatchString(str)
}

// IsUUIDv5 check if the string is a UUID version 5.
func IsUUIDv5(str string) bool {
	return rxUUID5.MatchString(str)
}

// IsUUID check if the string is a UUID (version 3, 4 or 5).
func IsUUID(str string) bool {
	return rxUUID.MatchString(str)
}

// IsCreditCard check if the string is a credit card.
func IsCreditCard(str string) bool {
	r, _ := regexp.Compile("[^0-9]+")
	sanitized := r.ReplaceAll([]byte(str), []byte(""))
	if !rxCreditCard.MatchString(string(sanitized)) {
		return false
	}
	var sum int64
	var digit string
	var tmpNum int64
	var shouldDouble bool
	for i := len(sanitized) - 1; i >= 0; i-- {
		digit = string(sanitized[i:(i + 1)])
		tmpNum, _ = ToInt(digit)
		if shouldDouble {
			tmpNum *= 2
			if tmpNum >= 10 {
				sum += ((tmpNum % 10) + 1)
			} else {
				sum += tmpNum
			}
		} else {
			sum += tmpNum
		}
		shouldDouble = !shouldDouble
	}

	if sum%10 == 0 {
		return true
	}
	return false
}

// IsISBN10 check if the string is an ISBN version 10.
func IsISBN10(str string) bool {
	return IsISBN(str, 10)
}

// IsISBN13 check if the string is an ISBN version 13.
func IsISBN13(str string) bool {
	return IsISBN(str, 13)
}

// IsISBN check if the string is an ISBN (version 10 or 13).
// If version value is not equal to 10 or 13, it will be check both variants.
func IsISBN(str string, version int) bool {
	r, _ := regexp.Compile("[\\s-]+")
	sanitized := r.ReplaceAll([]byte(str), []byte(""))
	var checksum int32
	var i int32
	if version == 10 {
		if !rxISBN10.MatchString(string(sanitized)) {
			return false
		}
		for i = 0; i < 9; i++ {
			checksum += (i + 1) * int32(sanitized[i]-'0')
		}
		if sanitized[9] == 'X' {
			checksum += 10 * 10
		} else {
			checksum += 10 * int32(sanitized[9]-'0')
		}
		if checksum%11 == 0 {
			return true
		}
		return false
	} else if version == 13 {
		if !rxISBN13.MatchString(string(sanitized)) {
			return false
		}
		factor := []int32{1, 3}
		for i = 0; i < 12; i++ {
			checksum += factor[i%2] * int32(sanitized[i]-'0')
		}
		if (int32(sanitized[12]-'0'))-((10-(checksum%10))%10) == 0 {
			return true
		}
		return false
	}
	return IsISBN(str, 10) || IsISBN(str, 13)
}

// IsJSON check if the string is valid JSON (note: uses json.Unmarshal).
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// IsMultibyte check if the string contains one or more multibyte chars.
func IsMultibyte(str string) bool {
	return rxMultibyte.MatchString(str)
}

// IsASCII check if the string contains ASCII chars only.
func IsASCII(str string) bool {
	return rxASCII.MatchString(str)
}

// IsFullWidth check if the string contains any full-width chars.
func IsFullWidth(str string) bool {
	return rxFullWidth.MatchString(str)
}

// IsHalfWidth check if the string contains any half-width chars.
func IsHalfWidth(str string) bool {
	return rxHalfWidth.MatchString(str)
}

// IsVariableWidth check if the string contains a mixture of full and half-width chars.
func IsVariableWidth(str string) bool {
	return rxHalfWidth.MatchString(str) && rxFullWidth.MatchString(str)
}

// IsBase64 check if a string is base64 encoded.
func IsBase64(str string) bool {
	return rxBase64.MatchString(str)
}

// IsDataURI checks if a string is base64 encoded data URI such as an image
func IsDataURI(str string) bool {
	dataURI := strings.Split(str, ",")
	if !rxDataURI.MatchString(dataURI[0]) {
		return false
	}
	return IsBase64(dataURI[1])
}

// IsIP checks if a string is either IP version 4 or 6.
func IsIP(str string) bool {
	return net.ParseIP(str) != nil
}

// IsIPv4 check if the string is an IP version 4.
func IsIPv4(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && ip.To4() != nil
}

// IsIPv6 check if the string is an IP version 6.
func IsIPv6(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && ip.To4() == nil
}

// IsMAC check if a string is valid MAC address.
// Possible MAC formats:
// 01:23:45:67:89:ab
// 01:23:45:67:89:ab:cd:ef
// 01-23-45-67-89-ab
// 01-23-45-67-89-ab-cd-ef
// 0123.4567.89ab
// 0123.4567.89ab.cdef
func IsMAC(str string) bool {
	_, err := net.ParseMAC(str)
	return err == nil
}

// IsMongoId check if the string is a valid hex-encoded representation of a MongoDB ObjectId.
func IsMongoId(str string) bool {
	return rxHexadecimal.MatchString(str) && (len(str) == 24)
}

// IsLatitude check if a string is valid latitude.
func IsLatitude(str string) bool {
	return rxLatitude.MatchString(str)
}

// IsLongitude check if a string is valid longitude.
func IsLongitude(str string) bool {
	return rxLongitude.MatchString(str)
}

// ValidateStruct use tags for fields
func ValidateStruct(s interface{}) (bool, error) {
	if s == nil {
		return true, nil
	}
	result := true
	var err error
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	// we only accept structs
	if val.Kind() != reflect.Struct {
		return false, fmt.Errorf("function only accepts structs; got %T", val)
	}
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		if typeField.PkgPath != "" {
			continue // Private field
		}
		resultField, err := typeCheck(valueField, typeField)
		if err != nil {
			return false, err
		}
		result = result && resultField
	}
	return result, err
}

// parseTag splits a struct field's tag into its
// comma-separated options.
func parseTag(tag string) tagOptions {
	split := strings.SplitN(tag, ",", -1)
	return tagOptions(split)
}

func isValidTag(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		switch {
		case strings.ContainsRune("!#$%&()*+-./:<=>?@[]^_{|}~ ", c):
			// Backslash and quote chars are reserved, but
			// otherwise any punctuation chars are allowed
			// in a tag name.
		default:
			if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
				return false
			}
		}
	}
	return true
}

// IsSSN will validate the given string as a U.S. Social Security Number
func IsSSN(str string) bool {
	if str == "" || len(str) != 11 {
		return false
	}
	return rxSSN.MatchString(str)
}

// Contains returns whether checks that a comma-separated list of options
// contains a particular substr flag. substr must be surrounded by a
// string boundary or commas.
func (opts tagOptions) contains(optionName string) bool {
	for i := range opts {
		tagOpt := opts[i]
		if tagOpt == optionName {
			return true
		}
	}
	return false
}

func typeCheck(v reflect.Value, t reflect.StructField) (bool, error) {
	if !v.IsValid() {
		return false, nil
	}
	switch v.Kind() {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.String:
		tag := t.Tag.Get(tagName)
		// Check if the field should be ignored
		if tag == "-" || tag == "" {
			return true, nil
		}
		options := parseTag(tag)
		if options.contains("required") {
			result := func(field interface{}) bool {
				//test is underlying type is not: nil, 0, ""
				return field != nil || field != reflect.Zero(reflect.TypeOf(v)).Interface()
			}(v)
			if result == false {
				err := fmt.Errorf("non zero value required for type %s", t.Name)
				return result, err
			}
		} else if isEmptyValue(v) { //not required and empty is valid
			return true, nil
		}
		//for each tag option check the map of validator functions
		for i := range options {
			tagOpt := options[i]
			if ok := isValidTag(tagOpt); !ok {
				continue
			}
			if validatefunc, ok := TagMap[tagOpt]; ok {
				if v.Kind() == reflect.String { //TODO:other options/types to string
					field := fmt.Sprint(v) //make value into string, then validate with regex
					if result := validatefunc(field); !result {
						err := fmt.Errorf("value: %s=%s does not validate as %s", t.Name, field, tagOpt)
						return result, err
					}
				}
			}
		}
		return true, nil
	case reflect.Map:
		if v.Type().Key().Kind() != reflect.String {
			return false, &UnsupportedTypeError{v.Type()}
		}
		//an empty map is not validated, always true
		if v.IsNil() {
			return true, nil
		}
		var sv stringValues
		sv = v.MapKeys()
		sort.Sort(sv)
		result := true
		for _, k := range sv {
			resultItem, err := ValidateStruct(v.MapIndex(k).Interface())
			if err != nil {
				return false, err
			}
			result = result && resultItem
		}
		return result, nil
	case reflect.Slice:
		//an empty slice is not validated, always true
		if v.IsNil() {
			return true, nil
		}
		result := true
		for i := 0; i < v.Len(); i++ {
			var resultItem bool
			var err error
			if v.Index(i).Kind() != reflect.Struct {
				resultItem, err = typeCheck(v.Index(i), t)
				if err != nil {
					return false, err
				}
			} else {
				resultItem, err = ValidateStruct(v.Index(i).Interface())
				if err != nil {
					return false, err
				}
			}
			result = result && resultItem
		}
		return result, nil
	case reflect.Array:
		result := true
		for i := 0; i < v.Len(); i++ {
			var resultItem bool
			var err error
			if v.Index(i).Kind() != reflect.Struct {
				resultItem, err = typeCheck(v.Index(i), t)
				if err != nil {
					return false, err
				}
			} else {
				resultItem, err = ValidateStruct(v.Index(i).Interface())
				if err != nil {
					return false, err
				}
			}
			result = result && resultItem
		}
		return result, nil

	case reflect.Interface, reflect.Ptr:
		// If the value is an interface or pointer then encode its element
		if v.IsNil() {
			return true, nil
		}
		return ValidateStruct(v.Interface())
	case reflect.Struct:
		return ValidateStruct(v.Interface())
	default:
		return false, &UnsupportedTypeError{v.Type()}
	}
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}

// Error returns string equivalent for reflect.Type
func (e *UnsupportedTypeError) Error() string {
	return "validator: unsupported type: " + e.Type.String()
}

func (sv stringValues) Len() int           { return len(sv) }
func (sv stringValues) Swap(i, j int)      { sv[i], sv[j] = sv[j], sv[i] }
func (sv stringValues) Less(i, j int) bool { return sv.get(i) < sv.get(j) }
func (sv stringValues) get(i int) string   { return sv[i].String() }
