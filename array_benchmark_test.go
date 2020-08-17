package govalidator

import "testing"

func BenchmarkEach(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		acc := 0
		data := []interface{}{1, 2, 3, 4, 5}
		var fn Iterator = func(value interface{}, index int) {
			acc = acc + value.(int)
		}
		Each(data, fn)
	}
}

func BenchmarkMap(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		data := []interface{}{1, 2, 3, 4, 5}
		var fn ResultIterator = func(value interface{}, index int) interface{} {
			return value.(int) * 3
		}
		_ = Map(data, fn)
	}
}

func BenchmarkFind(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		findElement := 96
		data := []interface{}{1, 2, 3, 4, findElement, 5}
		var fn1 ConditionIterator = func(value interface{}, index int) bool {
			return value.(int) == findElement
		}
		_ = Find(data, fn1)
	}
}

func BenchmarkFilter(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		var fn ConditionIterator = func(value interface{}, index int) bool {
			return value.(int)%2 == 0
		}
		_ = Filter(data, fn)
	}
}

func BenchmarkCount(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		var fn ConditionIterator = func(value interface{}, index int) bool {
			return value.(int)%2 == 0
		}
		_ = Count(data, fn)
	}
}
