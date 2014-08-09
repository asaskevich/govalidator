package govalidator

import "testing"

func TestToInt(t *testing.T) {
	tests := []string{"1000", "-123", "abcdef", "100000000000000000000000000000000000000000000"}
	expected := []int64{1000, -123, 0, 0}
	for i := 0; i < len(tests); i++ {
		result, _ := ToInt(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func TestToBoolean(t *testing.T) {
	tests := []string{"true", "1", "True", "false", "0", "abcdef"}
	expected := []bool{true, true, true, false, false, false}
	for i := 0; i < len(tests); i++ {
		res, _ := ToBoolean(tests[i])
		if res != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", expected)
			t.FailNow()
		}
	}
}

func TestToString(t *testing.T) {
	tests := []interface{}{"string", 100, -1.23, []int32{1, 2, 3}, struct{ Keys map[int]int }{Keys: map[int]int{1: 2, 3: 4}}}
	expected := []string{"\"string\"", "100", "-1.23", "[1,2,3]", ""}
	for i := 0; i < len(tests); i++ {
		res, _ := ToString(tests[i])
		if res != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", expected)
			t.FailNow()
		}
	}
}

func TestToFloat(t *testing.T) {
	tests := []string{"", "123", "-.01", "10.", "string", "1.23e3", ".23e10"}
	expected := []float64{0, 123, -0.01, 10.0, 0, 1230, 0.23e10}
	for i := 0; i < len(tests); i++ {
		res, _ := ToFloat(tests[i])
		if res != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", expected)
			t.FailNow()
		}
	}
}
