package validator

import (
	"regexp"
	"testing"
)

func TestValidatorValid(t *testing.T) {
	v := New()
	if !v.Valid() {
		t.Error("want valid to be true")
	}

	v.AddError("key", "message")
	if v.Valid() {
		t.Error("want valid to be false")
	}
}

func TestValidatorAddError(t *testing.T) {
	v := New()
	v.AddError("key", "message 1")
	v.AddError("key", "message 2")

	if len(v.Errors) != 1 {
		t.Errorf("want map length 1; got %d", len(v.Errors))
	}

	if v.Errors["key"] != "message 1" {
		t.Errorf("want message 1; got %s", v.Errors["key"])
	}
}

func TestValidatorCheck(t *testing.T) {
	v := New()
	v.Check(true, "key", "message")
	if !v.Valid() {
		t.Error("want valid to be true")
	}

	v.Check(false, "key", "message")
	if v.Valid() {
		t.Error("want valid to be false")
	}
}

func TestMatches(t *testing.T) {
	rx := regexp.MustCompile("^[a-z]+$")

	if !Matches("abc", rx) {
		t.Error("want match")
	}
	if Matches("123", rx) {
		t.Error("want no match")
	}
}

func TestPermittedValue(t *testing.T) {
	if !PermittedValue("a", "a", "b", "c") {
		t.Error("want permitted")
	}
	if PermittedValue("d", "a", "b", "c") {
		t.Error("want not permitted")
	}
	if !PermittedValue(1, 1, 2, 3) {
		t.Error("want permitted int")
	}
}
