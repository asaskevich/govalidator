package govalidator

// Basic regular expressions for validating strings
const (
	Email string         = "^((([a-z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])+(\\.([a-z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])+)*)|((\x22)((((\x20|\x09)*(\x0d\x0a))?(\x20|\x09)+)?(([\x01-\x08\x0b\x0c\x0e-\x1f\x7f]|\x21|[\x23-\x5b]|[\x5d-\x7e]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(\\([\x01-\x09\x0b\x0c\x0d-\x7f]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]))))*(((\x20|\x09)*(\x0d\x0a))?(\x20|\x09)+)?(\x22)))@((([a-z]|\\d|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(([a-z]|\\d|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])([a-z]|\\d|-|\\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])*([a-z]|\\d|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])))\\.)+(([a-z]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(([a-z]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])([a-z]|\\d|-|\\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])*([a-z]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])))\\.?$"
	URL string           = "^(https?):\\/\\/(([a-z0-9$_\\.\\+!\\*\\'\\(\\),;\\?&=-]|%[0-9a-f]{2})+(:([a-z0-9$_\\.\\+!\\*\\'\\(\\),;\\?&=-]|%[0-9a-f]{2})+)?@)?(?#)((([a-z0-9]\\.|[a-z0-9][a-z0-9-]*[a-z0-9]\\.)*[a-z][a-z0-9-]*[a-z0-9]|((\\d|[1-9]\\d|1\\d{2}|2[0-4][0-9]|25[0-5])\\.){3}(\\d|[1-9]\\d|1\\d{2}|2[0-4][0-9]|25[0-5]))(:\\d+)?)(((\\/+([a-z0-9$_\\.\\+!\\*\\'\\(\\),;:@&=-]|%[0-9a-f]{2})*)*(\\?([a-z0-9$_\\.\\+!\\*\\'\\(\\),;:@&=-]|%[0-9a-f]{2})*?)?)?(#([a-z0-9$_\\.\\+!\\*\\'\\(\\),;:@&=-]|%[0-9a-f]{2})*)?$"
	CreditCard string    = "^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$"
	ISBN10 string        = "^(?:[0-9]{9}X|[0-9]{10})$"
	ISBN13 string        = "^(?:[0-9]{13})$"
	IPv4 string          = "^(\\d?\\d?\\d)\\.(\\d?\\d?\\d)\\.(\\d?\\d?\\d)\\.(\\d?\\d?\\d)$"
	IPv6 string          = "^::|^::1|^([a-fA-F0-9]{1,4}::?){1,7}([a-fA-F0-9]{1,4})$"
	UUID3 string         = "^[0-9A-F]{8}-[0-9A-F]{4}-3[0-9A-F]{3}-[0-9A-F]{4}-[0-9A-F]{12}$"
	UUID4 string         = "^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"
	UUID5 string         = "^[0-9A-F]{8}-[0-9A-F]{4}-5[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"
	UUID string          = "^[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}$"
	Alpha string         = "^[a-zA-Z]+$"
	Alphanumeric string  = "^[a-zA-Z0-9]+$"
	Numeric string       = "^-?[0-9]+$"
	Int string           = "^(?:-?(?:0|[1-9][0-9]*))$"
	Float string         = "^(?:-?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$"
	Hexadecimal string   = "^[0-9a-fA-F]+$"
	Hexcolor string      = "^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"
	ASCII string         = "^[\x00-\x7F]+$"
	Multibyte string     = "[^\x00-\x7F]"
	FullWidth string     = "[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	HalfWidth string     = "[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
)
