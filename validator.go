// Package govalidator is package of validators and sanitizers for strings, structs and collections.
package govalidator

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"golang.org/x/net/idna"
	"io/ioutil"
	"net"
	"net/url"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

var (
	fieldsRequiredByDefault bool
	nilPtrAllowedByRequired = false
	notNumberRegexp         = regexp.MustCompile("[^0-9]+")
	whiteSpacesAndMinus     = regexp.MustCompile(`[\s-]+`)
	paramsRegexp            = regexp.MustCompile(`\(.*\)$`)
)

const maxURLRuneCount = 2083
const minURLRuneCount = 3
const rfc3339WithoutZone = "2006-01-02T15:04:05"

// SetFieldsRequiredByDefault causes validation to fail when struct fields
// do not include validations or are not explicitly marked as exempt (using `valid:"-"` or `valid:"email,optional"`).
// This struct definition will fail govalidator.ValidateStruct() (and the field values do not matter):
//     type exampleStruct struct {
//         Name  string ``
//         Email string `valid:"email"`
// This, however, will only fail when Email is empty or an invalid email address:
//     type exampleStruct2 struct {
//         Name  string `valid:"-"`
//         Email string `valid:"email"`
// Lastly, this will only fail when Email is an invalid email address but not when it's empty:
//     type exampleStruct2 struct {
//         Name  string `valid:"-"`
//         Email string `valid:"email,optional"`
func SetFieldsRequiredByDefault(value bool) {
	fieldsRequiredByDefault = value
}

// SetNilPtrAllowedByRequired causes validation to pass for nil ptrs when a field is set to required.
// The validation will still reject ptr fields in their zero value state. Example with this enabled:
//     type exampleStruct struct {
//         Name  *string `valid:"required"`
// With `Name` set to "", this will be considered invalid input and will cause a validation error.
// With `Name` set to nil, this will be considered valid by validation.
// By default this is disabled.
func SetNilPtrAllowedByRequired(value bool) {
	nilPtrAllowedByRequired = value
}

// IsURL checks if the string is an URL.
// Now supports Unicode/IDN domain names by converting them to Punycode before validation.
func IsURL(str string) bool {
	if str == "" || utf8.RuneCountInString(str) >= maxURLRuneCount || len(str) <= minURLRuneCount || strings.HasPrefix(str, ".") {
		return false
	}
	strTemp := str
	if strings.Contains(str, ":") && !strings.Contains(str, "://") {
		// support no indicated urlscheme but with colon for port number
		// http:// is appended so url.Parse will succeed, strTemp used so it does not impact rxURL.MatchString
		strTemp = "http://" + str
	}
	u, err := url.Parse(strTemp)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}
	
	// Convert Unicode/IDN domains to ASCII Punycode for validation
	if u.Host != "" {
		asciiHost, err := idna.ToASCII(u.Host)
		if err != nil {
			return false
		}
		// Replace the host in the URL with the ASCII version for regex validation
		if asciiHost != u.Host {
			// Construct URL with ASCII host for proper validation
			originalHost := u.Host
			u.Host = asciiHost
			strTemp = u.String()
			// If original URL didn't have scheme, remove the http:// we added
			if !strings.Contains(str, "://") {
				strTemp = strings.TrimPrefix(strTemp, "http://")
			}
		}
	}
	
	return rxURL.MatchString(strTemp)
}
