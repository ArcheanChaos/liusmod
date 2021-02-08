package pkg1

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	ret := Hello("Gladys")
	if ret != name {
		t.Errorf("hello name is not equal")
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	ret := Hello("")
	if ret != "" {
		t.Errorf("hello name is not empty")
	}
}
