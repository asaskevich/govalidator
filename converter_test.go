package govalidator

import "testing"

func Test_ToInt(t *testing.T) {
	if ToInt("1000") != 1000 {
		t.FailNow()
	}
	if ToInt("-123") != -123 {
		t.FailNow()
	}
	if ToInt("abcdef") != 0 {
		t.FailNow()
	}
	if ToInt("1000000000000000000000000000000000000000000000000000000000") != 0 {
		t.FailNow()
	}
}

func Test_ToBoolean(t * testing.T) {
	if !ToBoolean("true") {
		t.FailNow()
	}
	if !ToBoolean("1") {
		t.FailNow()
	}
	if !ToBoolean("True") {
		t.FailNow()
	}
	if ToBoolean("false") {
		t.FailNow()
	}
	if ToBoolean("0") {
		t.FailNow()
	}
	if ToBoolean("abcdef") {
		t.FailNow()
	}
}

func Test_ToString(t *testing.T) {
	if ToString("string") != "\"string\"" {
		t.FailNow()
	}
	if ToString(100) != "100" {
		t.FailNow()
	}
	if ToString(-1.23) != "-1.23" {
		t.FailNow()
	}
	if ToString([]int32{1, 2, 3}) != "[1,2,3]" {
		t.FailNow()
	}
}

func Test_ToFloat(t *testing.T) {
	if ToFloat("") != 0 {
		t.FailNow()
	}
	if ToFloat("123") != 123 {
		t.FailNow()
	}
	if ToFloat("-.01") != -0.01 {
		t.FailNow()
	}
	if ToFloat("10.") != 10.0 {
		t.FailNow()
	}
	if ToFloat("string") != 0 {
		t.FailNow()
	}
	if ToFloat("1.23e3") != 1230 {
		t.FailNow()
	}
}
