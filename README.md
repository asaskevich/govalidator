govalidator
===========
[![GoDoc](https://godoc.org/github.com/asaskevich/govalidator?status.png)](https://godoc.org/github.com/asaskevich/govalidator) [![Coverage Status](https://img.shields.io/coveralls/asaskevich/govalidator.svg)](https://coveralls.io/r/asaskevich/govalidator?branch=master) [![views](https://sourcegraph.com/api/repos/github.com/asaskevich/govalidator/.counters/views.png)](https://sourcegraph.com/github.com/asaskevich/govalidator)
[![wercker status](https://app.wercker.com/status/1ec990b09ea86c910d5f08b0e02c6043/s "wercker status")](https://app.wercker.com/project/bykey/1ec990b09ea86c910d5f08b0e02c6043)

A package of string validators and sanitizers for Go lang. Based on [validator.js](https://github.com/chriso/validator.js).

#### Installation
Make sure that Go is installed on your computer.
Type the following command in your terminal:

	go get github.com/asaskevich/govalidator
	
After it the package is ready to use.

#### Import package in your project
Add following line in your `*.go` file:
```go
import "github.com/asaskevich/govalidator"
```
If you unhappy to use long `govalidator`, you can do something like this:
```go
import (
	valid "github.com/asaskevich/govalidator"
)
```

#### List of functions:
```go
func BlackList(str, chars string) string
func CamelCaseToUnderscore(str string) string
func Contains(str, substring string) bool
func Escape(str string) string
func GetLine(s string, index int) (string, error)
func GetLines(s string) []string
func IsASCII(str string) bool
func IsAlpha(str string) bool
func IsAlphanumeric(str string) bool
func IsBase64(str string) bool
func IsByteLength(str string, min, max int) bool
func IsCreditCard(str string) bool
func IsDataURI(str string) bool
func IsDivisibleBy(str, num string) bool
func IsEmail(str string) bool
func IsFloat(str string) bool
func IsFullWidth(str string) bool
func IsHalfWidth(str string) bool
func IsHexadecimal(str string) bool
func IsHexcolor(str string) bool
func IsIP(str string, version int) bool
func IsIPv4(str string) bool
func IsIPv6(str string) bool
func IsISBN(str string, version int) bool
func IsISBN10(str string) bool
func IsISBN13(str string) bool
func IsInt(str string) bool
func IsJSON(str string) bool
func IsLatitude(str string) bool
func IsLongitude(str string) bool
func IsLowerCase(str string) bool
func IsMAC(str string) bool
func IsMultibyte(str string) bool
func IsNull(str string) bool
func IsNumeric(str string) bool
func IsRGBcolor(str string) bool
func IsURL(str string) bool
func IsUUID(str string) bool
func IsUUIDv3(str string) bool
func IsUUIDv4(str string) bool
func IsUUIDv5(str string) bool
func IsUpperCase(str string) bool
func IsVariableWidth(str string) bool
func LeftTrim(str, chars string) string
func Matches(str, pattern string) bool
func RemoveTags(s string) string
func ReplacePattern(str, pattern, replace string) string
func Reverse(s string) string
func RightTrim(str, chars string) string
func SafeFileName(str string) string
func StripLow(str string, keepNewLines bool) string
func ToBoolean(str string) (bool, error)
func ToFloat(str string) (float64, error)
func ToInt(str string) (int64, error)
func ToString(obj interface{}) (string, error)
func Trim(str, chars string) string
func UnderscoreToCamelCase(s string) string
func ValidateStruct(s interface{}) (bool, error)
func WhiteList(str, chars string) string
type UnsupportedTypeError
	func (e *UnsupportedTypeError) Error() string
type Validator
```

#### Examples
###### IsURL
```go
println(govalidator.IsURL(`http://user@pass:domain.com/path/page`))
```
###### ToString
```go
type User struct {
	FirstName string
	LastName string
}

str,_ := govalidator.ToString(&User{"John", "Juan"})
println(str)
```
###### ValidateStruct [#2](https://github.com/asaskevich/govalidator/pull/2)
If you want to validate structs, you can use tag `valid` for any field in your structure. All validators used with this field in one tag are separated by comma. If you want to ignore tag, place `-` in your tag. If you think, that package has no necessary validators, you can add it:
```go
govalidator.TagMap["duck"] = govalidator.Validator(func(str string) bool {
    return str == "duck"
})
```
Here is a list of available validators for struct fields (validator - used function):
```go
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
"longitude":     IsLongitude
```
And here is small example of usage:
```go
type Post struct {
    Title    string `valid:"alphanum,required"`
    Message  string `valid:"duck,ascii"`
    AuthorIP string `valid:"ipv4"`
    Date     string `valid:"-"`	
}
post := &Post{"My Example Post", "duck", "123.234.54.3"}

//Add your own struct validation tags
govalidator.TagMap["duck"] = govalidator.Validator(func(str string) bool {
    return str == "duck"
})

result, err := govalidator.ValidateStruct(post)
if err != nil {
    println("error: " + err.Error())
}
println(result)
```
###### WhiteList
```go
// Remove all characters from string ignoring characters between "a" and "z"
println(govalidator.WhiteList("a3a43a5a4a3a2a23a4a5a4a3a4", "a-z") == "aaaaaaaaaaaa")
```

#### Notes
Documentation is available here: [godoc.org](https://godoc.org/github.com/asaskevich/govalidator).
Full information about code coverage is also available here: [govalidator on gocover.io](http://gocover.io/github.com/asaskevich/govalidator).

#### Support
If you do have a contribution for the package feel free to put up a Pull Request or open Issue.


[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/asaskevich/govalidator/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

