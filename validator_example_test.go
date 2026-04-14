package govalidator

import "regexp"

func ExampleIsEmail() {
	_ = IsEmail("foo@example.com") // true
	_ = IsEmail("foo@bar")         // false
}

func ExampleValidateStruct_basic() {
	type User struct {
		Name  string `valid:"required,alpha"`
		Email string `valid:"required,email"`
		Age   int    `valid:"range(18|99)"`
	}

	user := User{
		Name:  "Bob",
		Email: "bob@example.com",
		Age:   21,
	}

	_, _ = ValidateStruct(user) // true, nil
}

func ExampleValidateMap_basic() {
	inputMap := map[string]interface{}{
		"name":  "Alice",
		"email": "alice@example.com",
	}

	validationMap := map[string]interface{}{
		"name":  "required,alpha",
		"email": "required,email",
	}

	_, _ = ValidateMap(inputMap, validationMap) // true, nil
}

func ExampleValidateStructAsync_basic() {
	type User struct {
		Email string `valid:"required,email"`
	}

	res, errs := ValidateStructAsync(User{Email: "foo@example.com"})

	_, _ = <-res, <-errs // true, nil
}

func ExampleValidateStruct_customTagMap() {
	TagMap["duck"] = Validator(func(str string) bool {
		return str == "duck"
	})

	type Post struct {
		Message string `valid:"duck"`
	}

	_, _ = ValidateStruct(Post{Message: "duck"}) // true, nil
}

func ExampleValidateStruct_customParamTagMap() {
	ParamTagMap["animal"] = ParamValidator(func(str string, params ...string) bool {
		return len(params) > 0 && str == params[0]
	})
	ParamTagRegexMap["animal"] = regexp.MustCompile("^animal\\((\\w+)\\)$")

	type Post struct {
		Message string `valid:"animal(dog)"`
	}

	_, _ = ValidateStruct(Post{Message: "dog"}) // true, nil
}
