package govalidator

func ExampleToBoolean() {
	// Returns the boolean value represented by the string.
	// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
	// Any other value returns an error.
	_, _ = ToBoolean("false")  // false, nil
	_, _ = ToBoolean("T")      // true, nil
	_, _ = ToBoolean("123123") // false, error
}

//ToString
//ToJSON
//ToFloat
//ToInt
