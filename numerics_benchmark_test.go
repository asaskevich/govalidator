package govalidator

import "testing"

func BenchmarkAbs(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = Abs(-123.3e1)
	}
}

func BenchmarkSign(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = Sign(-123.3e1)
	}
}

func BenchmarkIsNegative(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = IsNegative(-123.3e1)
	}
}

func BenchmarkIsPositive(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = IsPositive(-123.3e1)
	}
}

func BenchmarkIsNonNegative(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = IsNonNegative(-123.3e1)
	}
}

func BenchmarkIsNonPositive(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = IsNonPositive(-123.3e1)
	}
}

func BenchmarkInRangeInt(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = InRangeInt(10, -100, 100)
	}
}

func BenchmarkInRangeFloat32(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = InRangeFloat32(10, -100, 100)
	}
}

func BenchmarkInRangeFloat64(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = InRangeFloat64(10, -100, 100)
	}
}

func BenchmarkInRange(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = InRange("abc", "a", "cba")
	}
}

func BenchmarkIsWhole(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = IsWhole(123.132)
	}
}

func BenchmarkIsNatural(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = IsNatural(123.132)
	}
}
