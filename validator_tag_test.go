package govalidator

import (
	"reflect"
	"testing"
)

func TestParseTagIntoMapCustomMessages(t *testing.T) {
	tests := []struct {
		name string
		tag  string
		want tagOptionsMap
	}{
		{"empty tag", "", tagOptionsMap{}},
		{"no message", "required", tagOptionsMap{"required": {"required", "", 0}}},
		{"empty message", "required~", tagOptionsMap{"required": {"required", "", 0}}},
		{"plain message", "required~cannot be blank", tagOptionsMap{"required": {"required", "cannot be blank", 0}}},
		{"embedded tilde", "required~cannot~be blank", tagOptionsMap{"required": {"required", "cannot~be blank", 0}}},
		{"leading tilde", "required~~cannot be blank", tagOptionsMap{"required": {"required", "~cannot be blank", 0}}},
		{"trailing tilde", "required~cannot be blank~", tagOptionsMap{"required": {"required", "cannot be blank~", 0}}},
		{"multiple tildes", "required~cannot~~be~blank", tagOptionsMap{"required": {"required", "cannot~~be~blank", 0}}},
		{"multiple rules", "required~cannot~be blank,length(2|5)~must~fit", tagOptionsMap{
			"required":    {"required", "cannot~be blank", 0},
			"length(2|5)": {"length(2|5)", "must~fit", 1},
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := parseTagIntoMap(test.tag); !reflect.DeepEqual(got, test.want) {
				t.Errorf("parseTagIntoMap(%q) = %#v, want %#v", test.tag, got, test.want)
			}
		})
	}
}

func TestValidateStructCustomMessagesWithTildes(t *testing.T) {
	type input struct {
		Embedded string `valid:"required~cannot~be blank"`
		Leading  string `valid:"required~~cannot be blank"`
		Trailing string `valid:"required~cannot be blank~"`
		Multiple string `valid:"required~cannot~~be~blank"`
		Plain    string `valid:"required~cannot be blank"`
		Empty    string `valid:"required~"`
		Default  string `valid:"required"`
		Length   string `valid:"length(2|5)~must~fit"`
	}

	valid, err := ValidateStruct(input{Length: "x"})
	if valid || err == nil {
		t.Fatalf("expected validation errors, got valid=%v, err=%v", valid, err)
	}
	want := map[string]string{
		"Embedded": "cannot~be blank",
		"Leading":  "~cannot be blank",
		"Trailing": "cannot be blank~",
		"Multiple": "cannot~~be~blank",
		"Plain":    "cannot be blank",
		"Empty":    "non zero value required",
		"Default":  "non zero value required",
		"Length":   "must~fit",
	}
	if got := ErrorsByField(err); !reflect.DeepEqual(got, want) {
		t.Errorf("ErrorsByField() = %#v, want %#v", got, want)
	}

	valid, err = ValidateStruct(input{
		Embedded: "value", Leading: "value", Trailing: "value", Multiple: "value",
		Plain: "value", Empty: "value", Default: "value", Length: "ok",
	})
	if !valid || err != nil {
		t.Errorf("expected valid input, got valid=%v, err=%v", valid, err)
	}
}
