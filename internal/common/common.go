package common

import "github.com/atotto/clipboard"

// checks if a value exists in a slice of strings
func IncludesString(s []string, val string) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}
	return false
}

// returns the index of given value in a slice if exists
func FindIndexOfString(s []string, val string) int {
	for i, v := range s {
		if v == val {
			return i
		}
	}
	return -1
}

// copies string to clipboard
func CopyToClipboard(s string) error {
	if err := clipboard.WriteAll(s); err != nil {
		return err
	}
	return nil
}
