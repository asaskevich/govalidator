govalidator
===========
[![Build Status](https://travis-ci.org/asaskevich/govalidator.svg?branch=master)](https://travis-ci.org/asaskevich/govalidator) [![GoDoc](https://godoc.org/github.com/asaskevich/govalidator?status.png)](https://godoc.org/github.com/asaskevich/govalidator)
[![wercker status](https://app.wercker.com/status/1ec990b09ea86c910d5f08b0e02c6043/s "wercker status")](https://app.wercker.com/project/bykey/1ec990b09ea86c910d5f08b0e02c6043)

A package of string validators and sanitizers for Go lang. Based on [validator.js](https://github.com/chriso/validator.js).

Installation
===========
Type it in your terminal:

	go get github.com/asaskevich/govalidator
	
List of functions:
===========
* func BlackList(str, chars string) string
* func Contains(str, substr string) bool
* func IsASCII(str string) bool
* func IsAlpha(str string) bool
* func IsAlphanumeric(str string) bool
* func IsByteLength(str string, min, max int) bool
* func IsCreditCard(str string) bool
* func IsDivisibleBy(str, num string) bool
* func IsEmail(str string) bool
* func IsFloat(str string) bool
* func IsFullWidth(str string) bool
* func IsHalfWidth(str string) bool
* func IsHexadecimal(str string) bool
* func IsHexcolor(str string) bool
* func IsISBN(str string, version int) bool
* func IsInt(str string) bool
* func IsJSON(str string) bool
* func IsLowerCase(str string) bool
* func IsMultibyte(str string) bool
* func IsNull(str string) bool
* func IsNumeric(str string) bool
* func IsURL(str string) bool
* func IsUUID(str string, version int) bool
* func IsUpperCase(str string) bool
* func IsVariableWidth(str string) bool
* func LeftTrim(str, chars string) string
* func Matches(str, pattern string) bool
* func ReplacePattern(str, pattern, replace string) string
* func RightTrim(str, chars string) string
* func StripLow(str string, KeepNewLines bool) string
* func ToBoolean(str string) bool
* func ToFloat(str string) float64
* func ToInt(str string) int64
* func ToString(obj interface{}) string
* func Trim(str, chars string) string
* func WhiteList(str, chars string) string

Documentation is available here: [godoc.org](https://godoc.org/github.com/asaskevich/govalidator).
Information about code coverage is available here: [govalidator on gocover.io](http://gocover.io/github.com/asaskevich/govalidator).

If you do have a contribution for the package feel free to put up a Pull Request or open Issue.
