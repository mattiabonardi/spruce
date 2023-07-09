package utils

import "testing"

func TestExadecimalIdRegex(t *testing.T) {
	regexManager := RegexManager{}
	if !regexManager.Match("507f191e810c19729de86dea", EXADECIMAL_PATTERN) {
		t.Fatal("Regex not match")
	}
	if regexManager.Match("507f191e810c19729de86deX", EXADECIMAL_PATTERN) {
		t.Fatal("Regex match but character X isn't valid")
	}
}
