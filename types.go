package govalidator

import "reflect"

// Validator is a wrapper for validator functions, that returns bool and accepts string.
type Validator func(str string) bool
type tagOptions []string

// UnsupportedTypeError is a wrapper for reflect.Type
type UnsupportedTypeError struct {
	Type reflect.Type
}

// stringValues is a slice of reflect.Value holding *reflect.StringValue.
// It implements the methods to sort by string.
type stringValues []reflect.Value

// TagMap is a map of functions, that can be used as tags for ValidateStruct function.
var TagMap = map[string]Validator{
	"email":         IsEmail,
	"url":           IsURL,
	"alpha":         IsAlpha,
	"alphanum":      IsAlphanumeric,
	"numeric":       IsNumeric,
	"hexadecimal":   IsHexadecimal,
	"hexcolor":      IsHexcolor,
	"rgbcolor":      IsRGBcolor,
	"lowercase":     IsLowerCase,
	"uppercase":     IsUpperCase,
	"int":           IsInt,
	"float":         IsFloat,
	"null":          IsNull,
	"uuid":          IsUUID,
	"uuidv3":        IsUUIDv3,
	"uuidv4":        IsUUIDv4,
	"uuidv5":        IsUUIDv5,
	"creditcard":    IsCreditCard,
	"isbn10":        IsISBN10,
	"isbn13":        IsISBN13,
	"json":          IsJSON,
	"multibyte":     IsMultibyte,
	"ascii":         IsASCII,
	"fullwidth":     IsFullWidth,
	"halfwidth":     IsHalfWidth,
	"variablewidth": IsVariableWidth,
	"base64":        IsBase64,
	"datauri":       IsDataURI,
	"ipv4":          IsIPv4,
	"ipv6":          IsIPv6,
	"mac":           IsMAC,
	"latitude":      IsLatitude,
	"longitude":     IsLongitude,
}
