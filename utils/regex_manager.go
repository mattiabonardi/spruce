package utils

import "regexp"

const EXADECIMAL_PATTERN = "^[0-9a-fA-F]{24}$"

type RegexManager struct{}

func (h *RegexManager) Match(value string, pattern string) bool {
	match, _ := regexp.MatchString(pattern, value)
	if match {
		return true
	} else {
		return false
	}
}
