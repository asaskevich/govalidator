package govalidator

// Benchmark testing is produced with randomly filled array of 1 million elements

import (
	"math/rand"
	"testing"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomArray(n int) (res []interface{}) {
	res = make([]interface{}, n)

	for i := 0; i < n; i++ {
		res[i] = randomInt(-1000, 1000)
	}

	return
}

func BenchmarkSome(b *testing.B) {
	data := randomArray(1000000)
	var fn ConditionIterator = func(value interface{}, index int) bool {
		return value.(int)%2 == 0
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = Some(data, fn)
	}
}

func BenchmarkEvery(b *testing.B) {
	data := randomArray(1000000)
	var fn ConditionIterator = func(value interface{}, index int) bool {
		return value.(int)%2 == 0 || value.(int)%2 == 1
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = Every(data, fn)
	}
}

func BenchmarkReduce(b *testing.B) {
	data := randomArray(100000)
	var fn ReduceIterator = func(result interface{}, current interface{}) interface{} {
		return result.(int) + current.(int)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = Reduce(data, fn, 0)
	}
}

func BenchmarkEach(b *testing.B) {
	data := randomArray(1000000)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		acc := 0
		var fn Iterator = func(value interface{}, index int) {
			acc = acc + value.(int)
		}
		Each(data, fn)
	}
}

func BenchmarkMap(b *testing.B) {
	data := randomArray(1000000)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var fn ResultIterator = func(value interface{}, index int) interface{} {
			return value.(int) * 3
		}
		_ = Map(data, fn)
	}
}

func BenchmarkFind(b *testing.B) {
	data := randomArray(1000000)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		findElement := 96
		var fn1 ConditionIterator = func(value interface{}, index int) bool {
			return value.(int) == findElement
		}
		_ = Find(data, fn1)
	}
}

func BenchmarkFilter(b *testing.B) {
	data := randomArray(1000000)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var fn ConditionIterator = func(value interface{}, index int) bool {
			return value.(int)%2 == 0
		}
		_ = Filter(data, fn)
	}
}

func BenchmarkCount(b *testing.B) {
	data := randomArray(1000000)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var fn ConditionIterator = func(value interface{}, index int) bool {
			return value.(int)%2 == 0
		}
		_ = Count(data, fn)
	}
}
