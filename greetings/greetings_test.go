package greetings

import (
	"regexp"
	"testing"
)


func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := Hello("Gladys")
	if !want.MatchString(message) || err != nil {
		t.Fatalf("error")
	}
}
