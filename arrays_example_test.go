package govalidator

func ExampleSome() {
	data := []interface{}{1, 3, 5, 8}
	var fn ConditionIterator = func(value interface{}, index int) bool {
		return value.(int)%2 == 0
	}
	_ = Some(data, fn) // result = true
}

func ExampleEvery() {
	data := []interface{}{2, 4, 6, 8}
	var fn ConditionIterator = func(value interface{}, index int) bool {
		return value.(int)%2 == 0
	}
	_ = Every(data, fn) // result = true
}

func ExampleReduce() {
	data := []interface{}{1, 2, 3, 4}
	var fn ReduceIterator = func(result interface{}, current interface{}) interface{} {
		return result.(int) + current.(int)
	}
	_ = Reduce(data, fn, 0) // result = 10
}

func ExampleFilter() {
	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var fn ConditionIterator = func(value interface{}, index int) bool {
		return value.(int)%2 == 0
	}
	_ = Filter(data, fn) // result = []interface{}{2, 4, 6, 8, 10}
}

func ExampleCount() {
	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var fn ConditionIterator = func(value interface{}, index int) bool {
		return value.(int)%2 == 0
	}
	_ = Count(data, fn) // result = 5
}

func ExampleMap() {
	data := []interface{}{1, 2, 3, 4, 5}
	var fn ResultIterator = func(value interface{}, index int) interface{} {
		return value.(int) * 3
	}
	_ = Map(data, fn) // result = []interface{}{1, 6, 9, 12, 15}
}

func ExampleEach() {
	data := []interface{}{1, 2, 3, 4, 5}
	var fn Iterator = func(value interface{}, index int) {
		println(value.(int))
	}
	Each(data, fn)
}

func ExampleFind() {
	data := []interface{}{1, 2, 3, 4, 5}
	var fn ConditionIterator = func(value interface{}, index int) bool {
		return value.(int) == 4
	}
	_ = Find(data, fn) // result = 4
}
