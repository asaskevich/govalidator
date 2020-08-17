package govalidator

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
