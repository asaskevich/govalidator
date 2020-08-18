package govalidator

import "testing"

func BenchmarkToBoolean(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, _ = ToBoolean("True")
		_, _ = ToBoolean("False")
		_, _ = ToBoolean("")
	}
}

//ToString
//ToJSON
//ToFloat
//ToInt
