package govalidator

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

// Fixed inputs (valid where applicable) for stable microbenchmarks.
var (
	benchEmail          = "foo@example.com"
	benchEmailLocal     = "user@localhost"
	benchURL            = "http://example.com/path/to/page"
	benchRequestURL     = "http://user:pass@example.com/foo?bar=baz"
	benchRequestURI     = "/path/to/resource"
	benchAlpha          = "abcdefGHIJ"
	benchNumeric        = "1234567890"
	benchHex            = "deadbeef"
	benchHexColor       = "#aabbcc"
	benchRGB            = "rgb(10,20,30)"
	benchUUID           = "a987fbc9-4bed-3078-cf07-9141ba07c9f3"
	benchUUIDv3         = "a987fbc9-4bed-3078-cf07-9141ba07c9f3"
	benchUUIDv4         = "57b73598-8764-4ad0-a76a-679bb6640eb1"
	benchUUIDv5         = "987fbc97-4bed-5078-af07-9141ba07c9f3"
	benchULID           = "0123456789zzzzzzzzzzzzzzzz"
	benchCreditCard     = "4220855426222389"
	benchISBN10         = "3836221195"
	benchISBN13         = "9784873113685"
	benchJSON           = `{"a":1,"b":"x"}`
	benchBase64         = "TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4="
	benchJWT            = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJqb2UiLCJleHAiOjEzMDA4MTkzODAsImh0dHA6Ly9leGFtcGxlLmNvbS9pc19yb290Ijp0cnVlfQ.dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk"
	benchIPv4           = "127.0.0.1"
	benchIPv6           = "::1"
	benchCIDR           = "192.168.0.1/24"
	benchMAC            = "3D:F2:C9:A6:B3:4F"
	benchMongoID        = "507f1f77bcf86cd799439011"
	benchSemver         = "v1.2.3"
	benchRFC3339        = "2016-12-31T11:00:00Z"
	benchRFC3339NoZ     = "2016-12-31T11:00:00"
	benchDial           = "127.0.0.1:8080"
	benchWinPath        = `c:\path\file`
	benchUnixPath       = "/path/file/"
	benchSSN            = "191-60-2869"
	benchLat            = "23.123"
	benchLon            = "45.67"
	benchIMEI           = "123456789012345"
	benchIMSI           = "310150123456789"
	benchE164           = "+14155552671"
	benchYYYYMMDD       = "2000-01-01"
	benchSHA1           = "3ca25ae354e192b26879f651a51d92aa8a34d8d3"
	benchSHA256         = "579282cfb65ca1f109b78536effaf621b853c9f7079664a3fbe2b519f435898c"
	benchSHA384         = "bf547c3fc5841a377eb1519c2890344dbab15c40ae4150b4b34443d2212e5b04aa9d58865bf03d8ae27840fef430b891"
	benchSHA512         = "45bc5fa8cb45ee408c04b6269e9f1e1c17090c5ce26ffeeda2af097735b29953ce547e40ff3ad0d120e5361cc5f9cee35ea91ecd4077f3f589b4d439168f91b9"
	benchSHA3224        = "b87f88c72702fff1748e58b87e9141a42c0dbedc29a78cb0d4a5cd81"
	benchSHA3256        = "3338be694f50c5f338814986cdf0686453a888b84f424d792af4b9202398f392"
	benchSHA3384        = "720aea11019ef06440fbf05d87aa24680a2153df3907b23631e7177ce620fa1330ff07c0fddee54699a4c3ee0ee9d887"
	benchSHA3512        = "75d527c368f2efe848ecf6b073a36767800805e9eef2b1857d5f984f036eb6df891d75f72d9b154518c1cd58835286d1da9a38deba3de98b5a53e5ed78a84976"
	benchTiger192       = "46fc0125a148788a3ac1d649566fc04eb84a746f1a6e4fa7"
	benchTiger160       = "3ca25ae354e192b26879f651a51d92aa8a34d8d3"
	benchTiger128       = "579282cfb65ca1f109b78536effaf621"
	benchRipeMD160      = "3ca25ae354e192b26879f651a51d92aa8a34d8d3"
	benchRipeMD128      = "579282cfb65ca1f109b78536effaf621"
	benchCRC32          = "deadbeef"
	benchMD5            = "579282cfb65ca1f109b78536effaf621"
	benchMD4            = "579282cfb65ca1f109b78536effaf621"
	benchMagnet         = "magnet:?xt=urn:btih:06E2A9683BF4DA92C73A661AC56F0ECC9C63C5B4&dn=helloword2000&tr=udp://helloworld:1337/announce"
	benchDataURI        = "data:image/png;base64,TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4="
	benchDNS            = "example.com"
	benchHost           = "example.com:8080"
	benchPort           = "8080"
	benchISO4217        = "USD"
	benchISO3166A2      = "US"
	benchISO3166A3      = "USA"
	benchISO693A2       = "de"
	benchISO693A3b      = "mac"
	benchMultibyte      = "你好世界"
	benchFullWidth      = "ＡＢＣ"
	benchHalfWidth      = "ABC"
	benchVariableWidth  = "AＢC"
	benchPrintableASCII = "Hello World 123"
	benchNullStr        = ""
	benchNotNull        = "x"
	benchWhitespace     = "a b\tc"
	benchWhitespaceOnly = "   \t\n"
	benchIntStr         = "-42"
	benchFloatStr       = "3.14"
	benchDivA           = "100"
	benchDivB           = "10"
	benchByteLenStr     = "hello"
	benchInList         = "a|b|c"
	benchRegexPat       = "^[a-z]+$"
	benchRSAPublicKey   = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvncDCeibmEkabJLmFec7
x9y86RP6dIvkVxxbQoOJo06E+p7tH6vCmiGHKnuuXwKYLq0DKUE3t/HHsNdowfD9
+NH8caLzmXqGBx45/Dzxnwqz0qYq7idK+Qff34qrk/YFoU7498U1Ee7PkKb7/VE9
BmMEcI3uoKbeXCbJRIHoTp8bUXOpNTSUfwUNwJzbm2nsHo2xu6virKtAZLTsJFzT
UmRd11MrWCvj59lWzt1/eIMN+ekjH8aXeLOOl54CL+kWp48C+V9BchyKCShZB7uc
imFvjHTtuxziXZQRO7HlcsBOa0WwvDJnRnskdyoD31s4F4jpKEYBJNWTo63v6lUv
bQIDAQAB
-----END PUBLIC KEY-----`
)

type benchValidateStruct struct {
	Title    string `valid:"alphanum,required"`
	AuthorIP string `valid:"ipv4"`
}

var (
	benchValidateStructOK = benchValidateStruct{Title: "MyPost", AuthorIP: "192.168.1.1"}
	benchMapInput         = map[string]interface{}{
		"name":    "Bob",
		"family":  "Smith",
		"email":   "foo@bar.baz",
		"address": map[string]interface{}{"line1": "123 Main", "line2": "", "postal-code": "12345"},
	}
	benchMapSchema = map[string]interface{}{
		"name":    "required,alpha",
		"family":  "required,alpha",
		"email":   "required,email",
		"address": map[string]interface{}{"line1": "required,alphanum", "line2": "alphanum", "postal-code": "numeric"},
	}
	benchArray = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	benchCond  = func(value interface{}, index int) bool { return value.(int)%2 == 0 }
)

func BenchmarkSetFieldsRequiredByDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetFieldsRequiredByDefault(true)
		SetFieldsRequiredByDefault(false)
	}
}

func BenchmarkSetNilPtrAllowedByRequired(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetNilPtrAllowedByRequired(true)
		SetNilPtrAllowedByRequired(false)
	}
}

func BenchmarkIsEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsEmail(benchEmail)
	}
}

func BenchmarkIsExistingEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsExistingEmail(benchEmailLocal)
	}
}

func BenchmarkIsURL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsURL(benchURL)
	}
}

func BenchmarkIsRequestURL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsRequestURL(benchRequestURL)
	}
}

func BenchmarkIsRequestURI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsRequestURI(benchRequestURI)
	}
}

func BenchmarkIsAlpha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsAlpha(benchAlpha)
	}
}

func BenchmarkIsUTFLetter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsUTFLetter("こんにちは")
	}
}

func BenchmarkIsAlphanumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsAlphanumeric("abc123")
	}
}

func BenchmarkIsUTFLetterNumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsUTFLetterNumeric("αβγ123")
	}
}

func BenchmarkIsNumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsNumeric(benchNumeric)
	}
}

func BenchmarkIsUTFNumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsUTFNumeric("１２３")
	}
}

func BenchmarkIsUTFDigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsUTFDigit("１２３")
	}
}

func BenchmarkIsHexadecimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsHexadecimal(benchHex)
	}
}

func BenchmarkIsHexcolor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsHexcolor(benchHexColor)
	}
}

func BenchmarkIsRGBcolor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsRGBcolor(benchRGB)
	}
}

func BenchmarkIsLowerCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsLowerCase("abcxyz")
	}
}

func BenchmarkIsUpperCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsUpperCase("ABCXYZ")
	}
}

func BenchmarkHasLowerCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HasLowerCase("AbC")
	}
}

func BenchmarkHasUpperCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HasUpperCase("aBc")
	}
}

func BenchmarkIsInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsInt(benchIntStr)
	}
}

func BenchmarkIsFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsFloat(benchFloatStr)
	}
}

func BenchmarkIsDivisibleBy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsDivisibleBy(benchDivA, benchDivB)
	}
}

func BenchmarkIsNull(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsNull(benchNullStr)
	}
}

func BenchmarkIsNotNull(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsNotNull(benchNotNull)
	}
}

func BenchmarkHasWhitespaceOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HasWhitespaceOnly(benchWhitespaceOnly)
	}
}

func BenchmarkHasWhitespace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HasWhitespace(benchWhitespace)
	}
}

func BenchmarkIsByteLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsByteLength(benchByteLenStr, 1, 10)
	}
}

func BenchmarkIsUUIDv3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsUUIDv3(benchUUIDv3)
	}
}

func BenchmarkIsUUIDv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsUUIDv4(benchUUIDv4)
	}
}

func BenchmarkIsUUIDv5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsUUIDv5(benchUUIDv5)
	}
}

func BenchmarkIsUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsUUID(benchUUID)
	}
}

func BenchmarkIsULID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsULID(benchULID)
	}
}

func BenchmarkIsCreditCard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsCreditCard(benchCreditCard)
	}
}

func BenchmarkIsISBN10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsISBN10(benchISBN10)
	}
}

func BenchmarkIsISBN13(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsISBN13(benchISBN13)
	}
}

func BenchmarkIsISBN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsISBN(benchISBN13, 13)
	}
}

func BenchmarkIsJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsJSON(benchJSON)
	}
}

func BenchmarkIsMultibyte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsMultibyte(benchMultibyte)
	}
}

func BenchmarkIsASCII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsASCII(benchAlpha)
	}
}

func BenchmarkIsPrintableASCII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsPrintableASCII(benchPrintableASCII)
	}
}

func BenchmarkIsFullWidth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsFullWidth(benchFullWidth)
	}
}

func BenchmarkIsHalfWidth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsHalfWidth(benchHalfWidth)
	}
}

func BenchmarkIsVariableWidth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsVariableWidth(benchVariableWidth)
	}
}

func BenchmarkIsBase64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsBase64(benchBase64)
	}
}

func BenchmarkIsJWT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsJWT(benchJWT)
	}
}

func BenchmarkIsFilePath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = IsFilePath(benchUnixPath)
	}
}

func BenchmarkIsWinFilePath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsWinFilePath(benchWinPath)
	}
}

func BenchmarkIsUnixFilePath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsUnixFilePath(benchUnixPath)
	}
}

func BenchmarkIsDataURI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsDataURI(benchDataURI)
	}
}

func BenchmarkIsMagnetURI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsMagnetURI(benchMagnet)
	}
}

func BenchmarkIsISO3166Alpha2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsISO3166Alpha2(benchISO3166A2)
	}
}

func BenchmarkIsISO3166Alpha3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsISO3166Alpha3(benchISO3166A3)
	}
}

func BenchmarkIsISO693Alpha2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsISO693Alpha2(benchISO693A2)
	}
}

func BenchmarkIsISO693Alpha3b(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsISO693Alpha3b(benchISO693A3b)
	}
}

func BenchmarkIsDNSName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsDNSName(benchDNS)
	}
}

func BenchmarkIsHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsHash(benchSHA256, "sha256")
	}
}

func BenchmarkIsSHA3224(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsSHA3224(benchSHA3224)
	}
}

func BenchmarkIsSHA3256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsSHA3256(benchSHA3256)
	}
}

func BenchmarkIsSHA3384(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsSHA3384(benchSHA3384)
	}
}

func BenchmarkIsSHA3512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsSHA3512(benchSHA3512)
	}
}

func BenchmarkIsSHA512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsSHA512(benchSHA512)
	}
}

func BenchmarkIsSHA384(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsSHA384(benchSHA384)
	}
}

func BenchmarkIsSHA256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsSHA256(benchSHA256)
	}
}

func BenchmarkIsTiger192(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsTiger192(benchTiger192)
	}
}

func BenchmarkIsTiger160(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsTiger160(benchTiger160)
	}
}

func BenchmarkIsRipeMD160(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsRipeMD160(benchRipeMD160)
	}
}

func BenchmarkIsSHA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsSHA1(benchSHA1)
	}
}

func BenchmarkIsTiger128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsTiger128(benchTiger128)
	}
}

func BenchmarkIsRipeMD128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsRipeMD128(benchRipeMD128)
	}
}

func BenchmarkIsCRC32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsCRC32(benchCRC32)
	}
}

func BenchmarkIsCRC32b(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsCRC32b(benchCRC32)
	}
}

func BenchmarkIsMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsMD5(benchMD5)
	}
}

func BenchmarkIsMD4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsMD4(benchMD4)
	}
}

func BenchmarkIsDialString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsDialString(benchDial)
	}
}

func BenchmarkIsIP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsIP(benchIPv4)
	}
}

func BenchmarkIsPort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsPort(benchPort)
	}
}

func BenchmarkIsIPv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsIPv4(benchIPv4)
	}
}

func BenchmarkIsIPv6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsIPv6(benchIPv6)
	}
}

func BenchmarkIsCIDR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsCIDR(benchCIDR)
	}
}

func BenchmarkIsMAC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsMAC(benchMAC)
	}
}

func BenchmarkIsHost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsHost(benchHost)
	}
}

func BenchmarkIsMongoID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsMongoID(benchMongoID)
	}
}

func BenchmarkIsLatitude(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsLatitude(benchLat)
	}
}

func BenchmarkIsLongitude(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsLongitude(benchLon)
	}
}

func BenchmarkIsIMEI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsIMEI(benchIMEI)
	}
}

func BenchmarkIsIMSI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsIMSI(benchIMSI)
	}
}

func BenchmarkIsRsaPublicKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsRsaPublicKey(benchRSAPublicKey, 2048)
	}
}

func BenchmarkIsRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsRegex(benchRegexPat)
	}
}

func BenchmarkValidateArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ValidateArray(benchArray, benchCond)
	}
}

func BenchmarkValidateMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ValidateMap(benchMapInput, benchMapSchema)
	}
}

func BenchmarkValidateStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ValidateStruct(benchValidateStructOK)
	}
}

func BenchmarkValidateStructAsync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		okc, errc := ValidateStructAsync(benchValidateStructOK)
		<-okc
		<-errc
	}
}

func BenchmarkValidateMapAsync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		okc, errc := ValidateMapAsync(benchMapInput, benchMapSchema)
		<-okc
		<-errc
	}
}

func BenchmarkIsSSN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsSSN(benchSSN)
	}
}

func BenchmarkIsSemver(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsSemver(benchSemver)
	}
}

func BenchmarkIsType(b *testing.B) {
	v := 42
	for i := 0; i < b.N; i++ {
		_ = IsType(v, "int")
	}
}

func BenchmarkIsTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsTime(benchRFC3339, time.RFC3339)
	}
}

func BenchmarkIsUnixTime(b *testing.B) {
	s := "1609459200"
	for i := 0; i < b.N; i++ {
		_ = IsUnixTime(s)
	}
}

func BenchmarkIsRFC3339(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsRFC3339(benchRFC3339)
	}
}

func BenchmarkIsRFC3339WithoutZone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsRFC3339WithoutZone(benchRFC3339NoZ)
	}
}

func BenchmarkIsISO4217(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsISO4217(benchISO4217)
	}
}

func BenchmarkByteLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ByteLength(benchByteLenStr, "1", "10")
	}
}

func BenchmarkRuneLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RuneLength(benchByteLenStr, "1", "10")
	}
}

func BenchmarkStringLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = StringLength(benchByteLenStr, "1", "10")
	}
}

func BenchmarkMinStringLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MinStringLength(benchByteLenStr, "1")
	}
}

func BenchmarkMaxStringLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MaxStringLength(benchByteLenStr, "20")
	}
}

func BenchmarkRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Range("5", "1", "10")
	}
}

func BenchmarkIsInRaw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsInRaw("b", benchInList)
	}
}

func BenchmarkIsIn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsIn("b", "a", "b", "c")
	}
}

func BenchmarkIsRsaPub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsRsaPub(benchRSAPublicKey, "2048")
	}
}

func BenchmarkStringMatches(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = StringMatches("abc", benchRegexPat)
	}
}

func BenchmarkErrorByField(b *testing.B) {
	type bad struct {
		Email string `valid:"email"`
	}
	_, err := ValidateStruct(bad{Email: "x"})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ErrorByField(err, "Email")
	}
}

func BenchmarkErrorsByField(b *testing.B) {
	type bad struct {
		Email string `valid:"email"`
	}
	_, err := ValidateStruct(bad{Email: "x"})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ErrorsByField(err)
	}
}

func BenchmarkUnsupportedTypeError_Error(b *testing.B) {
	e := &UnsupportedTypeError{Type: reflect.TypeOf(map[int]string{})}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = e.Error()
	}
}

func BenchmarkErrors_Error(b *testing.B) {
	es := Errors{errors.New("a"), errors.New("b")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = es.Error()
	}
}

func BenchmarkError_Error(b *testing.B) {
	e := Error{Name: "field", Err: errors.New("oops"), Validator: "email"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = e.Error()
	}
}

func BenchmarkIsE164(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsE164(benchE164)
	}
}

func BenchmarkIsYYYYMMDD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsYYYYMMDD(benchYYYYMMDD)
	}
}
