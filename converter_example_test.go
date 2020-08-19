package govalidator

import "time"

func ExampleToBoolean() {
	// Returns the boolean value represented by the string.
	// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
	// Any other value returns an error.
	_, _ = ToBoolean("false")  // false, nil
	_, _ = ToBoolean("T")      // true, nil
	_, _ = ToBoolean("123123") // false, error
}

func ExampleToInt() {
	_, _ = ToInt(1.0)     // 1, nil
	_, _ = ToInt("-124")  // -124, nil
	_, _ = ToInt("false") // 0, error
}

func ExampleToFloat() {
	_, _ = ToFloat("-124.2e123") // -124.2e123, nil
	_, _ = ToFloat("false")      // 0, error
}

func ExampleToString() {
	_ = ToString(new(interface{}))        // 0xc000090200
	_ = ToString(time.Second + time.Hour) // 1h1s
	_ = ToString(123)                     // 123
}

func ExampleToJSON() {
	_, _ = ToJSON([]int{1, 2, 3})          // [1, 2, 3]
	_, _ = ToJSON(map[int]int{1: 2, 2: 3}) // { "1": 2, "2": 3 }
	_, _ = ToJSON(func() {})               // error
}
