// Package govalidator is package of string validators and sanitizers.
package govalidator

import (
	"encoding/json"
	"regexp"
	"strings"
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
func IsDivisibleBy(str, num string) bool {
	p := int64(ToFloat(str))
	q := ToInt(num)
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

// IsUUID check if the string is a UUID (version 3, 4 or 5).
func IsUUID(str string, version int) bool {
	switch version {
	case 0:
		return Matches(str, UUID)
	case 3:
		return Matches(str, UUID3)
	case 4:
		return Matches(str, UUID4)
	case 5:
		return Matches(str, UUID5)
	}
	return false
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
		tmpNum = ToInt(digit)
		if shouldDouble {
			tmpNum *= 2
			if tmpNum >= 10 {
				sum += ((tmpNum%10)+1)
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

// IsISBN check if the string is an ISBN (version 10 or 13).
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
			checksum += (i+1)*int32(sanitized[i])
		}
		if sanitized[9] == 'X' {
			checksum += 10*10
		} else {
			checksum += 10*int32(sanitized[9])
		}
		if checksum%11 == 0 {
			return true
		}
	}
	if version == 13 {
		if !Matches(string(sanitized), ISBN13) {
			return false
		}
		factor := []int32{1, 3}
		for i = 0; i < 12; i++ {
			checksum += factor[i%2]*int32(sanitized[i])
		}
		if int32(sanitized[12])-((10-(checksum%10))%10) == 0 {
			return true
		}
	}
	return false
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
