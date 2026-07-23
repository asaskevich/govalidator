package govalidator

import (
	"strings"
	"testing"
)

func TestToIntInvalidFormatMessage(t *testing.T) {
	_, err := ToInt("not-a-number")
	if err == nil {
		t.Fatal("expected error")
	}
	// Previously used %g on a string and printed "%!g(string=...)".
	if strings.Contains(err.Error(), "%!") {
		t.Fatalf("broken format verb in error: %v", err)
	}
	if !strings.Contains(err.Error(), "not-a-number") {
		t.Fatalf("error should include input, got %v", err)
	}
}

func TestToFloatUnknownTypeMessage(t *testing.T) {
	_, err := ToFloat(true)
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "ToFloat") {
		t.Fatalf("error should say ToFloat, got %v", err)
	}
	if strings.Contains(err.Error(), "ToInt:") {
		t.Fatalf("wrong function name in error: %v", err)
	}
}
