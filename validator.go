// Package govalidator is package of string validators and sanitizers.
// ver. 0.0.1
package govalidator

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

// IsEmail check if the string is an email.
func IsEmail(str string) bool {
	return Matches(str, Email)
}

// IsURL check if the string is an URL.
func IsURL(str string) bool {
	if str == "" || len(str) >= 2083 {
		return false
	}
	return Matches(str, URL)
}

// IsAlpha check if the string contains only letters (a-zA-Z).
func IsAlpha(str string) bool {
	return Matches(str, Alpha)
}

// IsAlphanumeric check if the string contains only letters and numbers.
func IsAlphanumeric(str string) bool {
	return Matches(str, Alphanumeric)
}

// IsNumeric check if the string contains only numbers.
func IsNumeric(str string) bool {
	return Matches(str, Numeric)
}

// IsHexadecimal check if the string is a hexadecimal number.
func IsHexadecimal(str string) bool {
	return Matches(str, Hexadecimal)
}

// IsHexcolor check if the string is a hexadecimal color.
func IsHexcolor(str string) bool {
	return Matches(str, Hexcolor)
}

// IsRGBcolor check if the string is a valid RGB color in form rgb(RRR, GGG, BBB).
func IsRGBcolor(str string) bool {
	return Matches(str, RGBcolor)
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
	return Matches(str, Int)
}

// IsFloat check if the string is a float.
func IsFloat(str string) bool {
	return str != "" && Matches(str, Float)
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
	return Matches(str, UUID3)
}

// IsUUIDv4 check if the string is a UUID version 4.
func IsUUIDv4(str string) bool {
	return Matches(str, UUID4)
}

// IsUUIDv5 check if the string is a UUID version 5.
func IsUUIDv5(str string) bool {
	return Matches(str, UUID5)
}

// IsUUID check if the string is a UUID (version 3, 4 or 5).
func IsUUID(str string) bool {
	return Matches(str, UUID)
}

// IsCreditCard check if the string is a credit card.
func IsCreditCard(str string) bool {
	r, _ := regexp.Compile("[^0-9]+")
	sanitized := r.ReplaceAll([]byte(str), []byte(""))
	if !Matches(string(sanitized), CreditCard) {
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
		if !Matches(string(sanitized), ISBN10) {
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
		if !Matches(string(sanitized), ISBN13) {
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
	var js map[string]interface{}
	return json.Unmarshal([]byte(str), &js) == nil
}

// IsMultibyte check if the string contains one or more multibyte chars.
func IsMultibyte(str string) bool {
	return Matches(str, Multibyte)
}

// IsASCII check if the string contains ASCII chars only.
func IsASCII(str string) bool {
	return Matches(str, ASCII)
}

// IsFullWidth check if the string contains any full-width chars.
func IsFullWidth(str string) bool {
	return Matches(str, FullWidth)
}

// IsHalfWidth check if the string contains any half-width chars.
func IsHalfWidth(str string) bool {
	return Matches(str, HalfWidth)
}

// IsVariableWidth check if the string contains a mixture of full and half-width chars.
func IsVariableWidth(str string) bool {
	return Matches(str, HalfWidth) && Matches(str, FullWidth)
}

// IsBase64 check if a string is base64 encoded.
func IsBase64(str string) bool {
	return Matches(str, Base64)
}

// IsDataURI checks if a string is base64 encoded data URI such as an image
func IsDataURI(str string) bool {
	dataURI := strings.Split(str, ",")
	if !Matches(dataURI[0], DataURI) {
		return false
	}
	return IsBase64(dataURI[1])
}

// IsIPv4 check if the string is an IP version 4.
func IsIPv4(str string) bool {
	return IsIP(str, 4)
}

// IsIPv6 check if the string is an IP version 6.
func IsIPv6(str string) bool {
	return IsIP(str, 6)
}

// IsIP check if the string is an IP (version 4 or 6).
// If version value is not equal to 6 or 4, it will be check both variants.
func IsIP(str string, version int) bool {
	if version == 4 {
		if !Matches(str, IPv4) {
			return false
		}
		parts := strings.Split(str, ".")
		isIPv4 := true
		for i := 0; i < len(parts); i++ {
			partI, _ := ToInt(parts[i])
			isIPv4 = isIPv4 && ((partI >= 0) && (partI <= 255))
		}
		return isIPv4
	} else if version == 6 {
		return Matches(str, IPv6)
	}
	return (IsIP(str, 4) || IsIP(str, 6))
}

// IsMAC check if a string is valid MAC address.
// Possible MAC formats:
// 3D:F2:C9:A6:B3:4F
// 3D-F2-C9-A6-B3:4F
// 3d-f2-c9-a6-b3:4f
func IsMAC(str string) bool {
	return Matches(str, MAC)
}

// IsLatitude check if a string is valid latitude.
func IsLatitude(str string) bool {
	return Matches(str, Latitude)
}

// IsLongitude check if a string is valid longitude.
func IsLongitude(str string) bool {
	return Matches(str, Longitude)
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
			resultItem, err := ValidateStruct(v.Index(i).Interface())
			if err != nil {
				return false, err
			}
			result = result && resultItem
		}
		return result, nil
	case reflect.Array:
		result := true
		for i := 0; i < v.Len(); i++ {
			resultItem, err := ValidateStruct(v.Index(i).Interface())
			if err != nil {
				return false, err
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
